package harbor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bingbaba/harbor-go/models"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	PATH_FMT_REPO_LIST     = "/api/repositories"
	PATH_FMT_REPOTAGS_LIST = "/api/repositories/%s/tags"
	PATH_FMT_REPOTAG       = "/api/repositories/%s/tags/%s"
)

type RepoSortType = string

const (
	REPO_SORT_BYNAME_ASC        = "name"
	REPO_SORT_BYNAME_DESC       = "-name"
	REPO_SORT_BYCREATION_ASC    = "creation_time"
	REPO_SORT_BYCREATION_DESC   = "-creation_time"
	REPO_SORT_BYUPDATETIME_ASC  = "+update_time"
	REPO_SORT_BYUPDATETIME_DESC = "-update_time"
)

type RepoOption struct {
	Name    string
	Sort    RepoSortType
	LabelId string
	PageOption
}

func (opt *RepoOption) Urls() url.Values {
	v := opt.PageOption.Urls()

	if opt.Name != "" {
		v.Set("q", opt.Name)
	}

	if opt.Sort != "" {
		v.Set("sort", opt.Sort)
	}

	if opt.LabelId != "" {
		v.Set("label_id", opt.LabelId)
	}

	return v
}

func (c *Client) ListRepos(ctx context.Context, project_id int64, opt *RepoOption) (int, []*models.Repo, error) {
	ret := make([]*models.Repo, 0)

	var values url.Values
	if opt != nil {
		values = opt.Urls()
	} else {
		values = make(map[string][]string)
	}
	values.Set("project_id", fmt.Sprintf("%d", project_id))

	path := PATH_FMT_REPO_LIST + "?" + values.Encode()
	//fmt.Println("path: ",path)
	req, err := http.NewRequest(http.MethodGet, c.host+path, nil)
	if err != nil {
		return 0, ret, err
	}

	total, code, err := c.doJsonWithTotal(ctx, req, &ret)
	switch code {
	case 400, 404:
		return 0, ret, NotFoundError
	case 403:
		return 0, ret, NotAllowError
	case 500:
		return 0, ret, CallApiError
	case 200:
		return total, ret, nil
	default:
		return 0, ret, err
	}
}

func (c *Client) ListReposByProjectName(ctx context.Context, name string) (total int, repos []*models.Repo, err error) {
	p, err := c.GetProjectByName(ctx, name)
	if err != nil {
		return total, repos, err
	}

	return c.ListRepos(ctx, p.ProjectID, nil)
}

func (c *Client) ListRepoTags(ctx context.Context, project_name, repo_name string) ([]*models.TagDetail, error) {
	ret := make([]*models.TagDetail, 0)

	path := fmt.Sprintf(PATH_FMT_REPOTAGS_LIST, repo_name)
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

func (c *Client) GetRepoTag(ctx context.Context, project_name, repo_name, tag_name string) (*models.TagDetail, error) {
	ret := new(models.TagDetail)

	path := fmt.Sprintf(PATH_FMT_REPOTAG, repo_name, tag_name)
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
func (c *Client) GetTagManifest(ctx context.Context, project_name, repo_name, tag_name string) (*models.ManifestInfo, error) {
	ret := new(models.ManifestInfo)

	path := fmt.Sprintf(PATH_FMT_REPOTAG+"/manifest", repo_name, tag_name)
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
func (c *Client) DeleteRepoTag(ctx context.Context, project_name, repo_name, tag_name string) (bool, error) {

	path := fmt.Sprintf(PATH_FMT_REPOTAG, repo_name, tag_name)
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
	case 400, 404:
		return false, NotFoundError
	case 401, 403:
		return false, NotAllowError
	case 200:
		return true, nil
	default:
		return false, ServerInternalError
	}
}

func (c *Client) UpdateRepoDesc(ctx context.Context, project_name, repo_name, descString string) error {
	desc := struct {
		Description string `json:"description"`
	}{}
	desc.Description = descString
	bytesData, err := json.Marshal(&desc)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	reader := bytes.NewReader(bytesData)
	path := fmt.Sprintf(PATH_FMT_REPO_LIST+"/%s", repo_name)
	req, err := http.NewRequest(http.MethodPut, c.host+path, reader)
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
func (c *Client) DeleteRepo(ctx context.Context, project_name, repo_name string) (bool, error) {

	path := fmt.Sprintf(PATH_FMT_REPO_LIST+"/%s", repo_name)
	req, err := http.NewRequest(http.MethodDelete, c.host+path, nil)
	if err != nil {
		return false, err
	}
	code, body, err := c.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()
	if code >= 400 {
		body_bytes, _ := ioutil.ReadAll(body)
		fmt.Printf("http request failed(%d): %s\n", code, string(body_bytes))
	}
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
