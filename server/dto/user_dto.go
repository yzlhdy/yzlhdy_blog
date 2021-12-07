package dto

type UserCreateDto struct {
	UserName string `json:"user_name" form:"user_name" binding:"required,min=3,max=30"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=30"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type UserUpdateDto struct {
	ID       int    `json:"id" form:"id" binding:"required"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
}
