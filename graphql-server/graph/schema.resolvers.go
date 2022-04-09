package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LeugimAtreides/professional_portfolio/graph/generated"
	"github.com/LeugimAtreides/professional_portfolio/graph/model"
)

// Create Project will be completed as part of a later effort to add content editing on the site itself
func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	project := &model.Project{
		Name:        input.Name,
		Description: &input.Description,
	}

	// CreateProject(input)

	return project, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {

	authorsJson := GetAuthors()

	fmt.Print(authorsJson)

	var authorsRes AuthorsResponse

	json.Unmarshal([]byte(authorsJson), &authorsRes)

	for _, v := range authorsRes.AuthorsConnection.Edges {
		node := v.Node
		r.authors = append(r.authors, &node)
	}

	return r.authors, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	projectsJson := GetProjects()

	fmt.Print(projectsJson)

	var projectsRes ProjectsResponse

	json.Unmarshal([]byte(projectsJson), &projectsRes)

	for _, v := range projectsRes.ProjectsConnection.Edges {
		node := v.Node
		r.projects = append(r.projects, &node)
	}

	return r.projects, nil
}

func (r *queryResolver) Author(ctx context.Context) (*model.Author, error) {
	authorJSON := GetAuthor()

	fmt.Print(authorJSON)

	var authorRes AuthorResponse

	json.Unmarshal([]byte(authorJSON), &authorRes)

	r.author = &model.Author{
		ID:          authorRes.Author.ID,
		Name:        authorRes.Author.Name,
		Intro:       authorRes.Author.Intro,
		Bio:         authorRes.Author.Bio,
		Slug:        authorRes.Author.Slug,
		Picture:     authorRes.Author.Picture,
		CreatedAt:   authorRes.Author.CreatedAt,
		UpdatedAt:   authorRes.Author.UpdatedAt,
		PublishedAt: authorRes.Author.PublishedAt,
	}
	return r.author, nil
}

func (r *queryResolver) Project(ctx context.Context, where *model.ProjectWhereUniqueInput) (*model.Project, error) {
	projectJSON := GetProject(*where.ID)

	fmt.Print(projectJSON)

	var projectRes ProjectResponse

	json.Unmarshal([]byte(projectJSON), &projectRes)

	r.project = &model.Project{
		ID:          projectRes.Project.ID,
		Name:        projectRes.Project.Name,
		Slug:        projectRes.Project.Slug,
		Description: projectRes.Project.Description,
		Images:      projectRes.Project.Images,
		CreatedAt:   projectRes.Project.CreatedAt,
		UpdatedAt:   projectRes.Project.UpdatedAt,
		PublishedAt: projectRes.Project.PublishedAt,
		Tags:        projectRes.Project.Tags,
	}
	return r.project, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
