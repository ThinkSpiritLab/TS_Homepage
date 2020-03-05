package models

type User struct {
	Uid int `gorm:"primary_key" json:"uid"`
	Stid string `json:"stid"`
	Name string `json:"name"`
	Grade int `json:"grade"`
	Identity int `json:"identity"`
	Privilege int `json:"privilege"`
}

type UserDetail struct {
	Uid int `gorm:"primary_key" json:"uid"`
	Email string `json:"email"`
	QQ string `gorm:"column:QQ" json:"QQ"`
	URL string `gorm:"column:URL" json:"URL"`
	Mobile string `json:"mobile"`
	Introduction string `gorm:"type:text" json:"introduction"` // []byte
}

func GetUserInfo(uid int) (user User, err bool) {
	db.Where("uid = ?", uid).First(&user)
	err = user.Uid > 0
	return
}

func GetUserDetail(uid int) (userDetail UserDetail, err bool) {
	db.Where("uid = ?", uid).First(&userDetail)
	err = userDetail.Uid > 0
	return
}