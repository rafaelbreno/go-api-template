package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}
