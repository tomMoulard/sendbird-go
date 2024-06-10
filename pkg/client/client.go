// Package client package provides the interface for the client of the sendbird
// API. It provides the methods to call the sendbird API.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client is the interface for the client of the sendbird API.
type Client interface {
	// Get sends a GET request to the sendbird API.
	Get(ctx context.Context, path string, obj any, resp any) (any, error)
	// Post sends a POST request to the sendbird API.
	Post(ctx context.Context, path string, obj any, resp any) (any, error)
	// Put sends a PUT request to the sendbird API.
	Put(ctx context.Context, path string, obj any, resp any) (any, error)
	// Delete sends a DELETE request to the sendbird API.
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

// do send a request to the sendbird API.
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

	u := c.getURL(path)

	logger = logger.With("url", u.Redacted())

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
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

func (c *client) getURL(path string) *url.URL {
	uu, err := url.Parse(path)
	if err != nil {
		c.logger.Warn("failed to parse path", "path", path)
		uu = &url.URL{Path: path}
	}

	u := &url.URL{
		Scheme: c.baseURL.Scheme,
		Host:   c.baseURL.Host,
		Path:   c.baseURL.Path + uu.Path,
	}
	u.Path = u.EscapedPath()
	u.RawQuery = uu.Query().Encode()

	return u
}
