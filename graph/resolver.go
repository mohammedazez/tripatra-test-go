package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"errors"
	"tripatra-test-go/db"
	"tripatra-test-go/graph/generated"
	"tripatra-test-go/graph/model"
	"tripatra-test-go/models"
)

type Resolver struct{}

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, name string, email string) (*model.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email must be provided")
	}

	user := &models.User{
		Name:  name,
		Email: email,
	}
	createdUser, err := db.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    createdUser.ID.Hex(),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name *string, email *string) (*model.User, error) {
	updatedUser, err := db.UpdateUser(id, name, email)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    updatedUser.ID.Hex(),
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	err := db.DeleteUser(id)
	return err == nil, err
}

// AddProduct is the resolver for the addProduct field.
func (r *mutationResolver) AddProduct(ctx context.Context, name string, price float64, stock int) (*model.Product, error) {
	product := &models.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	// Create the product in the database (models.Product)
	createdProduct, err := db.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	// Convert to model.Product (GraphQL model)
	return &model.Product{
		ID:    createdProduct.ID.Hex(),
		Name:  createdProduct.Name,
		Price: createdProduct.Price,
		Stock: createdProduct.Stock,
	}, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, name *string, price *float64, stock *int) (*model.Product, error) {
	updatedProduct, err := db.UpdateProduct(id, name, price, stock)
	if err != nil {
		return nil, err
	}

	// Convert to model.Product (GraphQL model)
	return &model.Product{
		ID:    updatedProduct.ID.Hex(),
		Name:  updatedProduct.Name,
		Price: updatedProduct.Price,
		Stock: updatedProduct.Stock,
	}, nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	err := db.DeleteProduct(id)
	return err == nil, err
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := db.GetUser(id)
	if err != nil {
		return nil, err
	}

	// Convert to model.User (GraphQL model)
	return &model.User{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := db.GetUsers()
	if err != nil {
		return nil, err
	}

	// Convert to []*model.User (GraphQL model)
	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:    user.ID.Hex(),
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return result, nil
}

// GetProduct is the resolver for the getProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	product, err := db.GetProduct(id)
	if err != nil {
		return nil, err
	}

	// Convert to model.Product (GraphQL model)
	return &model.Product{
		ID:    product.ID.Hex(),
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}, nil
}

// GetProducts is the resolver for the getProducts field.
func (r *queryResolver) GetProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := db.GetProducts()
	if err != nil {
		return nil, err
	}

	// Convert to []*model.Product (GraphQL model)
	var result []*model.Product
	for _, product := range products {
		result = append(result, &model.Product{
			ID:    product.ID.Hex(),
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		})
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
