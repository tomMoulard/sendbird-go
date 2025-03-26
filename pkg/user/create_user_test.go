package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	createUserRequest := CreateUserRequest{
		UserID:           "user-id",
		Nickname:         "nickname",
		ProfileURL:       "profile-url",
		IssueAccessToken: true,
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

	client := client.NewClientMock(t).
		OnPost("/users", createUserRequest, &CreateUserResponse{}).TypedReturns(createUserResponse, nil).Once().
		Parent
	user := NewUser(client)

	cur, err := user.CreateUser(t.Context(), createUserRequest)
	require.NoError(t, err)
	assert.Equal(t, createUserResponse, cur)
}
