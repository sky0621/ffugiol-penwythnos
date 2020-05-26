package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sky0621/fs-mng-backend/src/gcp"

	"github.com/sky0621/fs-mng-backend/src/auth"

	"github.com/google/uuid"
	"github.com/sky0621/fs-mng-backend/src/graph/generated"
	"github.com/sky0621/fs-mng-backend/src/graph/model"
	. "github.com/sky0621/fs-mng-backend/src/models"
	"github.com/sky0621/fs-mng-backend/src/util"
	"github.com/volatiletech/sqlboiler/boil"
	_ "gocloud.dev/pubsub/gcppubsub"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (*model.MutationResponse, error) {
	user := auth.GetAuthenticatedUser(ctx)
	if !user.HasCreatePermission("content") {
		err := errors.New("no permissions")
		log.Print(err)
		return nil, err
	}

	// トランザクションを貼る
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	defer func() {
		if tx != nil {
			// コミット発行されてなければ必ずロールバックされる
			if err := tx.Rollback(); err != nil {
				log.Printf("%+v", err)
			}
		}
	}()

	// DB登録
	id := uuid.New().String()
	m := Movie{
		ID:       id,
		Name:     input.Name,
		Filename: input.MovieFile.Filename,
		Scale:    input.Scale,
	}
	if err := m.Insert(ctx, r.DB, boil.Infer()); err != nil {
		log.Printf("%+v", err)
		// トランザクションロールバックされる
		return nil, err
	}

	// ファイルをCloud Storageにアップ
	if err := r.GCSClient.ExecUploadObject(ctx, input.MovieFile.Filename, input.MovieFile.File); err != nil {
		log.Printf("%+v", err)
		// トランザクションロールバックされる
		return nil, err
	}

	metadata := map[string]string{
		"facility-id": "fid:0001",
	}
	jsonFormat := `
		{
			"id": "%s",
			"name": "%s",
			"filename": "%s",
			"scale": %d
		}
	`
	bodyJSON := fmt.Sprintf(jsonFormat, id, input.Name, input.MovieFile.Filename, input.Scale)

	// イベント発生を通知
	if err := r.PubSubClient.SendTopic(ctx, gcp.CreateMovieTopic, metadata, bodyJSON); err != nil {
		log.Printf("%+v", err)
		// トランザクションロールバックされる
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("%+v", err)
		// トランザクションロールバックされる
		return nil, err
	}

	return &model.MutationResponse{
		ID: &m.ID,
	}, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	user := auth.GetAuthenticatedUser(ctx)
	if !user.HasReadMinePermission("content") {
		err := errors.New("no permissions")
		log.Print(err)
		return nil, err
	}

	records, err := Movies().All(ctx, r.DB)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	var results []*model.Movie
	for _, record := range records {
		url, err := r.GCSClient.ExecSignedURL(record.Filename, util.GetExpire(30*time.Second))
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
		results = append(results, &model.Movie{
			ID:       record.ID,
			Name:     record.Name,
			MovieURL: url,
			Scale:    record.Scale,
		})
	}
	return results, nil
}

func (r *movieResolver) ViewingHistories(ctx context.Context, obj *model.Movie) ([]*model.ViewingHistory, error) {
	records, err := For(ctx).ViewingHistoriesByMovieID.Load(obj.ID)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	var results []*model.ViewingHistory
	for _, record := range records {
		results = append(results, &model.ViewingHistory{
			ID: record.ID,
			Viewer: &model.Viewer{
				ID: record.Viewer.ID,
			},
			Movie:     obj,
			CreatedAt: record.CreatedAt,
		})
	}
	return results, nil
}

// Movie returns generated.MovieResolver implementation.
func (r *Resolver) Movie() generated.MovieResolver { return &movieResolver{r} }

type movieResolver struct{ *Resolver }
