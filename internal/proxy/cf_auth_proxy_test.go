package proxy_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/pivotal/cf-auth-proxy/internal/auth"
	. "github.com/pivotal/cf-auth-proxy/internal/proxy"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFAuthProxy", func() {
	It("proxies requests to gateways", func() {
		var called bool
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				called = true
			}))

		proxy := NewCFAuthProxy(testServer.URL, "127.0.0.1:0")
		proxy.Start()

		resp, err := http.Get("http://" + proxy.Addr())
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(called).To(BeTrue())
	})

	It("delegates to the auth middleware", func() {
		var middlewareCalled bool
		middleware := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			middlewareCalled = true
			w.WriteHeader(http.StatusNotFound)
		})

		proxy := NewCFAuthProxy(
			"https://127.0.0.1",
			"127.0.0.1:0",
			WithAuthMiddleware(func(http.Handler) http.Handler {
				return middleware
			}),
		)
		proxy.Start()

		resp, err := http.Get("http://" + proxy.Addr())
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		Expect(middlewareCalled).To(BeTrue())
	})

	It("delegates to the access middleware", func() {
		var middlewareCalled bool
		middleware := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			middlewareCalled = true
			w.WriteHeader(http.StatusNotFound)
		})

		proxy := NewCFAuthProxy(
			"https://127.0.0.1",
			"127.0.0.1:0",
			WithAccessMiddleware(func(http.Handler) *auth.AccessHandler {
				return auth.NewAccessHandler(middleware, auth.NewNullAccessLogger(), "0.0.0.0", "1234")
			}),
		)
		proxy.Start()

		resp, err := http.Get("http://" + proxy.Addr())
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		Expect(middlewareCalled).To(BeTrue())
	})
})
