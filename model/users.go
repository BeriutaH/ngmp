package model

import (
	"gorm.io/gorm"
	"ngmp/config"
)

// User 用户表
type User struct {
	BaseModel
	Username   string  `gorm:"type:varchar(255)" json:"username"`             // 用户名
	Password   string  `gorm:"type:varchar(255);not null" json:"password"`    // 密码
	SecretCode string  `gorm:"type:varchar(255);not null" json:"secret_code"` // 密码key
	Remark     *string `json:"remark"`                                        // 备注
	RoleId     string  `gorm:"type:varchar(255)" json:"role_id"`              // 角色ID
}

// NewUser 初始化用户
func NewUser() *User {
	return &User{}
}

// FindUserByName 基于角色名字
func (p *User) FindUserByName(name string) (user *User, err error) {
	err = config.DBDefault.Where("name = ?", name).First(&user).Error
	return
}

// Role 角色表
type Role struct {
	BaseModel
	Name        string       `gorm:"type:varchar(255);unique" json:"name"` // 角色名
	Permissions []Permission `gorm:"many2many:role_permissions;"`          // 权限
}

// NewRole 初始化角色
func NewRole() *Role {
	return &Role{}
}

// FindRoleByName 基于角色名字
func (p *Role) FindRoleByName(name string) (role *Role, err error) {
	err = config.DBDefault.Where("name = ?", name).First(&role).Error
	return
}

// FindRoleById 基于角色id
func (p *Role) FindRoleById(roleId string) (role *Role, err error) {
	err = config.DBDefault.Preload("Permissions").First(&role, "id = ?", roleId).Error
	return
}

// FindRoleAndPermissions 基于角色查询角色以及对应的权限
func FindRoleAndPermissions(PerRemovedFields []string) ([]map[string]interface{}, error) {
	roleModel := config.DBDefault.Model(NewRole())
	var roles []Role
	var results []map[string]interface{}
	// 获取角色表中的所有数据
	err := roleModel.Preload("Permissions", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, chinese_name")
	}).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		var permissionsMap []map[string]interface{}
		for _, per := range role.Permissions {
			perMap := map[string]interface{}{
				"id":           per.ID,
				"chinese_name": per.ChineseName,
			}
			permissionsMap = append(permissionsMap, perMap)
		}
		roleMap := map[string]interface{}{
			"name":        role.Name,
			"id":          role.ID,
			"permissions": permissionsMap,
			"create_time": role.CreateTime,
			"modify_time": role.ModifyTime,
		}
		results = append(results, roleMap)
	}
	return results, nil
}

// Permission 权限表
type Permission struct {
	BaseModel
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`         // 权限英文名
	ChineseName string `gorm:"type:varchar(255);unique;not null" json:"chinese_name"` // 权限中文名
	Path        string `gorm:"type:varchar(255);unique;not null" json:"path"`         // 对应路由
}

// NewPermission 初始化权限
func NewPermission() *Permission {
	return &Permission{}
}

// FindByName 查找权限名称或path
func (p *Permission) FindByName(name string, path string) (permission *Permission, err error) {
	err = config.DBDefault.Where("name = ? OR path = ?", name, path).First(&permission).Error
	return
}

// FindByIdList 查找权限Id
func (p *Permission) FindByIdList(idList []string) (permissions []Permission, err error) {
	err = config.DBDefault.Where("id IN (?)", idList).Find(&permissions).Error
	return
}
