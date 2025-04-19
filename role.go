package forecast

import (
	"errors"
	"fmt"
)

type rolesContainer struct {
	Roles Roles `json:"roles"`
}

type roleContainer struct {
	Role Role `json:"role"`
}

// Roles is a list of roles
type Roles []Role

// Role is a role that can be assigned to a person in Forecast
type Role struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	PlaceholderIDs []int  `json:"placeholder_ids"`
	PersonIDs      []int  `json:"person_ids"`
	HarvestRoleID  int    `json:"harvest_role_id"`
}

// Roles returns all roles in Forecast
func (api *API) Roles() (Roles, error) {
	container, err := get[rolesContainer](api, "roles")
	if err != nil {
		return nil, err
	}
	return container.Roles, nil
}

// Role returns the role with the requested ID
func (api *API) Role(id int) (*Role, error) {
	if id == 0 {
		return nil, errors.New("cannot retrieve a rolee with an id of 0")
	}
	container, err := get[roleContainer](api, fmt.Sprintf("roles/%v", id))
	if err != nil {
		return nil, err
	}
	return &container.Role, nil
}
