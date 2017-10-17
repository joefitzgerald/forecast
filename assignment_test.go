package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/joefitzgerald/forecast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assignment", func() {
	Context("Assignments()", func() {
		var (
			server *httptest.Server
			api    *API
		)

		AfterEach(func() {
			api = nil
			if server == nil {
				return
			}
			server.Close()
			server = nil
		})

		Context("when a response is returned from the server", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := ReadFile("assignments.json")
					fmt.Fprintf(w, "%s", response)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusOK)
				}))

				api = New(server.URL, "test-token", "987654")
			})

			It("should return assignments and a nil error", func() {
				assignments, err := api.Assignments()
				Ω(assignments).ShouldNot(BeNil())
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("when an error is returned from the server", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := "error"
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "%s", response)
				}))

				api = New(server.URL, "test-token", "987654")
			})

			It("should return an error", func() {
				assignments, err := api.Assignments()
				Ω(assignments).Should(BeNil())
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Context("Weekdays()", func() {
		doTestAssignmentDays := func(start string, end string, expected int) {
			assignment := Assignment{
				StartDate: start,
				EndDate:   end,
			}

			actual := assignment.Weekdays()
			Ω(actual).Should(Equal(expected))
		}
		It("calculates assignment weekdays appropriately", func() {
			doTestAssignmentDays("2017-06-20", "2017-07-01", 9)
			doTestAssignmentDays("2017-06-20", "2017-06-20", 1)
			doTestAssignmentDays("2017-06-20", "2017-06-21", 2)
		})

		It("returns an error when an invalid date is supplied", func() {
			assignment := Assignment{
				StartDate: "-------%^#$@",
				EndDate:   "",
			}
			actual := assignment.Weekdays()
			Ω(actual).Should(Equal(0))
			assignment = Assignment{
				StartDate: "2017-06-20",
				EndDate:   "-------%^#$@",
			}
			actual = assignment.Weekdays()
			Ω(actual).Should(Equal(0))
		})
	})
})
