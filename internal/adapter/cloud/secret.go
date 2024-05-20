package cloud

import "context"

type SecretService interface {
	GetSecret(ctx context.Context, secretName string) (string, error)
}
