package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fs-mng-backend/src/graph/generated"
	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *mutationResolver) CreateWorkHolder(ctx context.Context, input model.WorkHolderInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWorkHolder(ctx context.Context, input model.WorkHolderInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteWorkHolder(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WorkHolder(ctx context.Context, id string) (*model.WorkHolder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WorkHolders(ctx context.Context, condition *model.WorkHolderCondition) ([]*model.WorkHolder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *workHolderResolver) Organizations(ctx context.Context, obj *model.WorkHolder) ([]*model.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *workHolderResolver) HoldWorks(ctx context.Context, obj *model.WorkHolder) ([]*model.Work, error) {
	panic(fmt.Errorf("not implemented"))
}

// WorkHolder returns generated.WorkHolderResolver implementation.
func (r *Resolver) WorkHolder() generated.WorkHolderResolver { return &workHolderResolver{r} }

type workHolderResolver struct{ *Resolver }
