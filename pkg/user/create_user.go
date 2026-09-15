package user

import (
	"context"
	"fmt"
)

// CreateUserRequest is the request to create a user.
type CreateUserRequest struct {
	// UserID is the unique identifier of the user.
	// Maximum length is 80 characters.
	UserID string `json:"user_id"`
	// Nickname is the nickname of the user.
	// Maximum length is 80 characters.
	Nickname string `json:"nickname"`
	// ProfileURL is the URL of the user's profile image.
	// If left empty, no profile image is set for the user. Maximum length is
	// 2,048 characters.
	ProfileURL string `json:"profile_url"`

	// IssueAccessToken determines whether to create an access token for the
	// user.
	IssueAccessToken bool `json:"issue_access_token,omitempty"`
	// Metadata is the custom data of the user.
	// Specifies a JSON object to store up to five key-value items for additional
	// user information such as their preference settings. The key must not have
	// a comma (,), and the value must be a string.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreateUserResponse is the response of the create user request.
type CreateUserResponse struct {
	UserID                     string         `json:"user_id"`
	Nickname                   string         `json:"nickname"`
	ProfileURL                 string         `json:"profile_url"`
	AccessToken                string         `json:"access_token"`
	IsOnline                   bool           `json:"is_online"`
	IsActive                   bool           `json:"is_active"`
	IsCreated                  bool           `json:"is_created"`
	PhoneNumber                string         `json:"phone_number"`
	RequireAuthForProfileImage bool           `json:"require_auth_for_profile_image"`
	SessionTokens              []any          `json:"session_tokens"`
	LastSeenAt                 int            `json:"last_seen_at"`
	DiscoveryKeys              []string       `json:"discovery_keys"`
	PreferredLanguages         []any          `json:"preferred_languages"`
	HasEverLoggedIn            bool           `json:"has_ever_logged_in"`
	Metadata                   map[string]any `json:"metadata"`
}

// CreateUser creates a user.
// See https://sendbird.com/docs/chat/platform-api/v3/user/creating-users/create-a-user
func (u *user) CreateUser(ctx context.Context, createUserRequest CreateUserRequest) (*CreateUserResponse, error) {
	cur, err := u.client.Post(ctx, "/users", createUserRequest, &CreateUserResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	createUserResponse, ok := cur.(*CreateUserResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to CreateUserResponse: %+v", cur)
	}

	return createUserResponse, nil
}
