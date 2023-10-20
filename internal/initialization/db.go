package initialization

import (
	"fmt"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabaseConnection() {
	fmt.Println("正在初始化数据库连接...")

	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.DbUsername,
		AppConfig.DbPassword,
		AppConfig.DbHost,
		AppConfig.DbPort,
		AppConfig.DbDatabase,
	)
	fmt.Println("数据库连接地址是：", dsn)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}