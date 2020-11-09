package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	BankName string `gorm:"not null"`
	Number uint64 `gorm:"not null"`
	Balance int `gorm:"not null"`
}

func(a *Account) Create(db *gorm.DB) (*Account, error){
	
	result := db.Create(&a)

	if result.Error == nil {
		return nil,result.Error
	}


	return a, nil
}

func(a *Account) ReadByBankName(db *gorm.DB) (*Account, error){
	
	
	result := db.Where(&Account{BankName: a.BankName}).First(&a)

	if result.Error != nil {
		return nil,result.Error
	}


	return a, nil
}

func(a *Account) ReadAll(db *gorm.DB) ([]Account, error){
	accounts := []Account{}
	result := db.Model(&Account{}).Find(&accounts)

	if result.Error != nil {
		return nil,result.Error
	}


	return accounts, nil
}

func(a *Account) Update(db *gorm.DB, newAccount Account) (*Account, error){
	
	a.BankName = newAccount.BankName
	a.Number = newAccount.Number
	a.Balance = newAccount.Balance

	result := db.Save(&a)

	if result.Error == nil {
		return nil,result.Error
	}


	return a, nil
}

func(a *Account) Delete(db *gorm.DB) (*Account, error){
	
	result := db.Delete(&a)

	if result.Error == nil {
		return nil,result.Error
	}


	return a, nil
}