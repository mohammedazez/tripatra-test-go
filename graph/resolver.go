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

	// Convert to model.User (GraphQL model)
	return &model.User{
		ID:    createdUser.ID.Hex(),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
