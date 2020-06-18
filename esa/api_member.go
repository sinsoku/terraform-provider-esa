package esa

import (
	"net/http"
)

type Member struct {
	Name           string `json:"name"`
	ScreenName     string `json:"screen_name"`
	Icon           string `json:"icon"`
	Email          string `json:"email"`
	PostsCount     int    `json:"posts_count"`
	JoinedAt       string `json:"joined_at"`
	LastAccessedAt string `json:"last_accessed_at"`
}

type Members struct {
	Members []Member
}

type MembersWithPagination struct {
	Members []Member
}

func (api *Api) Members() (*MembersWithPagination, *http.Response, error) {
	resp, err := api.Get("/v1/teams/"+api.Team+"/members", nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}

	var members MembersWithPagination
	if err := api.unmarshal(resp, &members); err != nil {
		return nil, resp, err
	}

	return &members, resp, nil
}

func (api *Api) DeleteMember(screen_name string) (*http.Response, error) {
	resp, err := api.Delete("/v1/teams/"+api.Team+"/members/"+screen_name, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
