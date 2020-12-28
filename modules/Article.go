package modules

import "time"

type Article struct {
	Id        int64 `orm:"pk;auto"`
	Name      string
	Price     float32
	Author    string
	UpdatedAt *time.Time `orm:"auto_now;null;"`
	CreatedAt *time.Time `orm:"auto_now_add;null"`
	DeletedAt *time.Time `orm:"null;"`
}
