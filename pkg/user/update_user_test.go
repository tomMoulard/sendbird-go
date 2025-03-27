package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	updateUserRequest := UpdateUserRequest{
		Nickname:                "nickname",
		ProfileURL:              "profile-url",
		IssueAccessToken:        true,
		IsActive:                true,
		LastSeenAt:              0,
		DiscoveryKeys:           []string{"discovery-key"},
		PreferredLanguages:      []string{"en"},
		LeaveAllWhenDeactivated: true,
	}

	updateUserResponse := &UpdateUserResponse{
		UserID:                     "user-id",
		Nickname:                   "nickname",
		ProfileURL:                 "profile-url",
		AccessToken:                "access-token",
		IsOnline:                   true,
		IsActive:                   true,
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
		OnPut("/users/42", updateUserRequest, &UpdateUserResponse{}).TypedReturns(updateUserResponse, nil).Once().
		Parent
	user := NewUser(client)

	cur, err := user.UpdateUser(t.Context(), "42", updateUserRequest)
	require.NoError(t, err)
	assert.Equal(t, updateUserResponse, cur)
}
