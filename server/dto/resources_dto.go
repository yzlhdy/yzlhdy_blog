package dto

type CreateResourcesDto struct {
	Title     string `json:"title" form:"title" binding:"required,min=2,max=255"`
	SubTitle  string `json:"sub_title" form:"sub_title" binding:"required,min=2,max=255"`
	Recommend int    `json:"recommend" form:"recommend" binding:"required"`
	Rid       int    `json:"rid" form:"rid" binding:"required"`
}

type UpdateResourcesDto struct {
	Title     string `json:"title" form:"title"`
	SubTitle  string `json:"sub_title" form:"sub_title"`
	Recommend int    `json:"recommend" form:"recommend"`
	Rid       int    `json:"rid" form:"rid"`
}
