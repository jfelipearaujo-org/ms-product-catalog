package loader

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) GetEnvironmentFromFile(ctx context.Context, fileName string) (*environment.Config, error) {
	if err := godotenv.Load(fileName); err != nil {
		return nil, err
	}

	return l.GetEnvironment(ctx)
}

func (l *Loader) GetEnvironment(ctx context.Context) (*environment.Config, error) {
	var env environment.Config
	if err := envconfig.Process(ctx, &env); err != nil {
		return nil, err
	}

	return &env, nil
}
