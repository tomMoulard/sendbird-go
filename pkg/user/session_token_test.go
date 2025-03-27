package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestGetSessionToken(t *testing.T) {
	t.Parallel()

	getSessionTokenRequest := GetSessionTokenRequest{
		ExpiresAt: 0,
	}

	getSessionTokenResponse := &GetSessionTokenResponse{
		Token:     "token",
		ExpiresAt: 0,
	}

	client := client.NewClientMock(t).
		OnPost("/users/42/token", getSessionTokenRequest, &GetSessionTokenResponse{}).TypedReturns(getSessionTokenResponse, nil).Once().
		Parent
	user := NewUser(client)

	cur, err := user.GetSessionToken(t.Context(), "42", getSessionTokenRequest)
	require.NoError(t, err)
	assert.Equal(t, getSessionTokenResponse, cur)
}
