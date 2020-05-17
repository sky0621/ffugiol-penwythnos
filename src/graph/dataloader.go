package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/sky0621/fs-mng-backend/src/models"

	"github.com/sky0621/fs-mng-backend/src/graph/model"
)

const loadersKey = "dataLoaders"

type DataLoaders struct {
	ViewingHistoriesByMovieID *model.ViewingHistoryLoader
}

func DataLoaderMiddleware(resolver *Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &DataLoaders{
			ViewingHistoriesByMovieID: model.NewViewingHistoryLoader(model.ViewingHistoryLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(keys []string) ([][]*model.ViewingHistory, []error) {
					results, err := models.ViewingHistories(
						qm.Load(qm.Rels(models.ViewingHistoryRels.User)),
						qm.Load(qm.Rels(models.ViewingHistoryRels.Movie)),
						models.ViewingHistoryWhere.MovieID.IN(keys),
					).All(r.Context(), resolver.DB)
					errors := make([]error, len(keys))
					if err != nil {
						errors = append(errors, err)
					}

					var viewingHistorySlice []*model.ViewingHistory
					for _, result := range results {
						rec := &model.ViewingHistory{
							ID:        result.ID,
							CreatedAt: result.CreatedAt,
						}
						if result.R.User != nil {
							rec.Viewer = &model.Viewer{
								ID:       result.R.User.ID,
								Name:     result.R.User.Name,
								Nickname: result.R.User.Nickname.Ptr(),
							}
						}
						if result.R.Movie != nil {
							rec.Movie = &model.Movie{
								ID:       result.R.Movie.ID,
								Name:     result.R.Movie.Name,
								MovieURL: result.R.Movie.Filename,
								Scale:    result.R.Movie.Scale,
							}
						}
						viewingHistorySlice = append(viewingHistorySlice, rec)
					}
					return [][]*model.ViewingHistory{viewingHistorySlice}, nil
				},
			}),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *DataLoaders {
	return ctx.Value(loadersKey).(*DataLoaders)
}
