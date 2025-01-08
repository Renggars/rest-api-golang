package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"optional"`
	Email    string `json:"email" validate:"optional"`
	Password string `json:"password" validate:"optional"`
	Address  string `json:"address" validate:"optional"`
	Phone    string `json:"phone" validate:"optional"`
}

type UserUpdateEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
