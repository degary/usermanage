package forms

type UpdateUserForm struct {
	Name string `form:"username"`
	Password1 string `form:"password1"`
	Password2 string `form:"password2"`
	Address string `form:"address"`
}


