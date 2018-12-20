package models

type SearchResult struct {
	Projects    []*Project          `json:"project"`
	Repositorys []*SearchRepository `json:"repository"`
}

type SearchRepository struct {

	// The ID of the project that the repository belongs to
	ProjectId int32 `json:"project_id,omitempty"`

	// The name of the project that the repository belongs to
	ProjectName string `json:"project_name,omitempty"`

	// The flag to indicate the publicity of the project that the repository belongs to
	ProjectPublic interface{} `json:"project_public,omitempty"`

	// The name of the repository
	RepositoryName string `json:"repository_name,omitempty"`

	// The count of pull
	PullCount int64 `json:"pull_count,omitempty"`

	// The count of tag
	TagsCount int `json:"tags_count,omitempty"`
}

func (sr *SearchRepository) IsProjectPublic() bool {
	return getBool(sr.ProjectPublic)
}
