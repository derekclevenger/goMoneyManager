package models

import (
	"github.com/derekclevenger/accountapi/utils"
	"github.com/jinzhu/gorm"
)

type Categories struct {
	gorm.Model
	Category string `json:"category"`
	UserId uint `json:"user_id"`
}

func(category *Categories) Validate() (map[string]interface{}, bool) {
	if category.Category == "" {
		return utils.Message(false, "Category must have a value"), false
	}

	return utils.Message(true, "Success"), true
}

func(category *Categories) Create() (map[string]interface{}) {
	if resp, ok := category.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(category).Error
	if err != nil {
		return utils.Message(false, "Error creating category")
	}
	resp := utils.Message(true, "Successfully Created")
	resp["category"] = category
	return resp
}

func(category *Categories) GetAll(user uint) (map[string]interface{}) {
	categories := make([]*Categories, 0)
	err := GetDB().Table("categories").Where("user_id = ?" , user).Find(&categories).Error
	if err != nil {
		return utils.Message(false, "Error retrieving categories")
	}
	resp := utils.Message(true, "Success")
	resp["categories"] = categories
	return resp
}

func(category *Categories) Get(id uint) (map[string]interface{}) {
	cat := &Categories{}
	err := GetDB().Table("categories").Where("ID = ?", id).First(cat).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return utils.Message(false, "Error getting record")
	}

	resp := utils.Message(true, "Success")
	resp["category"] = cat
	return resp
}

func(category *Categories) Update(id uint) (map[string]interface{}) {
	if resp, ok := category.Validate(); !ok {
		return resp
	}

	temp := &Categories{}

	err := GetDB().Table("categories").Where("ID = ?", id).First(temp).Error
	if err != nil {
		return utils.Message(false, "Error getting record")
	}

	temp.Category = category.Category

	err = GetDB().Table("categories").Where("ID = ?", id).Update(temp).Error
	if err != nil {
		return utils.Message(false, "Error updating record")
	}

	resp := utils.Message(true, "Successfully updated")
	resp["category"] = temp
	return resp
}

func(category *Categories) Delete(id uint) (map[string]interface{}) {
	err := GetDB().Table("categories").Where("ID = ?", id).Delete(category).Error
	if err != nil {
		return utils.Message(false, "Error deleting record")
	}

	return utils.Message(true, "Success")
}
