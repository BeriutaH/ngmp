package model

import (
	"gorm.io/gorm"
	"ngmp/config"
)

// User 用户表
type User struct {
	BaseModel
	Username   string  `gorm:"type:varchar(255)" json:"username"`                       // 用户名
	Password   string  `gorm:"type:varchar(255);not null" json:"password,omitempty"`    // 密码
	SecretCode string  `gorm:"type:varchar(255);not null" json:"secret_code,omitempty"` // 密码key
	Remark     *string `json:"remark"`                                                  // 备注
	Roles      []Role  `gorm:"many2many:user_roles;" json:"roles,omitempty"`            // 用户跟角色对应关系
}

// NewUser 初始化用户
func NewUser() *User {
	return &User{}
}

// FindUserById 基于用户id查询
func (p *User) FindUserById(userId string) (user *User, err error) {
	err = config.DBDefault.First(&user, "id = ?", userId).Error
	return
}

// FindUserByIdList 基于id查询多个角色，全部或者指定的列表
func (p *User) FindUserByIdList(userIds interface{}) (users []User, err error) {
	omitList := []string{"password", "secret_code"}
	if userIds == "all" {
		err = config.DBDefault.Preload("Roles").Omit(omitList...).Find(&users).Error
	} else {
		err = config.DBDefault.Preload("Roles").Omit(omitList...).Find(&users, userIds).Error
	}
	return
}

// Role 角色表
type Role struct {
	BaseModel
	Name        string       `gorm:"type:varchar(255);unique" json:"name"`                     // 角色名
	Users       []User       `gorm:"many2many:user_roles" json:"users,omitempty"`              // 角色跟用户对应关系
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"` // 权限
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

// FindRoleByIdList 基于id查询多个角色，全部或者指定的列表
func (p *Role) FindRoleByIdList(roleIds interface{}) (roles []Role, err error) {
	if roleIds == "all" {
		err = config.DBDefault.Preload("Permissions").Find(&roles).Error
	} else {
		err = config.DBDefault.Preload("Permissions").Find(&roles, roleIds).Error
	}
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
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`           // 权限英文名
	ChineseName string `gorm:"type:varchar(255);unique;not null" json:"chinese_name"`   // 权限中文名
	Path        string `gorm:"type:varchar(255);unique;not null" json:"path,omitempty"` // 对应路由
	Roles       []Role `gorm:"many2many:role_permissions" json:"roles,omitempty"`       // 权限跟角色的对应关系
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
