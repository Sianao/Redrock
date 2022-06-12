package main

import (
	"Redrock/api"
	"Redrock/controller"
	"Redrock/dao"
)

func main() {
	dao.SqlInit()
	dao.PoolInitRedis()
	api.Init()
	controller.Entrance()
}
