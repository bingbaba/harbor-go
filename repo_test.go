package harbor

import (
	"context"
	"log"
	"testing"
)

func TestRepo(t *testing.T) {
	c, err := NewClientFromEnv(nil)
	if err != nil {
		return
	}

	// list project
	total, repos, err := c.ListRepos(context.Background(), 143, &RepoOption{Name: "cmdb"})
	if err != nil {
		t.Fatal(err)
	}
	if total == 0 {
		t.Fatalf("can't get any repo")
	}

	for i, repo := range repos {
		log.Printf("[%d] %+v", i, repo)
	}

	// list repo tag
	project_name := "console_gitlab"
	repo_name := "cmdb"
	tags, err := c.ListRepoTags(context.Background(), project_name, repo_name)
	if err != nil {
		t.Fatal(err)
	}

	if len(tags) < 13 {
		t.Fatalf("expect 13 tags at least, but get %d tags", len(tags))
	}

	for i, tag := range tags {
		log.Printf("[%d] %s/%s:%s %+v", i, project_name, repo_name, tag.Name, tag)
	}
}
