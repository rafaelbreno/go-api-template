package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `db:"username" json:"username" gorm:"varchar(128);unique;not null"`
	Password string `db:"password" json:"password" gorm:"varchar(128);not null"`
}

type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (u *User) EncryptPassword() error {
	var byteEncPwd []byte
	var err error

	if byteEncPwd, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return err
	}

	u.Password = string(byteEncPwd)
	return nil
}

func (u *User) CheckPassword(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
}

// Data transfer object
func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:       u.ID,
		Username: u.Username,
	}
}
