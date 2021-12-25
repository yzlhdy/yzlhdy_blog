package entity

type ResourceCategory struct {
	Model
	Icon string `json:"icon" form:"icon"`
	Name string `json:"name" form:"name"`
}
