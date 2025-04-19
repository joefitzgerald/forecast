package forecast

import "time"

type clientsContainer struct {
	Clients Clients `json:"clients"`
}

// Clients is a list of clients
type Clients []Client

// Client is a client may have one or more projects
type Client struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	HarvestID   int       `json:"harvest_id"`
	Archived    bool      `json:"archived"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID int       `json:"updated_by_id"`
}

// Clients retrieves all clients in the Forecast account
func (api *API) Clients() (Clients, error) {
	container, err := get[clientsContainer](api, "clients")
	if err != nil {
		return nil, err
	}
	return container.Clients, nil
}
