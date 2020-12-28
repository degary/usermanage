package forms

type UpdateArticleForm struct {
	Id     int64   `form:"id"`
	Author string  `form:"author"`
	Name   string  `form:"name"`
	Price  float32 `form:"price"`
}
