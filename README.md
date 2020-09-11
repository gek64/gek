# mypkg

## 错误处理

- 文件名：error.go

```go
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
```

## 命令行参数处理(FLAG)

- 文件名：flag.go

```go
// 显示所有的命令行传入参数
// ShowArgs show all args
func ShowArgs() {
    for i, args := range os.Args {
        fmt.Printf("args[%d]=%s\n", i, args)
    }
}

// 判断flag是否被传入
// IsFlagPassed Is Flag Passed
func IsFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}
```
