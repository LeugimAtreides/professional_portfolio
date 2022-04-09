package createproject

import (
	"fmt"

	"github.com/LeugimAtreides/professional_portfolio/graph/model"
)

// this mutation query requires addition of image with content and collection fields to work
const createProjectQuery string = `
  mutation {
	  createProject(data: { name: "%[1]s", description: "%[2]s" }) {
		  name
		  slug
		  description
	  }
  }
`

func ProjectMutation(input model.NewProject) string {
	response := fmt.Sprintf(createProjectQuery, input.Name, input.Description)
	return response
}
