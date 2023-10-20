package initialization

import (
	"log"
	// "os"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	fmt.Println("正在执行数据库迁移...")

	// 这是你的MySQL数据库连接字符串
	dbURL := fmt.Sprintf(
        "mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        AppConfig.DbUsername,
        AppConfig.DbPassword,
        AppConfig.DbHost,
        AppConfig.DbPort,
        AppConfig.DbDatabase,
    )

	migrationsDir := "file://./migrations"

	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// 执行迁移
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully")
}


func Rollback() {
	fmt.Println("正在执行数据库迁移...")

	// 这是你的MySQL数据库连接字符串
	dbURL := fmt.Sprintf(
        "mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        AppConfig.DbUsername,
        AppConfig.DbPassword,
        AppConfig.DbHost,
        AppConfig.DbPort,
        AppConfig.DbDatabase,
    )
	log.Println("数据库迁移地址: ", dbURL)

	migrationsDir := "file://./migrations"

	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// 执行迁移
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully")
}
