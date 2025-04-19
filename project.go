package forecast

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"
)

type projectsContainer struct {
	Projects Projects `json:"projects"`
}

// Projects is a list of projects
type Projects []Project

// Project is a Forecast project
type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Code        string    `json:"code"`
	Notes       string    `json:"notes"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	HarvestID   int       `json:"harvest_id"`
	Archived    bool      `json:"archived"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID int       `json:"updated_by_id"`
	ClientID    int       `json:"client_id"`
	Tags        []string  `json:"tags"`
}

// Projects returns the list of projects in the Forecast Account
func (api *API) Projects() (Projects, error) {
	container, err := get[projectsContainer](api, "projects")
	if err != nil {
		return nil, err
	}
	return container.Projects, nil
}

// ToCSV writes the projects to the supplied writer in CSV
// format
func (projects Projects) ToCSV(w io.Writer) error {
	writer := csv.NewWriter(w)
	header := []string{
		"id",
		"name",
		"color",
		"code",
		"notes",
		"start_date",
		"end_date",
		"harvest_id",
		"archived",
		"updated_at",
		"updated_by_id",
		"client_id",
		"tags",
	}
	err := writer.Write(header)
	if err != nil {
		return err
	}

	for _, project := range projects {
		var record []string
		record = append(record, strconv.Itoa(project.ID))
		record = append(record, project.Name)
		record = append(record, project.Color)
		record = append(record, project.Code)
		record = append(record, project.Notes)
		record = append(record, project.StartDate)
		record = append(record, project.EndDate)
		record = append(record, strconv.Itoa(project.HarvestID))
		record = append(record, strconv.FormatBool(project.Archived))
		record = append(record, project.UpdatedAt.UTC().Format(time.RFC822Z))
		record = append(record, strconv.Itoa(project.UpdatedByID))
		record = append(record, strconv.Itoa(project.ClientID))
		record = append(record, strings.Join(project.Tags, "|"))
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
