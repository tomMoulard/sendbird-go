package client

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleError_notAnError(t *testing.T) {
	t.Parallel()

	c := &client{}
	c.SetDefault()

	body := `{"code":418,"message":"I'm a teapot","error":false}`
	err := c.handleError(http.StatusOK, strings.NewReader(body))
	assert.NoError(t, err)
}

func TestHandleError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		statusCode int
		body       string
		err        error
	}{
		{
			name: "no body",
			err:  io.EOF,
		},
		{
			name:       "unknown status code",
			statusCode: http.StatusTeapot,
			body:       `{"code":418,"message":"I'm a teapot","error":true}`,
			err:        ErrAPIDefault,
		},
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			body:       `{"code":418,"message":"I'm a teapot","error":true}`,
			err:        ErrAPIBadRequest,
		},
		{
			name:       "InternalError is an ErrAPIInternalServerError",
			statusCode: http.StatusBadRequest,
			body:       `{"code":500901,"message":"InternalError","error":true}`,
			err:        ErrAPIInternalServerError,
		},
		{
			name:       "InternalError",
			statusCode: http.StatusBadRequest,
			body:       `{"code":500901,"message":"InternalError","error":true}`,
			err:        ErrInternalError,
		},
		{
			name:       "User not found",
			statusCode: http.StatusBadRequest,
			body:       `{"message": "\"User\" not found.", "code": 400201, "error": true}`,
			err:        ErrResourceNotFound,
		},
		{
			name:       "User not found is also an ErrAPIBadRequest",
			statusCode: http.StatusBadRequest,
			body:       `{"message": "\"User\" not found.", "code": 400201, "error": true}`,
			err:        ErrAPIBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := &client{}
			c.SetDefault()

			err := c.handleError(test.statusCode, strings.NewReader(test.body))
			assert.ErrorIs(t, err, test.err)
		})
	}
}
