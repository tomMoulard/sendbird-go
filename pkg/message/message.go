// Package message package provides the interface for the message service.
// It provides the methods to interact with the sendbird API.
// See https://sendbird.com/docs/chat/platform-api/v3/message/message-overview.
package message

import (
	"context"

	"github.com/tomMoulard/sendbird-go/pkg/client"
)

type Message interface {
	// SendMessage sends a message to a channel.
	// See https://sendbird.com/docs/chat/platform-api/v3/message/messaging-basics/send-a-message
	SendMessage(ctx context.Context, channelType ChannelType, channelURL string, sendMessageRequest SendMessageRequest) (*SendMessageResponse, error)

	// ListMessages retrieves a list of messages in a channel.
	// See https://sendbird.com/docs/chat/platform-api/v3/message/messaging-basics/list-messages
	ListMessages(ctx context.Context, channelType ChannelType, channelURL string, listMessagesRequest ListMessagesRequest) (*ListMessagesResponse, error)

	// MigrateMessages migrates messages to a channel.
	// See https://sendbird.com/docs/chat/platform-api/v3/message/migration/migrate-messages
	MigrateMessages(ctx context.Context, channelURL string, migrateMessagesRequest MigrateMessagesRequest) error
}

type message struct {
	client client.Client
}

func NewMessage(c client.Client) Message {
	return &message{client: c}
}
