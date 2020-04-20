package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *queryResolver) Viewer(ctx context.Context, id string) (*model.Viewer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Viewers(ctx context.Context) ([]*model.Viewer, error) {
	panic(fmt.Errorf("not implemented"))
}
