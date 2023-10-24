package service

import (
	"api-server/bean/dto"
	"api-server/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func CreateAccountService(ctx *gin.Context) {
	var account dto.AccountDTO

	// ctx.ShouldBindJSON(&account) 用于将 JSON 数据解析为 account 结构体，并将数据填充到 account 变量中。
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.New()
	createTime := time.Now()
	updateTime := createTime

	account.ID = newUUID.String()
	account.CreateTime = createTime
	account.UpdateTime = updateTime

	model.CreateAccountModel(&account)
}
