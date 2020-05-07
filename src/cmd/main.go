package main

import (
	"context"
	"database/sql"
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
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	e := loadEnv()

	/*
	 * setup db client
	 */
	var db *sql.DB
	{
		var err error
		// MEMO: ひとまずローカルのコンテナ相手の接続前提なので、べたに書いておく。
		dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%s", e.DBName, e.DBUser, e.DBPassword, e.DBSSLMode, e.DBPort)
		db, err = sql.Open(e.DBDriverName, dataSourceName)
		if err != nil {
			panic(err)
		}
		defer func() {
			if db != nil {
				if err := db.Close(); err != nil {
					panic(err)
				}
			}
		}()

		boil.DebugMode = true

		var loc *time.Location
		loc, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(loc)
	}

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

		c := cors.AllowAll()
		router.Use(c.Handler)

		// 認証認可チェック用（今はJWTのチェックのみ実装）
		a := auth.New(e.Auth0Domain, e.Auth0Audience, e.AuthDebug, e.AuthCredentialsOptional)
		router.Use(a.CheckJWTHandlerFunc())
		router.Use(a.HoldPermissionsHandler)

		router.Handle("/", playground.Handler("fs-mng-backend", "/query"))

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{
				DB:          db,
				GCSClient:   gcsClient,
				Auth0Client: auth0Client,
			},
		}))
		var mb int64 = 1 << 20
		srv.AddTransport(transport.MultipartForm{
			MaxMemory:     128 * mb,
			MaxUploadSize: 100 * mb,
		})

		router.Handle("/query", srv)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", e.ServerPort)
	if err := http.ListenAndServe(":"+e.ServerPort, router); err != nil {
		fmt.Println(err)
	}
}
