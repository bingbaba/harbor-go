package harbor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (c *Client) PutConfigures(ctx context.Context, maps map[string]interface{}) error {
	bytesData, err := json.Marshal(maps)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	reader := bytes.NewReader(bytesData)

	req, err := http.NewRequest(http.MethodPut, c.host+PATH_FMT_CONFIGURE_LIST, reader)
	if err != nil {
		return err
	}
	code, body, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer body.Close()

	if code >= 400 {
		body_bytes, _ := ioutil.ReadAll(body)
		return fmt.Errorf("http request failed(%d): %s", code, string(body_bytes))
	}

	return nil
}
