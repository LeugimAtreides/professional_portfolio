package projects

const ProjectsQuery string = `
	{
		projectsConnection(stage: PUBLISHED) {
		edges {
			node {
			id
			name
			slug
			description
			image {
				... on Asset {
				url
				height
				width
				}
			}
			}
		}
		}
	}
`
