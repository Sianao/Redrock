package dao

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func NewRoom(rid int, nick string) error {
	con := RedisPool.Get()
	defer con.Close()
	do, err := con.Do("SADD", rid, nick)
	if err != nil {
		log.Println(err)

	}
	_, err = redis.String(do, err)
	return err

}
