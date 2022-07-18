package dao

import "testing"

func TestSqlInit(t *testing.T) {
	SqlInit()

}
func TestPoolInitRedis(t *testing.T) {
	PoolInitRedis()
}
