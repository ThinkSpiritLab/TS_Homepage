package models

import "backend/routers/api/interfaceDataStruct"

type Bulletin struct {
	Bid            int       `gorm:"primary_key" json:"bid"`
	BulletinTitle  string     `json:"title"`
	BulletinDetail string     `json:"detail"`
	Stid           string     `json:"date"`
	CreatedAt      string     `json:"stid"`
}

func AddBulletin(bulletin Bulletin) bool {
	db.Create(&bulletin)
	return bulletin.Bid>0
}

func GetBulletinBriefList(rule interfaceDataStruct.RecordRule) (bulletin []Bulletin, cnt int) {
	db.Omit("detail").Offset(rule.Offset).Limit(rule.Size).Order("created_at desc").Find(&bulletin)
	db.Model(&Bulletin{}).Count(&cnt)
	return
}

func DeleteBulletin(bid int) {
	db.Where("bid=?", bid).Delete(&Bulletin{})
}

func GetBulletinInfo(bid int) (info interfaceDataStruct.EditBulletinForm){
	var bulletin Bulletin
	db.Where("bid=?", bid).First(&bulletin)
	info.Bid = bulletin.Bid
	info.BulletinDetail = bulletin.BulletinDetail
	info.BulletinTitle = bulletin.BulletinTitle
	return
}

func EditBulletin(editBulletinForm interfaceDataStruct.EditBulletinForm) {
	db.Model(Bulletin{}).Where("bid=?", editBulletinForm.Bid).Updates(Bulletin{BulletinTitle: editBulletinForm.BulletinTitle, BulletinDetail: editBulletinForm.BulletinDetail})
}