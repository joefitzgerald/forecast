package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/joefitzgerald/forecast"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testAssignment(t *testing.T, when spec.G, it spec.S) {
	var (
		server  *httptest.Server
		handler http.Handler
		api     *forecast.API
	)

	it.Before(func() {
		RegisterTestingT(t)
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if handler != nil {
				handler.ServeHTTP(w, r)
			}
		}))

		api = forecast.New(server.URL, "test-token", "987654")
	})

	it.After(func() {
		api = nil
		if server == nil {
			return
		}
		server.Close()
		server = nil
	})

	when("Weekdays()", func() {
		doTestAssignmentDays := func(start string, end string, expected int) {
			assignment := forecast.Assignment{
				StartDate: start,
				EndDate:   end,
			}

			actual := assignment.Weekdays()
			Expect(actual).Should(Equal(expected))
		}

		it("calculates assignment weekdays appropriately", func() {
			doTestAssignmentDays("2017-06-20", "2017-07-01", 9)
			doTestAssignmentDays("2017-06-20", "2017-06-20", 1)
			doTestAssignmentDays("2017-06-20", "2017-06-21", 2)
		})

		it("returns an error when an invalid date is supplied", func() {
			assignment := forecast.Assignment{
				StartDate: "-------%^#$@",
				EndDate:   "",
			}
			actual := assignment.Weekdays()
			Expect(actual).Should(Equal(0))
			assignment = forecast.Assignment{
				StartDate: "2017-06-20",
				EndDate:   "-------%^#$@",
			}
			actual = assignment.Weekdays()
			Expect(actual).Should(Equal(0))
		})
	})

	when("WorkingDaysBetween()", func() {
		doTestAssignmentDays := func(assignmentStart string, assignmentEnd string, startDate string, endDate string, expected int) {
			assignment := forecast.Assignment{
				StartDate: assignmentStart,
				EndDate:   assignmentEnd,
			}

			actual := assignment.WorkingDaysBetween(startDate, endDate)
			Expect(actual).Should(Equal(expected))
		}

		it("calculates assignment weekdays appropriately", func() {
			doTestAssignmentDays("2017-10-30", "2017-11-07", "2017-11-04", "2017-11-10", 2)
			doTestAssignmentDays("2017-10-30", "2017-11-13", "2017-11-04", "2017-11-10", 5)
			doTestAssignmentDays("2017-10-30", "2017-11-13", "", "2017-11-10", 10)
			doTestAssignmentDays("2017-10-30", "2017-11-13", "", "", 11)
			doTestAssignmentDays("2017-10-30", "2017-11-13", "2017-11-04", "", 6)
		})
	})

	when("AssignmentsWithFilter()", func() {
		var (
			validator func(values url.Values)
		)

		it.Before(func() {
			validator = func(values url.Values) {}
		})

		when("when a response is returned from the server", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := ReadFile("assignments.json")
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, "%s", response)
					validator(r.URL.Query())
				})
			})

			it("should return assignments and a nil error", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).To(BeZero())
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			it("should supply the project id filter", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).NotTo(BeZero())
					Expect(values.Get("project_id")).To(Equal("1"))
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{ProjectID: 1})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			it("should supply the person id filter", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).NotTo(BeZero())
					Expect(values.Get("person_id")).To(Equal("1155"))
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{PersonID: 1155})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			it("should supply the start date filter", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).NotTo(BeZero())
					Expect(values.Get("start_date")).To(Equal("2018-12-31"))
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{StartDate: "2018-12-31"})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			it("should supply the end date filter", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).NotTo(BeZero())
					Expect(values.Get("end_date")).To(Equal("2019-12-31"))
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{EndDate: "2019-12-31"})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			it("should supply the repeated assignment set id filter", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).NotTo(BeZero())
					Expect(values.Get("repeated_assignment_set_id")).To(Equal("10"))
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{RepeatedAssignmentSetID: 10})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})

			it("should supply the state filter", func() {
				validator = func(values url.Values) {
					Expect(values.Encode()).NotTo(BeZero())
					Expect(values.Get("state")).To(Equal("active"))
				}
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{State: "active"})
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		when("when an error is returned from the server", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := "error"
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "%s", response)
				})
			})

			it("should return an error", func() {
				assignments, err := api.AssignmentsWithFilter(forecast.AssignmentFilter{})
				Expect(assignments).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	when("Assignments()", func() {
		when("when a response is returned from the server", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := ReadFile("assignments.json")
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, "%s", response)
				})
			})

			it("should return assignments and a nil error", func() {
				assignments, err := api.Assignments()
				Expect(assignments).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		when("when an error is returned from the server", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := "error"
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "%s", response)
				})
			})

			it("should return an error", func() {
				assignments, err := api.Assignments()
				Expect(assignments).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})
		})
	})
}
