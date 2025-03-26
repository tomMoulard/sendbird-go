package channel

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestMarkAsRead(t *testing.T) {
	t.Parallel()

	client := client.NewClientMock(t).
		OnPut("/group_channels/channel-url/messages/mark_as_read", markAsReadRequest{UserID: "user-id"}, nil).TypedReturns(nil, nil).Once().
		Parent
	channel := NewChannel(client)

	err := channel.MarkAsRead(t.Context(), "channel-url", "user-id")
	require.NoError(t, err)
}
