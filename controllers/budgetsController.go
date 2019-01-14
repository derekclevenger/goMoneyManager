package controllers

import (
	"encoding/json"
	"github.com/derekclevenger/accountapi/models"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateBudget = func(w http.ResponseWriter, r *http.Request) {
	budget := &models.Budget{}
	err := json.NewDecoder(r.Body).Decode(budget)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request body"))
		return
	}

	resp := budget.Create()
	utils.Respond(w, resp)
}

var GetBudget = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	budget := &models.Budget{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := budget.Get(uint(id))
	utils.Respond(w, resp)
}

var GetBudgets = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	budget := &models.Budget{}
	user, err := strconv.Atoi(params["user_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := budget.GetAll(uint(user))
	utils.Respond(w, resp)
}

var UpdateBudget = func(w http.ResponseWriter, r *http.Request) {
	budget := &models.Budget{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(budget)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request body"))
		return
	}

	resp := budget.Update(uint(id))
	utils.Respond(w, resp)


}

var DeleteBudget = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	budget := &models.Budget{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := budget.Delete(uint(id))
	utils.Respond(w, resp)
}