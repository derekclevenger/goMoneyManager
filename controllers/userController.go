package controllers

import (
	"encoding/json"
	"github.com/derekclevenger/accountapi/models"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	response := user.Create()
	utils.Respond(w, response)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	response := models.Login(user.Email, user.Password)
	utils.Respond(w, response)
}

var UpdateUser = func (w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	resp := user.Update(uint(id))
	utils.Respond(w, resp)
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	resp := user.Delete(uint(id))
	utils.Respond(w, resp)
}