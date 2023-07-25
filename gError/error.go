package gError

import (
    "encoding/json"
    "fmt"
)

// Err 包含错误代码与错误信息的结构体
type Err struct {
    Code int
    Msg  string
}

// Error 将Err结构体转换为string输出
func (e *Err) Error() string {
    return fmt.Sprintf("code: %d\nmsg: %s", e.Code, e.Msg)
}

// Json 将Err结构体转换为json输出
func (e *Err) Json() string {
    err, _ := json.Marshal(e)
    return string(err)
}

// New 对当前Err重新赋值
func (e *Err) New(code int, msg string) {
    e.Code = code
    e.Msg = msg
}
