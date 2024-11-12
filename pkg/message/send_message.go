package message

import (
	"context"
	"errors"
	"fmt"
)

// File is a fiel to send.
type File struct {
	// URL specifies the URL of the file hosted on the server of your own or
	// other third-party companies.
	URL string `json:"url"`
	// FileName specifies the name of the file. The file_name can be any string
	// you've set to this property. If no file_name value is supplied, the
	// files.file_name property defaults to an empty string.
	FileName string `json:"file_name,omitempty"`
	// FileSize specifies the size of the file in bytes. The file_size can be set
	// to be any int you've set to this property. If no files.file_size value is
	// supplied, it defaults to 0. File size is stored in bytes.
	FileSize int `json:"file_size,omitempty"`
	// FileType specifies the type of the file. The file_type can be set to any
	// string you've set to this property. If no files.file_type value is
	// supplied, it defaults to an empty string.
	FileType string `json:"file_type,omitempty"`
}

// SendMessageRequest is the request to send a message.
type SendMessageRequest struct {
	// MessageType specifies the type of the message.
	MessageType Type `json:"message_type"`
	// UserID specifies the user ID of the sender.
	UserID string `json:"user_id"`

	// Message specifies the content of the message.
	Message string `json:"message,omitempty"`
	// Files specifies the data of files to upload to the Sendbird server by
	// their location.
	Files []File `json:"files,omitempty"`

	// Thumbnails specifies an array of external thumbnail image URLs to store a
	// reference to those images in the Sendbird server.
	Thumbnails []string `json:"thumbnails,omitempty"`
	// CustomType specifies a custom message type used for message grouping. The
	// length is limited to 128 characters.
	CustomType string `json:"custom_type,omitempty"`
	// Data specifies additional message information. This property serves as a
	// container for a long text of any type of characters which can also be a
	// JSON-formatted string like {"font-size": "24px"}.
	Data string `json:"data,omitempty"`
	// SendPush determines whether to send a push notification of the message to
	// the channel members. This property only applies to group channels.
	// (Default: true)
	SendPush *bool `json:"send_push,omitempty"`
	// PushMessageTemplate specifies the content of a push notification
	// customized for the message. This property only applies to group channels.
	PushMessageTemplate string `json:"push_message_template,omitempty"`
	// MentionType specifies whether to mention specific users or all users in
	// the channel. Acceptable values are users and channel.
	// If set to MentionTypeUsers, up to ten users in the mentioned_user_ids
	// property below are notified of the mention.
	// If set to MentionTypeChannels, up to ten users in the channel are notified
	// of the mention.
	// (Default: MentionTypeUsers)
	MentionType MentionType `json:"mention_type,omitempty"`
	// MentionUserIDs specifies an array of IDs of the users to mention in the
	// message. This property is used only when mention_type is users.
	MentionUserIDs []string `json:"mentioned_user_ids,omitempty"`
	// IsSilent d etermines whether to send a message without updating some of
	// the following channel properties. If set to true, the channel's
	// last_message is updated only for the sender while its unread_message_count
	// remains unchanged for all channel members. Also, a push notification isn't
	// sent to the users receiving the message. If the message is sent to a
	// hidden channel, the channel still remains hidden. (Default: false)
	IsSilent *bool `json:"is_silent,omitempty"`
	// MarkAsRead determines whether to mark the message as read for the sender.
	// If set to false, the sender's unread_count and read_receipt remain
	// unchanged after the message is sent. (Default: true)
	MarkAsRead *bool `json:"mark_as_read,omitempty"`
	// SortedMetaArray specifies an array of JSON objects consisting of
	// key-values items that store additional message information to be used for
	// classification and filtering. Items are saved and returned in the order
	// they've been specified. More details on what can be stored in this field
	// are available here:
	// https://sendbird.com/docs/chat/platform-api/v3/channel/channel-overview#2-manage-channel-information-3-information-types
	SortedMetaArray []MetaArray `json:"sorted_meta_array,omitempty"`
	// CreatedAt specifies the time when the message was sent in Unix
	// milliseconds format.
	CreatedAt int64 `json:"created_at,omitempty"`
	// PollID specifies the unique ID of the poll to be included in a message.
	// To use this property, the polls feature should be turned on in
	// Settings > Chat > Features.
	PollID int `json:"poll_id,omitempty"`
	// IncludePollDetails determines whether to include all properties of a poll
	// resource with a full list of options in the results. To use this property,
	// the polls feature should be turned on in Settings > Chat > Features.
	// If set to false, a selection of poll resource properties consisting of id,
	// title, close_at, created_at, updated_at, status, and message_id are
	// returned.
	// (Default: false)
	IncludePollDetails *bool `json:"include_poll_details,omitempty"`
	// DedupID specifies a unique ID for the message created by another system.
	// In general, this property is used to prevent the same message data from
	// getting inserted when migrating messages from another system to the
	// Sendbird server. If specified, the server performs a duplicate check using
	// the property value.
	DedupID string `json:"dedup_id,omitempty"`
	// ApnsBundleID s pecifies the bundle ID of the client app in order to send a
	// push notification to iOS devices. You can find this in Settings > Chat >
	// Push notifications > Push notification credentials.
	ApnsBundleID string `json:"apns_bundle_id,omitempty"`
	// Sound specifies the name of a sound file that is used for critical alerts.
	Sound string `json:"sound,omitempty"`
	// Volume specifies the volume of the critical alert sound. The volume ranges
	// from 0.0 to 1.0, which indicates silent and full volume, respectively.
	// (Default: 1.0)
	Volume float32 `json:"volume,omitempty"`
}

func (smr *SendMessageRequest) Validate() error {
	switch {
	case smr.MessageType == "":
		return errors.New("message type is required")
	case smr.UserID == "":
		return errors.New("user ID is required")
	}

	switch {
	case smr.MessageType == TypeText && smr.Message == "":
		return errors.New("message is required for text message")
	case smr.MessageType == TypeFile && len(smr.Files) == 0:
		return errors.New("files are required for file message")
	case smr.MessageType == TypeFile:
		for _, file := range smr.Files {
			if file.URL == "" {
				return errors.New("file URL is required")
			}
		}
	}

	return nil
}

// SendMessageResponse is the response to send a message.
type SendMessageResponse Resource

// SendMessage sends a message to a channel.
// See https://sendbird.com/docs/chat/platform-api/v3/message/messaging-basics/send-a-message
func (m *message) SendMessage(ctx context.Context, channelType ChannelType, channelURL string, sendMessageRequest SendMessageRequest) (*SendMessageResponse, error) {
	if err := sendMessageRequest.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate send message request: %w", err)
	}

	path := fmt.Sprintf("/%s/%s/messages", channelType, channelURL)

	smr, err := m.client.Post(ctx, path, sendMessageRequest, &SendMessageResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	sendMessageResponse, ok := smr.(*SendMessageResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to SendMessageResponse: %+v", smr)
	}

	return sendMessageResponse, nil
}
