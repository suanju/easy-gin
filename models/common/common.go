package common

import (
	"time"
)

type Model struct {
	ID        uint `json:"id" gorm:"column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
