package vo

// SuccessResponse represents a success response with a message.
type SuccessResponse struct {
	Message string `json:"message" example:"success"`
}

// ErrorResponse represents an error response with a message.
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// LoginResponse represents a successful login response.
type LoginResponse struct {
	Message string `json:"message" example:"登录成功"`
	User    UserVO `json:"user"`
	Token   string `json:"token" example:"some.jwt.token"`
}

// RegisterResponse represents a successful registration response.
type RegisterResponse struct {
	Message string `json:"message" example:"注册成功"`
	User    UserVO `json:"user"`
}

// UpdateUserResponse represents a successful user update response.
type UpdateUserResponse struct {
	Message string `json:"message" example:"更新成功"`
	User    UserVO `json:"user"`
}
