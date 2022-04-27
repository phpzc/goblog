package user

import (
	"fmt"
	"goblog/app/models"
	"goblog/pkg/password"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm:"-" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (user *User) ComparePassword(_password string) bool {
	fmt.Println(_password)
	fmt.Println(user.Password)

	return password.CheckHash(_password, user.Password)
}

//Link方法生成用户链接
func (user *User) Link() string {
	return ""
}
