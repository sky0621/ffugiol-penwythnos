package main

import (
	"context"
	"fmt"
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
	var gcsClient gcp.CloudStorageClient
	{
		var err error
		gcsClient, err = gcp.NewCloudStorageClient(context.Background(), e.MovieBucket)
		if err != nil {
			panic(err)
		}
	}

	/*
	 * setup Auth0 client
	 */
	var auth0Client auth.Auth0Client
	{
		tm, err := time.ParseDuration(e.Auth0Timeout)
		if err != nil {
			panic(err)
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
				DB:          db,
				GCSClient:   gcsClient,
				Auth0Client: auth0Client,
			}

			// GraphQLエンドポイント（DataLoaderでラップ）
			r.Handle("/query", graph.DataLoaderMiddleware(resolver, graphQlServer(resolver)))

			// GraphQLドキュメント
			r.Handle("/", playground.Handler("fs-mng-backend", "/query"))
		})
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", e.ServerPort)
	if err := http.ListenAndServe(":"+e.ServerPort, router); err != nil {
		fmt.Println(err)
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
