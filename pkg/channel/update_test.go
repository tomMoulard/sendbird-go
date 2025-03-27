package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestUpdateGroupChannel(t *testing.T) {
	t.Parallel()

	updateChannelRequest := UpdateGroupChannelRequest{
		Name:        "channel-name",
		CoverURL:    "cover-url",
		CustomType:  "custom-type",
		Data:        `{"key": "value"}`,
		IsDistinct:  true,
		IsPublic:    true,
		IsSuper:     true,
		AccessCode:  "access-code",
		OperatorIDs: []string{"42", "43"},
	}

	updateChannelResponse := &UpdateGroupChannelResponse{
		Name: "channel-name",
	}

	client := client.NewClientMock(t).
		OnPut("/group_channels/channel-url", updateChannelRequest, &UpdateGroupChannelResponse{}).TypedReturns(updateChannelResponse, nil).Once().
		Parent
	channel := NewChannel(client)

	cur, err := channel.UpdateGroupChannel(t.Context(), "channel-url", updateChannelRequest)
	require.NoError(t, err)
	assert.Equal(t, updateChannelResponse, cur)
}
