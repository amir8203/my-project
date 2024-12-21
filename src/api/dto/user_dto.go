package dto

type RegisterUserByUsernameRequest struct {
	Name     string `json:"name" binding:"required,min=6"`
	Username string `json:"username" binding:"required,min=5"`
	Phone    string `json:"phone" binding:"min=6,mobile"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserProfileRequest struct {
	Name     string `json:"name,omitempty" binding:"omitempty,min=6"`
	Username string `json:"username,omitempty" binding:"omitempty,min=5"`
	Phone    string `json:"phone,omitempty" binding:"omitempty,min=6,mobile"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6"`
}