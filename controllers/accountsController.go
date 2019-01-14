package controllers

import (
	"encoding/json"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/derekclevenger/accountapi/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"fmt"
)

var CreateAccount = func (w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	resp := account.Create()
	fmt.Println(resp)
	utils.Respond(w, resp)
}

var DeleteAccount = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := models.Delete(uint(id))
	utils.Respond(w, resp)
}

var UpdateAccount = func (w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	resp := account.Update(uint(id))
	utils.Respond(w, resp)
}

var GetAccounts = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := strconv.Atoi(params["user_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}
	resp := models.GetAll(uint(user))
	utils.Respond(w, resp)
}

var GetAccount = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := models.Get(uint(id))
	utils.Respond(w, resp)
}