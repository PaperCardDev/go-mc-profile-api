package go_mc_profile_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type McProfile struct {
	PlayerId string `json:"id"`
	Name     string `json:"name"`
}

type errorResponse struct {
	Path         string `json:"path"`
	ErrorMessage string `json:"errorMessage"`
}

func handleByNameResp(resp *http.Response, v any) error {

	if resp.StatusCode != http.StatusOK {

		errResponse := errorResponse{}

		err := json.NewDecoder(resp.Body).Decode(&errResponse)

		if err != nil {
			return errors.New(resp.Status)
		}

		return errors.New(errResponse.ErrorMessage)
	}

	err := json.NewDecoder(resp.Body).Decode(v)

	if err != nil {
		return err
	}

	return nil
}

func (p *McProfile) FetchByName(name string) error {
	get, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(get.Body)

	return handleByNameResp(get, p)
}

func (p *McProfile) FetchById(id string) error {

	get, err := http.Get("https://sessionserver.mojang.com/session/minecraft/profile/" + id)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(get.Body)

	return handleByNameResp(get, p)
}
