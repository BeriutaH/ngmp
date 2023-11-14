package utils

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
