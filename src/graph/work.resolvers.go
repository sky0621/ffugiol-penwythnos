package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fs-mng-backend/src/models"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *mutationResolver) CreateWork(ctx context.Context, input model.WorkInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWork(ctx context.Context, input model.WorkInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteWork(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Work(ctx context.Context, id string) (*model.Work, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Works(ctx context.Context, condition *model.WorkCondition) ([]*model.Work, error) {
	works, err := models.Works().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}
	var results []*model.Work
	for _, work := range works {
		results = append(results, &model.Work{
			ID:   work.ID,
			Name: work.Name,
		})
	}
	return results, nil
}
