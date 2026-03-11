package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name           string
		requestHeaders map[string]string
		expected       string
		error          error
	}{
		{
			name: "Correct APIKey Pulled",
			requestHeaders: map[string]string{
				"Authorization": "ApiKey 12345",
			},
			expected: "12345",
			error:    nil,
		},
		{
			name: "no auth included",
			requestHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expected: "",
			error:    ErrNoAuthHeaderIncluded,
		},
	}
	req := httptest.NewRequest("GET", "/", nil)

	for _, c := range cases {
		req.Header = make(http.Header)
		for key, value := range c.requestHeaders {
			req.Header.Set(key, value)
		}

		str, err := GetAPIKey(req.Header)
		if str != c.expected {
			t.Errorf("returned wrong str. requestHeaders: %v, expected: %v, got: %v", c.requestHeaders, c.expected, str)
			t.Fail()
		}
		if err != c.error {
			t.Errorf("unexpected error value. expected error: %v, got: %v", c.error, err)
		}
	}
}
