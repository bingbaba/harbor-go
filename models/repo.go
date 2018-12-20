package models

import "time"

type Repo struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	ProjectID    int64     `json:"project_id"`
	Description  string    `json:"description"`
	PullCount    int64     `json:"pull_count"`
	StarCount    int64     `json:"star_count"`
	TagsCount    int64     `json:"tags_count"`
	Labels       []*Label  `json:"labels"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}

type TagDetail struct {
	Digest        string     `json:"digest"`
	Name          string     `json:"name"`
	Size          int64      `json:"size"`
	Architecture  string     `json:"architecture"`
	OS            string     `json:"os"`
	DockerVersion string     `json:"docker_version"`
	Author        string     `json:"author"`
	Created       time.Time  `json:"created"`
	Config        *TagConfig `json:"config"`
}

type TagConfig struct {
	Labels map[string]string `json:"labels"`
}
