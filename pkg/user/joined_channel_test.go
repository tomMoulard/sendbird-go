package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestGetGroupChannelCount(t *testing.T) {
	t.Parallel()

	getGroupChannelCountRequest := GetGroupChannelCountRequest{
		CustomTypes: []string{"custom-type"},
		HiddenMode:  ModeAll,
		State:       StateInvited,
		SuperMode:   SuperModeAll,
	}

	getGroupChannelCountResponse := &GetGroupChannelCountResponse{
		GroupChannelCount: 42,
	}

	client := client.NewClientMock(t).
		OnGet("/users/42/group_channel_count", getGroupChannelCountRequest, &GetGroupChannelCountResponse{}).TypedReturns(getGroupChannelCountResponse, nil).Once().
		Parent
	user := NewUser(client)

	cur, err := user.GetGroupChannelCount(t.Context(), "42", getGroupChannelCountRequest)
	require.NoError(t, err)
	assert.Equal(t, getGroupChannelCountResponse, cur)
}
