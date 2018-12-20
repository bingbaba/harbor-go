package harbor

import (
	"context"
	"testing"
)

func TestListConfigures(t *testing.T) {
	c, err := NewClientFromEnv(nil)
	if err != nil {
		return
	}

	cfgs, err := c.ListConfigures(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(cfgs) == 0 {
		t.Fatal("can't get any configure")
	}
}
