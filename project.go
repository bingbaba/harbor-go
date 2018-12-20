package harbor

import (
	"context"
	"fmt"
	"github.com/bingbaba/harbor-go/models"
	"net/http"
	"net/url"
)

const (
	PATH_FMT_PROJECT_LIST = "/api/projects"
	PATH_FMT_PROJECT_GET  = "/api/projects/%d"
)

type ProjectOption struct {
	Name   string
	Public string
	Owner  string
	PageOption
}

func (opt *ProjectOption) Urls() url.Values {
	v := opt.PageOption.Urls()

	switch opt.Public {
	case "1", "true", "True":
		v.Set("public", "1")
	case "0", "false", "False":
		v.Set("public", "0")
	default:

	}

	if opt.Name != "" {
		v.Set("name", opt.Name)
	}
	if opt.Owner != "" {
		v.Set("owner", opt.Owner)
	}

	return v
}

func (c *Client) ListProjects(ctx context.Context, opt *ProjectOption) ([]*models.Project, error) {
	var ret []*models.Project

	path := PATH_FMT_PROJECT_LIST
	if opt != nil {
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

func (c *Client) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	ret := new(models.Project)

	path := fmt.Sprintf(PATH_FMT_PROJECT_GET, id)
	req, err := http.NewRequest(http.MethodGet, c.host+path, nil)
	if err != nil {
		return ret, err
	}

	err = c.doJson(ctx, req, ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *Client) DeleteProject(ctx context.Context, id int64) (deleted bool, err error) {

	path := fmt.Sprintf(PATH_FMT_PROJECT_GET, id)
	req, err := http.NewRequest(http.MethodDelete, c.host+path, nil)
	if err != nil {
		return false, err
	}

	code, body, err := c.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()

	switch code {
	case 200:
		return true, nil
	case 400, 404:
		return false, NotFoundError
	case 403:
		return true, UserNotLoginError
	case 412:
		return false, NotAllowError
	case 500:
		return false, CallApiError
	default:
		return true, ServerInternalError
	}
}

func (c *Client) ProjectIsExist(ctx context.Context, name string) (bool, error) {
	//projects, err := c.ListProjects(ctx, &ProjectOption{Name: name})
	//if err != nil {
	//	return false, err
	//}
	//if len(projects) > 0 {
	//	return true, nil
	//} else {
	//	return false, nil
	//}

	path := fmt.Sprintf("%s?project_name=%s", PATH_FMT_PROJECT_LIST, name)
	//log.Print(path)
	req, err := http.NewRequest(http.MethodHead, c.host+path, nil)
	if err != nil {
		return false, err
	}

	code, body, err := c.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()

	switch code {
	case 200:
		return true, nil
	case 401:
		return true, UserNotLoginError
	case 404:
		return false, nil
	case 500:
		return true, CallApiError
	default:
		return true, ServerInternalError
	}
}

func (c *Client) GetProjectByName(ctx context.Context, name string) (*models.Project, error) {
	projects, err := c.ListProjects(ctx, &ProjectOption{Name: name})
	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, NotFoundError
	}

	return projects[0], nil
}
