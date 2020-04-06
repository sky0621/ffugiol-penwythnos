package system

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/sky0621/fs-mng-backend/adapter/controller"
)

func NewApp(cfg Config, resolver controller.ResolverRoot) App {
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

	//r.Handle("/", playgroundHandler())
	//r.Handle("/graphql", graphqlHandler(resolver, logger))

	return App{cfg: cfg, router: r}
}

type App struct {
	cfg    Config
	router chi.Router
}

func (a App) Start() error {
	lp := a.cfg.ListenPort
	if err := http.ListenAndServe(lp, a.router); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a App) Shutdown() {
	if a != nil && a.
}
