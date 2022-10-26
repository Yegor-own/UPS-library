package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Yegor-own/ghqllibrary/graph"
	"github.com/Yegor-own/ghqllibrary/graph/generated"
)

const defaultPort = "8080"

//var (
//	DBLib *gorm.DB
//)
//
//func init() {
//	DBLib = database.ConnectDB("host=localhost user=postgres password=root dbname=library port=5432 sslmode=disable")
//	database.MigrateDB(DBLib)
//}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
