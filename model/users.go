package model

import "ngmp/config"

// User 用户表
type User struct {
	BaseModel
	Username string `gorm:"type:varchar(255)" json:"username"`          // 用户名
	Password string `gorm:"type:varchar(255);not null" json:"password"` // 密码
	Other    string `json:"other"`                                      // 备注
	RoleId   string `gorm:"type:varchar(255)" json:"role_id"`           // 角色ID
}

// Role 角色表
type Role struct {
	BaseModel
	Name       string       `gorm:"type:varchar(255)" json:"name"` // 角色名
	Permission []Permission `gorm:"many2many:role_permissions;"`   // 权限
}

// Permission 权限表
type Permission struct {
	BaseModel
	Name   string `gorm:"type:varchar(255);unique;not null" json:"name"`         // 权限英文名
	ChName string `gorm:"type:varchar(255);unique;not null" json:"chinese_name"` // 权限中文名
	Path   string `gorm:"type:varchar(255);unique;not null" json:"path"`         // 对应路由
}

// FindByName 基于权限查找角色
func (p *Permission) FindByName(name string, path string) (permission *Permission, err error) {
	err = config.DBDefault.Where("name = ? OR path = ?", name, path).First(&permission).Error
	return
}

// NewPermission 初始化权限
func NewPermission() *Permission {
	return &Permission{}
}
