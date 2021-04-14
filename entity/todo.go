package entity

import "time"

type Todo struct {
	Id          int64     `xorm:"not null 'id'" json:"id"`
	Id_user     int64     `xorm:"not null 'id_user'" json:"id_user"`
	Name        string    `xorm:"not null 'name'" json:"name"`
	Description string    `xorm:"not null 'description'" json:"description"`
	Created_at  time.Time `xorm:"created" json:"created_at"`
	Updated_at  time.Time `xorm:"updated" json:"updated_at"`
	Deleted_at  time.Time `xorm:"deleted_at" json:"deleted_at"`
}
