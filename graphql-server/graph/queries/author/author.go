package author

const AuthorQuery string = `
	{
		author(where: { name: "Miguel Villarreal"}, stage: PUBLISHED) {
			name
			intro
			bio
			slug
			stage
			id
			createdAt
			updatedAt
			publishedAt
		}
	}
`
