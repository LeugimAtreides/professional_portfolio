package project

import "fmt"

const projectQuery string = `
	{
		project(where:{id:"%s"}, stage: PUBLISHED) {
			name
			slug
			description
			demo
			image {
				... on Asset {
						width
						height
						url
				}
			}
			createdAt
			updatedAt
			publishedAt
			tags
			stage
	    }
	}
`

func ProjectString(id string) string {
	response := fmt.Sprintf(projectQuery, id)
	return response
}
