package models

import (
	"fmt"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"sql:""`
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Email address required"), false
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password isn't long enough"), false
	}

	temp := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Error connecting try again"), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email already in use"), false
	}

	return utils.Message(true, "Requirements Passed"), true
}

func (user *User) Create() (map[string]interface{}) {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	GetDB().Create(user)

	if user.ID < 0 {
		return utils.Message(false, "Failed to create user")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""
	response := utils.Message(true, "User created successfully")
	response["user"] = user
	return response
}

func Login(email, password string) (map[string]interface{}) {
	user := &User{}

	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email not found")
		}
		return utils.Message(false, "Connection error, Please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid Credentials. Please try again")
	}

	user.Password = ""
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	resp := utils.Message(true, "Success")
	resp["user"] = user
	return resp
}

func (userToUpdate *User) Update(id uint) (map[string] interface{}) {



	user := &User{}
	err := GetDB().Table("users").Where("ID = ?", id).First(user).Error
	if err != nil {
		return utils.Message(false, "error retrieving account")
	}
	if userToUpdate.Email != "" {
		temp := &User{}

		err := GetDB().Table("users").Where("email = ?", userToUpdate.Email).First(temp).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return utils.Message(false, "Error connecting try again")
		}

		if temp.Email != "" {
			return utils.Message(false, "Email already in use")
		}

			user.Email = userToUpdate.Email
	}

	if userToUpdate.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userToUpdate.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
	}
	fmt.Println(user)
	err = GetDB().Table("users").Where("ID = ?", id).Update(user).Error
	if err != nil {
		fmt.Println(err)
		return utils.Message(false, "Error updating record")
	}
	user.Password = ""
	resp := utils.Message(true, "success")
	resp["user"] = user
	fmt.Println(user)
	return resp
}


func(user *User) Delete(id uint) (map[string] interface{}) {

	err := GetDB().Table("users").Where("ID = ?", id).Delete(user).Error
	if err != nil {
		fmt.Println(err)
		return utils.Message(false, "Error deleting record")
	}
	return utils.Message(true, "Successfully deleted")
}

func GetUser(u uuid.UUID) (map[string]interface{}) {
	user := &User{}

	err := GetDB().Table("users").Where("ID = ?", u).First(user).Error
	if err != nil {
		return utils.Message(false, "Error getting user")
	}

	user.Password = ""
	response := utils.Message(true, "Success")
	response["user"] = user
	return response
}