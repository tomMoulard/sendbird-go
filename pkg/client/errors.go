package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// https://sendbird.com/docs/chat/platform-api/v3/error-codes

var (
	ErrAPIDefault = errors.New("API error")

	ErrAPIBadRequest                          = errors.New("bad request")
	ErrUnexpectedParameterType                = fmt.Errorf("%w: the request specifies one or more parameters in an unexpected data type", ErrAPIBadRequest)
	ErrUnexpectedParameterTypeString          = fmt.Errorf("%w: The data type of the parameters should be string", ErrUnexpectedParameterType)                                                         // "400100",
	ErrUnexpectedParameterTypeNumber          = fmt.Errorf("%w: The data type of the parameters should be number", ErrUnexpectedParameterType)                                                         // "400101",
	ErrUnexpectedParameterTypeList            = fmt.Errorf("%w: The data type of the parameters should be list", ErrUnexpectedParameterType)                                                           // "400102",
	ErrUnexpectedParameterTypeJSON            = fmt.Errorf("%w: The data type of the parameters should be JSON", ErrUnexpectedParameterType)                                                           // "400103",
	ErrUnexpectedParameterTypeBoolean         = fmt.Errorf("%w: The data type of the parameters should be boolean", ErrUnexpectedParameterType)                                                        // "400104",
	ErrMissingRequiredParameters              = fmt.Errorf("%w: the request is missing one or more required parameters", ErrAPIBadRequest)                                                             // "400105",
	ErrNegativeNumberNotAllowed               = fmt.Errorf("%w: the request specifies a negative number. It should specify a positive number", ErrAPIBadRequest)                                       // "400106",
	ErrUnauthorizedRequest                    = fmt.Errorf("%w: the request isn't authorized and can't access the requested resource", ErrAPIBadRequest)                                               // "400108",
	ErrParameterValueLengthExceeded           = fmt.Errorf("%w: the length of the parameter value is too long", ErrAPIBadRequest)                                                                      // "400110",
	ErrInvalidValue                           = fmt.Errorf("%w: the request specifies an invalid value", ErrAPIBadRequest)                                                                             // "400111",
	ErrIncompatibleValues                     = fmt.Errorf("%w: two parameters of the request, which should have unique values, specify the same value", ErrAPIBadRequest)                             // "400112",
	ErrParameterValueOutOfRange               = fmt.Errorf("%w: the request specifies one or more parameters outside the accepted value range", ErrAPIBadRequest)                                      // "400113",
	ErrInvalidURLOfResource                   = fmt.Errorf("%w: the resource identified with the URL in the request can't be found", ErrAPIBadRequest)                                                 // "400114",
	ErrNotAllowedCharacter                    = fmt.Errorf("%w: the request specifies an unacceptable value containing special character, empty string, or white space", ErrAPIBadRequest)             // "400151",
	ErrResourceNotFound                       = fmt.Errorf("%w: the resource identified with the request's resourceId parameter can't be found", ErrAPIBadRequest)                                     // "400201",
	ErrResourceAlreadyExists                  = fmt.Errorf("%w: the resource identified with the request's resourceId parameter already exists", ErrAPIBadRequest)                                     // "400202",
	ErrTooManyItemsInParameter                = fmt.Errorf("%w: the parameter specifies more items than allowed", ErrAPIBadRequest)                                                                    // "400203",
	ErrDeactivatedUserNotAccessible           = fmt.Errorf("%w: the request can't retrieve the deactivated user data", ErrAPIBadRequest)                                                               // "400300",
	ErrUserNotFound                           = fmt.Errorf("%w: the user identified with the request's userId parameter can't be found", ErrAPIBadRequest)                                             // "400301",
	ErrInvalidAccessToken                     = fmt.Errorf("%w: the access token provided for the request specifies an invalid value", ErrAPIBadRequest)                                               // "400302",
	ErrInvalidSessionKeyValue                 = fmt.Errorf("%w: the session key provided for the request specifies an invalid value", ErrAPIBadRequest)                                                // "400303",
	ErrApplicationNotFound                    = fmt.Errorf("%w: the application identified with the request can't be found", ErrAPIBadRequest)                                                         // "400304",
	ErrUserIDLengthExceeded                   = fmt.Errorf("%w: the length of the userId parameter value is too long", ErrAPIBadRequest)                                                               // "400305",
	ErrPaidQuotaExceeded                      = fmt.Errorf("%w: the request can't be completed because you have exceeded your plan's paid quota", ErrAPIBadRequest)                                    // "400306",
	ErrDomainNotAllowed                       = fmt.Errorf("%w: the request can't be completed because it came from a restricted domain", ErrAPIBadRequest)                                            // "400307",
	ErrInvalidAPIToken                        = fmt.Errorf("%w: the API token provided for the request specifies an invalid value", ErrAPIBadRequest)                                                  // "400401",
	ErrMissingSomeParameters                  = fmt.Errorf("%w: the request is missing one or more necessary parameters", ErrAPIBadRequest)                                                            // "400402",
	ErrInvalidJSONRequestBody                 = fmt.Errorf("%w: the request body is an invalid JSON", ErrAPIBadRequest)                                                                                // "400403",
	ErrInvalidRequestURL                      = fmt.Errorf("%w: the request specifies an invalid HTTP request URL that can't be accessed", ErrAPIBadRequest)                                           // "400404",
	ErrTooManyUserWebsocketConnections        = fmt.Errorf("%w: the number of the user's websocket connections exceeds the allowed amount", ErrAPIBadRequest)                                          // "400500",
	ErrTooManyApplicationWebsocketConnections = fmt.Errorf("%w: the number of the application's websocket connections exceeds the allowed amount", ErrAPIBadRequest)                                   // "400501",
	ErrBlockedUserSendNotAllowed              = fmt.Errorf("%w: the request can't be completed due to being blocked by the recipient or deactivated", ErrAPIBadRequest)                                // "400700",
	ErrBlockedUserInvitedNotAllowed           = fmt.Errorf("%w: the request can't be completed because the blocking user is trying to invite the blocked user to a channel", ErrAPIBadRequest)         // "400701",
	ErrBlockedUserInviteNotAllowed            = fmt.Errorf("%w: a blocked user is trying to invite the user who blocked them to a channel", ErrAPIBadRequest)                                          // "400702",
	ErrBannedUserEnterChannelNotAllowed       = fmt.Errorf("%w: the request can't be completed because the user is trying to enter a channel that they are banned from", ErrAPIBadRequest)             // "400750",
	ErrBannedUserEnterCustomChannelNotAllowed = fmt.Errorf("%w: the request can't be completed because the user is trying to enter a custom type channel that they are banned from", ErrAPIBadRequest) // "400751",
	ErrUnacceptable                           = fmt.Errorf("%w: the request is unacceptable because the combination of parameter values is invalid", ErrAPIBadRequest)                                 // "400920",
	ErrInvalidEndpoint                        = fmt.Errorf("%w: the request failed because it is sent to an invalid endpoint", ErrAPIBadRequest)                                                       // "400930",

	ErrAPIForbidden            = errors.New("forbidden")
	ErrApplicationNotAvailable = fmt.Errorf("%w: the application identified with the request isn't available", ErrAPIForbidden) // "403100",

	ErrAPITooManyRequests = errors.New("too many requests")
	ErrRateLimitExceeded  = fmt.Errorf("%w: the request can't be completed because you have exceeded your rate limits", ErrAPITooManyRequests) // "500910",

	ErrAPIInternalServerError                = errors.New("internal server error")
	ErrInternalErrorPushTokenNotRegistered   = fmt.Errorf("%w: the server encounters an error while trying to register the user's push token", ErrAPIInternalServerError)     // "500601",
	ErrInternalErrorPushTokenNotUnregistered = fmt.Errorf("%w: the server encounters an error while trying to unregister the user's push token", ErrAPIInternalServerError)   // "500602",
	ErrInternalError                         = fmt.Errorf("%w: the server encounters an unexpected exception while trying to process the request", ErrAPIInternalServerError) // "500901",

	ErrAPIServiceUnavailable = errors.New("service unavailable")
	ErrServiceUnavailable    = fmt.Errorf("%w: the request failed due to a temporary failure of the server", ErrAPIServiceUnavailable) // N/A

	errorMap = map[int]error{ //nolint:gochecknoglobals
		400100: ErrUnexpectedParameterTypeString,
		400101: ErrUnexpectedParameterTypeNumber,
		400102: ErrUnexpectedParameterTypeList,
		400103: ErrUnexpectedParameterTypeJSON,
		400104: ErrUnexpectedParameterTypeBoolean,
		400105: ErrMissingRequiredParameters,
		400106: ErrNegativeNumberNotAllowed,
		400108: ErrUnauthorizedRequest,
		400110: ErrParameterValueLengthExceeded,
		400111: ErrInvalidValue,
		400112: ErrIncompatibleValues,
		400113: ErrParameterValueOutOfRange,
		400114: ErrInvalidURLOfResource,
		400151: ErrNotAllowedCharacter,
		400201: ErrResourceNotFound,
		400202: ErrResourceAlreadyExists,
		400203: ErrTooManyItemsInParameter,
		400300: ErrDeactivatedUserNotAccessible,
		400301: ErrUserNotFound,
		400302: ErrInvalidAccessToken,
		400303: ErrInvalidSessionKeyValue,
		400304: ErrApplicationNotFound,
		400305: ErrUserIDLengthExceeded,
		400306: ErrPaidQuotaExceeded,
		400307: ErrDomainNotAllowed,
		400401: ErrInvalidAPIToken,
		400402: ErrMissingSomeParameters,
		400403: ErrInvalidJSONRequestBody,
		400404: ErrInvalidRequestURL,
		400500: ErrTooManyUserWebsocketConnections,
		400501: ErrTooManyApplicationWebsocketConnections,
		400700: ErrBlockedUserSendNotAllowed,
		400701: ErrBlockedUserInvitedNotAllowed,
		400702: ErrBlockedUserInviteNotAllowed,
		400750: ErrBannedUserEnterChannelNotAllowed,
		400751: ErrBannedUserEnterCustomChannelNotAllowed,
		400920: ErrUnacceptable,
		400930: ErrInvalidEndpoint,
		403100: ErrApplicationNotAvailable,
		500910: ErrRateLimitExceeded,
		500601: ErrInternalErrorPushTokenNotRegistered,
		500602: ErrInternalErrorPushTokenNotUnregistered,
		500901: ErrInternalError,
	}
)

// Error is an error returned by the Sendbird API.
// See https://sendbird.com/docs/chat/platform-api/v3/error-codes#2-error-codes
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func (c *client) handleError(status int, body io.Reader) error {
	var handledError Error

	err := json.NewDecoder(body).Decode(&handledError)
	if err != nil {
		b, _ := io.ReadAll(body)
		c.logger.Warn("failed to decode error", "body", string(b))

		return fmt.Errorf("failed to decode error: %w", err)
	}

	if !handledError.Error {
		c.logger.Warn("error is not an error", "handledError", handledError)

		return nil
	}

	if err, ok := errorMap[handledError.Code]; ok {
		return fmt.Errorf("%w: %s", err, handledError.Message)
	}

	switch status {
	case http.StatusBadRequest:
		return ErrAPIBadRequest
	case http.StatusForbidden:
		return ErrAPIForbidden
	case http.StatusTooManyRequests:
		return ErrAPITooManyRequests
	case http.StatusInternalServerError:
		return ErrAPIInternalServerError
	case http.StatusServiceUnavailable:
		return ErrAPIServiceUnavailable
	}

	return fmt.Errorf("unknown error %w: status: %d, code: %d, message: %q",
		ErrAPIDefault, status, handledError.Code, handledError.Message)
}
