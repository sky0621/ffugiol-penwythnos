package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
	. "github.com/sky0621/fs-mng-backend/src/models"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	records, err := Movies().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	var results []*model.Movie
	for _, record := range records {
		results = append(results, &model.Movie{
			ID:       record.ID,
			Name:     record.Name,
			MovieURL: record.Filename,
		})
	}
	return results, nil
}
