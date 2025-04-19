package forecast

import "fmt"

type repeatedAssignmentSetsContainer struct {
	RepeatedAssignmentSets RepeatedAssignmentSets `json:"repeated_assignment_sets"`
}

// RepeatedAssignmentSets is a list of repeated assignment sets
type RepeatedAssignmentSets []RepeatedAssignmentSet

type repeatedAssignmentSetContainer struct {
	RepeatedAssignmentSet RepeatedAssignmentSet `json:"repeated_assignment_set"`
}

// RepeatedAssignmentSet is a repeated assignment set
type RepeatedAssignmentSet struct {
	ID             int    `json:"id"`
	FirstStartDate string `json:"first_start_date"`
	LastEndDate    string `json:"last_end_date"`
	AssignmentIds  []int  `json:"assignment_ids"`
}

// RepeatedAssignmentSets returns a list of repeated assignment sets
func (api *API) RepeatedAssignmentSets() (RepeatedAssignmentSets, error) {
	container, err := get[repeatedAssignmentSetsContainer](api, "repeated_assignment_sets")
	if err != nil {
		return nil, err
	}
	return container.RepeatedAssignmentSets, nil
}

// RepeatedAssignmentSet returns the repeated assignment set for the given id
func (api *API) RepeatedAssignmentSet(id int) (*RepeatedAssignmentSet, error) {
	container, err := get[repeatedAssignmentSetContainer](api, fmt.Sprintf("repeated_assignment_sets/%v", id))
	if err != nil {
		return nil, err
	}
	return &container.RepeatedAssignmentSet, nil
}
