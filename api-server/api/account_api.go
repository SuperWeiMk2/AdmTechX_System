package api

import (
	"api-server/service"
	"github.com/gin-gonic/gin"
)

func GetAllAccount() {

}

func createAccount(ctx *gin.Context) {
	service.CreateAccountService(ctx)
}

func SetupRouter() {
	api := gin.Default().Group("/api")

	api.POST("/create", createAccount)
}
