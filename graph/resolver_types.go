package graph

import (
	"context"
	"tripatra-test-go/models"
)

type QueryResolver interface {
	getUser(ctx context.Context, id string) (*models.User, error)
}

type MutationResolver interface {
	AddUser(ctx context.Context, name string, email string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, name *string, email *string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
}
