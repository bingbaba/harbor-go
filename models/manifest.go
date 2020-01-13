package models

type ManifestInfo struct {
	Manifest Manifest `json:"manifest"`
	Config   string   `json:"config"`
}

type Manifest struct {
	SchemaVersion int            `json:"schemaVersion"`
	MediaType     string         `json:"mediaType"`
	Config        ManifestConfig `json:"config"`
	Layers        []Layer        `json:"layers"`
}
type ManifestConfig struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}
type Layer struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}
