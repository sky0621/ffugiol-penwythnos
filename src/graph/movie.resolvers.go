package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/sky0621/fs-mng-backend/src/util"

	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/xerrors"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
	. "github.com/sky0621/fs-mng-backend/src/models"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (string, error) {
	// トランザクションを貼る
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return "", xerrors.Errorf("failed to BeginTx: %w", err)
	}
	defer func() {
		if tx != nil {
			// コミット発行されてなければ必ずロールバックされる
			if err := tx.Rollback(); err != nil {
				fmt.Println(err)
			}
		}
	}()

	// DB登録
	m := Movie{
		Name:     input.Name,
		Filename: input.MovieFile.Filename,
	}
	if err := m.Insert(ctx, r.DB, boil.Infer()); err != nil {
		// トランザクションロールバックされる
		return "", xerrors.Errorf("failed to Insert: %w", err)
	}

	// ファイルをCloud Storageにアップ
	if err := r.GCSClient.ExecUploadObject(input.MovieFile.Filename, input.MovieFile.File); err != nil {
		// トランザクションロールバックされる
		return "", xerrors.Errorf("failed to GCSClient.ExecUploadObject: %w", err)
	}

	if err := tx.Commit(); err != nil {
		// トランザクションロールバックされる
		return "", xerrors.Errorf("failed to Commit: %w", err)
	}

	return m.ID, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	records, err := Movies().All(ctx, r.DB)
	if err != nil {
		return nil, xerrors.Errorf("failed to Movies ALL: %w", err)
	}

	var results []*model.Movie
	for _, record := range records {
		url, err := r.GCSClient.ExecSignedURL(record.Filename, util.GetExpire(30*time.Second))
		if err != nil {
			return nil, xerrors.Errorf("failed to GCSClient.ExecSignedURL: %w", err)
		}
		results = append(results, &model.Movie{
			ID:       record.ID,
			Name:     record.Name,
			MovieURL: url,
		})
	}
	return results, nil
}
