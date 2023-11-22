package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"ngmp/config"
	"ngmp/model"
	"ngmp/utils/response"
	"time"
)

// MenuAdd 添加权限
func MenuAdd(c *gin.Context) {
	var menu struct {
		Name   string `json:"menu_name"`
		ChName string `json:"chinese_name"`
		Path   string `json:"menu_path"`
	}
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	log.Println("获取的权限", menu)

	// 判断名称或者路径是否已存在
	dbPer, err := model.NewPermission().FindByName(menu.Name, menu.Path)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.LogicExceptionJSON("系统出错了："+err.Error(), c)
		return
	}
	log.Println("权限ID", dbPer.ID)
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
			ModifyTime: time.Now(),
		},
		Name:   menu.Name,
		ChName: menu.ChName,
		Path:   menu.Path,
	}

	if err := config.DBDefault.Create(&newPer).Error; err != nil {
		response.InvalidArgumentJSON("创建权限失败", c)
		return
	}
	resp := map[string]string{"id": perId}
	response.SuccessJSON(resp, "创建权限成功", c)
}
