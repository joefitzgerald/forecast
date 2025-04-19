package forecast

import (
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

func get[T any](api *API, path string) (T, error) {
	var result T
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.URL, path), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.token))
	req.Header.Set("Forecast-Account-ID", api.AccountID)
	if api.client == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return result, err
		}
		api.client = &http.Client{
			Jar: jar,
		}
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
