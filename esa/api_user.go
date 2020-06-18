package esa

import (
	"net/http"
)

type User struct {
	Id         int
	Name       string
	ScreenName string
	CreatedAt  string
	UpdatedAt  string
	Icon       string
	Email      string
}

func (api *Api) User() (*User, *http.Response, error) {
	resp, err := api.Get("/v1/user", nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}

	var user User
	if err := api.unmarshal(resp, &user); err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}
