package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
	. "github.com/sky0621/fs-mng-backend/src/models"

	"cloud.google.com/go/storage"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (string, error) {
	// トランザクションを貼る
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return "", err
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
		return "", err // トランザクションロールバックされる
	}

	// ファイルをCloud Storageにアップ
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err // トランザクションロールバックされる
	}
	writer := client.Bucket(r.Bucket).Object(input.MovieFile.Filename).NewWriter(ctx)
	if _, err := io.Copy(writer, input.MovieFile.File); err != nil {
		return "", err // トランザクションロールバックされる
	}
	defer func() {
		if writer != nil {
			if err := writer.Close(); err != nil {
				fmt.Println(err)
			}
		}
	}()

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return m.ID, nil
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
			MovieURL: record.Filename, // TODO: 署名付きURLに差し替える
		})
	}
	return results, nil
}
