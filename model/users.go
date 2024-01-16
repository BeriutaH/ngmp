package model

import (
	"ngmp/config"
)

// User 用户表
type User struct {
	BaseModel
	Username   string  `gorm:"type:varchar(255);unique" json:"username"`     // 用户名
	Password   string  `gorm:"type:varchar(255);not null" json:"-"`          // 密码
	SecretCode string  `gorm:"type:varchar(255);not null" json:"-"`          // 密码key
	Remark     *string `json:"remark"`                                       // 备注
	Roles      []Role  `gorm:"many2many:user_roles;" json:"roles,omitempty"` // 用户跟角色对应关系
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

// FindUserByName 基于用户name查询
func (p *User) FindUserByName(userName string) (user *User, err error) {
	err = config.DBDefault.First(&user, "username = ?", userName).Error
	return
}

// FindUserByIdList 基于id查询多个角色，全部或者指定的列表
func (p *User) FindUserByIdList(userIds any) (users []*User, err error) {
	if userIds == "all" {
		err = config.DBDefault.Preload("Roles").Find(&users).Error
	} else {
		err = config.DBDefault.Preload("Roles").Find(&users, userIds).Error
	}
	return
}

// FindUserList 按照分页展示查询相应的数据
func (p *User) FindUserList(params BasePageParams) (*BasePageResult[User], error) {
	// 需要过滤的字段
	pr := &BasePageResult[User]{Items: make([]*User, 0), Total: 0}
	err := config.DBDefault.Preload("Roles").Scopes(Paginate(params)).
		Scopes(SearchInfo(params.BaseOrderByParams)).Find(&pr.Items).Scopes(DBCount).Count(&pr.Total).Error
	if err != nil {
		return nil, err
	}
	// 计算总页数
	pr.CipherPage(params)
	return pr, nil
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
func (p *Role) FindRoleByIdList(roleIds any) (roles []Role, err error) {
	if roleIds == "all" {
		err = config.DBDefault.Preload("Permissions").Find(&roles).Error
	} else {
		err = config.DBDefault.Preload("Permissions").Find(&roles, roleIds).Error
	}
	return
}

// FindRoleList 按照分页展示查询相应的数据
func (p *Role) FindRoleList(params BasePageParams) (*BasePageResult[Role], error) {
	// 需要过滤的字段
	pr := &BasePageResult[Role]{Items: make([]*Role, 0), Total: 0}
	err := config.DBDefault.Preload("Permissions").Scopes(Paginate(params)).
		Scopes(SearchInfo(params.BaseOrderByParams)).Find(&pr.Items).Scopes(DBCount).Count(&pr.Total).Error
	if err != nil {
		return nil, err
	}
	// 计算总页数
	pr.CipherPage(params)
	return pr, nil
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

// FindPermissionById 查找权限名称或path
func (p *Permission) FindPermissionById(perId string) (permission *Permission, err error) {
	err = config.DBDefault.Preload("Roles").First(&permission, "id = ?", perId).Error
	return
}

// FindPermissionList 按照分页展示查询相应的数据
func (p *Permission) FindPermissionList(params BasePageParams) (*BasePageResult[Permission], error) {
	// 需要过滤的字段
	pr := &BasePageResult[Permission]{Items: make([]*Permission, 0), Total: 0}
	err := config.DBDefault.Scopes(Paginate(params)).Scopes(SearchInfo(params.BaseOrderByParams)).
		Find(&pr.Items).Scopes(DBCount).Count(&pr.Total).Error
	if err != nil {
		return nil, err
	}
	// 计算总页数
	pr.CipherPage(params)
	return pr, nil
}

// FindByIdList 查找权限Id
func (p *Permission) FindByIdList(idList []string) (permissions []Permission, err error) {
	err = config.DBDefault.Where("id IN (?)", idList).Find(&permissions).Error
	return
}
