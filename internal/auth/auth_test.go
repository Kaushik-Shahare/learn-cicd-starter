package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKey tests the GetAPIKey function for different scenarios
func TestGetAPIKey(t *testing.T) {
	// Define a slice of test cases
	testCases := []struct {
		name        string         // Name of the test case
		headers     http.Header    // Input headers for the test case
		expectedKey string         // Expected API key
		expectError bool           // Whether an error is expected
	}{
		{
			name:        "Valid API Key Header",
			headers:     http.Header{"Authorization": {"ApiKey valid-api-key"}},
			expectedKey: "valid-api-key",
			expectError: false,
		},
		{
			name:        "Empty Header",
			headers:     http.Header{"Authorization": {""}},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Malformed Authorization Header",
			headers:     http.Header{"Authorization": {"Bearer token"}},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Incomplete Authorization Header",
			headers:     http.Header{"Authorization": {"ApiKey"}},
			expectedKey: "",
			expectError: true,
		},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function with the test headers
			apiKey, err := GetAPIKey(tc.headers)

			// Check for errors
			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
			}

			// Check the returned API key
			if apiKey != tc.expectedKey {
				t.Errorf("expected API key %q, got %q", tc.expectedKey, apiKey)
			}
		})
	}
}
