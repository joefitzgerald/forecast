package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joefitzgerald/forecast"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testAccounts(t *testing.T, when spec.G, it spec.S) {
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
				response := ReadFile("accounts.json")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s", response)
			})
		})

		it("should return an account and a nil error", func() {
			account, err := api.Account()
			Expect(account).ShouldNot(BeNil())
			Expect(account.ID).Should(Equal(987654))
			Expect(account.Name).Should(Equal("Test"))
			Expect(account.WeeklyCapacity).Should(Equal(144000))
			Expect(len(account.ColorLabels)).Should(Equal(8))
			Expect(account.HarvestName).Should(Equal("Test"))
			Expect(account.HarvestSubdomain).Should(Equal("test"))
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
			account, err := api.Account()
			Expect(account).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})
	})
}
