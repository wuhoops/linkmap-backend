package response

type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func NewError(args1 any, args2 ...any) *ErrorResponse {
	if message, ok := args1.(string); ok {
		if len(args2) == 0 {
			return &ErrorResponse{
				Success: false,
				Message: message,
			}
		}
		if message2, ok := args2[0].(string); ok {
			return &ErrorResponse{
				Success: false,
				Message: message,
				Error:   message2,
			}
		} else {
			if ok := len(args2) == 1; ok {
				return &ErrorResponse{
					Success: false,
					Message: message,
					Error:   args2[0],
				}
			} else {
				return &ErrorResponse{
					Success: false,
					Message: message,
					Error:   args2,
				}
			}
		}
	}
	return &ErrorResponse{
		Success: true,
		Error:   args1,
	}
}
