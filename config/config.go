package config

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"github.com/xblymmx/huzhi123/utils"
)


var jsonData map[string]interface{}

type serverConfig struct {
	ENV string
	APIPrefix string

	// model
	UserMaxAge int64

	// image upload
	UploadImgDir string
	ImgHost string
	ImgPath string
}

var ServerConfig serverConfig

func initServer() {
	serverJson, ok := jsonData["server"].(map[string]interface{})
	if !ok {
		fmt.Println(`json["server"] should be sub-json`)
		os.Exit(-1)
	}

	err := utils.SetStructByJSON(&ServerConfig, serverJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}


func initJSON() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("reading json config file failed")
		os.Exit(-1)
	}

	err = json.Unmarshal(b, &jsonData)
	if err != nil {
		fmt.Println("invalid json file")
		os.Exit(-1)
	}
}


type dbConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConn int
	MaxOpenConn int
}

var DBConfig dbConfig

func initDB() {
	dbJson, ok := jsonData["database"].(map[string]interface{})
	if !ok {
		fmt.Println(`json["database"] should be sub-json`)
		os.Exit(-1)
	}

	err := utils.SetStructByJSON(&DBConfig, dbJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}

type redisConfig struct {
	Host string
	Port int
	Password string
	URL string
	MaxIdle int
	MaxActive int
}

var RedisConfig redisConfig

func initRedis() {
	redisJson, ok := jsonData["redis"].(map[string]interface{})
	if !ok {
		fmt.Println(`json["redis"] should be sub-json`)
		os.Exit(-1)
	}

	err := utils.SetStructByJSON(&RedisConfig, redisJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	url := fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	RedisConfig.URL = url
}

func init() {
	initJSON()
	initDB()
	//initRedis()
	initServer()
}