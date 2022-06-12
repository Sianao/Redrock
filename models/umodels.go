package models

type Userinfo struct {
	Uid      int
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"Nickname,omitempty" gorm:"nickname"`
	Sex      string `json:"Sex,omitempty" gorm:"sex"`
	Age      int32  `json:"Age,omitempty" gorm:"age"`
}
