package tests

import (
	"bytes"
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
type feature struct {
	URL    string
	Body   string
	Status int
}

func enrichContext(ctx context.Context, feat feature) (context.Context, error) {
	return context.WithValue(ctx, feature{}, feat), nil
}

func fromContext(ctx context.Context) (feature, error) {
	val := ctx.Value(feature{})
	if val == nil {
		return feature{}, fmt.Errorf("value not found in context")
	}

	return val.(feature), nil
}

func aCategory(ctx context.Context, categoryType string) (context.Context, error) {
	feat, err := fromContext(ctx)
	if err != nil {
		return ctx, err
	}

	if categoryType == "valid" {
		feat.Body = `{"title": "Test Category", "description": "Test Description"}`
	} else {
		feat.Body = `{"title": "Test Category"}`
	}

	return enrichContext(ctx, feat)
}

func iCreateTheCategory(ctx context.Context) (context.Context, error) {
	feat, err := fromContext(ctx)
	if err != nil {
		return ctx, err
	}

	route := fmt.Sprintf("%s/categories", feat.URL)
	req, err := http.NewRequest(http.MethodPost, route, bytes.NewBuffer([]byte(feat.Body)))
	if err != nil {
		return ctx, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ctx, err
	}

	feat.Status = resp.StatusCode

	return enrichContext(ctx, feat)
}

func theCategoryIsCreated(ctx context.Context) (context.Context, error) {
	feat, err := fromContext(ctx)
	if err != nil {
		return ctx, err
	}

	if feat.Status != http.StatusCreated {
		return ctx, fmt.Errorf("Expected status code 201, got %d", feat.Status)
	}

	return ctx, nil
}

func theCategoryIsNotCreated(ctx context.Context) (context.Context, error) {
	feat, err := fromContext(ctx)
	if err != nil {
		return ctx, err
	}

	if feat.Status != http.StatusUnprocessableEntity {
		return ctx, fmt.Errorf("Expected status code 422, got %d", feat.Status)
	}

	return ctx, nil
}

type testContext struct {
	network    *testcontainers.DockerNetwork
	containers []testcontainers.Container
}

var (
	containers = make(map[string]testContext)
)

func InitializeScenario(ctx *godog.ScenarioContext) {

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
				WaitingFor: wait.ForLog("Waiting for connections"),
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

		if len(ports["8080/tcp"]) == 0 {
			return ctx, fmt.Errorf("Port 8080/tcp not found")
		}

		port := ports["8080/tcp"][0].HostPort

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

		return enrichContext(ctx, feature{
			URL: fmt.Sprintf("http://localhost:%s/api/v1", port),
		})
	})

	ctx.Step(`^a "([^"]*)" category$`, aCategory)
	ctx.Step(`^I create the category$`, iCreateTheCategory)
	ctx.Step(`^the category is created$`, theCategoryIsCreated)
	ctx.Step(`^the category is not created$`, theCategoryIsNotCreated)

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
