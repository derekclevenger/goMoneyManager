package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/derekclevenger/accountapi/models"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateTransaction = func(w http.ResponseWriter, r *http.Request) {
	transaction := &models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		fmt.Println(transaction)
		utils.Respond(w, utils.Message(false, "Error decoding request body"))
		return
	}

	resp := transaction.Create()
	utils.Respond(w, resp)
}

var GetTransaction = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trans := &models.Transaction{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request"))
		return
	}

	resp := trans.Get(uint(id))
	utils.Respond(w, resp)
}

var GetTransactions = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trans := &models.Transaction{}
	user, err := strconv.Atoi(params["user_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding Request"))
		return
	}

	resp := trans.GetAll(uint(user))
	utils.Respond(w, resp)
}

var GetTransactionsByCategory = func(w http.ResponseWriter, r *http.Request) {
	trans := &models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(trans)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request body"))
		return
	}

	resp := trans.GetTransactionsByCategory(trans.UserId, trans.Category)
	utils.Respond(w, resp)
}

var UpdateTransaction = func(w http.ResponseWriter, r *http.Request) {
	trans := &models.Transaction{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(trans)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request body"))
		return
	}

	resp := trans.Update(uint(id))
	utils.Respond(w, resp)
}

var DeleteTransaction = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trans := &models.Transaction{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request"))
		return
	}

	resp := trans.Delete(uint(id))
	utils.Respond(w, resp)
}
