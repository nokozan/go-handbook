package service

import (
	"log"
	"testing"
	"time"
)

type testLogger struct {
	Logs []string
}

func (t *testLogger) Info(msg string)  { t.Logs = append(t.Logs, "[INFO] "+msg) }
func (t *testLogger) Error(msg string) { t.Logs = append(t.Logs, "[ERROR] "+msg) }

func TestFunctionalOptions(t *testing.T) {
	l := &testLogger{}
	svc, err := NewService(
		WithLogger(l),
		WithTimeout(2*time.Second),
		WithRetries(1),
	)
	if err != nil {
		log.Fatal("Failed to create service : ", err)
	}

	if svc.timeout != 2*time.Second {
		t.Errorf("expected timeout = 2s, got %v", svc.timeout)
	}

	if svc.retries != 1 {
		t.Errorf("expected retries = 1, got %v", svc.retries)
	}

	svc.Run()
	if len(l.Logs) == 0 {
		t.Error("expected logs to be captured")
	}
}
