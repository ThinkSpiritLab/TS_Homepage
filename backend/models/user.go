package models

import (
	"backend/routers/api/interfaceDataStruct"
	"time"
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
	Uid          int    `gorm:"primary_key" json:"uid"`
	Email        string `json:"email"`
	QQ           string `gorm:"column:QQ" json:"QQ"`
	URL          string `gorm:"column:URL" json:"URL"`
	Address      string `gorm:"column:address" json:"address"`
	Introduction string `gorm:"type:text" json:"introduction"` // []byte
	Education    string `gorm:"type:text" json:"education"` // []byte
	Career       string `gorm:"type:text" json:"career"` // []byte
	AvatarUrl    string `gorm:"column:avatar_url" json:"avatarUrl"`
}

func GetPrivilegeByUid(uid int) int {
	var tmp []int
	db.Where("uid = ?", uid).Model(&User{}).Pluck("Privilege", &tmp)
	return tmp[0]
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
	DB.Model(&User{}).Count(&total)
	if info.SortRule.Column != "" && info.SortRule.Order != ""{
		DB = DB.Order(info.SortRule.Column + " "+ info.SortRule.Order + ", uid asc")
	} else {
		DB = DB.Order("privilege asc, uid asc")
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

func DeleteUser(uid int) {
	db.Where("uid=?", uid).Delete(&UserPassword{}).Delete(&UserDetail{}).Delete(&User{})
}

func EditUser(user User) {
	db.Model(&user).Omit("uid", "stid").Updates(&user)
}

func EditUserDetail(ud UserDetail)  {
	db.Model(&ud).Omit("uid").Updates(&ud)
}

func ResetUserPswByUid(uid int) {
	var user User
	db.Select("stid, uid").Where("uid=?", uid).First(&user)
	db.Model(&UserPassword{}).Where("uid=?", uid).Update("psw", user.Stid)
}

func GetNameByDimStid(dimStid string) (users []User){
	db.Select("name, stid").Where("stid LIKE ?", dimStid+"%").Find(&users)
	return
}

func GetNameByStid(stid string) string {
	var user User
	db.Select("name").Where("stid=?", stid).First(&user)
	return user.Name
}

func GetUidByStid(stid string) int {
	var user User
	db.Select("uid").Where("stid=?", stid).First(&user)
	return user.Uid
}

func GetUserFaculty() (users []User) {
	db.Where("identity='1'").Find(&users)
	return
}
func GetUserStudents() (users []User) {
	year := time.Now().Year()
	month := time.Now().Month()
	if month < 9 {
		db.Where("identity='2' and grade>=?", year-4).Order("grade asc, privilege asc").Find(&users)
	} else {
		db.Where("identity='2' and grade>?", year-4).Order("grade asc, privilege asc").Find(&users)
	}
	return
}
func GetUserAlumni() (users []User) {
	year := time.Now().Year()
	month := time.Now().Month()
	if month < 9 {
		db.Where("identity='2' and grade<?", year-4).Order("grade desc, privilege asc").Find(&users)
	} else {
		db.Where("identity='2' and grade<=?", year-4).Order("grade desc, privilege asc").Find(&users)
	}
	return
}