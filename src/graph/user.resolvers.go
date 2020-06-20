package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.MutationResponse, error) {
	// FIXME:
	id := "xxxxxxxx"
	return &model.MutationResponse{
		ID: &id,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// FIXME:
	return []*model.User{}, nil
}
