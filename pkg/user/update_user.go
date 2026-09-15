package user

import (
	"context"
	"fmt"
)

// UpdateUserRequest is the request to update a user.
type UpdateUserRequest struct {
	// Nickname is the nickname of the user.
	// Maximum length is 80 characters.
	Nickname string `json:"nickname,omitempty"`
	// ProfileURL is the URL of the user's profile image.
	// If left empty, no profile image is set for the user. Maximum length is
	// 2,048 characters.
	ProfileURL string `json:"profile_url,omitempty"`
	// IssueAccessToken determines whether revoke the existing access token and
	// create a new one for the user.
	IssueAccessToken bool `json:"issue_access_token,omitempty"`
	// IsActive determines whether to activate or deactivate the user.
	IsActive bool `json:"is_active,omitempty"`
	// LastSeenAt is the time the user went offline in Unix milliseconds.
	LastSeenAt int64 `json:"last_seen_at,omitempty"`
	// DiscoveryKeys is the list of unique keys of the user, which is provided to
	// Sendbird server when searching for friends. The unique key acts as an
	// identifier for users' friends to find each other. The server uses
	// discovery keys to identify and match the user with other users.
	DiscoveryKeys []string `json:"discovery_keys,omitempty"`
	// PreferredLanguages is the list of one or more language codes to translate
	// notification messages to preferred languages. Up to four languages can be
	// set for the user. If messages are sent in one of the preferred languages,
	// notification messages won't be translated. If messages are sent in a
	// language other than the preferred languages, notification messages will be
	// translated into the first language in the array. Messages translated into
	// other preferred languages will be provided in the sendbird property of the
	// notification message payload.
	PreferredLanguages []string `json:"preferred_languages,omitempty"`
	// LeaveAllWhenDeactivatedDetermines whether the user leaves all joined group
	// channels upon deactivation. This property should be specified in
	// conjunction with the IsActive property above.
	LeaveAllWhenDeactivated bool `json:"leave_all_when_deactivated,omitempty"`
}

type UpdateUserResponse struct {
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

// UpdateUserRequest is the request to update a user.
// See https://sendbird.com/docs/chat/platform-api/v3/user/managing-users/update-a-user
func (u *user) UpdateUser(ctx context.Context, userID string, updateUserRequest UpdateUserRequest) (*UpdateUserResponse, error) {
	uur, err := u.client.Put(ctx, "/users/"+userID, updateUserRequest, &UpdateUserResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	updateUserResponse, ok := uur.(*UpdateUserResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to UpdateUserResponse: %+v", uur)
	}

	return updateUserResponse, nil
}
