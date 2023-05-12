package common

import (
	"easy-gin/global"
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"column:id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

type Crud struct {
}

//Find 根据id 查询
func (i *Crud) Find(id uint) {
	_ = global.Db.Where("id", id).Find(&i).Error
}
