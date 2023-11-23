package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"ngmp/config"
	"ngmp/model"
	"ngmp/utils/response"
	"time"
)

// RoleAdd 添加角色
func RoleAdd(c *gin.Context) {
	// 添加角色时要选择权限
	var role struct {
		Name        string   `json:"name"` // 角色名
		Permissions []string `json:"permissions"`
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
	response.SuccessJSON(resp, "创建角色成功", c)
}

// RoleSelect 查看角色
func RoleSelect(c *gin.Context) {
	omitFields := []string{"path", "name", "create_time", "modify_time"}
	results, err := model.FindRoleAndPermissions(omitFields)
	if err != nil {
		response.InvalidArgumentJSON("查询角色失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON(results, "", c)
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	roleID := c.Param("roleID")
	log.Println("roleID", roleID)
	roleModel := model.NewRole()

	// 查询角色是否存在
	exitRole, err := roleModel.FindRoleById(roleID)
	if err != nil {
		response.InvalidArgumentJSON("查询角色失败: "+err.Error(), c)
		return
	}
	var roleInfo struct {
		NewRoleName string   `json:"new_role_name"`
		Permissions []string `json:"permissions"`
	}
	if err := c.ShouldBindJSON(&roleInfo); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	// 开启数据库事务
	tx := config.DBDefault.Begin()
	//  更新角色名
	newRoleName := roleInfo.NewRoleName
	if newRoleName != "" {
		exitRole.Name = newRoleName
	}
	log.Println("roleInfo----", roleInfo)

	// 更新或删除权限
	permissions := roleInfo.Permissions
	//permissions := utils.GetMapValue(roleInfo, "permissions", 0)
	// 使用 Association 方法更新关联的权限
	//if err := tx.Model(&Role{}).Association("Permissions").Replace(updatedRole.Permissions); err != nil {
	//	tx.Rollback() // 回滚事务
	//	c.JSON(500, gin.H{"error": "Failed to update permissions"})
	//	return
	//}
	log.Println("permissions=======", permissions)
	if len(permissions) > 0 {
		exitRole.Permissions = []model.Permission{} // 清空已有的权限
		var permissionsList []model.Permission
		tx.Where("id IN (?)", permissions).Find(&permissionsList)
		//permissionsList, err := model.NewPermission().FindByIdList(permissions)
		if err != nil {
			response.InvalidArgumentJSON("查询权限失败: "+err.Error(), c)
			return
		}
		// 添加新的权限
		exitRole.Permissions = permissionsList
		// 替换关联的权限
		if err := tx.Model(&exitRole).Association("Permissions").Replace(permissionsList); err != nil {
			tx.Rollback() // 回滚事务
			c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to reload permissions: %s", err.Error())})
			return
		}
	}
	// 在事务中执行数据库操作
	if err := tx.Save(&exitRole).Error; err != nil {
		tx.Rollback() // 回滚事务
		c.JSON(500, gin.H{"error": "Failed to update role"})
		return
	}
	// 提交事务
	tx.Commit()

	response.SuccessJSON("", "", c)
}
