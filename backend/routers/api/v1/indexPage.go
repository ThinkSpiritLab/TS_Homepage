package v1

import (
	"backend/models"
	"backend/pkg/e"
	"backend/routers/api/interfaceDataStruct"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func GetMoments(c *gin.Context) {
	contests := models.GetContestMoments()
	news := models.GetNewsMoments()
	var moments []interfaceDataStruct.Moment
	for _, val := range contests {
		moments = append(moments, interfaceDataStruct.Moment{
			ID:    val.Cid,
			Type:  "c",
			Title: val.NameZh,
			Date:  val.Ctime,
		})
	}
	for _, val := range news {
		moments = append(moments, interfaceDataStruct.Moment{
			ID:    val.Nid,
			Type:  "n",
			Title: val.NewsTitle,
			Date:  val.CreatedAt,
		})
	}
	sort.Slice(moments, func(i, j int) bool {
		return moments[i].Date > moments[j].Date
	})
	data := moments[0:7]
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
