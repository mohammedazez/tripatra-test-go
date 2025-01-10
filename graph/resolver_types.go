package graph

import (
	"context"
	"tripatra-test-go/models"
)

type QueryResolver interface {
	getUser(ctx context.Context, id string) (*models.User, error)
	getUsers(ctx context.Context) ([]*models.User, error)
	getProduct(ctx context.Context, id string) (*models.Product, error)
	getProducts(ctx context.Context) ([]*models.Product, error)
}

type MutationResolver interface {
	AddUser(ctx context.Context, name string, email string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, name *string, email *string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)

	AddProduct(ctx context.Context, name string, price float64, stock int) (*models.Product, error)
	UpdateProduct(ctx context.Context, id string, name *string, price *float64, stock *int) (*models.Product, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
}
