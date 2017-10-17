package forecast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/joefitzgerald/forecast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API", func() {
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

	Context("when the access token is invalid", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := ReadFile("non-existent-token.json")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "%s", response)
			}))
			api = New(server.URL, "test-token", "000000")
		})

		It("should return an unauthorized error", func() {
			user, err := api.WhoAmI()
			Ω(user).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("401 Unauthorized: " + ReadFile("non-existent-token.json")))
		})
	})

	Context("when the access token is valid", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := ReadFile("whoami.json")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s", response)
			}))
			api = New(server.URL, "test-token", "000000")
		})

		It("should not return an unauthorized error", func() {
			user, err := api.WhoAmI()
			Ω(user).ShouldNot(BeNil())
			Ω(user.ID).Should(Equal(123456))
			Ω(len(user.AccountIds)).Should(Equal(2))
			Ω(user.AccountIds[0]).Should(Equal(111111))
			Ω(user.AccountIds[1]).Should(Equal(222222))
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
