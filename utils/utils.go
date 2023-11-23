package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
)

// ReturnMsg -----------对需要返回的信息进行封装,方便对数据进行进一步处理
type ReturnMsg struct {
	Status   int         `json:"status"`
	Messages string      `json:"message"`
	Result   interface{} `json:"result"`
}

// ReturnMsgFunc ------------对需要返回的信息进行赋值,并以结构体的形式返回
func ReturnMsgFunc(status int, msg string, result interface{}) *ReturnMsg {
	rm := new(ReturnMsg)
	rm.Status = status
	rm.Messages = msg
	rm.Result = result
	if _, ok := result.(int); ok == true {
		a := map[string]interface{}{}
		rm.Result = a
	} else {
		rm.Result = result
	}
	return rm
}

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
