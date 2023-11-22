package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"ngmp/utils"
	"ngmp/utils/response"
	"strconv"
	"time"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SelectInfo struct {
	Token string `form:"token" json:"token" binding:"required"`
}

type UserInfoData struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	Permission string `json:"permission"`
}

// LoginFunc 登录
func LoginFunc(c *gin.Context) {
	// 声明接收的变量
	var loginJson Login
	if err := c.ShouldBindJSON(&loginJson); err != nil {
		c.JSON(http.StatusOK, utils.ReturnMsgFunc(400, "未携带正确信息!", 0))
		return
	}
	// 判断用户名密码是否正确
	//user := model.User{}
	//// 将密码转换成md5
	//pwd := PassMd5(loginJson.Password)
	//db := config.DBDefault
	//db.Find(&user, db.Where("username = ? AND password = ?", loginJson.User, pwd))
	//if pwd == "" {
	//	c.JSON(http.StatusOK, utils.ReturnMsgFunc(400, "用户名或密码错误!", 0))
	//	return
	//} else {
	//	// 校验用户名和密码正确之后,生成token值
	//	token := TokenMd5()
	//	// 把token存到redis里,设置七天时效
	//	//userId := strconv.FormatInt(int64(user.ID), 10)
	//	//if err := model.Redis.Set(userId, token, 7).Err(); err != nil {
	//	//	c.JSON(http.StatusOK, api.ReturnMsgFunc(400, "redis错误!", 0))
	//	//	return
	//	//} else {
	//	//
	//	//}
	//	// 把用户名,用户权限,角色,token一并传入前端
	//	// 模拟权限数据
	//	var per = []map[string]interface{}{
	//		{"id": 1, "name": "create user"},
	//		{"id": 2, "name": "delete user"},
	//		{"id": 3, "name": "edit user"},
	//		{"id": 4, "name": "select user"},
	//		{"id": 5, "name": "create role"},
	//	}
	//	// [{"id": 1, "name": "create user"},{"id": 2, "name": "delete user"},{"id": 3, "name": "edit user"},{"id": 4, "name": "select user"},{"id": 5, "name": "create role"}]
	//	resData := map[string]interface{}{
	//		"username":   user.Username,
	//		"role":       user.RoleId,
	//		"token":      token,
	//		"permission": per,
	//	}
	//	c.JSON(http.StatusOK, utils.ReturnMsgFunc(200, "success", resData))
	c.JSON(http.StatusOK, gin.H{
		"message": "增加用户!!!!",
	})

}

// UserData 获取用户信息
func UserData(c *gin.Context) {
	//db := config.DBDefault
	// 声明接收的变量
	var token SelectInfo
	//var userListData []model.User
	//var userList []map[string]interface{}
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusOK, utils.ReturnMsgFunc(400, "请携带token!", 0))
		return
	} else {
		//
		//db.Table("users").Select("id", "username", "role_id", "permission").Scan(&userListData)
		//// 循环user信息，将角色跟权限展示出来
		//for _, v := range userListData {
		//	log.Println(v)
		//	r := model.Role{}
		//	p := model.Permission{}
		//	result := db.Where("id = ?", v.RoleId).Select("role_name").Take(&r)
		//	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//		log.Println("找不到记录")
		//		return
		//	}
		//
		//	//db.Where("id = ?", v.Permission).Select("per_name").Take(&p)
		//	user := map[string]interface{}{
		//		"id":          v.ID,
		//		"permissions": p.Name,
		//		"role":        r.Name,
		//		"username":    v.Username,
		//	}
		//	userList = append(userList, user)
		//}
		//resData := map[string]interface{}{
		//	"user_list": userList,
		//}
		//c.JSON(http.StatusOK, utils.ReturnMsgFunc(200, "success", resData))
		// 判断用户是否已经存在
		c.JSON(http.StatusOK, gin.H{
			"message": "增加用户!!!!",
		})
	}

}

// UserAdd 添加用户
func UserAdd(c *gin.Context) {
	var user struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Other      string `json:"other"`
		RoleId     string `json:"role_id"`
		Permission string `json:"permission"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}
	// 判断用户是否已经存在
	c.JSON(http.StatusOK, gin.H{
		"message": "增加用户!!!!",
	})
}

func BlogComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "增加用户!!!!",
	})
}

// PassMd5 md5加密
func PassMd5(str string) (md5str string) {
	pwdData := []byte(str)
	log.Println("data", pwdData)
	pwdMd5 := md5.Sum(pwdData)
	log.Println("pwdMd5", pwdMd5)
	md5str = fmt.Sprintf("%x", pwdMd5)
	return md5str
}

// TokenMd5 md5获取token
func TokenMd5() string {
	curTime := time.Now().Unix()
	log.Println("curTime", curTime)
	h := md5.New()
	log.Println("h-->", h)
	log.Println("strconv.FormatInt(curTime, 10)-->", strconv.FormatInt(curTime, 10))
	io.WriteString(h, strconv.FormatInt(curTime, 10))

	log.Println("h-->", h)

	token := fmt.Sprintf("%x", h.Sum(nil))
	log.Println("token--->", token)
	return token
}
