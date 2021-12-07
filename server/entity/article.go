package entity

type Article struct {
	Model
	Title          string `json:"title" type:"varchar(100)"`
	SubTitle       string `json:"sub_title" type:"varchar(100)"`
	Content        string `json:"content" type:"text"`
	Image          string `json:"image" type:"varchar(100)"`
	Cid            int    `json:"cid" type:"int"`
	Classification Classification
}
