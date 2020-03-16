package v1

import (
	"backend/models"
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexImgList(c *gin.Context) {
	data := models.GetIndexImgList()
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : data,
	})
}

func SetIndexImgList(c *gin.Context) {
	var imgs []models.IndexImg
	_ = c.BindJSON(&imgs)
	models.SetIndexImgList(imgs)
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : nil,
	})
}