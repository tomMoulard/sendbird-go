package channel

import (
	"context"
	"fmt"
)

// UpdateGroupChannelRequest is the request to update a channel.
type UpdateGroupChannelRequest struct {
	// Name specifies the name of the channel or the channel topic. The length is
	// limited to 191 characters.
	Name string `json:"name,omitempty"`
	// CoverURL specifies the unique URL of the channel's cover image. The length
	// is limited to 2,048 characters.
	CoverURL string `json:"cover_url,omitempty"`
	// CustomType Specifies the custom channel type which is used for channel
	// grouping. The length is limited to 128 characters.
	CustomType string `json:"custom_type,omitempty"`
	// Data specifies additional channel information such as a long description
	// of the channel or JSON formatted string.
	Data string `json:"data,omitempty"`
	// IsDistinct determines whether to reuse an existing channel or create a new
	// channel when creating a channel with the same group of members. If set to
	// true, returns an existing channel with the same members or creates a new
	// channel if no such channel exists. The Sendbird server can also use the
	// custom channel type in the custom_type property if specified along with
	// the users. If set to false, the Sendbird server always creates a new
	// channel with a combination of the users as well as the channel custom type
	// if specified. (Default: false)
	IsDistinct bool `json:"is_distinct,omitempty"`
	// IsPublic determines whether to allow a user to join the channel without an
	// invitation. (Default: false)
	IsPublic bool `json:"is_public,omitempty"`
	// IsSuper determines whether to allow the channel to accommodate 100 or more
	// members. If set to true, creates a Supergroup channel. (Default: false)
	IsSuper bool `json:"is_super,omitempty"`
	// Acchis parameter can only be used when the channel operator creates a
	// public group channel. If the channel operator uses an access code for the
	// corresponding type of channel, users will be required to enter the access
	// code to join the channel. If specified, the is_access_code_required
	// property of the channel resource is automatically set to true.
	AccessCode string `json:"access_code,omitempty"`
	// OperatorIDs specifies an array of one or more IDs of users to register as
	// operators of the channel. Users can be registered as operators regardless
	// of whether they are channel members. The maximum allowed number of
	// operators per channel is 100.
	OperatorIDs []string `json:"operator_ids,omitempty"`
}

// UpdateGroupChannelResponse is the response of the update channel request.
type UpdateGroupChannelResponse Resource

// UpdateGroupChannel updates a group channel.
// See https://sendbird.com/docs/chat/platform-api/v3/channel/managing-a-channel/update-a-group-channel
func (c *channel) UpdateGroupChannel(ctx context.Context, channelURL string, updateChannelRequest UpdateGroupChannelRequest) (*UpdateGroupChannelResponse, error) {
	ccr, err := c.client.Put(ctx, "/group_channels/"+channelURL, updateChannelRequest, &UpdateGroupChannelResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to update channel: %w", err)
	}

	updateChannelResponse, ok := ccr.(*UpdateGroupChannelResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to UpdateGroupChannelResponse: %+v", ccr)
	}

	return updateChannelResponse, nil
}
