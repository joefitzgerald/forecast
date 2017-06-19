package forecast

import "fmt"

type accountContainer struct {
	Account Account `json:"account"`
}

// Account is a Forecast account
type Account struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	WeeklyCapacity int    `json:"weekly_capacity"`
	ColorLabels    []struct {
		Name  string `json:"name"`
		Label string `json:"label"`
	} `json:"color_labels"`
	HarvestSubdomain string `json:"harvest_subdomain"`
	HarvestName      string `json:"harvest_name"`
}

// Account returns information about the current Forecast account
func (api *API) Account() (*Account, error) {
	var container accountContainer
	err := api.do(fmt.Sprintf("accounts/%v", api.AccountID), &container)
	if err != nil {
		return nil, err
	}
	return &container.Account, nil
}
