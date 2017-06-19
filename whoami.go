package forecast

type currentUserContainer struct {
	CurrentUser CurrentUser `json:"current_user"`
}

// CurrentUser contains information about the current Forecast user
type CurrentUser struct {
	ID         int   `json:"id"`
	AccountIds []int `json:"account_ids"`
}

// WhoAmI returns the CurrentUser for the logged in Forecast user
func (api *API) WhoAmI() (*CurrentUser, error) {
	var container currentUserContainer
	err := api.do("whoami", &container)
	if err != nil {
		return nil, err
	}
	return &container.CurrentUser, nil
}
