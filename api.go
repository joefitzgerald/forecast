package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

// Config holds the metadata necessary to authenticate with Forecast
type Config struct {
	Scheme    string `default:"https"`
	Host      string `default:"api.forecastapp.com"`
	AccountID string `required:"true" json:"forecast_accountid,omitempty"`
	Username  string `required:"true" json:"forecast_username,omitempty"`
	Password  string `required:"true" json:"forecast_password,omitempty"`
}

// Authenticate ensures the config has a valid Token
func (api *API) authenticate() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	api.client = &http.Client{
		Jar: jar,
	}

	authenticityToken, err := api.getAuthenticityToken()
	if err != nil {
		return err
	}
	err = api.createSession(authenticityToken)
	if err != nil {
		return err
	}
	return api.setToken()
}

// GetAuthenticityToken retrieves the Forecast CSRF authenticity-token
func (api *API) getAuthenticityToken() (string, error) {
	r, err := api.client.Get("https://id.getharvest.com/forecast/sign_in")
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	exp, err := regexp.Compile(`name="authenticity_token" value="(.*)"`)
	if err != nil {
		return "", err
	}
	authenticityToken := exp.FindStringSubmatch(body)
	if len(authenticityToken) == 0 {
		return "", errors.New("could not retrieve authenticity-token value")
	}

	return authenticityToken[1], nil
}

// CreateSession creates a Forecast session
func (api *API) createSession(authenticityToken string) error {
	values := url.Values{
		"authenticity_token": {authenticityToken},
		"email":              {api.Config.Username},
		"password":           {api.Config.Password},
		"product":            {"forecast"},
	}

	r, err := api.client.PostForm("https://id.getharvest.com/sessions", values)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	_, err = ioutil.ReadAll(r.Body)
	return err
}

// SetToken fetches a new token from Forecast and sets it on the config
func (api *API) setToken() error {
	api.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	r, err := api.client.Get(fmt.Sprintf("https://id.getharvest.com/accounts/%s", api.AccountID))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	location, err := r.Location()
	if err != nil {
		return err
	}
	url := location.String()
	token := strings.TrimSuffix(strings.TrimPrefix(url, fmt.Sprintf("https://forecastapp.com/%s/access_token/", api.AccountID)), "?")
	api.token = token
	return nil
}

// API provides access to the forecastapp.com API
type API struct {
	Config    *Config
	AccountID string
	token     string       `ignored:"true"`
	client    *http.Client `ignored:"true"`
}

// New returns a API that is authenticated with Forecast
func New(config *Config) (*API, error) {
	a := &API{
		Config:    config,
		AccountID: config.AccountID,
	}
	err := a.authenticate()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (api *API) do(path string, result interface{}) error {
	url := fmt.Sprintf("%s://%s/%s", api.Config.Scheme, api.Config.Host, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.token))
	req.Header.Set("Forecast-Account-ID", api.AccountID)
	if err != nil {
		return err
	}
	r, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(result)
}
