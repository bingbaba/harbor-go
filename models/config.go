package models

const (
	// Quota setting items for project
	CountPerProject   = "count_per_project"
	StoragePerProject = "storage_per_project"
)

type ConfigureValue struct {
	Value    interface{} `json:"value"`
	Editable bool        `json:"editable"`
}
