package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"math/rand"
	"regexp"
	"time"
)

const SALT = "demo_gin"

// md5加密
func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str+SALT)
	md5Srt := fmt.Sprintf("%x", w.Sum(nil))
	return md5Srt
}

// 匹配电子邮箱
func RegexpEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// RandString 生成随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// 生成UUID
func GetUUID() string {
	u2 := uuid.NewV4()
	return u2.String()
}
