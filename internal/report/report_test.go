package report_test

import (
	"enabling_tech/interview/internal/report"
	"errors"
	"testing"
	"time"
)

func TestGetCustomerOrders(t *testing.T) {
	customerOrders, err := report.GetCustomerOrders()
	if err != nil {
		t.Fatal(err)
	}

	if customerOrders == nil {
		t.Fatal(errors.New("customer orders cannot be nil"))
	}

	t.Logf("test successfully run on: %s", time.Now().Format(time.RFC3339))
}
