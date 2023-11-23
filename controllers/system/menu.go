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
		Name        string `json:"menu_name"`
		ChineseName string `json:"chinese_name"`
		Path        string `json:"menu_path"`
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
	resp := map[string]string{"id": perId}
	response.SuccessJSON(resp, "创建权限成功", c)
}

// MenuSelect 权限查询
func MenuSelect(c *gin.Context) {
	// 需要过滤的字段
	omitFields := []string{"path"}
	var results []map[string]interface{}
	// 获取权限表中的所有数据
	config.DBDefault.Model(model.NewPermission()).Omit(omitFields...).Find(&results)
	response.SuccessJSON(results, "", c)
}
