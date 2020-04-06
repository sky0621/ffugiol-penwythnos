package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fs-mng-backend/adapter/controller/model"
)

func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.OrganizationInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrganization(ctx context.Context, input model.OrganizationInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Organizations(ctx context.Context, condition *model.OrganizationCondition) ([]*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}
