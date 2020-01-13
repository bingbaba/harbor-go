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
	PATH_FMT_MEMBER_LIST = "/api/projects/%d/members?entityname=%s"
	PATH_FMT_MEMBER_ADD  = "/api/projects/%d/members"
	PATH_FMT_MEMBER      = "/api/projects/%d/members/%d"
)

func (c *Client) GetMemberByName(ctx context.Context, project_id int64, name string) (*models.Member, error) {
	members, err := c.ListMembers(ctx, project_id, name)
	if err != nil {
		return nil, err
	}
	if len(members) == 0 {
		return nil, NotFoundError
	}
	if members[0].EntityName == name {
		return members[0], nil
	}
	return nil, NotFoundError
}
func (c *Client) ListMembers(ctx context.Context, project_id int64, entityname string) ([]*models.Member, error) {
	var ret []*models.Member

	path := fmt.Sprintf(PATH_FMT_MEMBER_LIST, project_id, entityname)
	//log.Print(path)
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

func (c *Client) AddMember(ctx context.Context, project_id int64, create *models.AddMemberReq) error {

	path := fmt.Sprintf(PATH_FMT_MEMBER_ADD, project_id)
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

func (c *Client) DeleteMember(ctx context.Context, project_id, member_id int64) error {

	path := fmt.Sprintf(PATH_FMT_MEMBER, project_id, member_id)
	//log.Print(path)
	req, err := http.NewRequest(http.MethodDelete, c.host+path, nil)
	if err != nil {
		return err
	}

	code, body, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer body.Close()

	switch code {
	case 400, 404:
		return NotFoundError
	case 401, 403:
		return NotAllowError
	case 200:
		return nil
	default:
		return ServerInternalError
	}
}
