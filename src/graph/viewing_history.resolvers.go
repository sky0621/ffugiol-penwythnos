package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	. "github.com/sky0621/fs-mng-backend/src/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/xerrors"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

func (r *mutationResolver) RecordViewingHistory(ctx context.Context, input model.RecordViewingHistoryInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ViewingHistories(ctx context.Context, userID *string, movieID *string) ([]*model.ViewingHistory, error) {
	var mods []qm.QueryMod
	if userID != nil {
		mods = append(mods, ViewingHistoryWhere.UserID.EQ(*userID))
	}
	if movieID != nil {
		mods = append(mods, ViewingHistoryWhere.MovieID.EQ(*movieID))
	}
	records, err := ViewingHistories(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, xerrors.Errorf("failed to ViewingHistories ALL [user_id:%v][movie_id:%v]: %w", userID, movieID, err)
	}

	var results []*model.ViewingHistory
	for _, record := range records {
		results = append(results, &model.ViewingHistory{
			ID: record.ID,
			Viewer: &model.Viewer{
				ID: record.UserID,
			},
			Movie: &model.Movie{
				ID: record.MovieID,
			},
			CreatedAt: record.CreatedAt,
		})
	}
	return results, nil
}
