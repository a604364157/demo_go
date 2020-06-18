package models

type SmsCode struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	Phone      string `xorm:"varchar(14)" json:"phone"`
	Email      string `xorm:"varchar(14)" json:"email"`
	BizId      string `xorm:"varchar(30)" json:"biz_id"`
	Code       string `xorm:"varchar(6)" json:"code"`
	CreateTime int64  `xorm:"bigint" json:"create_time"`
}

func QuerySmsCodeById(id int) (SmsCode, error) {
	var sms SmsCode
	_, err := getDb().ID(id).Get(&sms)
	return sms, err
}

func QuerySmsCode(sms SmsCode) ([]SmsCode, error) {
	session := getDb().Cols("id", "phone", "email", "biz_id", "code", "create_time")
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
	var smses []SmsCode
	err := session.Find(&smses)
	return smses, err
}

func InsertSmsCode(sms SmsCode) error {
	_, err := getDb().Insert(&sms)
	return err
}

func UpdateSmsCode(sms SmsCode) error {
	if sms.Id == 0 {
		panic("修改操作时,主键不能为空")
	}
	_, err := getDb().ID(sms.Id).Update(&sms)
	return err
}

func DeleteSmsCode(id int) error {
	sms := new(SmsCode)
	_, err := getDb().ID(id).Delete(sms)
	return err
}
