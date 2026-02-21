package auth

import (
	"net/http"
	"testing"
)

func TestKey(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		test{"ApiKey goo", "goo"},
		test{"ApiKey g2o", "g2o"},
		test{"ApiKey g1o", "g1o"},
	}

	for _, tc := range tests {

		got, err := GetAPIKey(http.Header{"Authorization": {tc.input}})
		if err != nil {
			t.Fatalf("GetAPIKey failed: %v", err)
		}
		if got != tc.want {
			t.Fatalf("expected majik, got %s", got)
		}

	}

}
