package models

type UserPassword struct {
	Uid int `gorm:"primary_key" json:"uid"`
	//Name string `json:"name"`
	Psw string `json:"password"`
}

func CheckAuth(Stid, password string) (success bool, uid int) {
	var user User
	db.Select("uid").Where("Stid=?", Stid).First(&user)
	var auth UserPassword
	db.Select("uid").Where(UserPassword{Uid : user.Uid, Psw : password}).First(&auth)
	if auth.Uid > 0 {
		return true, user.Uid
	}
	return false, user.Uid
}
