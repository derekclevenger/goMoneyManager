package models

import (
	"github.com/derekclevenger/accountapi/utils"
	"github.com/jinzhu/gorm"
)

type Budget struct {
	gorm.Model
	Amount float64 `json:"amount"`
	Category string `json:"category"`
	Monthly bool `json:"monthly"`
	UserId uint `json:"user_id"`
}

func(budget *Budget) Validate() (map[string]interface{}, bool) {
	if budget.Amount <= 0 {
		return utils.Message( false,"Must have an amount greater than zero"), false
	}

	if budget.Category == "" {
		return utils.Message(false, "Must have a category"), false
	}

	return utils.Message(true, "success"), true
}

func(budget *Budget) Create() (map[string]interface{}) {
	if resp, ok := budget.Validate(); !ok {
		return resp
	}

	temp := &Budget{}
	err := GetDB().Table("budgets").Where("user_id = ? and category = ?", budget.UserId, budget.Category).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound  {
		return utils.Message(false, "error")
	}

	if temp.Category != "" {
		resp := utils.Message(false, "Budget with category exists already")
		resp["Existing Budget"] = temp
		return resp
	}

	err = GetDB().Create(budget).Error
	if err != nil {
		return utils.Message(false, "Error creating budget")
	}
	resp := utils.Message(true, "Successfully created")
	resp["budget"] = budget
	return resp
}

func(budget *Budget) GetAll(user uint) (map[string]interface{}) {
	budgets := make([]*Budget, 0)
	err := GetDB().Table("budgets").Where("user_id = ?", user).Find(&budgets).Error
	if err != nil {
		 return utils.Message(false, "Error getting budgets")
	}
	resp := utils.Message(true, "Success")
	resp["budgets"] = budgets
	return resp
}

func(budget *Budget) Get(id uint) (map[string]interface{}) {
	b := &Budget{}
	err := GetDB().Table("budgets").Where("ID = ?", id).First(b).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return utils.Message(false, "Error retrieving Budget")
	}

	resp := utils.Message(true, "Success")
	resp["budget"] = b
	return resp
}

func(budget *Budget) Update(id uint) (map[string]interface{}) {
	temp := &Budget{}
	err := GetDB().Table("budgets").Where("ID = ?", id).First(temp).Error
	if err != nil {
		return utils.Message(false, "Error retrieving record to update")
	}

	if budget.Category != temp.Category && budget.Category != "" {
		temp.Category = budget.Category
	}

	if budget.Amount != temp.Amount {
		temp.Amount = budget.Amount
	}

	if budget.Monthly != temp.Monthly {
		temp.Monthly = budget.Monthly
	}

	err = GetDB().Table("budgets").Where("ID = ?", id).Update(temp).Error
	if err != nil {
		return utils.Message(false, "Error updating budget")
	}

	resp := utils.Message(true,"Successfully updated")
	resp["budget"] = temp
	return resp
}

func(budget *Budget) Delete(id uint) (map[string]interface{}) {
	err := GetDB().Table("budgets").Where("ID = ?", id).Delete(budget).Error
	if err != nil {
		return utils.Message(false, "Error deleting record")
	}
	return utils.Message(true, "Successfully deleted")
}

