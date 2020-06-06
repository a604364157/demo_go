package models

import "fmt"

type SmsCode struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	Phone      string `xorm:"varchar(14)" json:"phone"`
	Email      string `xorm:"varchar(14)" json:"email"`
	BizId      string `xorm:"varchar(30)" json:"biz_id"`
	Code       string `xorm:"varchar(6)" json:"code"`
	CreateTime string `xorm:"bigint" json:"create_time"`
}

func QuerySmsCodeById(id int) SmsCode {
	var sms SmsCode
	_, err := GetDb().ID(id).Get(&sms)
	if err != nil {
		fmt.Println("查询验证码信息失败")
	}
	return sms
}

func QuerySmsCode(sms SmsCode) map[int64]SmsCode {
	session := GetDb().Cols("id", "phone", "email", "biz_id", "code", "create_time")
	if sms.Id != 0 {
		session = session.Where("id = ?", sms.Id)
	}
	if sms.Phone != "" {
		session = session.Where("phone = ?", sms.Phone)
	}
	if sms.Email != "" {
		session = session.Where("email = ?", sms.Email)
	}
	if sms.BizId != "" {
		session = session.Where("biz_id = ?", sms.BizId)
	}
	smses := make(map[int64]SmsCode)
	err := session.Find(&smses)
	if err != nil {
		fmt.Println("查询验证码信息失败")
	}
	return smses
}

func InsertSmsCode(sms SmsCode) {
	_, err := GetDb().Insert(&sms)
	if err != nil {
		fmt.Println("写入验证码信息失败")
	}
}

func UpdateSmsCode(sms SmsCode) {
	if sms.Id == 0 {
		fmt.Println("修改数据时,主键不能为空")
		return
	}
	_, err := GetDb().ID(sms.Id).Update(&sms)
	if err != nil {
		fmt.Println("修改验证码信息失败")
	}
}

func DeleteSmsCode(id int) {
	sms := new(SmsCode)
	_, err := GetDb().ID(id).Delete(sms)
	if err != nil {
		fmt.Println("删除验证码信息失败")
	}
}
