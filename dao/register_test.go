package dao

import (
	"Redrock/models"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

var mock sqlmock.Sqlmock
var gormDB *gorm.DB

func init() {
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalln("into sqlmock(mysql) db err:", err)
	}
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	DB = gormDB
	if err != nil {
		log.Fatal("init DB with sqlmock(gorm) fail err:", err)
	}
}

func TestRegister(t *testing.T) {
	var (
		success = "注册成功"
	)
	var u models.Userinfo
	sqlmock.New()
	u.Username = "sams"
	u.Password = "simple2002"

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `userinfos`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	want, err := Register(u)

	if want != success {
		fmt.Println(err)
		t.Errorf("register  err %#v", err)
	}

}
