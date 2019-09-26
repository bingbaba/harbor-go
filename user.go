package harbor

import (
	"context"
	"encoding/json"
	"fmt"
	"harbor-go/models"
	"net/http"
	"net/url"
	"strings"
)

const (
	PATH_FMT_USER_CREATE = "/api/users"
	PATH_FMT_USER_LIST   = "/api/users"
	PATH_FMT_USER_SEARCH = "/api/users/search"
)

type UserInfo struct {
	Username string
	Password string
	Realname string
	Email    string
}

type UserOption struct {
	Username string
	Email    string
	PageOption
}

func (opt *UserOption) Urls() url.Values {
	v := opt.PageOption.Urls()

	if opt.Username != "" {
		v.Set("username", opt.Username)
	}
	if opt.Email != "" {
		v.Set("email", opt.Email)
	}
	return v
}

func (c *Client) ListUser(ctx context.Context, opt *UserOption) ([]*models.Users, error) {
	var ret []*models.Users
	path := PATH_FMT_USER_LIST
	if opt != nil {
		// 记得这儿是一次重写
		path += "?" + opt.Urls().Encode()
	}
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

func (c *Client) UserIsExist(ctx context.Context, username string) (bool, error) {
	opt := &UserOption{
		Username: username,
	}
	list_user, err := c.ListUser(ctx, opt)
	if err != nil {
		return false, err
	}
	if len(list_user) < 1 {
		return false, fmt.Errorf("username=%s is not find", username)
	}
	return true, nil
}

func (c *Client) GetUserByName(ctx context.Context, username string) (*models.Users, error) {
	users, err := c.ListUser(ctx, &UserOption{Username: username})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, NotFoundError
	}
	return users[0], nil
}

func (c *Client) CreateUser(ctx context.Context, opt *UserInfo) (bool, error) {

	path := PATH_FMT_USER_CREATE
	if opt == nil {
		return false, ERROR_THE_GINSENG
	}

	ret := &models.InitUser{
		Username: opt.Username,
		Password: opt.Password,
		Realname: opt.Realname,
		Email:    opt.Email,
	}

	// struct to string
	out, err := json.Marshal(ret)
	if err != nil {
		return false, err
	}

	str_out := string(out)
	req_body := strings.NewReader(str_out)
	//log.Print(path)
	req, err := http.NewRequest(http.MethodPost, c.host+path, req_body)
	if err != nil {
		return false, err
	}
	code, body, err := c.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()

	// loacl use test
	//body_bite, err := ioutil.ReadAll(body)
	//
	//if err != nil {
	//	return false, errors.New("aaa")
	//}

	switch code {
	case 200, 201:
		return true, nil
	case 400:
		return false, ERROR_THE_FORMAT
	case 403:
		return false, ERROR_THE_PERMISSIONS
	case 415:
		return false, ERROR_THE_TYPE
	case 500:
		return false, ERROR_THE_SERVER
	default:
		return false, ERROR_THE_PKG
	}
}
