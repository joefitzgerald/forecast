package forecast

type currentUserContainer struct {
	CurrentUser CurrentUser `json:"current_user"`
}

// CurrentUser contains information about the current Forecast user
type CurrentUser struct {
	ID             int   `json:"id"`
	AccountIds     []int `json:"account_ids"`
	IdentityUserID int   `json:"identity_user_id"`
}

// WhoAmI returns the CurrentUser for the logged in Forecast user
func (api *API) WhoAmI() (*CurrentUser, error) {
	container, err := get[currentUserContainer](api, "whoami")
	if err != nil {
		return nil, err
	}
	return &container.CurrentUser, nil
}
