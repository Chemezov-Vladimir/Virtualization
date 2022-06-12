package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type client struct {
	http     *http.Client
	key      string
	lang     string
	pageSize int
}

type Client interface {
	FetchEverything(query, page string) (*Results, error)
	GetPageSize() int
}

var _ Client = &client{}

func (c *client) FetchEverything(query, page string) (*Results, error) {
	endpoint := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=%s", url.QueryEscape(query), c.pageSize, page, c.key, c.lang)
	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Results{}
	return res, json.Unmarshal(body, res)
}

func (c *client) GetPageSize() int{
	return c.pageSize
}

func NewClient(httpClient *http.Client, key, lang string, pageSize int) *client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &client{httpClient, key, lang,pageSize}
}
