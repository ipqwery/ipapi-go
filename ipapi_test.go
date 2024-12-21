package ipapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const mockIP = "1.1.1.1"

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

// go test -bench=BenchmarkQueryIP -benchmem
func BenchmarkQueryIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := QueryIP(mockIP)
		if err != nil {
			b.Fatalf("QueryIP failed: %v", err)
		}
	}
}

func mockServer(response string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	}))
}

// go test -bench=BenchmarkQueryOwnIP -benchmem
func BenchmarkQueryOwnIP(b *testing.B) {
	mockResp := `{"ip":"127.0.0.1"}`
	server := mockServer(mockResp, http.StatusOK)
	defer server.Close()

	baseURL = server.URL

	for i := 0; i < b.N; i++ {
		_, err := QueryOwnIP()
		if err != nil {
			b.Fatalf("QueryOwnIP failed: %v", err)
		}
	}
}
