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

// MenuAdd 添加权限
func MenuAdd(c *gin.Context) {
	var menu struct {
		Name        string `json:"menu_name" remark:"权限名"  binding:"required"`
		ChineseName string `json:"chinese_name" remark:"权限中文名"  binding:"required"`
		Path        string `json:"menu_path" remark:"路径"  binding:"required"`
	}
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}

	// 判断名称或者路径是否已存在
	dbPer, err := model.NewPermission().FindByName(menu.Name, menu.Path)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.LogicExceptionJSON("系统出错了："+err.Error(), c)
		return
	}
	if dbPer.ID > "" {
		response.InvalidArgumentJSON("权限名或路径已存在", c)
		return
	}
	perId := uuid.New().String()
	// 创建权限
	newPer := model.Permission{
		BaseModel: model.BaseModel{
			ID:         perId,
			CreateTime: time.Now(),
		},
		Name:        menu.Name,
		ChineseName: menu.ChineseName,
		Path:        menu.Path,
	}

	if err = config.DBDefault.Create(&newPer).Error; err != nil {
		response.InvalidArgumentJSON("创建权限失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON(model.BaseIdResult{ID: perId}, "创建权限成功", c)
}

// MenuSelect 权限查询
func MenuSelect(c *gin.Context) {
	var params model.BasePageParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	result, err := model.NewPermission().FindPermissionList(params)
	if err != nil {
		response.InvalidArgumentJSON("查询角色失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON(result, "", c)
}

// UpdateMenu 更新权限
func UpdateMenu(c *gin.Context) {
	roleModel := model.NewPermission()
	// 查询权限是否存在
	exitPer, err := roleModel.FindPermissionById(c.Param("perID"))
	if err != nil {
		response.InvalidArgumentJSON("查询权限失败: "+err.Error(), c)
		return
	}
	var permissionsInfo struct {
		NewName        string `json:"new_name" remark:"新权限名"`
		NewChineseName string `json:"new_chinese_name" remark:"新权限中文名"`
		NewPath        string `json:"new_path" remark:"新权限路径"`
	}
	if err = c.ShouldBindJSON(&permissionsInfo); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	//  更新角色名
	newRoleName := permissionsInfo.NewName
	if newRoleName != "" {
		exitPer.Name = newRoleName
	}
	newChName := permissionsInfo.NewChineseName
	if newChName != "" {
		exitPer.ChineseName = newChName
	}
	newPath := permissionsInfo.NewPath
	if newPath != "" {
		exitPer.Path = newPath
	}
	currentTime := time.Now()
	exitPer.ModifyTime = &currentTime
	if err = config.DBDefault.Save(&exitPer).Error; err != nil {
		response.InvalidArgumentJSON("更新权限失败: "+err.Error(), c)
		return
	}

	response.SuccessJSON("", "", c)
}

// DelMenu 删除权限
func DelMenu(c *gin.Context) {
	roleModel := model.NewPermission()
	// 查询权限是否存在
	exitPer, err := roleModel.FindPermissionById(c.Param("perID"))
	if err != nil {
		response.InvalidArgumentJSON("查询权限失败: "+err.Error(), c)
		return
	}
	// 事务处理
	err = config.DBDefault.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&exitPer).Association("Roles").Clear(); err != nil {
			return err
		}
		if err = tx.Delete(&exitPer).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		response.InvalidArgumentJSON("删除权限失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON("", "", c)
}
