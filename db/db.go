package db

import kkbot "github.kissvivi.kkbot"

type BaseDB interface {
	InitDB()                         //初始化链接驱动
	AutoMigrates(dst ...interface{}) //初始化表
	SetConfig(conf *kkbot.Config)
}

func NewBaseDB(t string) BaseDB {
	switch t {
	case "mysql":
		return &MysqlDB{}
	case "redis":
		return &RedisDB{}
	default:
		return &MysqlDB{}
	}
}

func main() {
	cfg, err := kkbot.InitConfig()
	if err != nil {
		panic(err)
	}
	baseDB := NewBaseDB("mysql")
	baseDB.SetConfig(cfg)
	baseDB.InitDB()
	baseDB.AutoMigrates()
}
