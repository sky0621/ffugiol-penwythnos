package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/fs-mng-backend/src/util"

	. "github.com/sky0621/fs-mng-backend/src/models"
	"golang.org/x/xerrors"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *queryResolver) Viewer(ctx context.Context, id string) (*model.Viewer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Viewers(ctx context.Context) ([]*model.Viewer, error) {
	records, err := Viewers().All(ctx, r.DB)
	if err != nil {
		log.Print(xerrors.Unwrap(err))
		return nil, err
	}

	var results []*model.Viewer
	for _, record := range records {

		results = append(results, &model.Viewer{
			ID:       record.ID,
			Name:     record.Name,
			Nickname: util.ToNullableString(record.Nickname),
		})
	}
	return results, nil
}
