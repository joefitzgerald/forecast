package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/joefitzgerald/forecast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Placeholder", func() {
	Context("Placeholders()", func() {
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
					response := ReadFile("placeholders.json")
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, "%s", response)
				}))

				api = New(server.URL, "test-token", "987654")
			})

			It("should return placeholders and a nil error", func() {
				placeholders, err := api.Placeholders()
				Expect(placeholders).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
				Expect(placeholders[0].Roles).ShouldNot(BeEmpty())
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
				placeholders, err := api.Placeholders()
				Expect(placeholders).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
