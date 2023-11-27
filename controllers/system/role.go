package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ngmp/config"
	"ngmp/model"
	"ngmp/utils/response"
	"time"
)

// RoleAdd 添加角色
func RoleAdd(c *gin.Context) {
	// 添加角色时要选择权限
	var role struct {
		Name        string   `json:"name" remark:"角色名"  binding:"required"` // 角色名
		Permissions []string `json:"permissions" remark:"权限ID列表"  binding:"required"`
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}

	// 判断名称是否已存在
	dbRole, err := model.NewRole().FindRoleByName(role.Name)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.LogicExceptionJSON("系统出错了："+err.Error(), c)
		return
	}
	if dbRole.ID > "" {
		response.InvalidArgumentJSON("角色名已存在", c)
		return
	}
	roleId := uuid.New().String()
	// 添加事务
	permissions, err := model.NewPermission().FindByIdList(role.Permissions)
	if err != nil {
		response.InvalidArgumentJSON("查询权限失败: "+err.Error(), c)
		return
	}
	// 创建角色跟权限对应关系
	newRole := model.Role{
		BaseModel: model.BaseModel{
			ID:         roleId,
			CreateTime: time.Now(),
		},
		Name:        role.Name,
		Permissions: permissions,
	}

	if err = config.DBDefault.Create(&newRole).Error; err != nil {
		response.InvalidArgumentJSON("创建角色失败: "+err.Error(), c)
		return
	}
	resp := map[string]string{"id": roleId}
	response.SuccessJSON(resp, "", c)
}

// RoleSelect 查看角色
func RoleSelect(c *gin.Context) {
	results2, err := model.NewRole().FindRoleByIdList("all")
	if err != nil {
		response.InvalidArgumentJSON("查询角色失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON(results2, "", c)
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	//roleID := c.Param("roleID")
	//log.Println("roleID", roleID)
	roleModel := model.NewRole()
	// 查询角色是否存在
	exitRole, err := roleModel.FindRoleById(c.Param("roleID"))
	if err != nil {
		response.InvalidArgumentJSON("查询角色失败: "+err.Error(), c)
		return
	}
	var roleInfo struct {
		NewRoleName string   `json:"new_role_name" remark:"新角色名"`
		Permissions []string `json:"permissions" remark:"最新的权限"`
	}
	if err = c.ShouldBindJSON(&roleInfo); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	// 开启数据库事务
	tx := config.DBDefault.Begin()
	defer tx.Rollback() // 回滚事务

	//  更新角色名
	newRoleName := roleInfo.NewRoleName
	if newRoleName != "" {
		exitRole.Name = newRoleName
	}

	// 更新或删除权限
	permissions := roleInfo.Permissions
	if len(permissions) > 0 {
		permissionsList, err := model.NewPermission().FindByIdList(permissions)
		if err != nil {

			response.InvalidArgumentJSON("查询权限失败: "+err.Error(), c)
			return
		}
		// 替换关联的权限
		if err = tx.Model(&exitRole).Association("Permissions").Replace(permissionsList); err != nil {
			tx.Rollback() // 回滚事务
			response.InvalidArgumentJSON("替换关联的权限失败: "+err.Error(), c)
			return
		}
	}
	currentTime := time.Now()
	exitRole.ModifyTime = &currentTime
	// 在事务中执行数据库操作
	if err = tx.Save(&exitRole).Error; err != nil {
		response.InvalidArgumentJSON("更新权限失败: "+err.Error(), c)
		return
	}
	// 提交事务
	tx.Commit()

	response.SuccessJSON("", "", c)
}

// DelRole 删除角色
func DelRole(c *gin.Context) {
	//roleID := c.Param("roleID")
	/*
		从关联表中删除角色与用户的关系。
		从关联表中删除角色与权限的关系。
		删除角色本身
	*/
	tx := config.DBDefault.Begin()
	// 回滚事务
	defer tx.Rollback()
	// 查询角色
	roleObj, err := model.NewRole().FindRoleById(c.Param("roleID"))
	if err != nil {
		response.InvalidArgumentJSON("查询角色失败: "+err.Error(), c)
		return
	}
	err = tx.Model(&roleObj).Association("Users").Clear()
	if err != nil {
		response.InvalidArgumentJSON("清空角色跟用户的关系失败: "+err.Error(), c)
		return
	}
	err = tx.Model(&roleObj).Association("Permissions").Clear()
	if err != nil {
		response.InvalidArgumentJSON("清空角色跟权限的关系失败: "+err.Error(), c)
		return
	}
	err = tx.Delete(&roleObj).Error
	if err != nil {
		response.InvalidArgumentJSON("清空角色跟权限的关系失败: "+err.Error(), c)
		return
	}
	tx.Commit()
	response.SuccessJSON("", "", c)
}
