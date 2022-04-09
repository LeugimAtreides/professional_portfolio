package authors

const AuthorsQuery string = `
  {
	authorsConnection(stage: PUBLISHED) {
	  pageInfo {
		  startCursor
		  endCursor
		  pageSize
	  }
	  edges {
		node {
		  id
		  name
		  slug
		}
	  }
	}
  }`
