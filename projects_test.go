package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joefitzgerald/forecast"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestProject(t *testing.T) {
	spec.Run(t, "Project", testProject, spec.Report(report.Terminal{}))
}

func testProject(t *testing.T, when spec.G, it spec.S) {
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

	when("when a response is returned from the server", func() {
		it.Before(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := ReadFile("projects.json")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s", response)
			})
		})

		it("should return projects and a nil error", func() {
			projects, err := api.Projects()
			Expect(projects).ShouldNot(BeNil())
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
			projects, err := api.Projects()
			Expect(projects).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})
	})
}
