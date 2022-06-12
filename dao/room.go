package dao

import "github.com/gomodule/redigo/redis"

func NewRoom(rid int, nick string) error {
	con := RedisPool.Get()
	defer con.Close()
	_, err := redis.String(con.Do("sadd", rid, nick))

	if err != nil {
		return err
	}
	return nil

}
