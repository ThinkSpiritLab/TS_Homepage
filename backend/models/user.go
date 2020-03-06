package models

type User struct {
	Uid int `gorm:"primary_key" json:"uid"`
	Stid string `json:"stid"`
	Name string `json:"name"`
	Grade string `json:"grade"`
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

func CheckUserExistByUid(uid int) (err bool) {
	var user User
	db.Where("uid = ?", uid).First(user)
	err = user.Uid > 0
	return
}

func CheckUserExistByStid(stid string) (flag bool) {
	var user User
	db.Where("stid = ?", stid).First(&user)
	return user.Uid > 0
}

func GetUserInfo(uid int) (user User, flag bool) {
	db.Where("uid = ?", uid).First(&user)
	flag = user.Uid > 0
	return
}

func GetUserDetail(uid int) (userDetail UserDetail, flag bool) {
	db.Where("uid = ?", uid).First(&userDetail)
	flag = userDetail.Uid > 0
	return
}

func AddUser(user User) (int, bool) {
	db.Create(&user)
	return user.Uid, user.Uid>0
}

func AddUserPassword(upsw UserPassword) bool {
	db.Create(&upsw)
	return upsw.Uid>0
}

func AddUserDetail(udetail UserDetail) bool {
	db.Create(&udetail)
	return udetail.Uid>0
}