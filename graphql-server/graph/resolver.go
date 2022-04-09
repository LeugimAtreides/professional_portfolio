package graph

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/LeugimAtreides/professional_portfolio/graph/model"
	"github.com/LeugimAtreides/professional_portfolio/graph/queries/author"
	"github.com/LeugimAtreides/professional_portfolio/graph/queries/authors"
	"github.com/LeugimAtreides/professional_portfolio/graph/queries/project"
	"github.com/LeugimAtreides/professional_portfolio/graph/queries/projects"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var client *graphql.Client

var CONTENT_API = goDotEnvVariable("CONTENT_API")

func setUpClient() {
	client = graphql.NewClient(CONTENT_API)
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type AuthorResponse struct {
	Author model.Author `json:"author"`
}

func GetAuthor() string {
	setUpClient()

	query := graphql.NewRequest(author.AuthorQuery)

	ctx := context.Background()

	var responseData AuthorResponse

	if err := client.Run(ctx, query, &responseData); err != nil {
		log.Fatal(err)
	}

	jsonResponse, _ := json.MarshalIndent(responseData, "", "\t")

	return string(jsonResponse)
}

type ProjectResponse struct {
	Project model.Project `json:"project"`
}

func GetProject(id string) string {
	setUpClient()

	queryStringFromId := project.ProjectString(id)

	query := graphql.NewRequest(queryStringFromId)

	ctx := context.Background()

	var responseData ProjectResponse

	if err := client.Run(ctx, query, &responseData); err != nil {
		log.Fatal(err)
	}

	jsonResponse, _ := json.MarshalIndent(responseData, "", "\t")

	return string(jsonResponse)
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	PageSize    int    `json:"pageSize"`
}

type ProjectEdge struct {
	Cursor string        `json:"cursor"`
	Node   model.Project `json:"node"`
}

type ProjectsConnection struct {
	PageInfo PageInfo      `json:"pageInfo"`
	Edges    []ProjectEdge `json:"edges"`
}

type ProjectsResponse struct {
	ProjectsConnection ProjectsConnection `json:"projectsConnection"`
}

func GetProjects() string {
	setUpClient()

	query := graphql.NewRequest(projects.ProjectsQuery)

	ctx := context.Background()

	var responseData ProjectsResponse

	if err := client.Run(ctx, query, &responseData); err != nil {
		log.Fatal(err)
	}

	jsonResponse, _ := json.MarshalIndent(responseData, "", "\t")

	return string(jsonResponse)
}

type AuthorEdge struct {
	Cursor string       `json:"cursor"`
	Node   model.Author `json:"node"`
}

type AuthorsConnection struct {
	PageInfo PageInfo     `json:"pageInfo"`
	Edges    []AuthorEdge `json:"edges"`
}

type AuthorsResponse struct {
	AuthorsConnection AuthorsConnection `json:"authorsConnection"`
}

func GetAuthors() string {
	setUpClient()

	query := graphql.NewRequest(authors.AuthorsQuery)

	ctx := context.Background()

	var responseData AuthorsResponse

	if err := client.Run(ctx, query, &responseData); err != nil {
		log.Fatal(err)
	}

	jsonResponse, _ := json.MarshalIndent(responseData, "", "\t")

	return string(jsonResponse)
}

type Resolver struct {
	authors  []*model.Author
	projects []*model.Project
	author   *model.Author
	project  *model.Project
}
