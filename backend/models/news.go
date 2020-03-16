package models

import "backend/routers/api/interfaceDataStruct"

type News struct {
	Nid            int       `gorm:"primary_key" json:"nid"`
	NewsTitle      string
	NewsDetail     string
	Stid           string
	CreatedAt      string
}

func AddNews(news News) bool {
	db.Create(&news)
	return news.Nid>0
}

func GetNewsMoments() (news []News) {
	db.Select("nid, news_title, created_at").Find(&news)
	return
}

func GetNewsBriefList(rule interfaceDataStruct.RecordRule) (news []News, cnt int) {
	db.Omit("detail").Offset(rule.Offset).Limit(rule.Size).Order("created_at desc").Find(&news)
	db.Model(&News{}).Count(&cnt)
	return
}

func DeleteNews(nid int) {
	db.Where("nid=?", nid).Delete(&News{})
}

func GetNewsInfo(nid int) (info interfaceDataStruct.EditNewsForm){
	var news News
	db.Where("nid=?", nid).First(&news)
	info.Nid = news.Nid
	info.NewsDetail = news.NewsDetail
	info.NewsTitle = news.NewsTitle
	info.NewsDate = news.CreatedAt
	return
}

func GetNewsDetail(nid int) (news News){
	db.Where("nid=?", nid).First(&news)
	return
}

func EditNews(editNewsForm interfaceDataStruct.EditNewsForm) {
	db.Model(News{}).Where("nid=?", editNewsForm.Nid).Updates(News{NewsTitle: editNewsForm.NewsTitle, NewsDetail: editNewsForm.NewsDetail, CreatedAt: editNewsForm.NewsDate})
}