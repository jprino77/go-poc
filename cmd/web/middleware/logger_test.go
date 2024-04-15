package middleware

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestLogger_ServeHTTP(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		url         string
		expectedLog string
	}{
		{
			name:        "Log GET request",
			method:      "GET",
			url:         "/test",
			expectedLog: "GET /test",
		},
		{
			name:        "Log POST request",
			method:      "POST",
			url:         "/another",
			expectedLog: "POST /another",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var buf bytes.Buffer
			log.SetOutput(&buf)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

			logger := NewLogger(handler)

			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			logger.ServeHTTP(nil, req)

			expectedLog := tt.expectedLog
			assert.Contains(t, buf.String(), expectedLog, "handler did not log the expected message")
		})
	}
}

func TestNewLogger(t *testing.T) {
	handlerToWrap := http.NewServeMux()
	type args struct {
		handlerToWrap http.Handler
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		{
			name: "Create a new logger instance",
			args: args{
				handlerToWrap,
			},
			want: &Logger{
				handlerToWrap,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger(tt.args.handlerToWrap)
			assert.Equal(t, tt.want, got)
		})
	}
}
