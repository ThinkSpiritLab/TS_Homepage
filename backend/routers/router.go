package routers

import (
	"backend/middleware/jwt"
	"backend/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/auth_login", v1.AuthLogin)

		apiV1.GET("/user_info", v1.UserInfo)
		apiV1.GET("/user_info_list", v1.UserInfoList)
		apiV1.GET("/user_detail", v1.UserDetail)
		apiV1.GET("/user_grade_list", v1.UserGradeList)

		apiV1.GET("/console/menu_list", jwt.JWT_console(), v1.ConsoleMenu)


		apiV1.POST("/console/user_regist", jwt.JWT_teamManage(), v1.UserAdd)
		apiV1.DELETE("/console/user_delete", jwt.JWT_teamManage(), v1.UserDelete)
		apiV1.POST("/console/user_edit", jwt.JWT_teamManage(), v1.UserEdit)
		apiV1.POST("/console/user_reset_psw", jwt.JWT_teamManage(), v1.UserResetPsw)

		apiV1.GET("/console/contest_test", jwt.JWT_contest(),func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}
	return r
}