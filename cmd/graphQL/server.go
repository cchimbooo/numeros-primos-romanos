package main

import (
	"log"
	"net/http"
	"os"
	"prova-programacao/cmd/graphQL/graph"
	"prova-programacao/pkg/core/menorPrimoRomano"
	"prova-programacao/pkg/helpers/primos"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

// main  função que instancia o GraphQL.
// Pessoalmente eu não tenho experiência nenhuma em produção.
// Tentei dar uma olhada e gerar os códigos usando gqlgen.
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	sievePrimos := primos.Sieve3999()
	romanoInterface := menorPrimoRomano.CriarProcessarStringInterface(sievePrimos)
	if romanoInterface == nil {
		log.Fatalln("interface vazia")
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Romano: romanoInterface}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
