package model

import (
	"time"
)

// LocalTime 自定义时间格式
//type LocalTime time.Time

// BaseIdResult ID结果响应
type BaseIdResult struct {
	ID string `json:"id"`
}

// BaseModel 基础Model
type BaseModel struct {
	ID         string     `json:"id" gorm:"type:varchar(255);primaryKey;unique"`
	CreateTime time.Time  `json:"create_time"`
	ModifyTime *time.Time `json:"modify_time"`
}

// BaseTimeFormat 方法，返回时间的年月日时分秒格式字符串
func (b *BaseModel) BaseTimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// CreateTimeFormatted 方法，返回 CreateTime 的格式化字符串
func (b *BaseModel) CreateTimeFormatted() string {
	return b.BaseTimeFormat(b.CreateTime)
}

// ModifyTimeFormatted 方法，返回 ModifyTime 的格式化字符串
func (b *BaseModel) ModifyTimeFormatted() string {
	if b.ModifyTime == nil {
		return ""
	}
	return b.BaseTimeFormat(*b.ModifyTime)
}
