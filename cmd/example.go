package main

import (
	"context"
	"fmt"
	harbor "harbor-go"
)

const (
	username = "xxxxx"
	password = "xxxxx"
	host     = "http://xxxx.xxxx.xxx:xxxx"
)

func main() {
	// create harbor client
	c, err := harbor.NewClient(nil, host, username, password, true)
	if err != nil {
		panic(err)
	}
	// list project
	//ps, err := c.ListProjects(context.Background(), nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// dump projects
	//for _, p := range ps {
	//	fmt.Printf("%+v\n", p)
	//}

	// todo: 查询 1. group is true. 2. user is new.
	// usergroup 使用上需要使用到LDAP
	//
	//// 用户自查
	//userSelect := &harbor.UserOption{
	//	Username: "harbort5",
	//	//Email:    "xhigeneral3@gmail.com",
	//}
	//psa, err := c.ListUser(context.Background(), userSelect)
	//if err != nil {
	//	panic(err)
	//}
	//// dump projects
	//for _, p := range psa {
	//	fmt.Printf("%+v\n", p)
	//}

	//// todo: create user
	//userinfo := &harbor.UserInfo{
	//	Username: "harbort5",
	//	Password: "abc123456AA",
	//	Realname: "PaasUser1",
	//	Email:    "xhigeneral3@gmail.com",
	//}
	//_, err = c.CreateUser(context.Background(), userinfo)
	//if err != nil {
	//	fmt.Println(err)
	//
	//} else {
	//	fmt.Println("success")
	//}

	// todo: create project (aipaas-[username])

	//projectinfo := &harbor.ProjectInit{
	//	ProjectName: "aipaas-test2",
	//	Public:      false,
	//}
	//
	//err = c.CreateProject(context.Background(), projectinfo)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("success")
	//}

	// todo: add new user to new project and public project

	// {role_id: 2, member_user: {username: "harbort5"}}
	//
	projectMember := &harbor.ProjectMember{
		RoleId:   2,
		Username: "xhtest",
		Project:  "aipaas-test2",
	}
	_, err = c.AddMemberToProject(context.Background(), projectMember)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("success")
	}
}
