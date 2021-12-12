package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/stevegood/kalmar/pkg/server/graph"
	"github.com/stevegood/kalmar/pkg/server/graph/generated"
	"github.com/stevegood/kalmar/web"
)

func Exec(port string) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", http.FileServer(http.FS(web.GetStaticFiles())))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("now running on http://localhost:%s/", port)
	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}
