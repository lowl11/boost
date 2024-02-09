package stat

import (
	"github.com/lowl11/boost/internal/healthcheck"
	"github.com/lowl11/boost/pkg/io/hard"
)

func Format(healthcheck *healthcheck.Healthcheck) map[string]any {
	healthStatus := "OK"
	if err := healthcheck.Trigger(); err != nil {
		healthStatus = err.Error()
	}

	return map[string]any{
		"healthcheck": healthStatus,
		"memory":      hard.MemoryFormat(),
		"goroutines":  hard.ActiveGoroutines(),
	}
}
