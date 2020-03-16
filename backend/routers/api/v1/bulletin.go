package v1

import (
	"backend/models"
	"backend/pkg/e"
	"backend/routers/api/interfaceDataStruct"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
	"time"
)

func BulletinAdd(c *gin.Context)  {
	code := e.ERROR
	var bulletinAddForm interfaceDataStruct.BulletinAddForm
	err := c.BindJSON(&bulletinAddForm)
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		stid := GetStid(c.Request.Header.Get("Authorization"))
		year, month, day := time.Now().Date()
		createdDate := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
		flag := models.AddBulletin(models.Bulletin{
			BulletinTitle:  bulletinAddForm.BulletinTitle,
			BulletinDetail: bulletinAddForm.BulletinDetail,
			Stid:           stid,
			CreatedAt:      createdDate,
		})
		if flag {
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

func BulletinListBrief(c *gin.Context) {
	var queryRecordRule interfaceDataStruct.RecordRule
	err := json.Unmarshal([]byte(c.Query("conf")), &queryRecordRule)
	code := e.ERROR
	var bulletinListBrief []interfaceDataStruct.BulletinListBrief
	total := 0
	if err == nil {
		records, totalNum := models.GetBulletinBriefList(queryRecordRule)
		total = totalNum
		for _,val := range records {
			name := models.GetNameByStid(val.Stid)
			bulletinListBrief = append(bulletinListBrief, interfaceDataStruct.BulletinListBrief{
				Bid:   val.Bid,
				Title: val.BulletinTitle,
				Date:  val.CreatedAt,
				Promulgator:  val.Stid + " " + name,
			})
		}
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : gin.H{
			"records": bulletinListBrief,
			"total": total,
		},
	})
}

func GetBulletinIndex(c *gin.Context)  {
	code := e.SUCCESS
	data := models.GetBulletinIndex(5)
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func BulletinDelete(c *gin.Context)  {
	bid := com.StrTo(c.Query("bid")).MustInt()
	models.DeleteBulletin(bid)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func BulletinEdit(c *gin.Context) {
	var editBulletinForm interfaceDataStruct.EditBulletinForm
	c.BindJSON(&editBulletinForm)
	models.EditBulletin(editBulletinForm)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func BulletinInfo(c *gin.Context) {
	code := e.SUCCESS
	bid := com.StrTo(c.Query("bid")).MustInt()
	info := models.GetBulletinInfo(bid)
	if info.Bid == 0 {
		code = e.ERROR_BULLETIN_NOT_EXIST
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : info,
	})
}

func BulletinDetail(c *gin.Context) {
	code := e.SUCCESS
	bid := com.StrTo(c.Query("bid")).MustInt()
	info := models.GetBulletinDetail(bid)
	var detail interfaceDataStruct.BulletinDetail
	if info.Bid == 0 {
		code = e.ERROR_BULLETIN_NOT_EXIST
	} else {
		detail.Bid = info.Bid
		detail.CreatedAt = info.CreatedAt
		name := models.GetNameByStid(info.Stid)
		detail.Promulgator = info.Stid + " " + name
		detail.BulletinTitle = info.BulletinTitle
		detail.BulletinDetail = info.BulletinDetail
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : detail,
	})
}