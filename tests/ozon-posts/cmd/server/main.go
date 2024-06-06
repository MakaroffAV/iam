package main

import (
	"log"
	"net/http"
	"ozon-posts/graph"
	"ozon-posts/internal/repo/memory"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					PostStorage:    memory.NewPostRepository(),
					CommentStorage: memory.NewCommentRepository(),
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL  playground", "/query"))
	http.Handle("/query", srv)

	log.Print("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
