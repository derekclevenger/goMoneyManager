package models

import (
	"github.com/derekclevenger/accountapi/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	UserID uint `json:"user_id"`
}

func (account *Account) Validate() (map[string] interface{}, bool) {
	if len(account.Type) <= 0 {
		return utils.Message( false,"Must have an account type"), false
	}

	if account.UserID <= 0 {
		return utils.Message(false, "User not recognized"), false
	}

	return utils.Message(true, "success"), true

}

func (account *Account) Create() (map[string] interface{}) {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(account).Error
	if err != nil {
		return utils.Message(false, "Error creating account")
	}

	resp := utils.Message(true, "Success")
	resp["account"] = account
	return resp

}

func (account *Account) Update(id uint) (map[string] interface{}) {
	acc := &Account{}

	err := GetDB().Table("accounts").Where("ID = ?", id).First(acc).Error
	if err != nil {
		return utils.Message(false, "Error retrieving account")
	}

	if resp, ok := account.Validate(); !ok {
		return resp
	}
	if account.Type != "" && account.Type != acc.Type {
		acc.Type = account.Type
	}
	if account.Amount != acc.Amount {
		acc.Amount = account.Amount
	}

	err = GetDB().Table("accounts").Where("ID = ?", id).Update(acc).Error
	if err != nil {
		fmt.Println(err)
		return utils.Message(false, "Error updating record")
	}
	resp := utils.Message(true, "Updated Successfully")
	resp["account"] = acc
	fmt.Println(acc)
	return resp
}

func Delete(id uint) (map[string] interface{}) {
	account := &Account{}
	err := GetDB().Table("accounts").Where("ID = ?", id).Delete(account).Error
	if err != nil {
		fmt.Println(err)
		return utils.Message(false, "Error deleting record")
	}

	return utils.Message(true, "Success")
}
//TODO fix below to match the rest of api
func GetAll(user uint) (map[string]interface{}) {
	accounts := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("user_id = ?", user).Find(&accounts).Error
	if err != nil {
		resp := utils.Message(false, "Error retrieving record")
		resp["accounts"] = accounts
	}
	resp := utils.Message(true, "Success")
	resp["accounts"] = accounts

	return resp
}

func Get(id uint) (map[string]interface{}) {
	account := &Account{}
	err := GetDB().Table("accounts").Where("ID = ?", id).First(account).Error
	if err != nil {
		fmt.Println(err)
		resp := utils.Message(false, "Error retrieving record")
		resp["account"] = account
		return resp
	}
	resp := utils.Message(true, "Success")
	resp["account"] = account
	return resp
}