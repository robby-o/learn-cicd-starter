package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		headers    http.Header
		want       string
		wantErrMsg string
	}{
		"proper": {
			headers:    http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:       "abc123",
			wantErrMsg: "",
		},
		"no header": {
			headers:    http.Header{},
			want:       "",
			wantErrMsg: "no authorization header included",
		},
		"bad format": {
			headers:    http.Header{"Authorization": []string{"InvalidFormat"}},
			want:       "",
			wantErrMsg: "malformed authorization headers",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected API key: %#v, got: %#v", tc.want, got)
			}
			if err != nil && err.Error() != tc.wantErrMsg {
				t.Fatalf("expected error message: %#v, got: %#v", tc.wantErrMsg, err.Error())
			}
		})
	}
}
