package model

type Books struct {
	Id       int    `gorm:"type:int;primary_key"`
	Judul    string `gorm:"type:varchar(255)"`
	Penerbit string `gorm:"type:varchar(255)"`
	Rating   int    `gorm:"type:int"`
}
