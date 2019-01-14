package main

import (
	"fmt"
	"github.com/derekclevenger/accountapi/models"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

var cat models.Categories
var acc models.Account
var bud models.Budget
var exp models.Expenditure
var tran models.Transaction
var user models.User

func TestCreateCategory(t *testing.T) {

	if cat != (models.Categories{}) {
		t.Error("Category should be empty")
	}

	cat = models.Categories{gorm.Model {uint(1), time.Now(), time.Now(), nil},
		"Entertainment", 5}

	fmt.Println(cat.Category)
	if cat.Category != "Entertainment" {
		t.Error("Category not created or something went wrong")
	}
}

func TestCreateAccount(t *testing.T) {
	if acc != (models.Account{}) {
		t.Error("Account should be empty")
	}

	acc = models.Account{gorm.Model {uint(1), time.Now(), time.Now(), nil},
		"City", "Checking", 5555,5}

	if acc.Type != "Checking" {
		t.Error("Account not created")
	}
}

func TestCreateBudget(t *testing.T) {
	if bud != (models.Budget{}) {
		t.Error("Budget should be empty")
	}

	bud = models.Budget{gorm.Model {uint(1), time.Now(), time.Now(), nil},
		555, "Checking", true,5}

	if bud.Category != "Checking" {
		t.Error("Budget not created")
	}
}


func TestCreateExpenditure(t *testing.T) {
	if exp != (models.Expenditure{}) {
		t.Error("Expenditure should be empty")
	}

	exp = models.Expenditure{gorm.Model {uint(1), time.Now(), time.Now(), nil},
		"FUNNNN", 25, 20,5, 5}

	if exp.Category != "FUNNNN" {
		t.Error("Expenditure not created")
	}
}


func TestCreateTransaction(t *testing.T) {
	if tran != (models.Transaction{}) {
		t.Error("Transaction should be empty")
	}

	tran = models.Transaction{gorm.Model {uint(1), time.Now(), time.Now(), nil},
		"You", time.Now(), 20,"Checking", 5, "Stuff"}

	if tran.Category != "Stuff" {
		t.Error("Transaction not created")
	}
}


func TestCreateUser(t *testing.T) {
	if user != (models.User{}) {
		t.Error("User should be empty")
	}

	user = models.User{gorm.Model {uint(1), time.Now(), time.Now(), nil},
		"me@Me.com", "secret", "SecretTOken" }

	if user.Email != "me@Me.com" {
		t.Error("User not created")
	}
}