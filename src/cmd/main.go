package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/sky0621/fs-mng-backend/src/graph"
	"github.com/sky0621/fs-mng-backend/src/graph/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

const defaultPort = "8080"

func main() {
	/*
	 * setup db
	 */
	var db *sql.DB
	{
		var err error
		// MEMO: ひとまずローカルのコンテナ相手の接続前提なので、べたに書いておく。
		db, err = sql.Open("postgres", "dbname=localdb user=postgres password=localpass sslmode=disable port=19999")
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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	bucket := os.Getenv("BUCKET")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB:     db,
		Bucket: bucket,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
