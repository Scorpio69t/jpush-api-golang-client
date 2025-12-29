package jpush

import "fmt"

// ValidatePushPayload checks required fields before sending.
func ValidatePushPayload(p *PayLoad) error {
	if p == nil {
		return &ValidationError{Reason: "payload is required"}
	}
	if p.Platform == nil || p.Platform.Interface() == nil {
		return &ValidationError{Reason: "platform is required"}
	}
	if p.Audience == nil || p.Audience.Interface() == nil {
		return &ValidationError{Reason: "audience is required"}
	}
	return nil
}

// ValidateSchedule checks schedule payload.
func ValidateSchedule(s *Schecule) error {
	if s == nil {
		return &ValidationError{Reason: "schedule is required"}
	}
	if s.Name == "" {
		return &ValidationError{Reason: "schedule name is required"}
	}
	if s.Push == nil {
		return &ValidationError{Reason: "schedule payload is required"}
	}
	if err := ValidatePushPayload(s.Push); err != nil {
		return fmt.Errorf("schedule push: %w", err)
	}
	return nil
}

// ValidateSMSPayload checks SMS payload.
func ValidateSMSPayload(p *SmsPayLoad) error {
	if p == nil {
		return &ValidationError{Reason: "sms payload is required"}
	}
	if p.Mobile == "" {
		return &ValidationError{Reason: "sms mobile is required"}
	}
	if p.TempId == 0 {
		return &ValidationError{Reason: "sms temp_id is required"}
	}
	return nil
}
