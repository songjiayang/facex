package facex

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/songjiayang/qclient"
)

type Facex struct {
	*Config

	*qclient.Client
	timeout time.Duration
}

func NewFacex(cfg *Config) *Facex {
	return &Facex{
		Config: cfg,

		Client:  qclient.New(qclient.NewAccessKey(cfg.AccessKey, cfg.SecretKey)),
		timeout: time.Duration(cfg.Timeout) * time.Second,
	}
}

const (
	groupNewAPI    = "/v1/face/group/%s/new"
	groupRemoveAPI = "/v1/face/group/%s/remove"
	groupAddAPI    = "/v1/face/group/%s/add"
	groupDeleteAPI = "/v1/face/group/%s/delete"
	groupSearchAPI = "/v1/face/group/%s/search"

	mimeType = "application/json"
)

func (facex *Facex) API(api string) string {
	return fmt.Sprintf(strings.TrimSuffix(facex.Endpoint, "/")+api, facex.GroupId)
}

func (facex *Facex) NewGroup(input FacexInput) (err error) {
	_, err = facex.Send(http.MethodPost, facex.API(groupNewAPI), mimeType, facex.timeout, toPayload(input))
	return
}

func (facex *Facex) RemoveGroup() (err error) {
	_, err = facex.Send(http.MethodPost, facex.API(groupRemoveAPI), mimeType, facex.timeout, nil)
	return
}

func (facex *Facex) AddFace(uri, id string) (err error) {
	input := NewFacexInput(uri, id)

	body, err := facex.Send(http.MethodPost, facex.API(groupAddAPI), mimeType, facex.timeout, toPayload(input))
	fmt.Println(string(body))
	return
}

func (facex *Facex) DeleteFace(ids []string) (err error) {
	in := map[string][]string{
		"faces": ids,
	}
	_, err = facex.Send(http.MethodPost, facex.API(groupDeleteAPI), mimeType, facex.timeout, toPayload(in))
	return
}

func (facex *Facex) Search(uri string) (result *SearchResult, err error) {
	input := NewSearchInput(uri)

	data, err := facex.Send(http.MethodPost, facex.API(groupSearchAPI), mimeType, facex.timeout, toPayload(input))
	if err != nil {
		return
	}

	result, err = NewSearchResult(data)
	return
}
