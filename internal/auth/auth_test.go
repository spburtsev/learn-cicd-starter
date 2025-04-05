package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	hdrWithAuth := http.Header{}
	want := "dummy"
	hdrWithAuth.Add("Authorization", fmt.Sprintf("ApiKey %s", want))

	got, err := GetAPIKey(hdrWithAuth)
	if err != nil {
		t.Errorf("%s", err)
	}
	if got != want {
		t.Errorf("expected '%s', got '%s'", want, got)
	}
}
