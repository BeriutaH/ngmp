package system

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"ngmp/config"
	"ngmp/model"
	"ngmp/utils"
	"ngmp/utils/response"
	"time"
)

// UserData 获取用户信息
func UserData(c *gin.Context) {
	// 需要过滤的字段
	result, err := model.NewUser().FindUserByIdList("all")
	if err != nil {
		response.LogicExceptionJSON("查询用户失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON(result, "", c)

}

// UserAdd 添加用户
func UserAdd(c *gin.Context) {
	db := config.DBDefault
	//校验角色id是否存在
	//加密密码，存入数据库
	var user struct {
		Username string   `json:"username"  remark:"用户名"  binding:"required"`
		Password string   `json:"password" remark:"密码" binding:"required"`
		Remark   string   `json:"remark" remark:"备注"`
		Roles    []string `json:"roles" remark:"角色id列表" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}

	// 判断角色
	roleList, err := model.NewRole().FindRoleByIdList(user.Roles)
	if err != nil {
		response.LogicExceptionJSON("查询角色失败: "+err.Error(), c)
		return
	}

	// 加密密码
	key := utils.GenerateRandomString()
	plaintext, err := utils.EncryptByAes(user.Password, key)
	if err != nil {
		response.LogicExceptionJSON("密码加密失败: "+err.Error(), c)
		return
	}

	// 创建用户，加密密码
	userId := uuid.New().String()
	//创建用户
	newUser := model.User{
		BaseModel: model.BaseModel{
			ID:         userId,
			CreateTime: time.Now(),
		},
		Username:   user.Username,
		Password:   plaintext,
		SecretCode: key,
		Remark:     &user.Remark,
		Roles:      roleList,
	}
	if err = db.Create(&newUser).Error; err != nil {
		response.InvalidArgumentJSON("创建用户失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON("", "", c)
}

// UpdateUser 修改用户
func UpdateUser(c *gin.Context) {
	//校验角色id是否存在，无法更改密码
	var user struct {
		Username string       `json:"username"  remark:"用户名"`
		Remark   string       `json:"remark" remark:"备注"`
		Roles    []model.Role `json:"roles" remark:"最新角色列表"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	userID := c.Param("userID")
	// 判断角色
	userList, err := model.NewUser().FindUserByIdList([]string{userID})
	if err != nil {
		response.LogicExceptionJSON("查询角色失败: "+err.Error(), c)
		return
	}
	// 开启数据库事务
	tx := config.DBDefault.Begin()
	userObj := userList[0]
	newUserName := user.Username
	if newUserName != "" {
		userObj.Username = newUserName
	}
	if user.Remark != "" {
		userObj.Remark = &user.Remark
	}
	if len(user.Roles) > 0 {
		roleStructs := user.Roles
		// 替换关联的角色
		if err := tx.Model(&userObj).Association("Roles").Replace(roleStructs); err != nil {
			tx.Rollback() // 回滚事务
			response.LogicExceptionJSON("替换关联的角色失败: "+err.Error(), c)
			return
		}
	}
	currentTime := time.Now()
	userObj.ModifyTime = &currentTime
	// 在事务中执行数据库操作
	if err := tx.Save(&userObj).Error; err != nil {
		tx.Rollback() // 回滚事务
		response.InvalidArgumentJSON("更新用户失败: "+err.Error(), c)
		return
	}
	// 提交事务
	tx.Commit()
	response.SuccessJSON("", "", c)
}

func DelUser(c *gin.Context) {
	userID := c.Param("userID")
	log.Println(userID)
	// 查询用户，删除用户跟角色的多对多关系，删除用户，提交事务
	userObj, err := model.NewUser().FindUserById(userID)
	if err != nil {
		response.LogicExceptionJSON("当前用户不存在"+err.Error(), c)
		return
	}
	// 开启数据库事务
	tx := config.DBDefault.Begin()
	err = tx.Model(&userObj).Association("Roles").Clear()
	if err != nil {
		tx.Rollback()
		response.LogicExceptionJSON("清空当前用户角色失败:"+err.Error(), c)
		return
	}
	err = tx.Delete(userObj).Error
	if err != nil {
		tx.Rollback()
		response.LogicExceptionJSON("删除用户失败:"+err.Error(), c)
		return
	}
	// 提交事务
	tx.Commit()
	response.SuccessJSON("", "", c)
}
