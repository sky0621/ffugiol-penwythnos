package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/sky0621/fs-mng-backend/src/auth"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql/handler/transport"

	"github.com/sky0621/fs-mng-backend/src/gcp"

	"github.com/rs/cors"

	"github.com/go-chi/chi"

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/lib/pq"
	"github.com/sky0621/fs-mng-backend/src/graph"
	"github.com/sky0621/fs-mng-backend/src/graph/generated"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered: %v", r)
		}
	}()

	e := loadEnv()

	/*
	 * setup db client
	 */
	db, closeDBFunc := newDB(e)
	defer func() {
		if err := closeDBFunc(); err != nil {
			log.Fatal(err)
		}
	}()

	/*
	 * setup GCP client
	 */
	var projectID string
	{
		var err error
		projectID, err = gcp.GetProjectID()
		if err != nil {
			log.Printf("%+v\n", err)
			return
		}
	}
	var gcsClient gcp.CloudStorageClient
	{
		var err error
		gcsClient, err = gcp.NewCloudStorageClient(context.Background(), e.MovieBucket)
		if err != nil {
			log.Printf("%+v\n", err)
			return
		}
	}
	var pubSubClient gcp.PubSubClient
	{
		pubSubClient = gcp.NewPubSubClient(e.Env, projectID, e.CreateMovieTopic)
	}

	/*
	 * setup Auth0 client
	 */
	var auth0Client auth.Auth0Client
	{
		tm, err := time.ParseDuration(e.Auth0Timeout)
		if err != nil {
			log.Printf("%+v\n", err)
			return
		}
		auth0Client = auth.NewAuth0Client(e.Auth0Domain, e.Auth0ClientID, e.Auth0ClientSecret, tm, e.AuthDebug)
	}

	/*
	 * setup web server
	 */
	var router *chi.Mux
	{
		router = chi.NewRouter()
		router.Group(func(r chi.Router) {
			// CORS
			r.Use(cors.AllowAll().Handler)

			// 認証認可チェック用（今はJWTのチェックのみ実装）
			a := auth.New(e.Auth0Domain, e.Auth0Audience, e.AuthDebug, e.AuthCredentialsOptional)
			r.Use(a.CheckJWTHandlerFunc())
			r.Use(a.HoldPermissionsHandler)

			// GraphQLリゾルバー
			resolver := &graph.Resolver{
				DB:           db,
				GCSClient:    gcsClient,
				PubSubClient: pubSubClient,
				Auth0Client:  auth0Client,
			}

			// GraphQLエンドポイント（DataLoaderでラップ）
			r.Handle("/query", graph.DataLoaderMiddleware(resolver, graphQlServer(resolver)))

			// GraphQLドキュメント
			r.Handle("/", playground.Handler("fs-mng-backend", "/query"))
		})
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", e.ServerPort)
	if err := http.ListenAndServe(":"+e.ServerPort, router); err != nil {
		log.Printf("%+v\n", err)
		return
	}
}

func graphQlServer(resolver *graph.Resolver) *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     128 * mb,
		MaxUploadSize: 100 * mb,
	})

	return srv
}
