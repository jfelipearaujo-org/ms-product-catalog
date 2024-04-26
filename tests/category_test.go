package tests

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var opts = godog.Options{
	Format:      "pretty",
	Paths:       []string{"features"},
	Output:      colors.Colored(os.Stdout),
	Concurrency: 4,
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &o,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

// Steps
type appFeature struct {
	url string
}

func (af *appFeature) aCategory(ctx context.Context, arg1 string) (context.Context, error) {
	res, err := http.Get(fmt.Sprintf("%s/categories", af.url))
	if err != nil {
		return ctx, err
	}

	if res.StatusCode != http.StatusOK {
		return ctx, fmt.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	return ctx, nil
}

func (af *appFeature) iCreateTheCategory(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func (af *appFeature) theCategoryIsCreated(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func (af *appFeature) theCategoryIsNotCreated(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

type testContext struct {
	network    *testcontainers.DockerNetwork
	containers []testcontainers.Container
}

var (
	containers = make(map[string]testContext)
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	af := &appFeature{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Create a network
		network, err := network.New(ctx, network.WithCheckDuplicate(), network.WithDriver("bridge"))
		if err != nil {
			return ctx, err
		}

		// Start a MongoDB container
		mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "mongo:7",
				ExposedPorts: []string{
					"27017",
				},
				Env: map[string]string{
					"MONGO_INITDB_ROOT_USERNAME": "product",
					"MONGO_INITDB_ROOT_PASSWORD": "product",
				},
				Networks: []string{
					network.Name,
				},
				NetworkAliases: map[string][]string{
					network.Name: {"mongo"},
				},
				WaitingFor: wait.ForListeningPort("27017"),
			},
			Started: true,
		})
		if err != nil {
			return ctx, err
		}

		mongoIp, err := mongoContainer.Host(ctx)
		if err != nil {
			return ctx, err
		}

		mongoPort, err := mongoContainer.MappedPort(ctx, "27017")
		if err != nil {
			return ctx, err
		}

		connStr := fmt.Sprintf("mongodb://product:product@%s:%s/", mongoIp, mongoPort.Port())

		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
		if err != nil {
			return ctx, err
		}

		err = mongoClient.Ping(ctx, nil)
		if err != nil {
			return ctx, err
		}

		// Start a API container
		apiContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				FromDockerfile: testcontainers.FromDockerfile{
					Context:    "../",
					Dockerfile: "Dockerfile",
					KeepImage:  true,
				},
				ExposedPorts: []string{
					"8080",
				},
				Env: map[string]string{
					"API_PORT":     "8080",
					"API_ENV_NAME": "development",
					"API_VERSION":  "v1",
					"DB_NAME":      "product_db",
					"DB_URL":       "mongodb://product:product@mongo:27017",
				},
				Networks: []string{
					network.Name,
				},
				WaitingFor: wait.ForLog("Server started address=:8080"),
			},
			Started: true,
		})

		if err != nil {
			return ctx, err
		}

		ports, err := apiContainer.Ports(ctx)
		if err != nil {
			return ctx, err
		}

		port := ports["8080/tcp"][0].HostPort

		af.url = fmt.Sprintf("http://localhost:%s/api/v1", port)

		// Run health check
		res, err := http.Get(fmt.Sprintf("http://localhost:%s/health", port))
		if err != nil {
			return ctx, err
		}

		if res.StatusCode != http.StatusOK {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return ctx, err
			}
			defer res.Body.Close()

			fmt.Printf("Body: %s", string(body))

			return ctx, fmt.Errorf("API health check failed with status: %d", res.StatusCode)
		}

		containers[sc.Id] = testContext{
			network: network,
			containers: []testcontainers.Container{
				mongoContainer,
				apiContainer,
			},
		}

		return ctx, nil
	})

	ctx.Step(`^a "([^"]*)" category$`, af.aCategory)
	ctx.Step(`^I create the category$`, af.iCreateTheCategory)
	ctx.Step(`^the category is created$`, af.theCategoryIsCreated)
	ctx.Step(`^the category is not created$`, af.theCategoryIsNotCreated)

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			return ctx, err
		}

		tc := containers[sc.Id]

		for _, c := range tc.containers {
			err := c.Terminate(ctx)
			if err != nil {
				return ctx, err
			}
		}

		err = tc.network.Remove(ctx)

		return ctx, err
	})
}
