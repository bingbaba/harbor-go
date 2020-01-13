package harbor

import (
	"context"
	"testing"
)

func TestProject(t *testing.T) {
	c, err := NewClientFromEnv(nil)
	if err != nil {
		return
	}

	// list project
	_, ps, err := c.ListProjects(context.Background(), &ProjectOption{Name: "library", Public: "1"})
	if err != nil {
		t.Fatal(err)
	}
	if len(ps) != 1 {
		t.Fatalf("expect get one project, but get %d project!", len(ps))
	}
	if ps[0].ProjectID != 1 {
		t.Fatalf("expect project id is 1,but get %d!", ps[0].ProjectID)
	}

	// by name
	p, err := c.GetProjectByName(context.Background(), "library")
	if err != nil {
		t.Fatal(err)
	}
	if p.ProjectID != 1 {
		t.Fatalf("expect project id is 1,but get %d!", p.ProjectID)
	}

	// check exist
	found, err := c.ProjectIsExist(context.Background(), "library")
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Fatalf("expect project was exists,but get not exist!")
	}

	// check exist
	found, err = c.ProjectIsExist(context.Background(), "aaa")
	if err != nil {
		t.Fatal(err)
	}
	if found {
		t.Fatalf("expect project was not exists,but has exist!")
	}
}
