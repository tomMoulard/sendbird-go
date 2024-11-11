package message

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestListMessages(t *testing.T) {
	t.Parallel()

	url := "/group_channels/url/messages"
	url += "?custom_types=%2A"
	url += "&include=true"
	url += "&include_parent_message_info=true"
	url += "&include_poll_details=true"
	url += "&include_reaction=true"
	url += "&include_reply_type=ALL"
	url += "&include_thread_info=true"
	url += "&including_removed=true"
	url += "&message_id=43"
	url += "&message_ts=42"
	url += "&message_type=MESG"
	url += "&next_limit=45"
	url += "&operator_filter=all"
	url += "&prev_limit=44"
	url += "&reverse=true"
	url += "&sender_id=46"
	url += "&sender_ids=47%2C48"
	url += "&show_subchannel_messages_only=true"
	url += "&user_id=49"
	url += "&with_sorted_meta_array=true"

	listMessagesRequest := ListMessagesRequest{
		MessageTS:                  42,
		MessageID:                  43,
		PrevLimit:                  ptr(44),
		NextLimit:                  ptr(45),
		Include:                    ptr(true),
		Reverse:                    ptr(true),
		SenderID:                   "46",
		SenderIDs:                  []string{"47", "48"},
		OperatorFilter:             OperatorFilterAll,
		MessageType:                TypeText,
		CustomTypes:                []string{"*"},
		IncludingRemoved:           ptr(true),
		IncludeParentMessageInfo:   ptr(true),
		IncludeThreadInfo:          ptr(true),
		IncludeReplyType:           ReplyTypeAll,
		IncludeReaction:            ptr(true),
		IncludePollDetails:         ptr(true),
		WithSortedMetaArray:        ptr(true),
		ShowSubchannelMessagesOnly: ptr(true),
		UserID:                     ptr("49"),
	}

	listMessagesResponse := &ListMessagesResponse{
		Messages: []Resource{{
			MessageID: 69,
		}},
	}

	client := client.NewClientMock(t).
		OnGet(url, nil, &ListMessagesResponse{}).TypedReturns(listMessagesResponse, nil).Once().
		Parent
	message := NewMessage(client)

	cur, err := message.ListMessages(context.Background(), ChannelTypeGroup, "url", listMessagesRequest)
	require.NoError(t, err)
	assert.Equal(t, listMessagesResponse, cur)
}
