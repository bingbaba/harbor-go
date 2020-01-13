package harbor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bingbaba/harbor-go/models"
	"io/ioutil"
	"net/http"
)

const (
	PATH_FMT_CREATE_USER = "/api/users"
	PATH_FMT_LIST_USER   = "/api/users/search?username=%s"
)

func (c *Client) CreateUser(ctx context.Context, create *models.UserCreate) error {

	path := PATH_FMT_CREATE_USER

	bytesData, err := json.Marshal(create)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest(http.MethodPost, c.host+path, reader)
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

func (c *Client) ExistUser(ctx context.Context, username string) (bool, error) {
	ret := make([]*models.UserInfo, 0)

	path := fmt.Sprintf(PATH_FMT_LIST_USER, username)
	req, err := http.NewRequest(http.MethodGet, c.host+path, nil)
	if err != nil {
		return false, err
	}

	err = c.doJson(ctx, req, &ret)
	if err != nil {
		return false, err
	}
	if len(ret) == 0 {
		return false, nil
	}
	if ret[0].Username == username {
		return true, nil
	}
	return false, nil
}
