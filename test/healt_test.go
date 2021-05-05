package test

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-resty/resty/v2"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Runnning E2E test for health check endopoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}