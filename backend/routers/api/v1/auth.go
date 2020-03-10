package v1

import (
	"backend/models"
	"backend/routers/api/interfaceDataStruct"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"backend/pkg/e"
	"backend/pkg/util"

)

func AuthLogin(c *gin.Context) {
	var auth interfaceDataStruct.Auth
	err := c.BindJSON(&auth)
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if err == nil {
		valid := validation.Validation{}
		ok, _ := valid.Valid(&auth)
		if ok {
			var user models.User
			isCorrect, uid:= models.CheckAuth(auth.Stid, auth.Psw)
			if isCorrect {
				user, _ = models.GetUserInfo(uid)
				data["userinfo"] = user
				data["token"], err = util.GenerateToken(auth.Stid, user.Privilege)
				if err != nil {
					data = nil
					code = e.ERROR_AUTH_OTHER_ERROR
				} else {
					code = e.SUCCESS
				}
			} else {
				code = e.ERROR_AUTH_STID_PSW
			}
		} else {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// privilege: 1: root     2: admin     3: team manage admin     4: contest admin     5: user
func GetPrivilege(token string) int {
	privilege, err := util.ParseToken(token)
	if err == nil {
		return privilege.Privilege
	} else {
		return -1
	}
}

func CheckTSMember(token string) bool {
	privilege := GetPrivilege(token)
	if privilege == 1 || privilege == 2 || privilege == 3 || privilege == 4 || privilege == 5 {
		return true
	} else {
		return false
	}
}

func CheckContestPrivilege(token string) bool {
	privilege := GetPrivilege(token)
	if privilege == 1 || privilege == 2 || privilege == 4 {
		return true
	} else {
		return false
	}
}

func CheckTeamManagePrivilege(token string) bool {
	privilege := GetPrivilege(token)
	if privilege == 1 || privilege == 2 || privilege == 3 {
		return true
	} else {
		return false
	}
}

func CheckConsolePrivilege(token string) bool {
	privilege := GetPrivilege(token)
	if privilege == 1 || privilege == 2 || privilege == 3 || privilege == 4 {
		return true
	} else {
		return false
	}
}

func CheckAdminPrivilege(token string) bool {
	privilege := GetPrivilege(token)
	if privilege == 1 || privilege == 2 {
		return true
	} else {
		return false
	}
}