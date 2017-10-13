package forecast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (api *API) do(path string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.URL, path), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.token))
	req.Header.Set("Forecast-Account-ID", api.AccountID)
	if api.client == nil {
		jar, e := cookiejar.New(nil)
		if e != nil {
			return e
		}
		api.client = &http.Client{
			Jar: jar,
		}
	}
	r, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("%s: %s", r.Status, string(body))
	}

	return json.NewDecoder(r.Body).Decode(result)
}
