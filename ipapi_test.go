package ipapi

import (
	"testing"
)

func TestQueryIP(t *testing.T) {
	ipInfo, err := QueryIP("8.8.8.8")
	if err != nil {
		t.Fatalf("Failed to query IP: %v", err)
	}
	if ipInfo.IP != "8.8.8.8" {
		t.Errorf("Expected IP to be '8.8.8.8', got %s", ipInfo.IP)
	}
}

func TestQueryOwnIP(t *testing.T) {
	ip, err := QueryOwnIP()
	if err != nil {
		t.Fatalf("Failed to fetch own IP: %v", err)
	}
	if ip == "" {
		t.Error("Expected non-empty IP")
	}
}
