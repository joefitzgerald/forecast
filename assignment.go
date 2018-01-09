package forecast

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type assignmentsContainer struct {
	Assignments Assignments `json:"assignments"`
}

// Assignments is a list of assignments
type Assignments []Assignment

// Assignment is a Forecast assignment
type Assignment struct {
	ID                      int       `json:"id"`
	StartDate               string    `json:"start_date"`
	EndDate                 string    `json:"end_date"`
	Allocation              int       `json:"allocation"`
	Notes                   string    `json:"notes"`
	UpdatedAt               time.Time `json:"updated_at"`
	UpdatedByID             int       `json:"updated_by_id"`
	ProjectID               int       `json:"project_id"`
	PersonID                int       `json:"person_id"`
	PlaceholderID           int       `json:"placeholder_id"`
	RepeatedAssignmentSetID int       `json:"repeated_assignment_set_id"`
	ActiveOnDaysOff         bool      `json:"active_on_days_off"`
}

// AssignmentFilter is used to filter assignments
type AssignmentFilter struct {
	ProjectID               int
	PersonID                int
	StartDate               string // Format: YYYY-MM-DD
	EndDate                 string // Format: YYYY-MM-DD
	RepeatedAssignmentSetID int
	State                   string // active or archived
}

// Weekdays returns the number of working days between the start date and end date
// of the assignment
func (a *Assignment) Weekdays() int {
	start, err := time.Parse("2006-01-02", a.StartDate)
	if err != nil {
		return 0
	}

	finish, err := time.Parse("2006-01-02", a.EndDate)
	if err != nil {
		return 0
	}

	next := start
	result := 0
	for {
		if finish.Sub(next).Seconds() < 0 {
			break
		}
		switch next.Weekday() {
		case time.Monday:
			result = result + 1
		case time.Tuesday:
			result = result + 1
		case time.Wednesday:
			result = result + 1
		case time.Thursday:
			result = result + 1
		case time.Friday:
			result = result + 1
		}
		next = next.Add(time.Hour * 24)
	}

	return result
}

func (a *Assignment) WorkingDaysBetween(startDate string, endDate string) int {
	var start *time.Time
	if strings.TrimSpace(startDate) != "" {
		parsedStart, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			start = &parsedStart
		} else {
			fmt.Println(err)
		}
	}
	var end *time.Time
	if strings.TrimSpace(endDate) != "" {
		parsedEnd, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			end = &parsedEnd
		} else {
			fmt.Println(err)
		}
	}
	assignmentStart, err := time.Parse("2006-01-02", a.StartDate)
	if err != nil {
		return 0
	}

	assignmentEnd, err := time.Parse("2006-01-02", a.EndDate)
	if err != nil {
		return 0
	}

	next := assignmentStart
	result := 0
	for {
		if end != nil {
			// Don't go past the requested end date
			if end.AddDate(0, 0, 1).Before(next) {
				break
			}
		}

		// Assignment has not ended
		if assignmentEnd.Sub(next).Seconds() < 0 {
			break
		}

		if start != nil {
			// Skip assignment dates that are prior to the requested start date
			if start.After(next.AddDate(0, 0, -1)) {
				next = next.Add(time.Hour * 24)
				continue
			}
		}
		switch next.Weekday() {
		case time.Monday:
			result = result + 1
		case time.Tuesday:
			result = result + 1
		case time.Wednesday:
			result = result + 1
		case time.Thursday:
			result = result + 1
		case time.Friday:
			result = result + 1
		}
		next = next.Add(time.Hour * 24)
	}

	return result
}

// Assignments retrieves all assignments for the Forecast account
func (api *API) Assignments() (Assignments, error) {
	var container assignmentsContainer
	err := api.do("assignments", &container)
	if err != nil {
		return nil, err
	}
	return container.Assignments, nil
}

// AssignmentsWithFilter retrieves all assignments for the Forecast account
func (api *API) AssignmentsWithFilter(filter AssignmentFilter) (Assignments, error) {
	params := ToParams(filter.Values())
	var container assignmentsContainer
	err := api.do("assignments"+params, &container)
	if err != nil {
		return nil, err
	}
	return container.Assignments, nil
}

// ToParams formats url.Values as a string
func ToParams(values url.Values) string {
	if len(values) == 0 {
		return ""
	}
	return "?" + values.Encode()
}

// Values returns the AssignmentFilter as a url.Values result
func (filter *AssignmentFilter) Values() url.Values {
	result := url.Values{}
	if filter.ProjectID != 0 {
		result.Set("project_id", strconv.Itoa(filter.ProjectID))
	}
	if filter.PersonID != 0 {
		result.Set("person_id", strconv.Itoa(filter.PersonID))
	}
	if strings.TrimSpace(filter.StartDate) != "" {
		result.Set("start_date", filter.StartDate)
	}
	if strings.TrimSpace(filter.EndDate) != "" {
		result.Set("end_date", filter.EndDate)
	}
	if filter.RepeatedAssignmentSetID != 0 {
		result.Set("repeated_assignment_set_id", strconv.Itoa(filter.RepeatedAssignmentSetID))
	}
	if strings.TrimSpace(filter.State) != "" {
		result.Set("state", filter.State)
	}
	return result
}

// ToCSV writes the projects to the supplied writer in CSV
// format
func (assignments Assignments) ToCSV(w io.Writer) error {
	writer := csv.NewWriter(w)
	header := []string{
		"id",
		"start_date",
		"end_date",
		"allocation",
		"notes",
		"updated_at",
		"updated_by_id",
		"project_id",
		"person_id",
		"placeholder_id",
		"repeated_assignment_set_id",
	}
	err := writer.Write(header)
	if err != nil {
		return err
	}

	for _, assignment := range assignments {
		var record []string
		record = append(record, strconv.Itoa(assignment.ID))
		record = append(record, assignment.StartDate)
		record = append(record, assignment.EndDate)
		record = append(record, strconv.Itoa(assignment.Allocation))
		record = append(record, assignment.Notes)
		record = append(record, assignment.UpdatedAt.UTC().Format(time.RFC822Z))
		record = append(record, strconv.Itoa(assignment.UpdatedByID))
		record = append(record, strconv.Itoa(assignment.ProjectID))
		record = append(record, strconv.Itoa(assignment.PersonID))
		record = append(record, strconv.Itoa(assignment.PlaceholderID))
		record = append(record, strconv.Itoa(assignment.RepeatedAssignmentSetID))
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
