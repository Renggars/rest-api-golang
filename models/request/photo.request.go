package request

type PhotoCreateRequest struct {
	CategoryID int `json:"category_id" form:"category_id" validate:"required"`
}
