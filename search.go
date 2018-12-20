package harbor

import (
	"context"
	"github.com/bingbaba/harbor-go/models"
	"net/http"
	"net/url"
)

const (
	PATH_FMT_SEARCH = "/api/search"
)

func (c *Client) Search(ctx context.Context, query string) (*models.SearchResult, error) {
	ret := new(models.SearchResult)

	v := url.Values{"q": []string{query}}
	path := PATH_FMT_SEARCH + "?" + v.Encode()
	req, err := http.NewRequest(http.MethodGet, c.host+path, nil)
	if err != nil {
		return ret, err
	}

	err = c.doJson(ctx, req, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
