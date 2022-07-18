package dao

import (
	"Redrock/models"
	"testing"
)

func TestLogin(t *testing.T) {

	rows := mock.NewRows(
		[]string{`uid`, `username`, `password`, `nickname`, `sex`, `age`}).
		AddRow("1", "sia", "simple2002", "siam", "man", 12)
	mock.ExpectQuery("^SELECT \\* FROM `userinfos` WHERE username=\\?").
		WithArgs("sia").
		WillReturnRows(rows)
	mock.ExpectQuery("^SELECT \\* FROM `userinfos` WHERE username=\\?").WithArgs("saimf").WillReturnRows()
	var u models.Userinfo
	u.Username = "sia"
	u.Password = "simple2002"
	_, err := Login(u)
	if err != nil {
		t.Errorf("err %#v", err)
	}

	u.Password = "saimf"
	_, err = Login(u)

}
