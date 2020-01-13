package models

type RoleType int

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Realname string `json:"realname"`
}

type UserInfo struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}
