package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Keijun-KUMAGAI/graphql-server/gqlgen"
	"github.com/Keijun-KUMAGAI/graphql-server/prisma-client"
	"github.com/Keijun-KUMAGAI/graphql-server/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := resolvers.Resolver{
		Prisma: client,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
