package dto

type ArticleCreateDto struct {
	Title    string `json:"title" form:"title" binding:"required,min=4,max=100"`
	SubTitle string `json:"sub_title" form:"sub_title" binding:"required,min=6,max=100"`
	Content  string `json:"content" form:"content" binding:"required,min=6,"`
	Image    string `json:"image" form:"image" binding:"required"`
	Cid      int    `json:"cid" form:"cid" binding:"required"`
}

type ArticleUpdateDto struct {
	Title    string `json:"title" form:"title"`
	SubTitle string `json:"sub_title" form:"sub_title"`
	Content  string `json:"content" form:"content"`
	Image    string `json:"image" form:"image"`
	Cid      int    `json:"cid" form:"cid"`
}
