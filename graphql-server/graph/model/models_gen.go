// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Intro       string  `json:"intro"`
	Bio         *string `json:"bio"`
	Slug        *string `json:"slug"`
	Picture     *Image  `json:"picture"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
	PublishedAt *string `json:"publishedAt"`
}

type Image struct {
	ID          string   `json:"id"`
	Height      *float64 `json:"height"`
	Width       *float64 `json:"width"`
	URL         string   `json:"url"`
	MimeType    *string  `json:"mimeType"`
	FileName    string   `json:"fileName"`
	CreatedAt   *string  `json:"createdAt"`
	UpdatedAt   *string  `json:"updatedAt"`
	PublishedAt *string  `json:"publishedAt"`
}

type NewProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Slug        *string  `json:"slug"`
	Description *string  `json:"description"`
	Demo        *string  `json:"demo"`
	Images      []*Image `json:"images"`
	CreatedAt   *string  `json:"createdAt"`
	UpdatedAt   *string  `json:"updatedAt"`
	PublishedAt *string  `json:"publishedAt"`
	Tags        []string `json:"tags"`
}

type ProjectWhereUniqueInput struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
	Slug *string `json:"slug"`
}
