package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

const SOLT = "demo_beego"

func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str+SOLT)
	md5Srt := fmt.Sprintf("%x", w.Sum(nil))
	return md5Srt
}
