package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"gopkg.in/auth0.v2/management"

	"github.com/sky0621/fs-mng-backend/src/auth"
	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

var (
	ConnectionDatabase = "Initial-Connection"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.MutationResponse, error) {
	user := auth.GetAuthenticatedUser(ctx)
	if !user.HasCreatePermission("user") {
		err := errors.New("no permissions")
		log.Print(err)
		return nil, err
	}

	m, err := r.Auth0Client.NewManagement(ctx)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	if m.User == nil {
		err := errors.New("no user")
		log.Print(err)
		return nil, err
	}

	users, err := m.User.ListByEmail(input.Email)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	if len(users) > 0 {
		err := errors.New("already exists")
		log.Print(err)
		return nil, err
	}

	u := &management.User{
		Name:       &input.Name,
		Email:      &input.Email,
		Connection: &ConnectionDatabase,
	}
	if err := m.User.Create(u); err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	return &model.MutationResponse{
		ID: u.ID,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := auth.GetAuthenticatedUser(ctx)
	if !user.HasReadAllPermission("user") {
		err := errors.New("no permissions")
		log.Print(err)
		return nil, err
	}

	m, err := r.Auth0Client.NewManagement(ctx)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	if m.User == nil {
		err := errors.New("no user")
		log.Print(err)
		return nil, err
	}

	users, err := m.User.List()
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	var results []*model.User
	for _, user := range users {
		mUser := &model.User{}
		if user.ID != nil {
			mUser.ID = *user.ID
		}
		if user.Name != nil {
			mUser.Name = *user.Name
		}
		if user.Email != nil {
			mUser.Email = *user.Email
		}
		results = append(results, mUser)
	}

	return results, nil
}
