package utils

import (
	"github.com/mojocn/base64Captcha"
	"time"
)

type Captcha struct {
	Id          string
	CaptchaType string
	VerifyValue string
}

// 验证码管理器[内存单体],验证码有效时间2分钟
var store = base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, 2*time.Minute)

// 生成验证码
func GenerateCaptcha(param *Captcha) error {
	var driver base64Captcha.Driver
	switch param.CaptchaType {
	case "audio":
		driver = new(base64Captcha.DriverAudio)
	case "string":
		driver = new(base64Captcha.DriverString).ConvertFonts()
	case "math":
		driver = new(base64Captcha.DriverChinese).ConvertFonts()
	case "chinese":
		driver = new(base64Captcha.DriverMath).ConvertFonts()
	default:
		driver = &base64Captcha.DriverDigit{
			Height: 50,
			Width:  120,
			Length: 6,
			// 歪斜因数
			MaxSkew: 1.0,
			// 干扰线数
			DotCount: 5,
		}
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	param.Id = id
	param.VerifyValue = b64s
	return err
}

// 验证验证码
func VerifyCaptcha(param *Captcha) bool {
	b := store.Verify(param.Id, param.VerifyValue, false)
	if b {
		store.Get(param.Id, true)
	}
	return b
}
