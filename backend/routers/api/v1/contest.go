package v1

import (
	"backend/models"
	"backend/pkg/e"
	"backend/routers/api/interfaceDataStruct"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ContestAdd(c *gin.Context)  {
	code := e.ERROR
	var contestAddForm interfaceDataStruct.ContestAddForm
	err := c.BindJSON(&contestAddForm)
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		fmt.Println(contestAddForm)
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
