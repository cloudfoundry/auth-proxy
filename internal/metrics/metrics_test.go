package metrics_test

import (
	"github.com/pivotal/cf-auth-proxy/internal/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metrics", func() {
	It("publishes the total of a counter", func() {
		m := metrics.New()
		c := m.NewCounter("some_counter")
		c(99)
		c(101)

		Expect(m.Registry).To(ContainCounterMetric("some_counter", 200))
	})

	It("publishes the value of a gauge", func() {
		m := metrics.New()
		c := m.NewGauge("some_gauge", "some_unit")
		c(99.9)
		c(101.1)

		Expect(m.Registry).To(ContainGaugeMetric("some_gauge", "some_unit", 101.1))
	})
})
