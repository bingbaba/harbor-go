package models

import (
	"time"
)

type Project struct {
	ProjectID    int64             `orm:"pk;auto;column(project_id)" json:"project_id"`
	OwnerID      int               `orm:"column(owner_id)" json:"owner_id"`
	Name         string            `orm:"column(name)" json:"name"`
	CreationTime time.Time         `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time         `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      interface{}       `orm:"column(deleted)" json:"deleted"`
	OwnerName    string            `orm:"-" json:"owner_name"`
	Togglable    bool              `orm:"-" json:"togglable"`
	Role         int               `orm:"-" json:"current_user_role_id"`
	RepoCount    int64             `orm:"-" json:"repo_count"`
	ChartCount   uint64            `orm:"-" json:"chart_count"`
	Metadata     map[string]string `orm:"-" json:"metadata"`
}

func (p *Project) IsDeleted() bool {
	return getBool(p.Deleted)
}

type CreateProject struct {
	ProjectName  string         `json:"project_name"`
	Metadata     CreateMetadata `json:"metadata"`
	CountLimit   int            `json:"count_limit"`
	StorageLimit int            `json:"storage_limit"`
}

type CreateMetadata struct {
	Public string `json:"public"`
}
