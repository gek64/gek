package gMath

import (
    "math/rand"
    "time"
)

func RandStringRunes(n int) string {

    rand.Seed(time.Now().UnixNano())

    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

// 快速产生随机数字符串 https://colobu.com/2018/09/02/generate-random-string-in-Go/
