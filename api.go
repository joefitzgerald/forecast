package forecast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

// API provides access to the forecastapp.com API
type API struct {
	URL       string
	AccountID string
	token     string       `ignored:"true"`
	client    *http.Client `ignored:"true"`
}

// New returns a API that is authenticated with Forecast
func New(url string, accountID string, accessToken string) *API {
	return &API{
		URL:       url,
		AccountID: accountID,
		token:     accessToken,
	}
}

func (api *API) initClient() error {
	if api.client == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return err
		}
		api.client = &http.Client{
			Jar: jar,
		}
	}
	return nil
}

func get[T any](api *API, path string) (T, error) {
	var result T
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.URL, path), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.token))
	req.Header.Set("Forecast-Account-ID", api.AccountID)
	err = api.initClient()
	if err != nil {
		return result, err
	}
	r, err := api.client.Do(req)
	if err != nil {
		return result, err
	}
	defer r.Body.Close()
	if r.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return result, err
		}

		return result, fmt.Errorf("%s: %s", r.Status, string(body))
	}

	err = json.NewDecoder(r.Body).Decode(&result)
	return result, err
}

func mutate[T any](api *API, method string, path string, content T) error {
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", api.URL, path), bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.token))
	req.Header.Set("Forecast-Account-ID", api.AccountID)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	return doRequest(api, req)
}

func mutateNoBody(api *API, method string, path string) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", api.URL, path), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.token))
	req.Header.Set("Forecast-Account-ID", api.AccountID)
	return doRequest(api, req)
}

func doRequest(api *API, req *http.Request) error {
	err := api.initClient()
	if err != nil {
		return err
	}
	r, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("%s: %s", r.Status, string(body))
	}

	return nil
}
