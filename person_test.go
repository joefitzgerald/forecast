package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/joefitzgerald/forecast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Person", func() {
	Context("People()", func() {
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
					response := ReadFile("people.json")
					fmt.Fprintf(w, "%s", response)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusOK)
				}))

				api = New(server.URL, "test-token", "987654")
			})

			It("should return people and a nil error", func() {
				people, err := api.People()
				Ω(people).ShouldNot(BeNil())
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("when an error is returned from the server", func() {
			BeforeEach(func() {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					response := "error"
					fmt.Fprintf(w, "%s", response)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusBadRequest)
				}))

				api = New(server.URL, "test-token", "987654")
			})

			It("should return an error", func() {
				people, err := api.People()
				Ω(people).Should(BeNil())
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
