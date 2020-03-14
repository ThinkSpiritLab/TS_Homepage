package models

import (
	"backend/routers/api/interfaceDataStruct"
	"fmt"
	"strconv"
)

type Contest struct {
	Cid        int    `gorm:"primary_key" json:"cid"`
	NameZh     string `gorm:"column:name_zh" json:"c_name_zh"`
	NameEn     string `gorm:"column:name_en" json:"c_name_en"`
	LocationZH string `gorm:"column:location_zh" json:"c_location_zh"`
	LocationEN string `gorm:"column:location_en" json:"c_location_en"`
	Ctime      string `gorm:"column:time" json:"c_time"`
	AwardType  string `gorm:"column:award_type" json:"r_type"`
	TotalTeams int    `gorm:"column:total_teams" json:"total_team_num"`
	CExtras    string `gorm:"column:c_extras" json:"total_team_num"`
}

type ContestTeam struct {
	Ctid     int    `gorm:"primary_key" json:"ctid"`
	Cid      int    `gorm:"column:cid" json:"cid"`
	NameZh   string `gorm:"column:name_zh" json:"t_name_zh"`
	NameEn   string `gorm:"column:name_en" json:"t_name_en"`
	Mem1Stid string `gorm:"column:mem1_stid" json:"t_mem1_stid"`
	Mem2Stid string `gorm:"column:mem2_stid" json:"t_mem2_stid"`
	Mem3Stid string `gorm:"column:mem3_stid" json:"t_mem3_stid"`
	Star     string    `gorm:"column:star" json:"t_star"`
	Rank     int    `gorm:"column:rank" json:"t_rank"`
	Awards   string `gorm:"column:awards" json:"t_awards"`
}

func AddContest(contest Contest) (bool, int) {
	db.Create(&contest)
	fmt.Println(contest)
	return contest.Cid>0, contest.Cid
}

func AddContestTeam(contestteam ContestTeam) bool {
	db.Create(&contestteam)
	return contestteam.Ctid>0
}

func GetContestBriefList(queryRecordRule interfaceDataStruct.RecordRule) (contests []Contest, total int) {
	db.Offset(queryRecordRule.Offset).Limit(queryRecordRule.Size).Order("time asc").Find(&contests)
	db.Model(&Contest{}).Count(&total)
	return
}

func GetContestByCid(cid int) (contest Contest) {
	db.Where("cid=?", cid).First(&contest)
	return
}

func GetContestResultBriefByCid(cid int) (result [5]int) {
	for i:=0; i<5; i++ {
		db.Where("cid=? and awards=?", cid, strconv.Itoa(i)).Model(&ContestTeam{}).Count(&result[i])
	}
	return
}

func DeleteContest(cid int) {
	db.Where("cid=?", cid).Delete(&ContestTeam{}).Delete(&Contest{})
}

func GetContestExtras(cid int) string{
	var data []string
	db.Where("cid=?", cid).Model(&Contest{}).Pluck("c_extras", &data)
	return data[0]
}

func EditContestExtras(editInfo interfaceDataStruct.ContestEditForm) {
	db.Model(Contest{}).Where("cid=?", editInfo.Cid).Update("c_extras", editInfo.Extras)
}

func GetUserContestByStid(stid string) (ct []ContestTeam){
	db.Where("mem1_stid=? or mem2_stid=? or mem3_stid=?", stid,stid,stid).Find(&ct)
	return
}