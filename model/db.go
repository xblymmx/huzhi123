package model

import (
	"github.com/garyburd/redigo/redis"
	//"github.com/globalsign/mgo"
	"github.com/jinzhu/gorm"
	"github.com/xblymmx/huzhi123/config"
	"fmt"
	"os"
	"time"
)

// todo
var DB *gorm.DB

var RedisPool *redis.Pool

func initDB() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if config.ServerConfig.ENV == "dev" {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConn)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConn)

	DB = db
}

func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle: config.RedisConfig.MaxIdle,
		MaxActive: config.RedisConfig.MaxActive,
		IdleTimeout: 120*time.Second,
		Wait: true,
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL, redis.DialPassword(config.RedisConfig.Password))
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

