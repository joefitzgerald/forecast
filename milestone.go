package forecast

import "time"

type milestonesContainer struct {
	Milestones Milestones `json:"milestones"`
}

// Milestones is a list of milestones
type Milestones []Milestone

// Milestone is a Forecast milestone
type Milestone struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Date        string    `json:"date"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID int       `json:"updated_by_id"`
	ProjectID   int       `json:"project_id"`
}

// Milestones returns all milestones in the Forecast account
func (api *API) Milestones() (Milestones, error) {
	container, err := get[milestonesContainer](api, "milestones")
	if err != nil {
		return nil, err
	}
	return container.Milestones, nil
}
