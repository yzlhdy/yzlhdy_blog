package dto

type CreateReCategory struct {
	Icon string `json:"icon" form:"icon" binding:"required"`
	Name string `json:"name" form:"name" binding:"required,min=3,max=30"`
}
type UpdateReCategory struct {
	Icon string `json:"icon" form:"icon"`
	Name string `json:"name" form:"name"`
}
