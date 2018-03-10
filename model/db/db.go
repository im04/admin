package db

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"admin/config"
)
var DB *xorm.Engine
var RedisPool *redis.Pool

func initDB() {
	var err error
	DB, err = xorm.NewEngine(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		panic("error when connect database!,err:" + err.Error())
	}
	DB.SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	DB.SetMaxOpenConns(config.DBConfig.MaxOpenConns)
}

func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	initDB()
	initRedis()
}