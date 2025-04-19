package forecast

import "time"

type userConnectionsContainer struct {
	UserConnections UserConnections `json:"user_connections"`
}

// UserConnections is a list of UserConnection items
type UserConnections []UserConnection

// UserConnection includes information about currently connected users
type UserConnection struct {
	ID           int       `json:"id"`
	PersonID     int       `json:"person_id"`
	LastActiveAt time.Time `json:"last_active_at"`
}

// UserConnections returns all current user connections
func (api *API) UserConnections() (UserConnections, error) {
	container, err := get[userConnectionsContainer](api, "user_connections")
	if err != nil {
		return nil, err
	}
	return container.UserConnections, nil
}
