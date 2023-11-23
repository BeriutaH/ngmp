package system

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"ngmp/config"
	"ngmp/model"
	"ngmp/utils"
	"ngmp/utils/response"
	"time"
)

type SelectInfo struct {
	Token string `form:"token" json:"token" binding:"required"`
}

type UserInfoData struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	Permission string `json:"permission"`
}

// UserData 获取用户信息
func UserData(c *gin.Context) {
	// 需要过滤的字段
	userModel := config.DBDefault.Model(model.NewUser())
	var result []map[string]interface{}
	err := userModel.
		Select("users.id as user_id, users.username, users.remark, users.role_id, roles.name as role_name").
		Joins("LEFT JOIN roles ON users.role_id = roles.id").
		Where("users.role_id = roles.id").
		Scan(&result).Error
	if err != nil {
		response.LogicExceptionJSON("查询用户失败: "+err.Error(), c)
		return
	}
	response.SuccessJSON(result, "", c)

}

// UserAdd 添加用户
func UserAdd(c *gin.Context) {
	//db := config.DBDefault
	// 校验角色id是否存在
	// 加密密码，存入数据库
	var user struct {
		Username string `json:"username"` // 用户名
		Password string `json:"password"` // 密码
		Remark   string `json:"remark"`   // 备注
		RoleId   string `json:"role_id"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}

	// 判断名称或者路径是否已存在
	_, err := model.NewRole().FindRoleById(user.RoleId)
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
		RoleId:     user.RoleId,
	}
	if err = config.DBDefault.Create(&newUser).Error; err != nil {
		response.InvalidArgumentJSON("创建用户失败: "+err.Error(), c)
		return
	}
	resp := map[string]string{"id": userId}

	response.SuccessJSON(resp, "创建用户成功", c)
}

func BlogComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "增加用户!!!!",
	})
}
