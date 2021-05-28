package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `db:"username" json:"username" gorm:"varchar(128);unique;not null"`
	Password string `db:"password" json:"password" gorm:"varchar(128);not null"`
}
