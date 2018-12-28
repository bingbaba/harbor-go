package harbor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	http_client *http.Client
	host        string
	hostname    string
	username    string
	password    string
}

func NewClientFromEnv(http_client *http.Client) (*Client, error) {
	host := os.Getenv("HARBOR_HOST")
	if host == "" {
		return nil, fmt.Errorf("read \"DRONE_HOST\" from env failed")
	}

	username := os.Getenv("HARBOR_USER")
	if username == "" {
		return nil, fmt.Errorf("read \"HARBOR_USER\" from env failed")
	}

	password := os.Getenv("HARBOR_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("read \"HARBOR_PASSWORD\" from env failed")
	}

	var sniffer bool
	sniffer_str := os.Getenv("HARBOR_SNIFFER")
	switch sniffer_str {
	case "0", "false", "False", "FALSE":
		sniffer = false
	default:
		sniffer = true
	}

	return NewClient(http_client, host, username, password, sniffer)
}

func NewClient(http_client *http.Client, host, username, password string, sniffer bool) (*Client, error) {
	if http_client == nil {
		http_client = http.DefaultClient
	}

	array := strings.SplitN(host, "//", 2)
	if len(array) != 2 {
		return nil, fmt.Errorf("parse host failed")
	}
	hostname := array[1]

	client := &Client{
		http_client: http_client,
		host:        host,
		hostname:    hostname,
		username:    username,
		password:    password,
	}

	// sniffer
	if sniffer {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := client.ListConfigures(ctx)
		if err != nil {
			return client, err
		}
	}

	return client, nil
}

func (c *Client) Image(project, repo, tag string) string {
	return fmt.Sprintf("%s/%s/%s:%s", c.hostname, project, repo, tag)
}
func (c *Client) do(ctx context.Context, req *http.Request) (int, io.ReadCloser, error) {
	resp, err := c._do(ctx, req)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, resp.Body, nil
}

func (c *Client) _do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req = req.WithContext(ctx)

	return c.http_client.Do(req)
}

func (c *Client) doJson(ctx context.Context, req *http.Request, v interface{}) error {
	code, body, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer body.Close()

	if code >= 400 {
		body_bytes, _ := ioutil.ReadAll(body)
		return fmt.Errorf("http request failed(%d): %s", code, string(body_bytes))
	}

	err = json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doJsonWithTotal(ctx context.Context, req *http.Request, v interface{}) (total int, code int, err error) {
	resp, err := c._do(ctx, req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if code >= 400 {
		body_bytes, _ := ioutil.ReadAll(resp.Body)
		return 0, resp.StatusCode, fmt.Errorf("http request failed(%d): %s", code, string(body_bytes))
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return 0, resp.StatusCode, err
	}

	total, _ = strconv.Atoi(resp.Header.Get("X-Total-Count"))

	return total, resp.StatusCode, nil
}
