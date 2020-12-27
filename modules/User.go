package modules

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID int64 `orm:"column(id);pk;auto"`
	Name string
	Password string
	Address string
}

func (user *User)CheckPassword(password string) bool  {
	if err :=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err != nil{
		return false
	}
	return true
}