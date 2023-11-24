package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math/big"
	"ngmp/utils/response"
	"strings"
)

// RemoveFields 结构体列表删除指定的键
func RemoveFields(dataList interface{}, omitFields ...string) ([]map[string]interface{}, error) {
	var jsonList []map[string]interface{}

	// 将结构体转为map
	jsonData, err := json.Marshal(dataList)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	}
	log.Println("jsonData--------", jsonData)
	// 将 JSON 数据解码为 map
	var structMapList []map[string]interface{}
	err = json.Unmarshal(jsonData, &structMapList)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	log.Println("structMapList-----", structMapList)
	// 删除指定的键
	for _, structMap := range structMapList {
		for _, key := range omitFields {
			delete(structMap, key)
			jsonList = append(jsonList, structMap)
		}
	}

	return jsonList, nil
}

func GenerateRandomString() string {
	// 定义包含的字符集合
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-=_+[]{}|;:'\",.<>/?`~"

	// 生成随机字符串
	var result strings.Builder
	for i := 0; i < 32; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		result.WriteByte(charSet[randomIndex.Int64()])
	}

	return result.String()
}

func GetMapValue(dict map[string]interface{}, key string, defaultValue int) interface{} {
	if value, ok := dict[key]; ok {
		return value
	}
	return defaultValue
}

// ErrorHandling 统一错误处理
func ErrorHandling(errMsg string, err error, c *gin.Context, tx *gorm.DB) {
	if err != nil {
		_ = tx.Rollback()
		response.InvalidArgumentJSON(errMsg+err.Error(), c)
		return
	}
	return
}
