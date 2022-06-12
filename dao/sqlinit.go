package dao

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

func SqlInit() {
	db, err := gorm.Open(mysql.Open("sia:sianao2002@tcp(110.42.216.125:3306)/Redrock?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		log.Println(err)
	}
	DB = db
}

var RedisPool *redis.Pool

func PoolInitRedis() *redis.Pool {
	server := "110.42.216.125:6379"
	password := ""
	redisPool := &redis.Pool{
		MaxIdle:     4, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   10, //最大数
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
	RedisPool = redisPool
	return RedisPool
}
