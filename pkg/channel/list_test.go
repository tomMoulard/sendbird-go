package channel

import (
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

func TestListGroupChannels(t *testing.T) {
	t.Parallel()

	url := "/group_channels"
	url += "?channel_urls=channel-url1%2Cchannel-url2"
	url += "&created_after=43"
	url += "&created_before=44"
	url += "&custom_type_starts_with=custom-type-starts-with"
	url += "&custom_types=custom-types1%2Ccustom-types2"
	url += "&distinct_mode=all"
	url += "&include_sorted_metaarray_in_last_message=true"
	url += "&limit=42"
	url += "&members_exactly_in=members-exactly-in1%2Cmembers-exactly-in2"
	url += "&members_include_in=members-include-in1%2Cmembers-include-in2"
	url += "&members_nickname=members-nickname"
	url += "&members_nickname_contains=members-nickname-contains"
	url += "&metacounter_key=metacounter-key"
	url += "&metacounter_values=metacounter-value1%2Cmetacounter-value2"
	url += "&metacounter_values_gt=metacounter-values-gt"
	url += "&metacounter_values_gte=metacounter-values-gte"
	url += "&metacounter_values_lt=metacounter-values-lt"
	url += "&metacounter_values_lte=metacounter-values-lte"
	url += "&metadata_key=metadata-key"
	url += "&metadata_order_key=metadata-order-key"
	url += "&metadata_value_startswith=metadata-value-starts-with"
	url += "&metadata_values=metadata-value1%2Cmetadata-value2"
	url += "&name=name"
	url += "&name_contains=name-contains"
	url += "&name_startswith=name-starts-with"
	url += "&order=chronological"
	url += "&public_mode=all"
	url += "&query_type=AND"
	url += "&show_delivery_receipt=true"
	url += "&show_empty=true"
	url += "&show_frozen=true"
	url += "&show_member=true"
	url += "&show_metadata=true"
	url += "&show_read_receipt=true"
	url += "&super_mode=all"
	url += "&token=token"

	listChannelsRequest := ListGroupChannelRequest{
		Token:                               "token",
		Limit:                               ptr(42),
		DistinctMode:                        DistinctModeAll,
		PublicMode:                          PublicModeAll,
		SuperMode:                           SuperModeAll,
		CreatedAfter:                        ptr(43),
		CreatedBefore:                       ptr(44),
		ShowEmpty:                           ptr(true),
		ShowMember:                          ptr(true),
		ShowDeliveryReceipt:                 ptr(true),
		ShowReadReceipt:                     ptr(true),
		ShowMetadata:                        ptr(true),
		ShowFrozen:                          ptr(true),
		Order:                               OrderChronological,
		MetadataOrderKey:                    "metadata-order-key",
		CustomTypes:                         []string{"custom-types1", "custom-types2"},
		CustomTypeStartsWith:                "custom-type-starts-with",
		ChannelURLs:                         []string{"channel-url1", "channel-url2"},
		Name:                                "name",
		NameContains:                        "name-contains",
		NameStartswith:                      "name-starts-with",
		MembersExactlyIn:                    []string{"members-exactly-in1", "members-exactly-in2"},
		MembersIncludeIn:                    []string{"members-include-in1", "members-include-in2"},
		QueryType:                           QueryTypeAnd,
		MembersNickname:                     "members-nickname",
		MembersNicknameContains:             "members-nickname-contains",
		MetadataKey:                         "metadata-key",
		MetadataValues:                      []string{"metadata-value1", "metadata-value2"},
		MetadataValueStartswith:             "metadata-value-starts-with",
		MetacounterKey:                      "metacounter-key",
		MetacounterValues:                   []string{"metacounter-value1", "metacounter-value2"},
		MetacounterValuesGT:                 "metacounter-values-gt",
		MetacounterValuesGTE:                "metacounter-values-gte",
		MetacounterValuesLT:                 "metacounter-values-lt",
		MetacounterValuesLTE:                "metacounter-values-lte",
		IncludeSortedMetaarrayInLastMessage: ptr(true),
	}

	listChannelsResponse := &ListGroupChannelResponse{
		Channels: []Resource{{
			Name: "channel-name",
		}},
	}

	client := client.NewClientMock(t).
		OnGet(url, nil, &ListGroupChannelResponse{}).TypedReturns(listChannelsResponse, nil).Once().
		Parent
	channel := NewChannel(client)

	cur, err := channel.ListGroupChannels(t.Context(), listChannelsRequest)
	require.NoError(t, err)
	assert.Equal(t, listChannelsResponse, cur)
}
