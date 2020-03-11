package file

import (
	"backend/pkg/e"
	"backend/pkg/setting"
	"backend/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	file, header, err := c.Request.FormFile("upfile")
	data := ""
	if err != nil {
		code = e.ERROR_UPLOAD_IMAGE_FAIL
	} else {
		if header.Size > int64(setting.AppSetting.ImageMaxSize)*int64(1048576) {
			code = e.ERROR_UPLOAD_IMAGE_TOO_LARGE
		}
		index := -1
		if strings.Contains(header.Filename, ".jpg") {
			index = 0
		} else if strings.Contains(header.Filename, ".jpeg") {
			index = 1
		} else if strings.Contains(header.Filename, ".png") {
			index = 2
		}
		if index == -1 {
			code = e.ERROR_UPLOAD_IMAGE_WRONG_TYPE
		} else {
			filename := util.EncodeMD5(header.Filename+time.Now().String()) + setting.AppSetting.ImageAllowExts[index]
			fmt.Println(setting.AppSetting.ImageSavePath)
			fmt.Println(setting.AppSetting.PublicDIR)
			out, err := os.Create(setting.AppSetting.ImageSavePath + "/" + filename)
			if err != nil {
				code = e.ERROR_UPLOAD_IMAGE_FAIL
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				code = e.ERROR_UPLOAD_IMAGE_FAIL
			} else {
				data = setting.AppSetting.ImagePrefixUrl + filename
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
