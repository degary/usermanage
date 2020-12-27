package services

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/degary/usermanage/forms"
	"github.com/degary/usermanage/modules"
	"golang.org/x/crypto/bcrypt"
)

var UserService *userService = newUserService()

type userService struct {

}

func newUserService() *userService {
	return &userService{}
}

func GetUsers() []*modules.User {
	ormer := orm.NewOrm()
	querySeter := ormer.QueryTable(&modules.User{})
	var users []*modules.User
	querySeter.All(&users)
	return users
}

func AddUser(name ,password,address string)  {
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := modules.User{
		Name: name,
		Password: string(fromPassword),
		Address: address,
	}
	ormer := orm.NewOrm()
	ormer.Insert(&user)
}

func DeleteUser(id int64)  {
	ormer := orm.NewOrm()
	ormer.Delete(&modules.User{ID: id})
}

func GetUserByName(name string) *modules.User {
	user := &modules.User{
		Name: name,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name");err != nil{
		return nil
	}
	return user
}

func GetUserById(pk int64) *modules.User {
	ormer := orm.NewOrm()
	user := &modules.User{ID: pk}
	if err := ormer.Read(user);err != nil{
		return nil
	}
	return user
}

func UpdateUserById(pk int64) *modules.User {
	return nil

}

func Auth(loginform *forms.LoginForm) *modules.User {
	if user := GetUserByName(loginform.Username);user == nil{
		logs.Error("用户名或密码错误")
		return nil
	}else {
		if user.CheckPassword(loginform.Password) {
			//成功
			return user
		} else {
			//用户名或密码错误
			return nil
		}

	}
}

