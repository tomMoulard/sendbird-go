package message

import (
	"context"
	"errors"
	"fmt"
)

type TextMessage struct {
	// UserID specifies the user ID of the sender - required.
	UserID string `json:"user_id"`

	// MessageType specifies the type of the message - required.
	MessageType MessageType `json:"message_type"`

	// Message specifies the content of the message - required.
	Message string `json:"message,omitempty"`

	// Timestamp specifies the time when the message was sent in Unix milliseconds format - required.
	Timestamp int64 `json:"timestamp,omitempty"`

	// CustomType specifies a custom message type used for message grouping. The
	// length is limited to 128 characters - optional.
	CustomType string `json:"custom_type,omitempty"`

	// MentionUserIDs specifies an array of IDs of the users to mention in the
	// message. This property is used only when mention_type is users - optional.
	MentionUserIDs []string `json:"mentioned_user_ids,omitempty"`

	// Data specifies additional message information. This property serves as a
	// container for a long text of any type of characters which can also be a
	// JSON-formatted string like {"font-size": "24px"} - optional.
	Data string `json:"data,omitempty"`

	// DedupID specifies a unique ID for the message created by another system.
	// In general, this property is used to prevent the same message data from
	// getting inserted when migrating messages from another system to the
	// Sendbird server. If specified, the server performs a duplicate check using
	// the property value - optional.
	DedupID string `json:"dedup_id,omitempty"`
}

// MigrateMessagesRequest is the request to migrate messages to a channel.
type MigrateMessagesRequest struct {
	// Messages specifies an array of messages to migrate - required.
	Messages []TextMessage `json:"messages"`

	// UpdateReadTS determines whether to update the read receipt time for all channel members when message.timestamp
	// of the latest migrated message is prior to their read receipt time
	UpdateReadTS bool `json:"update_read_ts,omitempty"`
}

func (smr *MigrateMessagesRequest) Validate() error {
	if len(smr.Messages) == 0 {
		return errors.New("messages cannot be empty")
	}

	for _, message := range smr.Messages {
		if message.UserID == "" {
			return errors.New("user_id cannot be empty")
		}

		if message.MessageType == "" {
			return errors.New("message_type cannot be empty")
		}

		if message.Message == "" {
			return errors.New("message cannot be empty")
		}

		if message.Timestamp == 0 {
			return errors.New("timestamp cannot be empty")
		}
	}

	return nil
}

// MigrateMessages migrates messages to a channel.
// See https://sendbird.com/docs/chat/platform-api/v3/message/migration/migrate-messages
func (m *message) MigrateMessages(ctx context.Context, channelURL string, migrateMessagesRequest MigrateMessagesRequest) error {
	if err := migrateMessagesRequest.Validate(); err != nil {
		return fmt.Errorf("failed to validate migrate messages request: %w", err)
	}

	path := "/migration/" + channelURL

	_, err := m.client.Post(ctx, path, migrateMessagesRequest, nil)
	if err != nil {
		return fmt.Errorf("failed to migrate messages: %w", err)
	}

	return nil
}
