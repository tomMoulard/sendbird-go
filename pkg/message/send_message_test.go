package message

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func ptr[T any](t T) *T {
	return &t
}

func TestPtr(t *testing.T) {
	t.Parallel()

	assert.True(t, *ptr(true))
	assert.False(t, *ptr(false))
	assert.Zero(t, *ptr(0))
	assert.Equal(t, "foo", *ptr("foo"))
}

func TestValidateSMR(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		smr       SendMessageRequest
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name:      "empty",
			smr:       SendMessageRequest{},
			assertErr: assert.Error,
		},
		{
			name: "without Message Type",
			smr: SendMessageRequest{
				UserID:  "42",
				Message: "Hello, World!",
			},
			assertErr: assert.Error,
		},
		{
			name: "without user id",
			smr: SendMessageRequest{
				MessageType: TypeText,
				Message:     "Hello, World!",
			},
			assertErr: assert.Error,
		},
		{
			name: "text message without message",
			smr: SendMessageRequest{
				MessageType: TypeText,
				UserID:      "42",
			},
			assertErr: assert.Error,
		},
		{
			name: "file message without files",
			smr: SendMessageRequest{
				MessageType: TypeFile,
				UserID:      "42",
			},
			assertErr: assert.Error,
		},
		{
			name: "file message with files, but not url",
			smr: SendMessageRequest{
				MessageType: TypeFile,
				UserID:      "42",
				Files:       []File{{}},
			},
			assertErr: assert.Error,
		},
		{
			name: "valid text message",
			smr: SendMessageRequest{
				MessageType: TypeText,
				UserID:      "42",
				Message:     "Hello, World!",
			},
			assertErr: assert.NoError,
		},
		{
			name: "valid file message",
			smr: SendMessageRequest{
				MessageType: TypeFile,
				UserID:      "42",
				Files: []File{
					{URL: "https://example.com/file1"},
					{URL: "https://example.com/file2"},
				},
			},
			assertErr: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := test.smr.Validate()
			test.assertErr(t, err)
		})
	}
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		smrq SendMessageRequest
		smrp SendMessageResponse
	}{
		{
			name: "Text Message",
			smrq: SendMessageRequest{
				MessageType:         TypeText,
				UserID:              "42",
				Message:             "Hello, World!",
				CustomType:          "custom-type",
				Data:                `{ "key": "value" }`,
				SendPush:            ptr(true),
				PushMessageTemplate: "push-notification-template",
				MentionType:         MentionTypeChannels,
				MentionUserIDs:      []string{"mention-user-id"},
				IsSilent:            ptr(true),
				MarkAsRead:          ptr(true),
				SortedMetaArray: []MetaArray{
					{Key: "key1", Value: []string{"value1"}},
					{Key: "key2", Value: []string{"value2"}},
				},
				CreatedAt:          42,
				PollID:             42,
				IncludePollDetails: ptr(true),
				DedupID:            "dedup-id",
				ApnsBundleID:       "apns-bundle-id",
				Sound:              "sound",
				Volume:             0.5,
			},
			smrp: SendMessageResponse{
				MessageID: 42,
			},
		},
		{
			name: "File Message",
			smrq: SendMessageRequest{
				MessageType: TypeFile,
				UserID:      "42",
				Files: []File{
					{URL: "https://example.com/file1"},
					{URL: "https://example.com/file2"},
				},
			},
			smrp: SendMessageResponse{
				MessageID: 42,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			client := client.NewClientMock(t).
				OnPost("/group_channels/url/messages", test.smrq, &SendMessageResponse{}).TypedReturns(&test.smrp, nil).Once().
				Parent
			message := NewMessage(client)

			cur, err := message.SendMessage(context.Background(), ChannelTypeGroup, "url", test.smrq)
			require.NoError(t, err)
			require.NotNil(t, cur)
			assert.Equal(t, test.smrp, *cur)
		})
	}
}
