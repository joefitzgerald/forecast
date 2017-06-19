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
	ProjectID    int     `json:"project_id"`
	BudgetBy     string  `json:"budget_by"`
	Hours        float64 `json:"hours"`
	ResponseCode int     `json:"response_code"`
}

// RemainingBudgetedHours returns the remaining budgeted hours for all
// Forecast projects
func (api *API) RemainingBudgetedHours() (RemainingBudgetedHours, error) {
	var container remainingBudgetedHoursContainer
	err := api.do("aggregate/remaining_budgeted_hours", &container)
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
	ProjectID  int     `json:"project_id"`
	PersonID   int     `json:"person_id"`
	Allocation float64 `json:"allocation"`
}

// FutureScheduledHours returns all future scheduled hours using the supplied
// date as the starting point
func (api *API) FutureScheduledHours(from string) (FutureScheduledHours, error) {
	var container futureScheduledHoursContainer
	err := api.do(fmt.Sprintf("aggregate/future_scheduled_hours/%s", from), &container)
	if err != nil {
		return nil, err
	}
	return container.FutureScheduledHours, nil
}

// FutureScheduledHoursForProject returns all future scheduled hours for the
// given project using the supplied date as the starting point
func (api *API) FutureScheduledHoursForProject(from string, projectid int) (FutureScheduledHours, error) {
	var container futureScheduledHoursContainer
	err := api.do(fmt.Sprintf("aggregate/future_scheduled_hours/%s?project_id=%v", from, projectid), &container)
	if err != nil {
		return nil, err
	}
	return container.FutureScheduledHours, nil
}
