package message

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomMoulard/sendbird-go/pkg/client"
)

func TestValidateMMR(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		request   MigrateMessagesRequest
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name: "invalid request - empty messages",
			request: MigrateMessagesRequest{
				Messages: []TextMessage{},
			},
			assertErr: assert.Error,
		},
		{
			name: "invalid request - missing user_id",
			request: MigrateMessagesRequest{
				Messages: []TextMessage{
					{
						MessageType: TypeText,
						Message:     "Hello, World!",
						Timestamp:   1609459200000,
					},
				},
			},
			assertErr: assert.Error,
		},
		{
			name: "invalid request - missing message_type",
			request: MigrateMessagesRequest{
				Messages: []TextMessage{
					{
						UserID:    "42",
						Message:   "Hello, World!",
						Timestamp: 1609459200000,
					},
				},
			},
			assertErr: assert.Error,
		},
		{
			name: "invalid request - missing message",
			request: MigrateMessagesRequest{
				Messages: []TextMessage{
					{
						UserID:      "42",
						MessageType: TypeText,
						Timestamp:   1609459200000,
					},
				},
			},
			assertErr: assert.Error,
		},
		{
			name: "invalid request - missing timestamp",
			request: MigrateMessagesRequest{
				Messages: []TextMessage{
					{
						UserID:      "42",
						MessageType: TypeText,
						Message:     "Hello, World!",
					},
				},
			},
			assertErr: assert.Error,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := test.request.Validate()
			test.assertErr(t, err)
		})
	}
}

func TestMigrateMessages(t *testing.T) {
	t.Parallel()

	validRequest := MigrateMessagesRequest{
		Messages: []TextMessage{
			{
				UserID:      "42",
				MessageType: TypeText,
				Message:     "Hello, World!",
				Timestamp:   1609459200000,
				CustomType:  "greeting",
				Data:        `{ "emotion": "happy" }`,
				DedupID:     "unique123",
			},
		},
	}

	tests := []struct {
		name      string
		request   MigrateMessagesRequest
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name:      "valid request",
			request:   validRequest,
			assertErr: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			client := client.NewClientMock(t).
				OnPost("/migration/channel_url", test.request, nil).TypedReturns(nil, nil).Once().
				Parent
			message := NewMessage(client)

			err := message.MigrateMessages(context.Background(), "channel_url", test.request)
			test.assertErr(t, err)
		})
	}
}
