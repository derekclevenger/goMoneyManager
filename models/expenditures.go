package models

import (
	"fmt"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/jinzhu/gorm"
)

type Expenditure struct {
	gorm.Model
	Category string `json:"category"`
	Budget float64 `json:"budget"`
	Spent float64 `json:"spent"`
	Total float64 `json:"total"`
	UserId uint `json:"user_id"`
}


func(expenditure *Expenditure) Get(user uint) (map[string]interface{}) {
	exps := make([]*Expenditure, 0)

	sql :=  fmt.Sprintf("select sum(transactions.amount) as spent, transactions.category," +
		 	"budgets.amount as budget, budgets.amount - sum(transactions.amount) as total,  transactions.user_id " +
		 	"from transactions join budgets on budgets.category = transactions.category " +
		 	"where transactions.user_id = %d and budgets.deleted_at ISNULL and transactions.deleted_at ISNULL " +
			"group by transactions.category, budget, transactions.user_id", user)

	err := db.Raw(sql).Scan(&exps).Error
	if err != nil && err == gorm.ErrRecordNotFound{
		return utils.Message(false, "Error retrieving record")
	}

	resp := utils.Message(true, "Success")
	resp["expenditures"] = exps
	return resp
}
