package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewResponseHeader(t *testing.T) {
	handlerToWrap := http.NewServeMux()

	type args struct {
		handlerToWrap http.Handler
		headerName    string
		headerValue   string
	}
	tests := []struct {
		name string
		args args
		want *ResponseHeader
	}{
		{
			name: "Should create a new response header",
			args: args{
				handlerToWrap: handlerToWrap,
				headerName:    "Content-Type",
				headerValue:   "application/json",
			},
			want: &ResponseHeader{
				handler:     handlerToWrap,
				headerName:  "Content-Type",
				headerValue: "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewResponseHeader(tt.args.handlerToWrap, tt.args.headerName, tt.args.headerValue)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponseHeader_ServeHTTP(t *testing.T) {
	tests := []struct {
		name          string
		headerName    string
		headerValue   string
		expectedValue string
	}{
		{
			name:          "Should add header to response",
			headerName:    "Content-Type",
			headerValue:   "application/json",
			expectedValue: "application/json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

			responseHeader := NewResponseHeader(handler, tt.headerName, tt.headerValue)

			req, err := http.NewRequest("GET", "/", nil)

			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			responseHeader.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedValue, rec.Header().Get(tt.headerName))
		})
	}
}
