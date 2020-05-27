package utils

import (
	"crypto/md5"
	"demo_beego/commom/constants"
	"fmt"
	"io"
	"time"
)

const SOLT = "demo_beego"

func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str+SOLT)
	md5Srt := fmt.Sprintf("%x", w.Sum(nil))
	return md5Srt
}

func TimeStampToData(timesTamp int64) string {
	datetime := time.Unix(timesTamp, 0).Format(constants.FORMT_DATE)
	return datetime
}
