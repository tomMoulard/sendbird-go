package message_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/message"
)

// iencli wraps the message.Message interface.
type iencli struct {
	message message.Message
}

// SendMessage wraps the sends message method.
func (c *iencli) SendMessage(ctx context.Context, channelType message.ChannelType, channelURL string, sendMessageRequest message.SendMessageRequest) (*message.SendMessageResponse, error) {
	got, err := c.message.SendMessage(ctx, channelType, channelURL, sendMessageRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return got, nil
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	req := message.SendMessageRequest{
		Message: "hello",
	}
	messageMock := message.NewMessageMock(t).
		OnSendMessage(message.ChannelTypeGroup, "channelURL", req).TypedReturns(&message.SendMessageResponse{MessageID: 42}, nil).Once().
		Parent

	c := &iencli{
		message: messageMock,
	}

	got, err := c.SendMessage(t.Context(), message.ChannelTypeGroup, "channelURL", req)
	require.NoError(t, err)
	assert.Equal(t, 42, got.MessageID)
}
