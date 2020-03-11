package routers

import (
	"backend/middleware/jwt"
	"backend/pkg/setting"
	"backend/routers/api/file"
	"backend/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/public", http.Dir(setting.AppSetting.PublicDIR))

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/auth_login", v1.AuthLogin)

		apiV1.POST("/upload_image",jwt.JWT_TS_member(), file.UploadImage)

		apiV1.GET("/user_info", v1.UserInfo)
		apiV1.GET("/user_info_list", v1.UserInfoList)
		apiV1.GET("/user_detail", v1.UserDetail)
		apiV1.GET("/user_grade_list", v1.UserGradeList)
		apiV1.GET("/user_name_by_dim_stid", v1.NameByDimStid)
		apiV1.GET("/contest_list_brief", v1.ContestListBrief)
		apiV1.GET("/contest_extras", v1.ContestExtras)

		apiV1.GET("/console/menu_list", jwt.JWT_console(), v1.ConsoleMenu)

		apiV1.POST("/console/user_regist", jwt.JWT_teamManage(), v1.UserAdd)
		apiV1.DELETE("/console/user_delete", jwt.JWT_teamManage(), v1.UserDelete)
		apiV1.POST("/console/user_edit", jwt.JWT_teamManage(), v1.UserEdit)
		apiV1.POST("/console/user_reset_psw", jwt.JWT_teamManage(), v1.UserResetPsw)

		apiV1.POST("/contest/contest_add", jwt.JWT_contest(), v1.ContestAdd)
		apiV1.DELETE("/contest/contest_delete", jwt.JWT_contest(), v1.ContestDelete)
		apiV1.POST("/contest/contest_extras_edit", jwt.JWT_contest(), v1.ContestExtrasEdit)
	}
	return r
}