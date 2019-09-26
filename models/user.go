package models

type Users struct {
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Realname     string `json:"realname"`
	Comment      string `json:"comment"`
	Deleted      bool   `json:"deleted"`
	RoleName     string `json:"role_name"`
	RoleId       int    `json:"role_id"`
	HasAdminRole bool   `json:"has_admin_role"`
	ResetUuid    string `json:"reset_uuid"`
	Salt         string `json:"Salt"`
	CreationTime string `json:"creation_time"`
	UpdateTime   string `json:"update_time"`
}

type InitUser struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Realname     string `json:"realname"`
	Comment      string `json:"comment"`
}
