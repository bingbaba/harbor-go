package models

const (
	ROLE_PROJECTMANAGER RoleType = 1 //1:项目管理员
	ROLE_DEV            RoleType = 2 // 2：开发人员
	ROLE_VISITOR        RoleType = 3 //3：访客
	ROLE_FIX            RoleType = 4 // 4：维护人员
)

type Member struct {
	ID         int64  `json:"id"`
	ProjectID  int64  `json:"project_id"`
	EntityName string `json:"entity_name"`
	RoleName   string `json:"role_name"`
	RoleID     int    `json:"role_id"`
	EntityID   int    `json:"entity_id"`
	EntityType string `json:"entity_type"`
}

type AddMemberReq struct {
	RoleID     RoleType   `json:"role_id"`
	MemberUser MemberUser `json:"member_user"`
}
type MemberUser struct {
	Username string `json:"username"`
}
