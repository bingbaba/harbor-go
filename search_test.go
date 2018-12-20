package harbor

import (
	"context"
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	c, err := NewClientFromEnv(nil)
	if err != nil {
		return
	}

	search_ret, err := c.Search(context.Background(), "library")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%d repo", len(search_ret.Repositorys))

	if len(search_ret.Projects) != 1 {
		t.Fatalf("search %s project failed, expect get only one, but get %d project!", "library", len(search_ret.Projects))
	}
	if search_ret.Projects[0].ProjectID != 1 {
		t.Fatalf("expect project id is 1,but get %d!", search_ret.Projects[0].ProjectID)
	}
}
