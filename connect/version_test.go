package connect

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_IsGreaterOrEqualThan(t *testing.T) {
	cases := map[string]struct {
		serverVersion  version
		minimumVersion version
		expected       bool
	}{
		"equal": {
			serverVersion:  version{1, 2, 3},
			minimumVersion: version{1, 2, 3},
			expected:       true,
		},
		"lower major": {
			serverVersion:  version{1, 2, 3},
			minimumVersion: version{2, 2, 3},
			expected:       false,
		},
		"higher major": {
			serverVersion:  version{2, 0, 0},
			minimumVersion: version{1, 2, 3},
			expected:       true,
		},
		"higher minor": {
			serverVersion:  version{1, 3, 0},
			minimumVersion: version{1, 2, 3},
			expected:       true,
		},
		"lower minor": {
			serverVersion:  version{1, 1, 3},
			minimumVersion: version{1, 2, 3},
			expected:       false,
		},
		"higher patch": {
			serverVersion:  version{1, 2, 4},
			minimumVersion: version{1, 2, 3},
			expected:       true,
		},
		"lower patch": {
			serverVersion:  version{1, 2, 2},
			minimumVersion: version{1, 2, 3},
			expected:       false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.serverVersion.IsGreaterOrEqualThan(tc.minimumVersion))
		})
	}
}

func TestGetServerVersion(t *testing.T) {
	cases := map[string]struct {
		version           string
		expectedVersion   version
		expectedOrEarlier bool
		expectErr         bool
	}{
		"header set": {
			version:         "1.2.3",
			expectedVersion: version{1, 2, 3},
		},
		"header not set": {
			version:           "",
			expectedVersion:   version{1, 2, 0},
			expectedOrEarlier: true,
		},
		"malformed header": {
			version:   "1111.2",
			expectErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Set("1Password-Connect-Version", tc.version)
			resp := &http.Response{
				Header: header,
			}

			res, err := getServerVersion(resp)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tc.expectedVersion, res.version)
				assert.Equal(t, tc.expectedOrEarlier, res.orEarlier)
			}
		})
	}

}

func TestExpectMinimumVersion(t *testing.T) {
	cases := map[string]struct {
		minimumVersion version
		headerValue    string
		expectErr      bool
	}{
		"above minimum version": {
			minimumVersion: version{1, 2, 3},
			headerValue:    "1.3.0",
			expectErr:      false,
		},
		"below minimum version": {
			minimumVersion: version{1, 2, 3},
			headerValue:    "1.1.0",
			expectErr:      true,
		},
		"illegal version provided": {
			minimumVersion: version{1, 2, 3},
			headerValue:    "a",
			expectErr:      false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Set("1Password-Connect-Version", tc.headerValue)
			resp := &http.Response{
				Header: header,
			}

			err := expectMinimumConnectVersion(resp, tc.minimumVersion)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
