package models

type ConfigureValue struct {
	Value    interface{} `json:"value"`
	Editable bool        `json:"editable"`
}
