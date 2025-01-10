package graph

import (
	"context"
	"tripatra-test-go/models"
)

type MutationResolver interface {
	AddUser(ctx context.Context, name string, email string) (*models.User, error)
}
