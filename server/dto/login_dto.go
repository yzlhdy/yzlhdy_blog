package dto

type LoginDTO struct {
	Email    string `json:"email" from:"email" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}
