package forecast_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joefitzgerald/forecast"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testRaw(t *testing.T, when spec.G, it spec.S) {
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
		if server != nil {
			server.Close()
			server = nil
		}
	})

	when("the server returns a body", func() {
		var gotPath string
		it.Before(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.RequestURI()
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"people":[{"id":1,"first_name":"A","brand_new_attr":true}]}`)
			})
		})

		it("GetRaw returns the body verbatim", func() {
			raw, err := api.GetRaw("people")
			Expect(err).ShouldNot(HaveOccurred())
			var m map[string][]map[string]any
			Expect(json.Unmarshal(raw, &m)).To(Succeed())
			Expect(m["people"][0]["brand_new_attr"]).To(Equal(true))
			Expect(gotPath).To(Equal("/people"))
		})

		it("AssignmentsRaw applies the filter", func() {
			_, err := api.AssignmentsRaw(forecast.AssignmentFilter{StartDate: "2026-01-01", EndDate: "2026-03-31"})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(gotPath).To(Equal("/assignments?end_date=2026-03-31&start_date=2026-01-01"))
		})
	})

	when("the server returns an error", func() {
		it.Before(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprint(w, `{"error":"nope"}`)
			})
		})

		it("GetRaw returns the error", func() {
			_, err := api.GetRaw("people")
			Expect(err).Should(HaveOccurred())
		})
	})
}
