package models

import (
	"fmt"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Payee string `json:"payee"`
	TransactionDate time.Time `json:"transactionDate,string"`
	Amount float64 `json:"amount"`
	AccountType string `json:"accountType"`
	UserId uint `json:"user_id"`
	Category string `json:"category"`
}


func(transaction *Transaction) Validate() (map[string]interface{}, bool) {
	if transaction.Payee == "" {
		return utils.Message(false, "Payee required"), false
	}

	if transaction.TransactionDate.IsZero()  {
		return utils.Message(false, "Date is required"), false
	}

	if transaction.Amount == 0 {
		return utils.Message(false, "Amount is required"), false
	}

	if transaction.AccountType == "" {
		return utils.Message(false, "Account type required"), false
	}

	if transaction.Category == "" {
		return utils.Message(false, "Category required"), false
	}

	return utils.Message(true, "Success"), true

}

func(transaction *Transaction) Create() (map[string]interface{}) {
	if resp, ok := transaction.Validate(); !ok {
		return resp
	}

	err := GetDB().Table("transactions").Create(transaction).Error
	if err != nil {
		return utils.Message(false, "Error creating transaction")
	}

	resp := utils.Message(true, "Successfully created")
	resp["transaction"] = transaction
	return resp
}

func(transaction *Transaction) GetAll(user uint) (map[string]interface{}) {
	transactions := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("user_id = ?", user).Find(&transactions).Error
	if err != nil {
		fmt.Println(transactions)
		return utils.Message(false, "Error getting records")
	}

	resp := utils.Message(true, "Success")
	resp["transactions"] = transactions
	return resp
}

func(transaction *Transaction) GetTransactionsByCategory(user uint, category string) (map[string]interface{}) {
	transactions := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("user_id = ? AND category = ?", user, category).Find(&transactions).Error
	if err != nil {
		return utils.Message(false, "Error getting records")
	}

	resp := utils.Message(true, "Success")
	resp["transactions"] = transactions
	return resp
}

func(transaction *Transaction) Get(id uint) (map[string]interface{}) {
	t := &Transaction{}
	err := GetDB().Table("transactions").Where("ID = ?", id).First(t).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return utils.Message(false, "Error getting record")
	}
	resp := utils.Message(true, "Success")
	resp["transaction"] = t
	return resp
}

func(transaction *Transaction) Update(id uint) (map[string]interface{}) {
	if resp, ok := transaction.Validate(); ok{
		return resp
	}
	temp := &Transaction{}
	err := GetDB().Table("transactions").Where("ID = ?", id).First(temp).Error
	if err != nil {
		return utils.Message(false, "Error retrieving record to update")
	}

	if transaction.Category != temp.Category && transaction.Category != "" {
		temp.Category = transaction.Category
	}

	if transaction.AccountType != temp.AccountType && transaction.AccountType != "" {
		temp.AccountType = transaction.AccountType
	}

	if transaction.TransactionDate != temp.TransactionDate && !transaction.TransactionDate.IsZero() {
		temp.TransactionDate = transaction.TransactionDate
	}

	if transaction.Amount != temp.Amount && transaction.Amount > 0.00 {
		temp.Amount = transaction.Amount
	}

	if transaction.Payee != temp.Payee && transaction.Payee != "" {
		temp.Payee = transaction.Payee
	}

	err = GetDB().Table("transactions").Where("ID = ?", id).Update(temp).Error
	if err != nil {
		return utils.Message(false, "Error updating record")
	}

	resp := utils.Message(true, "Successfully updated")
	resp["transaction"] = temp
	return resp
}

func(transaction *Transaction) Delete(id uint) (map[string]interface{}) {
	err:= GetDB().Table("transactions").Where("ID = ?", id).Delete(transaction).Error
	if err != nil {
		return utils.Message(false, "Error deleting record")
	}
	return utils.Message(true,"Successfully deleted")
}
