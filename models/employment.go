package models

import "gorm.io/gorm"

type Employment struct {
	gorm.Model
	Mode string `gorm:"not null"`
	Company int `gorm:"not null"`
	Person []*Person`gorm:"many2many:person_employments;"`
	Salary float32 `gorm:"not null"`
}

func(e *Employment) Create(db *gorm.DB) (*Employment, error){
	
	result := db.Create(&e)

	if result.Error == nil {
		return nil,result.Error
	}


	return e, nil
}

func(e *Employment) ReadByMode(db *gorm.DB) (*Employment, error){
	
	
	result := db.Where(&Employment{Mode: e.Mode}).First(&e)

	if result.Error != nil {
		return nil,result.Error
	}


	return e, nil
}

func(e *Employment) ReadAll(db *gorm.DB) ([]Employment, error){
	employments := []Employment{}
	result := db.Model(&Employment{}).Find(&employments)

	if result.Error != nil {
		return nil,result.Error
	}


	return employments, nil
}

func(e *Employment) Update(db *gorm.DB, newEmployment Employment) (*Employment, error){
	
	e.Mode = newEmployment.Mode
	e.Company = newEmployment.Company
	e.Person = newEmployment.Person
	e.Salary = newEmployment.Salary

	result := db.Save(&e)

	if result.Error == nil {
		return nil,result.Error
	}


	return e, nil
}

func(e *Employment) Delete(db *gorm.DB) (*Employment, error){
	
	result := db.Delete(&e)

	if result.Error == nil {
		return nil,result.Error
	}


	return e, nil
}