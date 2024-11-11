package message

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	strconvSlice "github.com/tomMoulard/sendbird-go/pkg/utils/strconv"
)

// ListMessagesRequest is the request to list messages.
type ListMessagesRequest struct {
	// MessageTS s pecifies the timestamp to be the reference point of the query
	// in Unix milliseconds. Either this or the message_id parameter below should
	// be specified in your query URL to retrieve a list. It fetches messages
	// that were sent prior to and after the specified message_ts and the default
	// value for both prev_limit and next_limit is 15.
	MessageTS int64
	// MessageID specifies the unique ID of the message to be the reference point
	// of the query. Either this or the message_ts parameter above should be
	// specified in your query URL to retrieve a list. It fetches messages that
	// were sent prior to and after the specified message_id and the default
	// value for both prev_limit and next_limit is 15.
	MessageID int

	// PrevLimit specifies the number of previously sent messages to retrieve
	// before message_ts.
	// Optional. Acceptable values range from 0 to 200. (Default: 15)
	PrevLimit *int
	// NextLimit specifies the number of sent messages to retrieve after
	// message_ts.
	// Optional. Acceptable values range from 0 to 200. (Default: 15)
	NextLimit *int
	// Include determines whether to include messages sent exactly on the
	// specified message_ts in the results.
	// Optional. (Default: true)
	Include *bool
	// Reverse determines whether to sort the results in reverse chronological
	// order. If set to true, messages appear in reverse chronological order
	// where the newest comes first and the oldest last.
	// Optional. (Default: false)
	Reverse *bool
	// SenderID restricts the search scope to only retrieve messages sent by the
	// user with the specified ID.
	// Optional.
	SenderID string
	// SenderIDs restricts the search scope to only retrieve messages sent by one
	// or more users with the specified IDs.
	// Optional.
	SenderIDs []string
	// OperatorFilter sestricts the search scope to only retrieve messages sent
	// by operators or non-operator users of the channel
	// Optional. (Default: OperatorFilterAll)
	OperatorFilter OperatorFilter
	// MessageType specifies a message type to retrieve. If not specified, all
	// messages are retrieved.
	// Optional.
	MessageType Type
	// CustomTypes specifies a list of one or more custom message types to
	// retrieve. The value set to this parameter can serve as a filter as
	// follows:
	// - A string of specific custom types: Messages with the corresponding
	// custom types are returned.
	// - Empty: Messages whose custom_type property is
	// empty or has a value of null are returned.
	// - An asterisk (`*`): All messages are returned regardless of their
	// custom_type.
	// Optional. (default '*')
	CustomTypes []string
	// IncludingRemoved determines whether to include messages removed from the
	// channel in the results.
	// Optional. (Default: false)
	IncludingRemoved *bool
	// IncludeParentMessageInfo determines whether to include information of the
	// parent message in the results.
	// Optional. (Default: false)
	IncludeParentMessageInfo *bool
	// IncludeThreadInfo determines whether to include thread information in the
	// results.
	// Optional. (Default: false)
	IncludeThreadInfo *bool
	// IncludeReplyType specifies the type of message to include in the results.
	// - ReplyTypeNone (default): All messages that are not replies. These
	// messages may or may not have replies in its thread.
	// - ReplyTypeAll: All messages including threaded and non-threaded parent
	// messages as well as its replies.
	// - ReplyTypeOnlyReplyToChannel: Messages that are not threaded. Only the
	// parent messages and replies that were sent to the channel are included.
	// Optional
	IncludeReplyType ReplyType
	// IncludeReaction determines whether to include reactions added to messages
	// in the channel in the results.
	// Optional. (Default: false)
	IncludeReaction *bool
	// IncludePollDetails determines whether to include all properties of a poll
	// resource with a full list of options in the results. If set to false, a
	// selection of poll resource properties consisting of id, title, close_at,
	// created_at, updated_at, status, and message_id are returned.
	// Optional. (Default: false)
	IncludePollDetails *bool
	// WithSortedMetaArray determines whether to include the sorted_metaarray
	// property in the response.
	// Optional. (Default: false)
	WithSortedMetaArray *bool
	// ShowSubchannelMessagesOnly determines whether to retrieve messages from a
	// subchannel where the user specified in the user_id parameter is currently
	// in. The show_subchannel_messages_only parameter takes effect only when the
	// channel_type parameter is set to open_channels, the subchannel's
	// subchannel_messages_lifetime property hasn't expired, and the value of the
	// max_recent_messages_count property is specified with the maximum number of
	// messages to be retrieved. If set to false, messages from all subchannels
	// are retrieved. This should be specified in conjunction with user_id below.
	// Optional. (Default: false)
	ShowSubchannelMessagesOnly *bool
	// UserID specifies the unique ID of a user to restrict the scope to the
	// user's subchannel. Messages in the user's subchannel are retrieved. This
	// parameter should be specified in conjunction with the
	// show_subchannel_messages_only parameter above.
	// Optional.
	UserID *string
}

// ListMessagesResponse is the response to list messages.
type ListMessagesResponse struct {
	Messages []Resource `json:"messages"`
}

func listMessagesRequestToMap(lmr ListMessagesRequest) map[string]string {
	m := make(map[string]string)

	m["message_ts"] = strconv.FormatInt(lmr.MessageTS, 10)
	m["message_id"] = strconv.Itoa(lmr.MessageID)

	// Optional fields
	if lmr.PrevLimit != nil {
		m["prev_limit"] = strconv.Itoa(*lmr.PrevLimit)
	}

	if lmr.NextLimit != nil {
		m["next_limit"] = strconv.Itoa(*lmr.NextLimit)
	}

	if lmr.Include != nil {
		m["include"] = strconv.FormatBool(*lmr.Include)
	}

	if lmr.Reverse != nil {
		m["reverse"] = strconv.FormatBool(*lmr.Reverse)
	}

	if lmr.SenderID != "" {
		m["sender_id"] = lmr.SenderID
	}

	if len(lmr.SenderIDs) > 0 {
		m["sender_ids"] = strconvSlice.FormatSliceToCSV(lmr.SenderIDs)
	}

	if lmr.OperatorFilter != "" {
		m["operator_filter"] = string(lmr.OperatorFilter)
	}

	if lmr.MessageType != "" {
		m["message_type"] = string(lmr.MessageType)
	}

	if len(lmr.CustomTypes) > 0 {
		m["custom_types"] = strconvSlice.FormatSliceToCSV(lmr.CustomTypes)
	}

	if lmr.IncludingRemoved != nil {
		m["including_removed"] = strconv.FormatBool(*lmr.IncludingRemoved)
	}

	if lmr.IncludeParentMessageInfo != nil {
		m["include_parent_message_info"] = strconv.FormatBool(*lmr.IncludeParentMessageInfo)
	}

	if lmr.IncludeThreadInfo != nil {
		m["include_thread_info"] = strconv.FormatBool(*lmr.IncludeThreadInfo)
	}

	if lmr.IncludeReplyType != "" {
		m["include_reply_type"] = string(lmr.IncludeReplyType)
	}

	if lmr.IncludeReaction != nil {
		m["include_reaction"] = strconv.FormatBool(*lmr.IncludeReaction)
	}

	if lmr.IncludePollDetails != nil {
		m["include_poll_details"] = strconv.FormatBool(*lmr.IncludePollDetails)
	}

	if lmr.WithSortedMetaArray != nil {
		m["with_sorted_meta_array"] = strconv.FormatBool(*lmr.WithSortedMetaArray)
	}

	if lmr.ShowSubchannelMessagesOnly != nil {
		m["show_subchannel_messages_only"] = strconv.FormatBool(*lmr.ShowSubchannelMessagesOnly)
	}

	if lmr.UserID != nil {
		m["user_id"] = *lmr.UserID
	}

	return m
}

// ListMessages retrieves a list of messages in a channel.
// See https://sendbird.com/docs/chat/platform-api/v3/message/messaging-basics/list-messages
func (m *message) ListMessages(ctx context.Context, channelType ChannelType, channelURL string, listMessagesRequest ListMessagesRequest) (*ListMessagesResponse, error) {
	u, err := url.Parse(fmt.Sprintf("/%s/%s/messages", channelType, channelURL))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	query := u.Query()
	for k, v := range listMessagesRequestToMap(listMessagesRequest) {
		query.Set(k, v)
	}

	u.RawQuery = query.Encode()

	lmr, err := m.client.Get(ctx, u.String(), nil, &ListMessagesResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	listMessagesResponse, ok := lmr.(*ListMessagesResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to ListMessagesResponse: %+v", lmr)
	}

	return listMessagesResponse, nil
}
