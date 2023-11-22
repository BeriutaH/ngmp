package model

import (
	"time"
)

// LocalTime 自定义时间格式
//type LocalTime time.Time

// BaseModel 基础Model
type BaseModel struct {
	ID         string     `json:"id" gorm:"type:varchar(255);primaryKey"`
	CreateTime time.Time  `json:"create_time"`
	ModifyTime *time.Time `json:"modify_time"`
}
