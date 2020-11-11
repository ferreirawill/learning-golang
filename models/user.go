package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	uuid string
	Name string `json:"name"`
	Username string `gorm:"type:varchar(100);unique_index" json:"Username"`
	Password string
}

func (u *User) CreateUser(db *gorm.DB) (error){

	pass,err :=bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	
	if err != nil {
		log.Fatalf("Erro ao gerar hash de senha: %v",err)
		return err
	}

	
	u.Password = string(pass)

	return nil

}


