package persistence

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kill8s/config"
	"log"
)

var globalMysql *gorm.DB

func InitMySQL(config *config.Config) {
	cfg := config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s&parseTime=True&loc=Local", cfg.Username, cfg.Password,
		cfg.Host, cfg.Port, cfg.Database, cfg.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("err init mysql:%v", err)
	}
	globalMysql = db
}

func GetMysql() *gorm.DB {
	return globalMysql
}
