package models

type Userinfo struct {
	Username string `form:"username," gorm:"username"`
	Nickname string `form:"nickname," gorm:"nickname"`
	Password string `form:"password," gorm:"password"`
	QQClient string `form:"qq_client" gorm:"qqclient"`
}
