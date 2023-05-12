package users

import (
	"crypto/md5"
	"easy-gin/global"
	"easy-gin/models/common"
	"fmt"
)

//User 表结构体
type User struct {
	common.CommmonModel
	Email     string `json:"email" gorm:"column:email"`
	Username  string `json:"username" gorm:"column:username"`
	Salt      string `json:"salt" gorm:"column:salt"`
	Password  string `json:"password" gorm:"column:password"`
	Signature string `json:"signature" gorm:"column:signature"`
}

type UserList []User

func (User) TableName() string {
	return "ey_users"
}

//Update 更新数据
func (us *User) Update() bool {
	err := global.Db.Where("id", us.ID).Updates(&us).Error
	if err != nil {
		return false
	}
	return true
}

//Create 添加数据
func (us *User) create() bool {
	err := global.Db.Create(&us).Error
	if err != nil {
		return false
	}
	return true
}

//IsExistByField 根据字段判断用户是否存在
func (us *User) IsExistByField(field string, value any) bool {
	err := global.Db.Where(field, value).Find(&us).Error
	if err != nil {
		return false
	}
	if us.ID <= 0 {
		return false
	}
	return true
}

//IfPasswordCorrect 判断密码
func (us *User) IfPasswordCorrect(password string) bool {
	passwordImport := fmt.Sprintf("%s%s%s", us.Salt, password, us.Salt)
	passwordImport = fmt.Sprintf("%x", md5.Sum([]byte(passwordImport)))
	if passwordImport != us.Password {
		return false
	}
	return true
}

//Find 根据id 查询
func (us *User) Find(id uint) {
	_ = global.Db.Where("id", id).Find(&us).Error
}
