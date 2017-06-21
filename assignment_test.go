package forecast_test

import (
	"testing"

	"github.com/joefitzgerald/forecast"
)

func TestAssignmentDays(t *testing.T) {
	doTestAssignmentDays(t, "2017-06-20", "2017-07-01", 9)
	doTestAssignmentDays(t, "2017-06-20", "2017-06-20", 1)
	doTestAssignmentDays(t, "2017-06-20", "2017-06-21", 2)
}

func doTestAssignmentDays(t *testing.T, start string, end string, expected int) {
	assignment := forecast.Assignment{
		StartDate: start,
		EndDate:   end,
	}

	days := assignment.Weekdays()
	if days != expected {
		t.Errorf("for start of %s and end of %s, expected %v working days, but got %v working days", start, end, expected, days)
	}
}
