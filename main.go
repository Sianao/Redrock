package main

import (
	"Redrock/controller"
	"Redrock/dao"
)

func main() {
	dao.SqlInit()
	dao.PoolInitRedis()
	controller.Entrance()
}
