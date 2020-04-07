package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/sky0621/fs-mng-backend/src/graph"
	"github.com/sky0621/fs-mng-backend/src/graph/generated"
)

func main() {
	cfg := InitConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%s", cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSSLMode, cfg.DBPort))
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(cfg.ServerPort, nil))
}
