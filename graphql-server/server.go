package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LeugimAtreides/professional_portfolio/graph"
	"github.com/LeugimAtreides/professional_portfolio/graph/generated"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

const defaultPort = "8080"

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var client *graphql.Client

var CONTENT_API = goDotEnvVariable("CONTENT_API")

var CMS_TOKEN = goDotEnvVariable("CMS_TOKEN")

func setUpClient() {
	client = graphql.NewClient(CONTENT_API)
	fmt.Print(client)
}

func main() {
	setUpClient()

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
