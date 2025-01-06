package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"optional"`
	Email   string `json:"email" validate:"optional,email"`
	Address string `json:"address" validate:"optional"`
	Phone   string `json:"phone" validate:"optional"`
}
