package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `gorm:"not null"`
	Document string `gorm:"not null"`
	Age int `gorm:"not null"`
	Employment []*Employment `gorm:"many2many:person_employments;"`
	BankAccount []Account `gorm:"foreignKey:ID"`
}

func(p *Person) PersonListToJsonList(personList []Person) (jsonList []*bytes.Buffer,err error){

	for _, person := range personList{
		var jsonPerson *bytes.Buffer

		jsonPerson, err = person.PersonToJson()

		if err != nil{
			log.Fatalf("Erro em passar para o json | Error: %e",err)
			return nil, err
		}

		jsonList = append(jsonList, jsonPerson)
	}

	if err != nil{
		log.Fatalf("Erro em passar para o json | Error: %e",err)
		return nil, err
	}
	
	return jsonList, nil
}

func(p *Person) PersonToJson() (*bytes.Buffer, error){
	
	personMarshal, err := json.Marshal(p)
	if err != nil{
		log.Fatalf("Erro em passar para o json | Error: %e",err)
		return nil, err
	}
	person := bytes.NewBuffer(personMarshal)

	return person, nil
}


func(p *Person) Create(db *gorm.DB) (*Person, error){
	
	result := db.Create(&p)

	if result.Error == nil {
		return nil,result.Error
	}


	return p, nil
}

func(p *Person) ReadByName(db *gorm.DB) (*Person, error){
	
	fmt.Println(p.Name)
	result := db.Where(&Person{Name: p.Name}).First(&p)

	if result.Error != nil {
		return nil,result.Error
	}


	return p, nil
}

func(p *Person) ReadAll(db *gorm.DB) ([]Person, error){
	people := []Person{}
	result := db.Model(&Person{}).Find(&people)

	if result.Error != nil {
		return nil,result.Error
	}


	return people, nil
}

func(p *Person) Update(db *gorm.DB, newPerson Person) (*Person, error){
	
	p.Name = newPerson.Name
	p.Age = newPerson.Age
	p.Document = newPerson.Document
	p.BankAccount = newPerson.BankAccount
	p.Employment = newPerson.Employment

	result := db.Save(&p)

	if result.Error == nil {
		return nil,result.Error
	}


	return p, nil
}

func(p *Person) Delete(db *gorm.DB) (*Person, error){
	
	result :=db.Where(&Person{Name: p.Name}).Delete(&p)

	if result.Error == nil {
		return nil,result.Error
	}


	return p, nil
}