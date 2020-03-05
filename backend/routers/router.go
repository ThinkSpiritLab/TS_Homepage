package routers

import (
	"backend/middleware/jwt"
	"backend/routers/api/v1/common"
	"backend/routers/api/v1/console"
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

	apiV1 := r.Group("/api/v1/common")
	{
		apiV1.POST("/auth_login", common.AuthLogin)

		apiV1.GET("/user_info", common.UserInfo)
		apiV1.GET("/user_detail", common.UserDetail)
	}

	r.GET("/api/v1/console/menu_list", jwt.JWT_Console(), console.ConsoleMenu)

	apiV1Contest := r.Group("/api/v1/contest")
	apiV1Contest.Use(jwt.JWT_contest())
	{
		apiV1Contest.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}

	apiV1Teammanage := r.Group("/api/v1/teammanage")
	apiV1Teammanage.Use(jwt.JWT_teamManage())
	{

	}

	return r
}