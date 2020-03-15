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

func NameByDimStid(c *gin.Context) {
	dimStid := c.Query("stid")
	code := e.SUCCESS
	lists := models.GetNameByDimStid(dimStid)
	var data []map[string]string
	for _, item := range lists {
		data = append(data, map[string]string{"value": item.Stid, "name": item.Name})
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

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

func UserAllInfo(c *gin.Context)  {
	uid := com.StrTo(c.Query("uid")).MustInt()
	valid := validation.Validation{}
	valid.Required(uid, "uid")
	code := e.INVALID_PARAMS
	var data map[string]interface{}
	if !valid.HasErrors() {
		userBrief, success := models.GetUserInfo(uid)
		if !success {
			code = e.ERROR_NOT_EXIST_USER
		} else {
			userDetail, _ := models.GetUserDetail(uid)
			res := interfaceDataStruct.UserAllInfo{
				Uid:          userBrief.Uid,
				Stid:         userBrief.Stid,
				Name:         userBrief.Name,
				Identity:     userBrief.Identity,
				Privilege:    userBrief.Privilege,
				Email:        userDetail.Email,
				Address:      userDetail.Address,
				QQ:           userDetail.QQ,
				URL:          userDetail.URL,
				Introduction: userDetail.Introduction,
				AvatarUrl:    userDetail.AvatarUrl,
				Education:    userDetail.Education,
				Career:       userDetail.Career,
			}
			userContest := models.GetUserContestByStid(userBrief.Stid)
			var userContestHistory []interfaceDataStruct.UserContestHistory
			for _, val := range userContest {
				contest := models.GetContestByCid(val.Cid)
				userContestHistory = append(userContestHistory, interfaceDataStruct.UserContestHistory{
					Cid:      val.Cid,
					CNameZh:   contest.NameZh,
					CNameEn:   contest.NameEn,
					Ctime:     contest.Ctime,
					AwardType: contest.AwardType,
					TNameZh:   val.NameZh,
					TNameEn:   val.NameEn,
					Star:      val.Star,
					Rank:      val.Rank,
					Awards:    val.Awards,
				})
			}
			data = gin.H {
				"userAllInfo": res,
				"userContest": userContestHistory,
			}
			code = e.SUCCESS
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
			privilege, _ := strconv.Atoi(user["privilege"])
			thisPrivilege := GetPrivilege(c.Request.Header.Get("Authorization"))
			if (privilege==1||privilege==2)&& thisPrivilege !=1 {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				uid, flag := models.AddUser(models.User{Stid: user["stid"], Name: user["name"],
					Grade: user["grade"], Identity: identity, Privilege: privilege})
				if flag {
					userPassword := addForm["user_password"]
					flag = models.AddUserPassword(models.UserPassword{Uid: uid, Psw: userPassword["psw"]})
					if flag {
						userDetail := addForm["user_detail"]
						flag = models.AddUserDetail(models.UserDetail{
							Uid: uid, Email: userDetail["email"], QQ: userDetail["QQ"], URL: userDetail["URL"],
							Address: userDetail["address"], Introduction: userDetail["introduction"]})
						if flag {
							code = e.SUCCESS
						}
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

func UserDelete(c *gin.Context)  {
	uid := com.StrTo(c.Query("uid")).MustInt()
	valid := validation.Validation{}
	valid.Required(uid, "uid")
	code := e.INVALID_PARAMS
	thisPrivilege := GetPrivilege(c.Request.Header.Get("Authorization"))
	deletedPrivilege := models.GetPrivilegeByUid(uid)
	if thisPrivilege >= deletedPrivilege {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else if !valid.HasErrors() {
			models.DeleteUser(uid)
			code = e.SUCCESS
		}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func UserEdit(c *gin.Context)  {
	var user models.User
	c.BindJSON(&user)
	code := e.SUCCESS
	thisPrivilege := GetPrivilege(c.Request.Header.Get("Authorization"))
	editedPrivilege := models.GetPrivilegeByUid(user.Uid)
	if thisPrivilege>=editedPrivilege && editedPrivilege!=user.Privilege {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		models.EditUser(user)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func UserAllEdit(c *gin.Context) {
	var uf interfaceDataStruct.UserAllInfo
	c.BindJSON(&uf)
	code := e.SUCCESS
	thisPrivilege := GetPrivilege(c.Request.Header.Get("Authorization"))
	thisStid := GetStid(c.Request.Header.Get("Authorization"))
	if (thisStid != uf.Stid) && (thisPrivilege != 1 && thisPrivilege != 2 && thisPrivilege != 3) {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		userDetail := models.UserDetail{
			Uid:          uf.Uid,
			Email:        uf.Email,
			QQ:           uf.QQ,
			URL:          uf.URL,
			Address:      uf.Address,
			Introduction: uf.Introduction,
			Education:    uf.Education,
			Career:       uf.Career,
			AvatarUrl:    uf.AvatarUrl,
		}
		models.EditUserDetail(userDetail)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func UserResetPsw(c *gin.Context)  {
	code := e.ERROR
	var tmp map[string]int
	c.BindJSON(&tmp)
	uid := tmp["uid"]
	thisPrivilege := GetPrivilege(c.Request.Header.Get("Authorization"))
	resetPrivilege := models.GetPrivilegeByUid(uid)
	if thisPrivilege >= resetPrivilege {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		models.ResetUserPswByUid(uid)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : nil,
	})
}

func UserMemberFacultyBrief(c *gin.Context)  {
	users := models.GetUserFaculty()
	var data []interfaceDataStruct.MemberListBrief
	for _, val := range users {
		userDetail, _ := models.GetUserDetail(val.Uid)
		data = append(data, interfaceDataStruct.MemberListBrief{
			Uid:       val.Uid,
			Name:      val.Name,
			Stid:      val.Stid,
			Grade:     val.Grade,
			Email:     userDetail.Email,
			Identity:  val.Identity,
			AvatarUrl: userDetail.AvatarUrl,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : data,
	})
}

func UserMemberStudentsBrief(c *gin.Context)  {
	users := models.GetUserStudents()
	var data []interfaceDataStruct.MemberListBrief
	for _, val := range users {
		userDetail, _ := models.GetUserDetail(val.Uid)
		data = append(data, interfaceDataStruct.MemberListBrief{
			Uid:       val.Uid,
			Name:      val.Name,
			Stid:      val.Stid,
			Grade:     val.Grade,
			Email:     userDetail.Email,
			Identity:  val.Identity,
			AvatarUrl: userDetail.AvatarUrl,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : data,
	})
}

func UserMemberAlumniBrief(c *gin.Context)  {
	users := models.GetUserAlumni()
	var data []interfaceDataStruct.MemberListBrief
	for _, val := range users {
		userDetail, _ := models.GetUserDetail(val.Uid)
		data = append(data, interfaceDataStruct.MemberListBrief{
			Uid:       val.Uid,
			Name:      val.Name,
			Stid:      val.Stid,
			Grade:     val.Grade,
			Email:     userDetail.Email,
			Identity:  val.Identity,
			AvatarUrl: userDetail.AvatarUrl,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : e.SUCCESS,
		"msg" : e.GetMsg(e.SUCCESS),
		"data" : data,
	})
}