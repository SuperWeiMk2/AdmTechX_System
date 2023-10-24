package model

import (
	"api-server/bean/dto"
	"api-server/core/mysqlSetting"
)

func CreateAccountModel(accDTO *dto.AccountDTO) {
	result, _ := mysqlSetting.LinkMySQL()
	err := result.AutoMigrate(dto.AccountDTO{})
	if err != nil {
		return
	}
}

func allAccount() {

}
