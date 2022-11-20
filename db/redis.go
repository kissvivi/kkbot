package db

import kkbot "github.kissvivi.kkbot"

type RedisDB struct {
	Url      string
	UserName string
	Password string
	DBName   string
}

func (r RedisDB) InitDB() {
	panic("implement me")
}

func (r RedisDB) AutoMigrates(dst ...interface{}) {
	//panic("implement me")
	//无
}

func (r RedisDB) SetConfig(conf *kkbot.Config) {
	panic("implement me")
}

