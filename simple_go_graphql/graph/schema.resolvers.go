package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/abbyb97/simple_go_graphql/graph/generated"
	"github.com/abbyb97/simple_go_graphql/graph/model"
	"github.com/abbyb97/simple_go_graphql/repository"
)

var userRepo repository.UserRepository = repository.New()

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:        strconv.Itoa(rand.Int()),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Age:       input.Age,
		Phone:     input.Phone,
		Address:   input.Address,
	}
	userRepo.Save(user)
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	user := &model.User{
		ID:        input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		Address:   input.Address,
	}
	userRepo.UpdateUser(input.ID, user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return userRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
