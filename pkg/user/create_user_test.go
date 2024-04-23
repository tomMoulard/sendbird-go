package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	client := newClientMock(t)
	user := NewUser(client)

	createUserRequest := CreateUserRequest{
		UserID:                "user-id",
		Nickname:              "nickname",
		ProfileURL:            "profile-url",
		IssueAccessToken:      true,
		SessionTokenExpiresAt: 0,
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}

	createUserResponse := &CreateUserResponse{
		UserID:                     "user-id",
		Nickname:                   "nickname",
		ProfileURL:                 "profile-url",
		AccessToken:                "access-token",
		IsOnline:                   true,
		IsActive:                   true,
		IsCreated:                  true,
		PhoneNumber:                "phone-number",
		RequireAuthForProfileImage: true,
		SessionTokens:              []interface{}{},
		LastSeenAt:                 0,
		DiscoveryKeys:              []string{},
		PreferredLanguages:         []interface{}{},
		HasEverLoggedIn:            true,
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}

	client.OnGet("/user", createUserRequest, &CreateUserResponse{}).Return(createUserResponse, nil)

	cur, err := user.CreateUser(context.Background(), createUserRequest)
	require.NoError(t, err)
	assert.Equal(t, createUserResponse, cur)
}
