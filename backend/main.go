package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/models"
	"backend/pkg/setting"
	"backend/pkg/util"
	"backend/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description TS_system backend

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}