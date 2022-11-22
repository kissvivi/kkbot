package db

import (
	"fmt"
	kkbot "github.kissvivi.kkbot"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConn *gorm.DB

type MysqlDB struct {
	DB *gorm.DB
}

func (m MysqlDB) InitDB(conf *kkbot.Config) {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.DataSource.Mysql.Username, conf.DataSource.Mysql.Password, conf.DataSource.Mysql.Url, conf.DataSource.Mysql.Dbname), // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	m.DB = db
	MysqlConn = db
}

var _ BaseDB = (*MysqlDB)(nil)
