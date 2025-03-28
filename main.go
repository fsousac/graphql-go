package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fsousac/graphql-go/graph"
)

func main() {
  // Inicializa o Resolver com a lista vazia
  resolver := &graph.Resolver{
    todos: make([]*graph.Todo, 0), // Use o tipo correto (graph.Todo)
  }

  srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
    Resolvers: resolver,
  }))

  http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
  http.Handle("/query", srv)

  log.Println("Servidor rodando em http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}