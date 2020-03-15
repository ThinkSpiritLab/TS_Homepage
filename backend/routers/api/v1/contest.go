package v1

import (
	"backend/models"
	"backend/pkg/e"
	"backend/routers/api/interfaceDataStruct"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func ContestAdd(c *gin.Context)  {
	code := e.ERROR
	var contestAddForm interfaceDataStruct.ContestAddForm
	err := c.BindJSON(&contestAddForm)
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		flag1, cid := models.AddContest(models.Contest{
			NameZh:     contestAddForm.BaseInfo.NameZh,
			NameEn:     contestAddForm.BaseInfo.NameEn,
			LocationZH: contestAddForm.BaseInfo.LocationZH,
			LocationEN: contestAddForm.BaseInfo.LocationEN,
			Ctime:      contestAddForm.BaseInfo.Ctime,
			AwardType:  contestAddForm.ResultInfo.AwardType,
			TotalTeams: contestAddForm.ResultInfo.TotalTeams,
			CExtras:    contestAddForm.Extras,
		})
		flag2 := true
		for _, val := range contestAddForm.ResultInfo.Results {
			flag2 = flag2 && models.AddContestTeam(models.ContestTeam{
				Cid:      cid,
				NameZh:   val.NameZh,
				NameEn:   val.NameEn,
				Mem1Stid: val.Mem1Stid,
				Mem2Stid: val.Mem2Stid,
				Mem3Stid: val.Mem3Stid,
				Star:     val.Star,
				Rank:     val.Rank,
				Awards:   val.Awards,
			})
		}
		if flag1 && flag2 {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func ContestListFront(c *gin.Context) {
	var contestListBrief []interfaceDataStruct.ContestListBrief
	records := models.GetContestList()
	for _,val := range records {
		awards := models.GetContestResultBriefByCid(val.Cid)
		contestListBrief = append(contestListBrief, interfaceDataStruct.ContestListBrief{
			Cid: val.Cid, NameZh: val.NameZh, NameEn: val.NameEn, LocationZH: val.LocationZH,
			LocationEN: val.LocationEN, Ctime: val.Ctime, AwardType: val.AwardType,
			Num0: awards[0], Num1: awards[1], Num2: awards[2], Num3: awards[3], Num4: awards[4],
		})
}
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : contestListBrief,
	})
}

func ContestListBrief(c *gin.Context) {
	var queryRecordRule interfaceDataStruct.RecordRule
	err := json.Unmarshal([]byte(c.Query("conf")), &queryRecordRule)
	code := e.ERROR
	var contestListBrief []interfaceDataStruct.ContestListBrief
	total := 0
	if err == nil {
		records, totalNum := models.GetContestBriefList(queryRecordRule)
		total = totalNum
		for _,val := range records {
			awards := models.GetContestResultBriefByCid(val.Cid)
			contestListBrief = append(contestListBrief, interfaceDataStruct.ContestListBrief{
				Cid: val.Cid, NameZh: val.NameZh, NameEn: val.NameEn, LocationZH: val.LocationZH,
				LocationEN: val.LocationEN, Ctime: val.Ctime, AwardType: val.AwardType,
				Num0: awards[0], Num1: awards[1], Num2: awards[2], Num3: awards[3], Num4: awards[4],
			})
		}
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : gin.H{
			"records": contestListBrief,
			"total": total,
		},
	})
}

func ContestDelete(c *gin.Context)  {
	cid := com.StrTo(c.Query("cid")).MustInt()
	models.DeleteContest(cid)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func ContestExtras(c *gin.Context) {
	cid := com.StrTo(c.Query("cid")).MustInt()
	data := models.GetContestExtras(cid)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func ContestExtrasEdit(c *gin.Context) {
	var editInfo interfaceDataStruct.ContestEditForm
	c.BindJSON(&editInfo)
	models.EditContestExtras(editInfo)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func ContestDetail(c *gin.Context) {
	code := e.SUCCESS
	cid := com.StrTo(c.Query("cid")).MustInt()
	contest := models.GetContestByCid(cid)
	var teamsDetail []interfaceDataStruct.ContestTeamDetail
	if contest.Cid <= 0 {
		code = e.ERROR_CONTEST_NOT_EXIST
	} else {
		teams := models.GetContestTeamByCid(cid)
		for _, val := range teams {
			mem1name := models.GetNameByStid(val.Mem1Stid)
			mem2name := models.GetNameByStid(val.Mem2Stid)
			mem3name := models.GetNameByStid(val.Mem3Stid)
			mem1uid := models.GetUidByStid(val.Mem1Stid)
			mem2uid := models.GetUidByStid(val.Mem2Stid)
			mem3uid := models.GetUidByStid(val.Mem3Stid)
			teamsDetail = append(teamsDetail, interfaceDataStruct.ContestTeamDetail{
				NameZh:   val.NameZh,
				NameEn:   val.NameEn,
				Mem1Stid: val.Mem1Stid,
				Mem2Stid: val.Mem2Stid,
				Mem3Stid: val.Mem3Stid,
				Mem1Name: mem1name,
				Mem2Name: mem2name,
				Mem3Name: mem3name,
				Mem1Uid:  mem1uid,
				Mem2Uid:  mem2uid,
				Mem3Uid:  mem3uid,
				Star:     val.Star,
				Rank:     val.Rank,
				Awards:   val.Awards,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : gin.H {
			"contest": contest,
			"teams": teamsDetail,
		},
	})
}