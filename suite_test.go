package forecast_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

var suite spec.Suite

func init() {
	suite = spec.New("forecast", spec.Report(report.Terminal{}))
	suite("Accounts", testAccounts)
	suite("Assignment", testAssignment)
	suite("Milestones", testMilestones)
	suite("Person", testPerson)
	suite("Placeholder", testPlaceholder)
	suite("Project", testProject)
	suite("WhoAmI", testWhoAmI)
}

func Test(t *testing.T) {
	suite.Run(t)
}
