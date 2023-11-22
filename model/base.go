package model

import (
	"time"
)

// LocalTime 自定义时间格式
type LocalTime time.Time

// BaseModel 基础Model
type BaseModel struct {
	ID         uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	CreateTime LocalTime `json:"create_time"`
	ModifyTime LocalTime `json:"modify_time"`
}
