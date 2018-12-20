package main

import (
	"context"
	"fmt"
	"github.com/bingbaba/harbor-go"
)

const (
	username = "admin"
	password = "admin"
	host     = "http://myharbor.company.com"
)

func main() {
	// create harbor client
	c, err := harbor.NewClient(nil, host, username, password, true)
	if err != nil {
		panic(err)
	}

	// list project
	ps, err := c.ListProjects(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	// dump projects
	for _, p := range ps {
		fmt.Printf("%+v\n", p)
	}
}
