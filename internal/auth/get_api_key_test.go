package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_EmptyHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_NoApiKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "Foo Bar")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetAPIKey_ApiKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey 12345")
	res, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res != "12345" {
		t.Fatalf("expected %v, got %v", "12345", res)
	}
}
