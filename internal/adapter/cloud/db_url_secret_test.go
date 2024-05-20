package cloud

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
	"github.com/stretchr/testify/assert"
)

func TestGetSecret(t *testing.T) {
	t.Run("Should return secret", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		stubber := testtools.NewStubber()

		stubber.Add(testtools.Stub{
			OperationName: "GetSecretValue",
			Input: &secretsmanager.GetSecretValueInput{
				SecretId: aws.String("my-secret"),
			},
			Output: &secretsmanager.GetSecretValueOutput{
				SecretString: aws.String("my-secret-value"),
			},
		})

		service := NewSecretService(*stubber.SdkConfig)

		// Act
		res, err := service.GetSecret(ctx, "my-secret")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "my-secret-value", res)
		testtools.ExitTest(stubber, t)
	})

	t.Run("Should return empty string if secret not found", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		stubber := testtools.NewStubber()

		stubber.Add(testtools.Stub{
			OperationName: "GetSecretValue",
			Input: &secretsmanager.GetSecretValueInput{
				SecretId: aws.String("my-secret"),
			},
			Output: &secretsmanager.GetSecretValueOutput{
				SecretString: aws.String(""),
			},
		})

		service := NewSecretService(*stubber.SdkConfig)

		// Act
		res, err := service.GetSecret(ctx, "my-secret")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "", res)
		testtools.ExitTest(stubber, t)
	})

	t.Run("Should return error if error", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		stubber := testtools.NewStubber()

		raiseErr := &testtools.StubError{Err: errors.New("ClientError")}

		stubber.Add(testtools.Stub{
			OperationName: "GetSecretValue",
			Input: &secretsmanager.GetSecretValueInput{
				SecretId: aws.String("my-secret"),
			},
			Error: raiseErr,
		})

		service := NewSecretService(*stubber.SdkConfig)

		// Act
		res, err := service.GetSecret(ctx, "my-secret")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "", res)
		testtools.ExitTest(stubber, t)
	})
}
