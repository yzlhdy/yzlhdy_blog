package dto

type UserCreateDto struct {
	UserName string `json:"userName" form:"userName" binding:"required,min=3,max=30"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=30"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type UserUpdateDto struct {
	ID       uint   `json:"id" form:"id" binding:"required"`
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
}
