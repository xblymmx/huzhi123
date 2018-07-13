package model

import (
	"time"
	"crypto/sha256"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"errors"
	"github.com/xblymmx/huzhi123/constant"
	"encoding/json"
	"github.com/xblymmx/huzhi123/config"
)

type User struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
	ActivatedAt  *time.Time `json:"activatedAt"`
	UserName     string     `json:"user_name"`
	Password     string     `json:"-"`
	Email        string     `json:"-"`
	Gender       int       `json:"gender"`
	Location     string     `json:"location"`
	Introduce    string     `json:"introduce"`
	PhoneNumber  string     `json:"-"`
	Score        uint       `json:"score"`
	ArticleCount uint       `json:"articleCount"`
	CommentCount uint       `json:"commentCount"`
	CollectCount uint       `json:"collectCount"`
	Signature    string     `json:"signature"`
	Role         int        `json:"role"`
	AvatarURL    string     `json:"avatarURL"`
	CoverURL     string     `json:"coverURL"`
	Status       int        `json:"status"`

	Salt string `json:"-"` // password salt
}


func (user *User) CheckPassword(pwd string) bool {
	b := []byte(pwd + user.Salt)
	h := sha256.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil)) == user.Password
}



// todo: user key
func GetUserFromRedis(uid int) (*User, error) {
	redisConn := RedisPool.Get()
	defer redisConn.Close()

	userKey := fmt.Sprintf("%s:%d", "user", uid)

	userBytes, err := redis.Bytes(redisConn.Do("GET", userKey))
	if err != nil {
		return nil, errors.New(constant.Msg.UserSignedOut)
	}

	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// todo: user key
func CacheUser(user *User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	userKey := fmt.Sprintf("%s:%d", "user", user.ID)

	redisConn := RedisPool.Get()
	defer redisConn.Close()

	_, err = redisConn.Do("SET", userKey, string(userBytes), "EX", config.ServerConfig.UserMaxAge)
	if err != nil {
		return err
	}

	return nil
}




