package channel

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	strconvSlice "github.com/tomMoulard/sendbird-go/pkg/utils/strconv"
)

// ListGroupChannelRequest is the request to list a channel.
type ListGroupChannelRequest struct {
	// Token specifies a page token that indicates the starting index of a chunk
	// of results. If not specified, the index is set as 0.
	// Optional.
	Token string
	// Limit specifies the number of results to return per page. Acceptable
	// values are 1 to 100, inclusive. (Default: 10)
	// Optional.
	Limit *int
	// DistinctMode restricts the search scope to only retrieve distinct or
	// nondistinct group channels. Acceptable values are the following:
	// - DistinctModeAll (default): All group channels are returned.
	// - DistinctModeDistinct: Only distinct group channels are returned.
	// - DistinctModeNonDistinct: Only group channels that aren't distinct are
	// returned.
	// Optional.
	DistinctMode DistinctMode
	// PublicMode restricts the search scope to only retrieve either public or
	// private group channels. Acceptable values are the following:
	// - PublicModeAll (default): All group channels are returned.
	// - PublicModeNonPublic: All private group channels are returned.
	// - PublicModePublic: All public group channels are returned.
	// Optional.
	PublicMode PublicMode
	// SuperMod&e specifies which type of group channels to retrieve. Acceptable
	// values are the following:
	// - SuperModeAll (default): All types of group channels including Supergroup
	// channels are returned.
	// - SuperModeSuper: Only Supergroup channels are returned.
	// - SuperModeNonSuper: Group channels excluding Supergroup channels are
	// returned.
	// Optional.
	SuperMode SuperMode
	// CreatedAfter restricts the search scope to only retrieve group channels
	// which have been created after the specified time, in Unix milliseconds
	// format.
	// Optional.
	CreatedAfter *int
	// CreatedBefore restricts the search scope to only retrieve group channels
	// which have been created before the specified time, in Unix milliseconds
	// format.
	// Optional.
	CreatedBefore *int
	// ShowEmpty determines whether to include empty channels in the response.
	// Empty channels are channels that have been created but that don't have any
	// messages. If set to true, empty channels are included in the response.
	// (Default: false)
	// Optional.
	ShowEmpty *bool
	// ShowMember determines whether to include information about the members of
	// each channel in the response. (Default: false)
	// Optional.
	ShowMember *bool
	// ShowDeliveryReceipt determines whether to include information about the
	// delivery receipts of each channel in the response. The delivery receipt
	// indicates the timestamp of when each member has last received messages
	// from the Sendbird server in the channel, in Unix milliseconds.
	// (Default: false)
	// Optional.
	ShowDeliveryReceipt *bool
	// ShowReadReceipt determines whether to include information about the read
	// receipts of each channel in the response. The read receipt indicates the
	// timestamp of when each member has last read the messages in the channel,
	// in Unix milliseconds. (Default: false)
	// Optional.
	ShowReadReceipt *bool
	// ShowMetadata determines whether to include channel metadata in the
	// response. (Default: false)
	// Optional.
	ShowMetadata *bool
	// ShowFrozen determines whether to include frozen channels in the response.
	// Frozen channels are channels where only operators are allowed to send
	// messages. (Default: true)
	// Optional.
	ShowFrozen *bool
	// Order specifies the method to sort a list of results. Acceptable values
	// are the following:
	// - OrderChronological (default): sorts by time of channel creation, from
	// most to least recent.
	// - OrderLatestLastMessage: sorts by the time of the last message in the
	// channel, from most to least recent. This is available only when user_id is
	// specified.
	// - OrderChannelNameAlphabetical: sorts by channel name in alphabetical
	// order.
	// - OrderMetadataValueAlphabetical: sorts by a value of metadata in
	// alphabetical order. This is available only when the metadata_order_key
	// parameter is specified.
	// Optional.
	Order Order
	// MetadataOrderKey specifies the key of an item in metadata. When a value of
	// the order parameter is set to metadata_value_alphabetical, the results are
	// alphabetically sorted by the value of the item specified by the key.
	// Optional.
	MetadataOrderKey string
	// CustomTypes specifies a list of one or more custom types to filter group
	// channels. URL encoding each type is recommended. If not specified, all
	// channels are returned, regardless of their custom type.
	// Optional.
	CustomTypes []string
	// CustomTypeStartsWith dearches for group channels with the custom type
	// which starts with the specified value.
	// Optional.
	CustomTypeStartsWith string
	// ChannelURLs specifies a list of one or more group channel URLs to restrict
	// the search scope.
	// Optional.
	ChannelURLs []string
	// Name specifies one or more group channel names.
	// Optional.
	Name string
	// NameContains searches for group channels whose names contain the specified
	// value. Note that this parameter is case-insensitive.
	// Optional.
	NameContains string
	// NameStartswith searches for group channels whose names start with the
	// specified value. Note that this parameter is case-insensitive.
	// Optional.
	NameStartswith string
	// MembersExactlyIn searches for group channels with all the specified users
	// as members.
	// Only user IDs that match those of existing users are used for channel
	// search.
	// Optional.
	MembersExactlyIn []string
	// MembersIncludeIn searches for group channels that include one or more
	// users as members among the specified users. You can specify up to 60 user
	// IDs. Only user IDs that match those of existing users are used for channel
	// search.
	// Optional.
	MembersIncludeIn []string
	// QueryType specifies a logical condition applied to the members_include_in
	// parameter. Acceptable values are either QueryTypeAnd or QueryTypeOr. For
	// example, if you specify three members, A, B, and C, in members_include_in,
	// the value of AND returns all channels that include every one of {A. B, C}
	// as members. The value of OR returns channels that include {A}, plus those
	// that include {B}, plus those that include {C}. (Default: AND)
	// Optional.
	QueryType QueryType
	// MembersNickname searches for group channels with members whose nicknames
	// match the specified value.
	// Optional.
	MembersNickname string
	// MembersNicknameContains sSearches for group channels with members whose
	// nicknames contain the specified value. Note that this parameter is
	// case-insensitive.
	// We recommend using at least three characters for the parameter value for
	// better search efficiency when you design and implement related features.
	// If you would like to allow one or two characters for searching, use
	// members_nickname instead to prevent performance issues.
	// Optional.
	MembersNicknameContains string
	// MetadataKey searches for group channels with metadata containing an item
	// with the specified value as its key. To use this parameter, either the
	// metadata_values parameter or the metadata_value_startswith parameter
	// should be specified.
	// Optional.
	MetadataKey string
	// MetadataValues searches for group channels with metadata containing an
	// item with the key specified by the metadata_key parameter, and the value
	// of that item matches one or more values specified by this parameter.
	// To use this parameter, the metadata_key parameter should be specified.
	// Optional.
	MetadataValues []string
	// MetadataValueStartswith searches for group channels with metadata
	// containing an item with the key specified by the metadata_key parameter,
	// and the values of that item that start with the specified value of this
	// parameter.
	// To use this parameter, the metadata_key parameter should be specified.
	// Optional.
	MetadataValueStartswith string
	// MetacounterKey searches for group channels with metacounter containing an
	// item with the specified value as its key. To use this parameter, either
	// the metacounter_values parameter or one of the metacounter_value_gt,
	// metacounter_value_gte, metacounter_value_lt, and metacounter_value_lte
	// parameters should be specified.
	// Optional.
	MetacounterKey string
	// MetacounterValues searches for group channels with metacounter containing
	// an item with the key specified by the metadata_key parameter, where the
	// value of that item is equal to one or more values specified by this
	// parameter. To use this parameter, the metacounter_key parameter should be
	// specified.
	// Optional.
	MetacounterValues []string
	// MetacounterValuesGT searches for group channels with metacounter
	// containing an item with the key specified by the metadata_key parameter,
	// where the value of that item is greater than the value specified by this
	// parameter. To use this parameter, the metacounter_key parameter should be
	// specified.
	// Optional.
	MetacounterValuesGT string
	// MetacounterValuesGTE searches for group channels with metacounter
	// containing an item with the key specified by the metadata_key parameter,
	// where the value of that item is greater than or equal to the value
	// specified by this parameter. To use this parameter, the metacounter_key
	// parameter should be specified.
	// Optional.
	MetacounterValuesGTE string
	// MetacounterValuesLT searches for group channels with metacounter
	// containing an item with the key specified by the metadata_key parameter,
	// where the value of that item is lower than the value specified by this
	// parameter. To use this parameter, the metacounter_key parameter should be
	// specified.
	// Optional.
	MetacounterValuesLT string
	// MetacounterValuesLTE searches for group channels with metacounter
	// containing an item with the key specified by the metadata_key parameter,
	// where the value of that item is lower than or equal to the value specified
	// by this parameter. To use this parameter, the metacounter_key parameter
	// should be specified.
	// Optional.
	MetacounterValuesLTE string
	// IncludeSortedMetaarrayInLastMessage determines whether to include the
	// sorted_metaarray as one of the last_message’s properties in the response.
	// Optional.
	IncludeSortedMetaarrayInLastMessage *bool
}

// ListGroupChannelResponse is the response of the list channel request.
type ListGroupChannelResponse struct {
	// Channels is the list of group channel objects that match the specified
	// optional parameters.
	Channels []Resource `json:"channels"`
	// Next is the value for the token parameter to retrieve the next page in the
	// result.
	Next string `json:"next"`
}

func listChannelRequestToMap(lcr ListGroupChannelRequest) map[string]string {
	m := make(map[string]string)

	if lcr.Token != "" {
		m["token"] = lcr.Token
	}

	if lcr.Limit != nil {
		m["limit"] = strconv.Itoa(*lcr.Limit)
	}

	if lcr.DistinctMode != "" {
		m["distinct_mode"] = string(lcr.DistinctMode)
	}

	if lcr.PublicMode != "" {
		m["public_mode"] = string(lcr.PublicMode)
	}

	if lcr.SuperMode != "" {
		m["super_mode"] = string(lcr.SuperMode)
	}

	if lcr.CreatedAfter != nil {
		m["created_after"] = strconv.Itoa(*lcr.CreatedAfter)
	}

	if lcr.CreatedBefore != nil {
		m["created_before"] = strconv.Itoa(*lcr.CreatedBefore)
	}

	if lcr.ShowEmpty != nil {
		m["show_empty"] = strconv.FormatBool(*lcr.ShowEmpty)
	}

	if lcr.ShowMember != nil {
		m["show_member"] = strconv.FormatBool(*lcr.ShowMember)
	}

	if lcr.ShowDeliveryReceipt != nil {
		m["show_delivery_receipt"] = strconv.FormatBool(*lcr.ShowDeliveryReceipt)
	}

	if lcr.ShowReadReceipt != nil {
		m["show_read_receipt"] = strconv.FormatBool(*lcr.ShowReadReceipt)
	}

	if lcr.ShowMetadata != nil {
		m["show_metadata"] = strconv.FormatBool(*lcr.ShowMetadata)
	}

	if lcr.ShowFrozen != nil {
		m["show_frozen"] = strconv.FormatBool(*lcr.ShowFrozen)
	}

	if lcr.Order != "" {
		m["order"] = string(lcr.Order)
	}

	if lcr.MetadataOrderKey != "" {
		m["metadata_order_key"] = lcr.MetadataOrderKey
	}

	if len(lcr.CustomTypes) > 0 {
		m["custom_types"] = strconvSlice.FormatSliceToCSV(lcr.CustomTypes)
	}

	if lcr.CustomTypeStartsWith != "" {
		m["custom_type_starts_with"] = lcr.CustomTypeStartsWith
	}

	if len(lcr.ChannelURLs) > 0 {
		m["channel_urls"] = strconvSlice.FormatSliceToCSV(lcr.ChannelURLs)
	}

	if lcr.Name != "" {
		m["name"] = lcr.Name
	}

	if lcr.NameContains != "" {
		m["name_contains"] = lcr.NameContains
	}

	if lcr.NameStartswith != "" {
		m["name_startswith"] = lcr.NameStartswith
	}

	if len(lcr.MembersExactlyIn) > 0 {
		m["members_exactly_in"] = strconvSlice.FormatSliceToCSV(lcr.MembersExactlyIn)
	}

	if len(lcr.MembersIncludeIn) > 0 {
		m["members_include_in"] = strconvSlice.FormatSliceToCSV(lcr.MembersIncludeIn)
	}

	if lcr.QueryType != "" {
		m["query_type"] = string(lcr.QueryType)
	}

	if lcr.MembersNickname != "" {
		m["members_nickname"] = lcr.MembersNickname
	}

	if lcr.MembersNicknameContains != "" {
		m["members_nickname_contains"] = lcr.MembersNicknameContains
	}

	if lcr.MetadataKey != "" {
		m["metadata_key"] = lcr.MetadataKey
	}

	if len(lcr.MetadataValues) > 0 {
		m["metadata_values"] = strconvSlice.FormatSliceToCSV(lcr.MetadataValues)
	}

	if lcr.MetadataValueStartswith != "" {
		m["metadata_value_startswith"] = lcr.MetadataValueStartswith
	}

	if lcr.MetacounterKey != "" {
		m["metacounter_key"] = lcr.MetacounterKey
	}

	if len(lcr.MetacounterValues) > 0 {
		m["metacounter_values"] = strconvSlice.FormatSliceToCSV(lcr.MetacounterValues)
	}

	if lcr.MetacounterValuesGT != "" {
		m["metacounter_values_gt"] = lcr.MetacounterValuesGT
	}

	if lcr.MetacounterValuesGTE != "" {
		m["metacounter_values_gte"] = lcr.MetacounterValuesGTE
	}

	if lcr.MetacounterValuesLT != "" {
		m["metacounter_values_lt"] = lcr.MetacounterValuesLT
	}

	if lcr.MetacounterValuesLTE != "" {
		m["metacounter_values_lte"] = lcr.MetacounterValuesLTE
	}

	if lcr.IncludeSortedMetaarrayInLastMessage != nil {
		m["include_sorted_metaarray_in_last_message"] = strconv.FormatBool(*lcr.IncludeSortedMetaarrayInLastMessage)
	}

	return m
}

// ListGroupChannels lists group channels.
// See https://sendbird.com/docs/chat/platform-api/v3/channel/listing-channels-in-an-application/list-group-channels
func (c *channel) ListGroupChannels(ctx context.Context, listChannelRequest ListGroupChannelRequest) (*ListGroupChannelResponse, error) {
	u := &url.URL{
		Path: "/group_channels",
	}

	query := u.Query()
	for k, v := range listChannelRequestToMap(listChannelRequest) {
		query.Set(k, v)
	}

	u.RawQuery = query.Encode()

	lgcr, err := c.client.Get(ctx, u.String(), nil, &ListGroupChannelResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to list channel: %w", err)
	}

	listChannelResponse, ok := lgcr.(*ListGroupChannelResponse)
	if !ok {
		return nil, fmt.Errorf("failed to cast body to ListGroupChannelResponse: %+v", lgcr)
	}

	return listChannelResponse, nil
}
