package model

import (
	"fmt"
	"gorm.io/gorm"
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
	//CreateTimeStr string     `json:"create_time" gorm:"-"`
	//ModifyTimeStr *string    `json:"modify_time" gorm:"-"`
}

// BaseOrderByParams 排序请求参数
type BaseOrderByParams struct {
	Field       string `json:"field,omitempty" remark:"排序字段"`
	Order       string `json:"order,omitempty" remark:"排序方向"` //asc,升序 desc降序
	SearchField string `json:"searchField,omitempty" remark:"查询的字段"`
	SearchValue string `json:"searchValue,omitempty" remark:"查询的值"`
}

// BasePageParams 分页请求参数
type BasePageParams struct {
	BaseOrderByParams
	Page     int `json:"page" remark:"页码" binding:"required,gt=0"`
	PageSize int `json:"pageSize" remark:"每页显示条数" binding:"required,gt=0"`
}

// BasePageResult 分页结果响应
type BasePageResult[T any] struct {
	Items     []*T  `json:"items" remark:"显示内容"`
	Total     int64 `json:"total" remark:"总记录数"`
	TotalPage int64 `json:"totalPage" remark:"总页数"`
}

// CipherPage 计算总页码
func (p *BasePageResult[T]) CipherPage(params BasePageParams) {
	p.TotalPage = (p.Total + int64(params.PageSize) - 1) / int64(params.PageSize)
	return
}

// Paginate 分页数据
func Paginate(pageInfo BasePageParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageInfo.Page <= 0 {
			pageInfo.Page = 1
		}
		if pageInfo.PageSize <= 0 {
			pageInfo.PageSize = 10
		}
		if pageInfo.PageSize > 100 {
			pageInfo.PageSize = 100
		}
		offset := (pageInfo.Page - 1) * pageInfo.PageSize
		return db.Offset(offset).Limit(pageInfo.PageSize).Scopes(OrderBy(pageInfo.BaseOrderByParams))
	}
}

// DBCount 总条数
func DBCount(db *gorm.DB) *gorm.DB {
	return db.Offset(-1).Limit(-1)
}

// OrderBy 排序处理
func OrderBy(oderByInfo BaseOrderByParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if oderByInfo.Field == "" || oderByInfo.Order == "" {
			return db
		}
		return db.Order(fmt.Sprintf("%s %s", oderByInfo.Field, oderByInfo.Order))
	}
}

// SearchInfo 条件查询处理
func SearchInfo(oderByInfo BaseOrderByParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if oderByInfo.SearchValue == "" || oderByInfo.SearchField == "" {
			return db
		}
		return db.Where(fmt.Sprintf("%s LIKE ?", oderByInfo.SearchField), "%"+oderByInfo.SearchValue+"%")
	}
}

//func (b *BaseModel) AfterFind() (err error) {
//	b.convertTimeToString()
//	return nil
//}

//func (b *BaseModel) convertTimeToString() {
//	b.CreateTimeStr = b.CreateTime.Format(config.TimeString)
//	if b.ModifyTime != nil {
//		modifyTimeString := b.ModifyTime.Format(config.TimeString)
//		b.ModifyTimeStr = &modifyTimeString
//	}
//}
