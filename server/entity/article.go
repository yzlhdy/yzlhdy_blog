package entity

type Article struct {
	Model
	Title          string         `gorm:"type:varchar(255)" json:"title"`
	SubTitle       string         `gorm:"type:varchar(255)" json:"sub_title" `
	Content        string         `gorm:"type:text" json:"content"`
	Image          string         `gorm:"type:varchar(255)" json:"image"`
	Cid            int            `json:"cid"`
	Classification Classification `gorm:"foreignkey:Cid" json:"classification"`
}
