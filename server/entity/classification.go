package entity

type Classification struct {
	Model
	Name string `json:"name" form:"name"`
}
