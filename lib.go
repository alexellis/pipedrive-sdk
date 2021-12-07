package pipedrivesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PipeDriveClient struct {
	Client       *http.Client
	APIKey       string
	Organization string
}

func NewPipeDriveClient(client *http.Client, apiKey, organization string) *PipeDriveClient {
	if client == nil {
		client = http.DefaultClient
	}

	return &PipeDriveClient{
		Client:       client,
		APIKey:       apiKey,
		Organization: organization,
	}
}

func (p *PipeDriveClient) CreateOrg(name string) (CreateOrgResponse, error) {
	r := CreateOrgResponse{}
	u, err := url.Parse("https://" + p.Organization + ".pipedrive.com/v1/organizations")
	if err != nil {
		return r, err
	}

	q := url.Values{
		"api_token": []string{p.APIKey},
	}
	u.RawQuery = q.Encode()

	f := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	out, _ := json.Marshal(f)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(out))
	if err != nil {
		return r, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := p.Client.Do(req)
	if err != nil {
		return r, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		return r, fmt.Errorf("createOrg status: %d, body: %s", res.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &r)
	return r, err
}

func (p *PipeDriveClient) CreatePerson(name, email string, orgID int) (CreatePersonResponse, error) {
	r := CreatePersonResponse{}
	u, err := url.Parse("https://" + p.Organization + ".pipedrive.com/v1/persons")
	if err != nil {
		return r, err
	}

	q := url.Values{
		"api_token": []string{p.APIKey},
	}
	u.RawQuery = q.Encode()

	f := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		OrgID int    `json:"org_id"`
	}{
		Name:  name,
		Email: email,
		OrgID: orgID,
	}

	out, _ := json.Marshal(f)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(out))
	if err != nil {
		return r, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := p.Client.Do(req)
	if err != nil {
		return r, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		return r, fmt.Errorf("createOrg status: %d, body: %s", res.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &r)
	return r, err
}

func (p *PipeDriveClient) SearchPerson(term string) (SearchPersonResponse, error) {
	r := SearchPersonResponse{}
	u, err := url.Parse("https://" + p.Organization + ".pipedrive.com/v1/persons/search")
	if err != nil {
		return r, err
	}
	q := url.Values{
		"term":      []string{term},
		"api_token": []string{p.APIKey},
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return r, err
	}

	res, err := p.Client.Do(req)
	if err != nil {
		return r, err
	}

	if res.StatusCode != http.StatusOK {
		return r, fmt.Errorf("searchPerson status: %d", res.StatusCode)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &r)
	return r, err
}

func (p *PipeDriveClient) SearchOrg(term string) (SearchResponse, error) {
	r := SearchResponse{}
	u, err := url.Parse("https://" + p.Organization + ".pipedrive.com/v1/organizations/search")
	if err != nil {
		return r, err
	}
	q := url.Values{
		"term":      []string{term},
		"api_token": []string{p.APIKey},
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return r, err
	}

	res, err := p.Client.Do(req)
	if err != nil {
		return r, err
	}

	if res.StatusCode != http.StatusOK {
		return r, fmt.Errorf("searchOrg status: %d", res.StatusCode)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &r)
	return r, err
}
