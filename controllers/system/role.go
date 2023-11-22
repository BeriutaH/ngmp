package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
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
	//permissions, err := model.NewPermission().FindByIdList(role.Permission)
	// 创建权限
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

func RoleData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "增加用户!!!!",
	})
}
