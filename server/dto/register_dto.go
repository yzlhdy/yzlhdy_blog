package dto

type RegisterDTO struct {
	UserName string `json:"userName" form:"userName" binding:"required,min=3,max=30"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=30"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Avatar   string `json:"avatar" form:"avatar"`
}
