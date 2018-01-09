package forecast

import (
	"errors"
	"fmt"
	"time"
)

type placeholdersContainer struct {
	Placeholders Placeholders `json:"placeholders"`
}

type placeholderContainer struct {
	Placeholder Placeholder `json:"placeholder"`
}

// Placeholders is a list of placeholders
type Placeholders []Placeholder

// Placeholder is a placeholder who is being scheduled in Forecast
type Placeholder struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Archived    bool      `json:"archived"`
	Teams       []string  `json:"teams"`
	Roles       []string  `json:"roles"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID int       `json:"updated_by_id"`
}

// Placeholders returns all placeholders being scheduled in Forecast
func (api *API) Placeholders() (Placeholders, error) {
	var container placeholdersContainer
	err := api.do("placeholders", &container)
	if err != nil {
		return nil, err
	}
	return container.Placeholders, nil
}

// Placeholder returns the placeholder with the requested ID
func (api *API) Placeholder(id int) (*Placeholder, error) {
	if id == 0 {
		return nil, errors.New("cannot retrieve a placeholder with an id of 0")
	}
	var container placeholderContainer
	err := api.do(fmt.Sprintf("placeholders/%v", id), &container)
	if err != nil {
		return nil, err
	}
	return &container.Placeholder, nil
}
