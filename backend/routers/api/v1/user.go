package v1

import (
	"backend/pkg/e"
	"backend/routers/api/interfaceDataStruct"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"

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

func UserInfoList(c *gin.Context)  {
	var listConf interfaceDataStruct.QueryRecordInfo
	err := json.Unmarshal([]byte(c.Query("conf")), &listConf)
	if err == nil {
		records, total := models.GetUserInfoList(listConf)
		c.JSON(http.StatusOK, gin.H{
			"code" : e.SUCCESS,
			"msg" : e.GetMsg(e.SUCCESS),
			"data" : gin.H{
				"records": records,
				"total": total,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code" : e.ERROR,
			"msg" : e.GetMsg(e.ERROR),
			"data" : nil,
		})
	}
}

func UserGradeList(c *gin.Context) {
	data := models.GetUserGradeList()
	var gradeList []interfaceDataStruct.GradeInterface
	for _, val := range data {
		if val == "" {
			continue
		}
		gradeList = append(gradeList, interfaceDataStruct.GradeInterface{Text: val, Value: val})
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : gradeList,
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

func UserAdd(c *gin.Context)  {
	code := e.ERROR
	var addForm map[string]map[string]string
	err := c.BindJSON(&addForm)
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		user := addForm["user"]
		if models.CheckUserExistByStid(user["stid"]) {
			code = e.ERROR_HAS_EXIST_USER
		} else {
			identity, _ := strconv.Atoi(user["identity"])
			uid, flag := models.AddUser(models.User{Stid: user["stid"], Name: user["name"],
				Grade: user["grade"], Identity: identity, Privilege: 5})
			if flag {
				userPassword := addForm["user_password"]
				flag = models.AddUserPassword(models.UserPassword{Uid: uid, Psw: userPassword["psw"]})
				if flag {
					userDetail := addForm["user_detail"]
					flag = models.AddUserDetail(models.UserDetail{
						Uid: uid, Email: userDetail["email"], QQ: userDetail["QQ"], URL: userDetail["URL"],
						Mobile: userDetail["mobile"], Introduction: userDetail["introduction"]})
					if flag {
						code = e.SUCCESS
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}