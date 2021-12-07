package dto

type CreateClassifDto struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type UpdateClassifDto struct {
	Name string `json:"name" form:"name" binding:"required" validate:"min=4,max=255"`
}
