package model

// 此文件为用户信息表格
// 用户表, id, username, password, role_id, permission:{"id":1,"name":"增加用户"}
// 角色表, id, name
// 权限表, id, name, url_path, per_menu(权限所属菜单)
// 权限所属菜单表, id, name,

// User 用户表
type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password"`
	Other      string `json:"other"`
	RoleId     int    `json:"role_id"`
	Permission string `json:"permission"`
}

// Role 角色表
type Role struct {
	BaseModel
	RoleName string `json:"name"`
}

// Permission 权限表
type Permission struct {
	BaseModel
	PerName string `json:"name"`
	UrlPath string `json:"url_path"`
	PerMenu int    `json:"per_menu"`
}

// PermissionMenu 权限菜单表
type PermissionMenu struct {
	BaseModel
	MenuName string `json:"name"`
}
