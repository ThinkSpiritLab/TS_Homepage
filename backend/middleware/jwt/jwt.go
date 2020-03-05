package jwt

import (
	"backend/routers/api/v1/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/pkg/e"
	"backend/pkg/util"
)

func checkTokenValid(token string) (code int){
	code = e.SUCCESS
	if token == "" {
		code = e.ERROR_AUTH_NOT_LOGIN
	} else {
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}
	return code
}

func JWT_contest() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		code := checkTokenValid(token)
		if code == e.SUCCESS && !common.CheckContestPrivilege(token) {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
		if code == e.SUCCESS {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : nil,
			})
			c.Abort()
		}
	}
}

func JWT_teamManage() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		code := checkTokenValid(token)
		if code == e.SUCCESS && !common.CheckTeamManagePrivilege(token) {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
		if code == e.SUCCESS {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : nil,
			})
			c.Abort()
		}
	}
}

func JWT_Console() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		code := checkTokenValid(token)
		if code == e.SUCCESS && !common.CheckConsolePrivilege(token){
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
		if code == e.SUCCESS {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : nil,
			})
			c.Abort()
		}
	}
}