package entity

type User struct {
	Model
	UserName string `gorm:"type:varchar(255)" json:"user_name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar" `
	Token    string `json:"token"`
}
