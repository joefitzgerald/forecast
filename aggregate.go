package forecast

import "fmt"

type remainingBudgetedHoursContainer struct {
	RemainingBudgetedHours RemainingBudgetedHours `json:"remaining_budgeted_hours"`
}

// RemainingBudgetedHours is a list of remaining budgeted hours items
type RemainingBudgetedHours []RemainingBudgetedHoursItem

// RemainingBudgetedHoursItem is an aggregate representing the remaining budgeted
// hours for the given Forecast project
type RemainingBudgetedHoursItem struct {
	ProjectID       int     `json:"project_id"`
	BudgetBy        string  `json:"budget_by"`
	BudgetIsMonthly bool    `json:"budget_is_monthly"`
	Hours           float64 `json:"hours"`
	ResponseCode    int     `json:"response_code"`
}

// RemainingBudgetedHours returns the remaining budgeted hours for all
// Forecast projects
func (api *API) RemainingBudgetedHours() (RemainingBudgetedHours, error) {
	container, err := get[remainingBudgetedHoursContainer](api, "aggregate/remaining_budgeted_hours")
	if err != nil {
		return nil, err
	}
	return container.RemainingBudgetedHours, nil
}

type futureScheduledHoursContainer struct {
	FutureScheduledHours FutureScheduledHours `json:"future_scheduled_hours"`
}

// FutureScheduledHours is a list of future scheduled hours items
type FutureScheduledHours []FutureScheduledHoursItem

// FutureScheduledHoursItem is a representation of the future scheduled hours for a project
type FutureScheduledHoursItem struct {
	ProjectID     int     `json:"project_id"`
	PersonID      int     `json:"person_id"`
	PlaceholderID int     `json:"placeholder_id"`
	Allocation    float64 `json:"allocation"`
}

// FutureScheduledHours returns all future scheduled hours using the supplied
// date as the starting point
func (api *API) FutureScheduledHours(from string) (FutureScheduledHours, error) {
	container, err := get[futureScheduledHoursContainer](api, fmt.Sprintf("aggregate/future_scheduled_hours/%s", from))
	if err != nil {
		return nil, err
	}
	return container.FutureScheduledHours, nil
}

// FutureScheduledHoursForProject returns all future scheduled hours for the
// given project using the supplied date as the starting point
func (api *API) FutureScheduledHoursForProject(from string, projectid int) (FutureScheduledHours, error) {
	container, err := get[futureScheduledHoursContainer](api, fmt.Sprintf("aggregate/future_scheduled_hours/%s?project_id=%v", from, projectid))
	if err != nil {
		return nil, err
	}
	return container.FutureScheduledHours, nil
}

// AssignedPeople returns a map of project ID to a slice of person IDs who are
// assigned to each project
func (api *API) AssignedPeople(from string, to string) (map[string][]int, error) {
	return get[map[string][]int](api, fmt.Sprintf("aggregate/projects/assigned_people?start_date=%s&end_date=%s", from, to))
}

type ProjectHeatmapItem struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// ProjectHeatmap returns an overview of projects with assignments for the given time period
func (api *API) ProjectHeatmap(from string, to string, projectID int, scale string) ([]ProjectHeatmapItem, error) {
	if scale == "" {
		scale = "daily"
	}
	return get[[]ProjectHeatmapItem](api, fmt.Sprintf("aggregate/heatmap/project?starting=%s&ending=%s&project_id=%d&scale=%s", from, to, projectID, scale))
}

type PersonHeatmapItem struct {
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	DailyAllocation int    `json:"daily_allocation"`
	DailyTimeOff    int    `json:"daily_time_off"`
}

// PersonHeatmap returns an overview of people with assignments for the given time period
func (api *API) PersonHeatmap(from string, to string, personID int, scale string) ([]PersonHeatmapItem, error) {
	if scale == "" {
		scale = "daily"
	}
	return get[[]PersonHeatmapItem](api, fmt.Sprintf("aggregate/heatmap/person?starting=%s&ending=%s&person_id=%d&scale=%s", from, to, personID, scale))
}

type PlaceholderHeatmapItem struct {
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	DailyAllocation int    `json:"daily_allocation"`
	DailyTimeOff    int    `json:"daily_time_off"`
}

// PlaceholderHeatmap returns an overview of placeholders with assignments for the given time period
func (api *API) PlaceholderHeatmap(from string, to string, placeholderID int, scale string) ([]PlaceholderHeatmapItem, error) {
	if scale == "" {
		scale = "daily"
	}
	return get[[]PlaceholderHeatmapItem](api, fmt.Sprintf("aggregate/heatmap/placeholder?starting=%s&ending=%s&placeholder_id=%d&scale=%s", from, to, placeholderID, scale))
}
