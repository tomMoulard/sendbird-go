package channel

import (
	"context"
	"fmt"
)

// CreateGroupChannelRequest is the request to create a channel.
type CreateGroupChannelRequest struct {
	// UserIDs specifies an array of one or more IDs of users to invite to the
	// channel. Up to 100 users can be invited at a time.
	UserIDs []string `json:"user_ids"`
	// Name specifies the name of the channel, or the channel topic.
	// The length is limited to 191 characters.
	// (Default: "group channel")
	Name string `json:"name,omitempty"`
	// ChannelURL specifies the URL of the channel. Only numbers, letters,
	// underscores, and hyphens are allowed. The allowed length is 4 to 100
	// characters, inclusive. If not specified, a URL is automatically generated.
	ChannelURL string `json:"channel_url,omitempty"`
	// CoverURL specifies the URL of the channel's cover image.
	// This should be no longer than 2,048 characters.
	CoverURL string `json:"cover_url,omitempty"`
	// CustomType specifies a custom channel type which is used for channel
	// grouping. Maximum length is 128 characters.
	CustomType string `json:"custom_type,omitempty"`
	// Data additional channel information such as a long description of the
	// channel or JSON formatted string. More detail on what can be stored in the
	// data field is available here:
	// https://sendbird.com/docs/chat/platform-api/v3/channel/channel-overview#2-manage-channel-information-3-information-types
	Data string `json:"data,omitempty"`
	// IsDistinct determines whether to reuse an existing channel or create a new
	// channel when attempting to create a channel with the same group of
	// members. By default, a new channel is created even if one already exists
	// for the same members. You can override this and force an existing channel
	// to be returned by setting a value of is_distinct to true.
	// You can also reuse channels if the new channel has the same members, and
	// the same custom_type, or create an entirely new channel if there are the
	// same members but had a different custom_type. (Default: false)
	IsDistinct bool `json:"is_distinct,omitempty"`
	// IsPublic determines whether to allow users to join the channel without an
	// invitation. By default, a user must be invited to join a group channel. It
	// is possible to allow any user to join a group channel without an invite by
	// setting is_public to true.
	// Public group channels are similar to open channels but include extra
	// features like typing indicators and are limited to 100 members.
	// (Default: false)
	IsPublic bool `json:"is_public,omitempty"`
	// IsSuper determines whether to allow the channel to accommodate 100 or more
	// members. By default, a group channel has a member limit of 100. A
	// Supergroup channel must be used in order to have more than 100 channel
	// members. A Supergroup channel has a member limit of 2000.
	// A group channel can be created as a Supergroup channel by setting is_super
	// to true.
	// The is_distinct property must be false in order to create a Supergroup.
	// (Default: false)
	IsSuper bool `json:"is_super,omitempty"`
	// IsEphemeral determines whether to preserve messages in the channel for the
	// purpose of retrieving chat history. By default, all chat messages will be
	// stored by Sendbird. This means that chat histories can easily be retrieved
	// when using the SDK or API.
	// If is_ephemeral is set to true then no chat history is stored. In this
	// case, the chat history can't be retrieved in the future.
	// This option may be helpful for exceptional data privacy requirements.
	// (Default: false)
	IsEphemeral bool `json:"is_ephemeral,omitempty"`
	// AccessCode specifies an access code that is only applicable to public
	// group channels. If a value is specified for access_code, users will be
	// required to enter this code when joining a group channel.
	// This parameter can only be used when the channel operator creates a public
	// group channel. If specified, the is_access_code_required property of the
	// channel resource is automatically set to true.
	AccessCode string `json:"access_code,omitempty"`
	// InviterID specifies the ID of a user who invites other users to the
	// channel. The inviter isn't automatically registered to the channel as a
	// member, so you should specify the ID of the inviter in the user_ids
	// property if needed.
	InviterID string `json:"inviter_id,omitempty"`
	// Strict determines whether to receive a 400111 error and cease channel
	// creation when there is at least one user that doesn't exist in the
	// specified user_ids or users property. Channel creation by default won't
	// fail if a user passed to user_ids or users doesn't exist in the Sendbird
	// application.
	// Setting strict to true causes a 400111 error during creation if any users
	// passed to user_ids or users haven't already been successfully created.
	// (Default: false)
	Strict bool `json:"strict,omitempty"`
	// OperatorIDs specifies an array of one or more IDs of users to register as
	// operators of the channel. You should also include these IDs in the
	// user_ids property to invite them to the channel as members. They can
	// delete any messages in the channel, and also view all messages without any
	// filtering or throttling. A channel may have up to 100 operators.
	OperatorIDs []string `json:"operator_ids,omitempty"`
	// BlocSDKUserChannelJoin determines whether to block users from joining the
	// channel through the Chat SDK. If set to true, users can only join the
	// channel using the Sendbird Platform API join a channel action. (Default:
	// false)
	BlocSDKUserChannelJoin bool `json:"bloc_sdk_user_channel_join,omitempty"`
}

// CreateGroupChannelResponse is the response of the create channel request.
type CreateGroupChannelResponse Resource

// CreateGroupChannel creates a group/super channel.
// https://sendbird.com/docs/chat/platform-api/v3/channel/creating-a-channel/create-a-group-channel
func (c *channel) CreateGroupChannel(ctx context.Context, createChannelRequest CreateGroupChannelRequest) (*CreateGroupChannelResponse, error) {
	cgcr, err := c.client.Post(ctx, "/group_channels", createChannelRequest, &CreateGroupChannelResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}

	createChannelResponse, ok := cgcr.(*CreateGroupChannelResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to CreateChannelResponse: %+v", cgcr)
	}

	return createChannelResponse, nil
}
