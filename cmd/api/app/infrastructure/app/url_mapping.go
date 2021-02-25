package app

import (
	"github.com/gin-gonic/gin"
	controller "github.com/giojimen3z/CashRegisterApi/cmd/api/app/infrastructure/controller/health"
)

func MapUrls(router *gin.Engine) {

	router.GET("/ping", controller.PingController.Ping)

}
