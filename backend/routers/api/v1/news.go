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

func NewsAdd(c *gin.Context)  {
	code := e.ERROR
	var newsAddForm interfaceDataStruct.NewsAddForm
	err := c.BindJSON(&newsAddForm)
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		stid := GetStid(c.Request.Header.Get("Authorization"))
		if newsAddForm.NewsDate == "" {
			year, month, day := time.Now().Date()
			s_month := strconv.Itoa(int(month))
			if len(s_month) == 1 {
				s_month = "0" + s_month
			}
			newsAddForm.NewsDate = strconv.Itoa(year) + "-" + s_month + "-" + strconv.Itoa(day)
		}
		flag := models.AddNews(models.News{
			NewsTitle:  newsAddForm.NewsTitle,
			NewsDetail: newsAddForm.NewsDetail,
			Stid:           stid,
			CreatedAt:      newsAddForm.NewsDate,
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

func NewsListBrief(c *gin.Context) {
	var queryRecordRule interfaceDataStruct.RecordRule
	err := json.Unmarshal([]byte(c.Query("conf")), &queryRecordRule)
	code := e.ERROR
	var newsListBrief []interfaceDataStruct.NewsListBrief
	total := 0
	if err == nil {
		records, totalNum := models.GetNewsBriefList(queryRecordRule)
		total = totalNum
		for _,val := range records {
			name := models.GetNameByStid(val.Stid)
			newsListBrief = append(newsListBrief, interfaceDataStruct.NewsListBrief{
				Nid:   val.Nid,
				Title: val.NewsTitle,
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
			"records": newsListBrief,
			"total": total,
		},
	})
}

func NewsDelete(c *gin.Context)  {
	nid := com.StrTo(c.Query("nid")).MustInt()
	models.DeleteNews(nid)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func NewsEdit(c *gin.Context) {
	var editNewsForm interfaceDataStruct.EditNewsForm
	c.BindJSON(&editNewsForm)
	if editNewsForm.NewsDate == "" {
		year, month, day := time.Now().Date()
		s_month := strconv.Itoa(int(month))
		if len(s_month) == 1 {
			s_month = "0" + s_month
		}
		editNewsForm.NewsDate = strconv.Itoa(year) + "-" + s_month + "-" + strconv.Itoa(day)
	}
	models.EditNews(editNewsForm)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func NewsInfo(c *gin.Context) {
	code := e.SUCCESS
	nid := com.StrTo(c.Query("nid")).MustInt()
	info := models.GetNewsInfo(nid)
	if info.Nid == 0 {
		code = e.ERROR_NEWS_NOT_EXIST
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : info,
	})
}

func NewsDetail(c *gin.Context) {
	code := e.SUCCESS
	nid := com.StrTo(c.Query("nid")).MustInt()
	info := models.GetNewsDetail(nid)
	var detail interfaceDataStruct.NewsDetail
	if info.Nid == 0 {
		code = e.ERROR_BULLETIN_NOT_EXIST
	} else {
		detail.Nid = info.Nid
		detail.CreatedAt = info.CreatedAt
		name := models.GetNameByStid(info.Stid)
		detail.Promulgator = info.Stid + " " + name
		detail.NewsTitle = info.NewsTitle
		detail.NewsDetail = info.NewsDetail
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : detail,
	})
}