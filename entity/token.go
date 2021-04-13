package entity

type Token struct {
	Id      int64  `xorm:"not null 'id'" json:"id"`
	Id_user int64  `xorm:"not null 'id_user'" json:"id_user"`
	Token   string `xorm:"not null 'token'" json:"token"`
}
