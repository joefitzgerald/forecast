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

func TestAPI(t *testing.T) {
	spec.Run(t, "API", testAPI, spec.Report(report.Terminal{}))
}

func testAPI(t *testing.T, when spec.G, it spec.S) {
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

		api = forecast.New(server.URL, "test-token", "000000")
	})

	it.After(func() {
		api = nil
		if server == nil {
			return
		}
		server.Close()
		server = nil
	})

	when("when the access token is invalid", func() {
		it.Before(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := ReadFile("non-existent-token.json")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "%s", response)
			})
		})

		it("should return an unauthorized error", func() {
			user, err := api.WhoAmI()
			Expect(user).Should(BeNil())
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(Equal("401 Unauthorized: " + ReadFile("non-existent-token.json")))
		})
	})

	when("when the access token is valid", func() {
		it.Before(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := ReadFile("whoami.json")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s", response)
			})
		})

		it("should not return an unauthorized error", func() {
			user, err := api.WhoAmI()
			Expect(user).ShouldNot(BeNil())
			Expect(user.ID).Should(Equal(123456))
			Expect(len(user.AccountIds)).Should(Equal(2))
			Expect(user.AccountIds[0]).Should(Equal(111111))
			Expect(user.AccountIds[1]).Should(Equal(222222))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
}
