package v1

import (
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConsoleMenu(c *gin.Context)  {
	token := c.Request.Header.Get("Authorization")
	privilege := GetPrivilege(token)
	code := e.SUCCESS
	var data []map[string]interface{}
	if privilege == 1 || privilege == 2 || privilege == 4 {
		menuContestItem1 := map[string]interface{} {"mid": 121, "mname": "添加记录", "path": "addContest"}
		menuContestItem2 := map[string]interface{} {"mid": 122, "mname": "竞赛列表", "path": "listContest"}
		meunContest := make(map[string]interface{})
		meunContest["mid"] = 120
		meunContest["mname"] = "竞赛历史"
		meunContest["children"] = []map[string]interface{} {menuContestItem1, menuContestItem2}
		data = append(data, meunContest)
	}
	if privilege == 1 || privilege == 2 || privilege == 3 {
		menuUserItem1 := map[string]interface{} {"mid": 111, "mname": "添加成员", "path": "addUser"}
		menuUserItem2 := map[string]interface{} {"mid": 112, "mname": "成员列表", "path": "listUser"}
		meunUser := make(map[string]interface{})
		meunUser["mid"] = 110
		meunUser["mname"] = "成员管理"
		meunUser["children"] = []map[string]interface{} {menuUserItem1, menuUserItem2}
		data = append(data, meunUser)

		menuBulletinItem1 := map[string]interface{} {"mid": 131, "mname": "新增公告", "path": "addBulletin"}
		menuBulletinItem2 := map[string]interface{} {"mid": 132, "mname": "公告列表", "path": "listBulletin"}
		meunBulletin := make(map[string]interface{})
		meunBulletin["mid"] = 130
		meunBulletin["mname"] = "公告管理"
		meunBulletin["children"] = []map[string]interface{} {menuBulletinItem1, menuBulletinItem2}
		data = append(data, meunBulletin)

		menuNewsItem1 := map[string]interface{} {"mid": 141, "mname": "新增新闻", "path": "addNews"}
		menuNewsItem2 := map[string]interface{} {"mid": 142, "mname": "新闻列表", "path": "listNews"}
		meunNews := make(map[string]interface{})
		meunNews["mid"] = 140
		meunNews["mname"] = "新闻管理"
		meunNews["children"] = []map[string]interface{} {menuNewsItem1, menuNewsItem2}
		data = append(data, meunNews)
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
	c.Abort()
}
