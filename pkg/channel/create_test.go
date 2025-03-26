package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestCreateGroupChannel(t *testing.T) {
	t.Parallel()

	createChannelRequest := CreateGroupChannelRequest{
		UserIDs:                []string{"42", "43"},
		Name:                   "name",
		ChannelURL:             "channel-url",
		CoverURL:               "cover-url",
		CustomType:             "custom-type",
		Data:                   `{"key": "value"}`,
		IsDistinct:             true,
		IsPublic:               true,
		IsSuper:                true,
		IsEphemeral:            true,
		AccessCode:             "access-code",
		InviterID:              "inviter-id",
		Strict:                 true,
		OperatorIDs:            []string{"44", "45"},
		BlocSDKUserChannelJoin: true,
	}

	createChannelResponse := &CreateGroupChannelResponse{
		Name: "name",
	}

	client := client.NewClientMock(t).
		OnPost("/group_channels", createChannelRequest, &CreateGroupChannelResponse{}).TypedReturns(createChannelResponse, nil).Once().
		Parent
	channel := NewChannel(client)

	cur, err := channel.CreateGroupChannel(t.Context(), createChannelRequest)
	require.NoError(t, err)
	assert.Equal(t, createChannelResponse, cur)
}
