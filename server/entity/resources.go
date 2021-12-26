package entity

type Resources struct {
	Model
	Title            string           `json:"title" form:"title"`
	SubTitle         string           `json:"sub_title" form:"sub_title"`
	Recommend        int              `json:"recommend" form:"recommend"`
	Rid              int              `json:"rid" form:"rid"`
	ResourceCategory ResourceCategory `gorm:"foreignkey:Rid"`
}
