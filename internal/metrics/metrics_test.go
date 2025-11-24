package metrics

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	serviceName := fmt.Sprintf("test_svc_%d", time.Now().UnixNano())

	var m *Metrics
	assert.NotPanics(t, func() {
		m = New(serviceName)
	})

	assert.NotNil(t, m)
	assert.NotNil(t, m.RequestsTotal)
	assert.NotNil(t, m.RequestDuration)

	t.Run("IncHttpRequestsTotal", func(t *testing.T) {
		assert.NotPanics(t, func() {
			m.IncHttpRequestsTotal(200, "GET", "/api/test")
		})
	})

	t.Run("ObserveHttpRequestDuration", func(t *testing.T) {
		assert.NotPanics(t, func() {
			m.ObserveHttpRequestDuration("POST", "/api/test", time.Second)
		})
	})
}
