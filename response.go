package jpush

import "encoding/json"

// APIError represents an error returned by the JPush API or HTTP layer.
type APIError struct {
	StatusCode int    `json:"status_code,omitempty"`
	Code       int    `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	RawBody    []byte `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

// ValidationError is returned when client-side validation fails.
type ValidationError struct {
	Reason string
}

func (e *ValidationError) Error() string {
	return e.Reason
}

// PushResponse represents a Push API v3 response.
type PushResponse struct {
	MsgID  int64        `json:"msg_id"`
	SendNo json.Number  `json:"sendno,omitempty"`
	Error  *APIError    `json:"error,omitempty"`
	Extra  interface{}  `json:"-"` // placeholder for unmodeled fields
	Raw    json.RawMessage `json:"-"`
}

// ScheduleResponse represents schedule operations.
type ScheduleResponse struct {
	ScheduleID string      `json:"schedule_id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Error      *APIError   `json:"error,omitempty"`
	Raw        interface{} `json:"-"`
}

// CIDResponse represents CID list response.
type CIDResponse struct {
	CIDList []string `json:"cidlist"`
	Error   *APIError `json:"error,omitempty"`
}

// SMSResponse represents SMS API response.
type SMSResponse struct {
	MsgID string     `json:"msg_id,omitempty"`
	Error *APIError  `json:"error,omitempty"`
	Raw   json.RawMessage `json:"-"`
}

type apiErrorEnvelope struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func parseAPIError(status int, body []byte) error {
	var env apiErrorEnvelope
	_ = json.Unmarshal(body, &env)
	msg := env.Error.Message
	code := env.Error.Code
	if msg == "" {
		msg = string(body)
	}
	return &APIError{
		StatusCode: status,
		Code:       code,
		Message:    msg,
		RawBody:    body,
	}
}
