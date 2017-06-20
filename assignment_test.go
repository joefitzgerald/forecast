package forecast_test

import (
	"testing"
	"time"

	"github.com/joefitzgerald/forecast"
)

func TestAssignmentDays(t *testing.T) {
	assignment := forecast.Assignment{
		StartDate: "2017-06-20",
		EndDate:   "2017-07-01",
	}
	start, _ := time.Parse("2006-01-02", assignment.StartDate)

	t.Log(start)

	days := assignment.Weekdays()
	if days != 9 {
		t.Errorf("expected %v, got %v", 9, days)
	}
}
