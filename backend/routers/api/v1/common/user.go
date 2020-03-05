package common

import (
	"backend/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

	"backend/models"
)

func UserInfo(c *gin.Context)  {
	uid := com.StrTo(c.Query("uid")).MustInt()
	valid := validation.Validation{}
	valid.Required(uid, "uid")
	code := e.INVALID_PARAMS
	var data interface {}
	if !valid.HasErrors() {
		var success bool
		data, success = models.GetUserInfo(uid)
		if success {
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_USER
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func UserDetail(c *gin.Context)  {
	uid := com.StrTo(c.Query("uid")).MustInt()
	valid := validation.Validation{}
	valid.Required(uid, "uid")
	code := e.INVALID_PARAMS
	var data interface {}
	if !valid.HasErrors() {
		var success bool
		data, success = models.GetUserDetail(uid)
		if success {
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_USER
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}