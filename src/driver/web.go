package driver

import (
	"net/http"

	app "github.com/sky0621/fs-mng-backend"

	"github.com/rs/cors"

	"github.com/go-chi/chi"

	"github.com/sky0621/fs-mng-backend/adapter/controller"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewWeb(cfg app.Config, resolver controller.ResolverRoot, logger app.Logger) Web {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(c.Handler)

	r.Handle("/", playgroundHandler())
	r.Handle("/graphql", graphqlHandler(resolver, logger))

	return &web{cfg: cfg, router: r, logger: logger}
}

type Web interface {
	Start() error
}

type web struct {
	cfg    app.Config
	router chi.Router
	logger app.Logger
}

func (w *web) Start() error {
	lgr := w.logger.NewWrappedLogger("Start")

	lp := w.cfg.WebConfig.ListenPort
	lgr.Info().Str("ListenPort", lp).Send()
	if err := http.ListenAndServe(lp, w.router); err != nil {
		lgr.Err(err)
		return err
	}
	return nil
}

func playgroundHandler() http.HandlerFunc {
	handler.NewDefaultServer()
	h := handler.Playground("fiktivt-handelssystem-graphql-playground", "/graphql")
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func graphqlHandler(resolver controller.ResolverRoot) http.HandlerFunc {
	h := handler.GraphQL(
		controller.NewExecutableSchema(controller.Config{Resolvers: resolver}),
	)
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
