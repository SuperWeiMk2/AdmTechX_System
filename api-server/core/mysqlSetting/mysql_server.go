package mysqlSetting

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlSetting struct {
	Username string
	Password string
	Host     string
	Port     int
	DB       string
	Timeout  string
}

func LinkMySQL() (*gorm.DB, error) {
	// 根据 mysqlSetting 结构，初始化一个 linkPrepare 变量:
	linkPrepare := mysqlSetting{
		Username: "admin",
		Password: "admin1234",
		Host:     "192.168.10.130",
		Port:     3306,
		DB:       "test",
		Timeout:  "30s",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s",
		linkPrepare.Username,
		linkPrepare.Password,
		linkPrepare.Host,
		linkPrepare.Port,
		linkPrepare.DB,
		linkPrepare.Timeout,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败！" + err.Error())
	}
	fmt.Print("连接成功！")

	return db, nil
	//// 执行 "SHOW DATABASES" 查询
	//var databases []struct {
	//	Database string `gorm:"column:Database"`
	//}
	//
	//if err := db.Raw("SHOW DATABASES").Scan(&databases).Error; err != nil {
	//	panic("查询数据库列表失败: " + err.Error())
	//}
	//
	//// 打印查询结果
	//for _, db := range databases {
	//	fmt.Println(db.Database)
	//}
}
