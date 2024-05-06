package client

import (
	"bytes"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefault(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	assert.Equal(t, slog.Default(), client.logger)
	assert.Equal(t, http.DefaultClient, client.httpClient)
	assert.Equal(t, "https", client.baseURL.Scheme)
	assert.Equal(t, "/v3", client.baseURL.Path)
	assert.Equal(t, "application/json; charset=utf-8", client.header.Get("Content-Type"))
	assert.Len(t, client.header, 1)
}

func TestWithLogger(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&b, &slog.HandlerOptions{Level: slog.LevelDebug}))

	client := &client{}
	client = WithLogger(logger)(client)
	assert.Equal(t, logger, client.logger)
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	httpClient := http.DefaultClient

	client := &client{}
	client = WithHTTPClient(httpClient)(client)
	assert.Equal(t, httpClient, client.httpClient)
}

func TestWithScheme(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	client = WithScheme("grcp")(client)
	assert.Equal(t, "grcp", client.baseURL.Scheme)
}

func TestWithHost(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	client = WithHost("example.org:8080")(client)
	assert.Equal(t, "example.org:8080", client.baseURL.Host)
}

func TestWithPath(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	client = WithPath("/foo/bar")(client)
	assert.Equal(t, "/foo/bar", client.baseURL.Path)
}

func TestWithURL(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	client = WithURL("grcp://example.org:8080/foo/bar")(client)
	assert.Equal(t, "grcp", client.baseURL.Scheme)
	assert.Equal(t, "example.org:8080", client.baseURL.Host)
	assert.Equal(t, "/foo/bar", client.baseURL.Path)
}

func TestWithAPIToken(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	client = WithAPIToken("api-token")(client)
	assert.Equal(t, "api-token", client.header.Get("Api-Token"))
	assert.Len(t, client.header, 2)
}

func TestWithAPPID(t *testing.T) {
	t.Parallel()

	client := &client{}
	client.SetDefault()

	client = WithAPPID("api-id")(client)
	assert.Equal(t, "api-api-id.sendbird.com", client.baseURL.Host)
}
