package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKet(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "")
	result, err := GetAPIKey(header)
	if result != "" || err != ErrNoAuthHeaderIncluded {
		t.Errorf("No error thrown for when authorization attribute is missing from header.")
	}
}
