package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *mutationResolver) RecordViewingHistory(ctx context.Context, input model.RecordViewingHistoryInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ViewingHistories(ctx context.Context, userID *string, movieID *string) ([]*model.ViewingHistory, error) {
	panic(fmt.Errorf("not implemented"))
}
