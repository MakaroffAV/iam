package main

import (
	"log"
	"net/http"
	"posts/internal/graphql/resolvers"
	"posts/internal/repos/memory"
	"posts/internal/service"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	repoPost := memory.NewPostRepository()
	repoComment := memory.NewCommentRepository()

	servicePost := service.NewPostService(*repoPost)
	serviceComment := service.NewCommentService(*repoComment)

	srv := handler.NewDefaultServer(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: &graphql.Resolver{
					PostResolver:    resolvers.NewPostResolver(servicePost, serviceComment),
					CommentResolver: resolvers.NewCommentResolver(serviceComment),
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
