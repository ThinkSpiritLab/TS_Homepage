package models

import (
	"backend/routers/api/interfaceDataStruct"
)

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

func GetUserInfoList(info interfaceDataStruct.QueryRecordInfo) (users []User, total int) {
	DB := db
	if len(info.FilterRule.Identity) != 0 {
		DB = DB.Where("identity in (?)", info.FilterRule.Identity)
	}
	if len(info.FilterRule.Privilege) != 0 {
		DB = DB.Where("privilege in (?)", info.FilterRule.Privilege)
	}
	if len(info.FilterRule.Grade) != 0 {
		DB = DB.Where("grade in (?)", info.FilterRule.Grade)
	}
	if len(info.RecordRule.Keyword) != 0 {
		DB = DB.Where("name LIKE ? or stid LIKE ?", "%"+info.RecordRule.Keyword+"%", "%"+info.RecordRule.Keyword+"%")
	}
	DB.Find(&users).Count(&total)
	if info.SortRule.Column != "" && info.SortRule.Order != ""{
		DB = DB.Order(info.SortRule.Column + " "+ info.SortRule.Order)
	} else {
		DB = DB.Order("privilege asc")
	}
	if info.RecordRule.Offset != 0 {
		DB = DB.Offset(info.RecordRule.Offset)
	}
	if info.RecordRule.Size != 0 {
		DB = DB.Limit(info.RecordRule.Size)
	}
	DB.Find(&users)
	return
}

func GetUserGradeList() (gradeList []string){
	db.Model(&User{}).Group("grade").Order("grade asc").Pluck("grade", &gradeList)
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