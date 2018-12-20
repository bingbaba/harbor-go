package harbor

import (
	"context"
	"net/http"
)

const (
	PATH_FMT_CONFIGURE_LIST = "/api/configurations"
)

func (c *Client) ListConfigures(ctx context.Context) (map[string]interface{}, error) {
	configs := make(map[string]interface{}, 0)

	req, err := http.NewRequest(http.MethodGet, c.host+PATH_FMT_CONFIGURE_LIST, nil)
	if err != nil {
		return configs, err
	}

	err = c.doJson(ctx, req, &configs)
	if err != nil {
		return configs, err
	}

	return configs, nil
}
