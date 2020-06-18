package esa

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Invitation struct {
	Email     string
	Code      string
	ExpiresAt string
	Url       string
}

type Invitations struct {
	Invitations []Invitation
}

type InvitationsWithPagination struct {
	Invitations []Invitation
}

func (api *Api) PendingInvitations() (*InvitationsWithPagination, *http.Response, error) {
	resp, err := api.Get("/v1/teams/"+api.Team+"/invitations", nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}

	var invitations InvitationsWithPagination
	if err := api.unmarshal(resp, &invitations); err != nil {
		return nil, resp, err
	}

	return &invitations, resp, nil
}

func (api *Api) SendInvitation(emails []string) (*Invitations, *http.Response, error) {
	params := map[string]map[string][]string{"member": {"emails": emails}}
	json, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}
	log.Printf("[DEBUG] json: %s", json)

	path := "/v1/teams/" + api.Team + "/invitations"
	resp, err := api.Post(path, bytes.NewBuffer(json))
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}

	var invitations Invitations
	if err := api.unmarshal(resp, &invitations); err != nil {
		return nil, resp, err
	}

	return &invitations, resp, nil
}

func (api *Api) CancelInvitation(code string) (*http.Response, error) {
	resp, err := api.Delete("/v1/teams/"+api.Team+"/invitations/"+code, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
