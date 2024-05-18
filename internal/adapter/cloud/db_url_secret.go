package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AwsDbUrlSecretService struct {
	SecretName string
	Client     *secretsmanager.Client
}

func NewSecretService(secretName string, config aws.Config) SecretService {
	return &AwsDbUrlSecretService{
		SecretName: secretName,
		Client:     secretsmanager.NewFromConfig(config),
	}
}

func (s *AwsDbUrlSecretService) GetSecret(ctx context.Context, secretName string) (string, error) {
	input := secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	output, err := s.Client.GetSecretValue(ctx, &input)
	if err != nil {
		return "", err
	}

	return *output.SecretString, nil
}
