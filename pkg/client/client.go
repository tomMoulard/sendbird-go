// This package is for the client of the sendbird API
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the interface for the client of the sendbird API.
type Client interface {
	Get(ctx context.Context, path string, obj any, resp any) (any, error)
	Post(ctx context.Context, path string, obj any, resp any) (any, error)
	Put(ctx context.Context, path string, obj any, resp any) (any, error)
	Delete(ctx context.Context, path string, obj any, resp any) (any, error)
}

// NewClient creates a new client for the sendbird API.
func NewClient(opts ...Option) Client {
	cfg := &client{}
	cfg.SetDefault()

	for _, opt := range opts {
		cfg = opt(cfg)
	}

	return cfg
}

// do sends a request to the sendbird API.
func (c *client) do(ctx context.Context, method, path string, obj any, resp any) (any, error) {
	logger := c.logger.With("method", method, "path", path)
	logger.Debug("do")

	var reqBody io.Reader

	if obj != nil {
		m, err := json.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal object: %w", err)
		}

		reqBody = bytes.NewReader(m)
	}

	url := c.baseURL
	url.Path += path
	url.Path = url.EscapedPath()

	logger = logger.With("url", url.Redacted())

	req, err := http.NewRequestWithContext(ctx, method, url.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header = c.header

	r, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer r.Body.Close()

	logger = logger.With("status", r.StatusCode)

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		logger.Error("request failed")

		return nil, c.handleError(r.StatusCode, r.Body)
	}

	if resp != nil {
		if err := json.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		logger = logger.With("response", resp)
	}

	logger.Debug("request succeeded")

	return resp, nil
}

// Get sends a GET request to the sendbird API.
func (c *client) Get(ctx context.Context, path string, obj any, resp any) (any, error) {
	return c.do(ctx, http.MethodGet, path, obj, resp)
}

// Post sends a POST request to the sendbird API.
func (c *client) Post(ctx context.Context, path string, obj any, resp any) (any, error) {
	return c.do(ctx, http.MethodPost, path, obj, resp)
}

// Put sends a PUT request to the sendbird API.
func (c *client) Put(ctx context.Context, path string, obj any, resp any) (any, error) {
	return c.do(ctx, http.MethodPut, path, obj, resp)
}

// Delete sends a DELETE request to the sendbird API.
func (c *client) Delete(ctx context.Context, path string, obj any, resp any) (any, error) {
	return c.do(ctx, http.MethodDelete, path, obj, resp)
}
