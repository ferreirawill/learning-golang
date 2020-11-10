package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string
}

func (u *User) CreateUser() (error){

	pass,err :=bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	
	if err != nil {
		log.Fatalf("Erro ao gerar hash de senha: %v",err)
		return err
	}

	
	u.Password = string(pass)

	return nil

}


