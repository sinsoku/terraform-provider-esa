package esa

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

type Api struct {
	Team        string
	ApiEndpoint string
	Client      *http.Client
}

func NewApi(token string, team string, api_endpoint string) *Api {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	c := oauth2.NewClient(ctx, ts)

	return &Api{
		Team:        team,
		ApiEndpoint: api_endpoint,
		Client:      c,
	}
}

func (api *Api) Get(path string, body io.Reader) (*http.Response, error) {
	return api.request("GET", path, body)
}

func (api *Api) Post(path string, body io.Reader) (*http.Response, error) {
	return api.request("POST", path, body)
}

func (api *Api) Patch(path string, body io.Reader) (*http.Response, error) {
	return api.request("PATCH", path, body)
}

func (api *Api) Delete(path string, body io.Reader) (*http.Response, error) {
	return api.request("DELETE", path, body)
}

func (api *Api) request(method string, path string, body io.Reader) (*http.Response, error) {
	url := api.ApiEndpoint + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.Client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return resp, errors.New(http.StatusText(resp.StatusCode))
	}

	return resp, nil
}

func (api *Api) unmarshal(resp *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}
