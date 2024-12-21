package service_errors

type ServiceError struct {
	EndUserMessage   string `json:"end_user_message"`
	TechnicalMessage string `json:"technical_message"`
	Err              error
}

func (s *ServiceError) Error() string {
	return s.EndUserMessage
}