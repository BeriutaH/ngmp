package model

// 此文件为用户信息表格
// 用户表, id, username, password, role_id, permission:{"id":1,"name":"增加用户"}
// 角色表, id, name
// 权限表, id, name, url_path, per_menu(权限所属菜单)
// 权限所属菜单表, id, name,

// User 用户表
type User struct {
	ID         uint   `db:"id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Other      string `db:"other"`
	RoleId     int    `db:"role_id"`
	Permission string `db:"permission"`
}

// Role 角色表
type Role struct {
	ID       uint   `db:"id"`
	RoleName string `db:"name"`
}

// Permission 权限表
type Permission struct {
	ID      uint   `db:"id"`
	PerName string `db:"name"`
	UrlPath string `db:"url_path"`
	PerMenu int    `db:"per_menu"`
}

// PermissionMenu 权限菜单表
type PermissionMenu struct {
	ID       uint   `db:"permission_id"`
	MenuName string `db:"name"`
}
