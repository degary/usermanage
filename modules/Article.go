package modules

import "time"

type Article struct {
	Id int64 `orm:"pk;auto"`
	Name string
	Price float32
	Author string
	UpdatedAt *time.Time `orm:"auto_now;"`
	CreatedAt *time.Time `orm:"auto_now_add;"`
	DeletedAt *time.Time `orm:"null;"`
}
