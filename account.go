package forecast

import (
	"fmt"
	"time"
)

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
	HarvestSubdomain   string    `json:"harvest_subdomain"`
	HarvestLink        string    `json:"harvest_link"`
	SamlSignInRequired bool      `json:"saml_sign_in_required"`
	HarvestName        string    `json:"harvest_name"`
	WeekendsEnabled    bool      `json:"weekends_enabled"`
	CreatedAt          time.Time `json:"created_at"`
	CreatorFirstName   string    `json:"creator_first_name"`
	CreatorLastName    string    `json:"creator_last_name"`
	BillingStatus      string    `json:"billing_status"`
	GDPR               bool      `json:"gdpr"`
}

// Account returns information about the current Forecast account
func (api *API) Account() (*Account, error) {
	container, err := get[accountContainer](api, fmt.Sprintf("accounts/%v", api.AccountID))
	if err != nil {
		return nil, err
	}
	return &container.Account, nil
}
