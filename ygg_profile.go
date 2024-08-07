package go_mc_profile_api

import (
	"io"
	"net/http"
)

type YggProfile struct {
	McProfile
	yggApiBaseUrl string
}

func NewYggProfile(yggApiBaseUrl string) YggProfile {
	return YggProfile{yggApiBaseUrl: yggApiBaseUrl}
}

func (p *YggProfile) FetchByName(name string) error {
	get, err := http.Get(p.yggApiBaseUrl + "api/users/profiles/minecraft/" + name)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(get.Body)

	return handleByNameResp(get, p)
}

func (p *YggProfile) FetchById(id string) error {

	get, err := http.Get(p.yggApiBaseUrl + "sessionserver/session/minecraft/profile/" + id)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(get.Body)

	return handleByNameResp(get, p)
}
