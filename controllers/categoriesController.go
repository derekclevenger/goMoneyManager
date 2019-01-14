package controllers

import (
	"encoding/json"
	"github.com/derekclevenger/accountapi/models"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateCategory = func(w http.ResponseWriter, r *http.Request) {
	category := &models.Categories{}
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request body"))
		return
	}

	resp := category.Create()
	utils.Respond(w, resp)
}

var GetCategory = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cat := &models.Categories{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp :=cat.Get(uint(id))
	utils.Respond(w, resp)
}

var GetCategories = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cat := &models.Categories{}
	user, err := strconv.Atoi(params["user_id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := cat.GetAll(uint(user))
	utils.Respond(w, resp)
}

var UpdateCategory = func(w http.ResponseWriter, r *http.Request) {
	category := &models.Categories{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding body"))
		return
	}

	resp := category.Update(uint(id))
	utils.Respond(w, resp)
}

var DeleteCategory= func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cat := &models.Categories{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in request"))
		return
	}

	resp := cat.Delete(uint(id))
	utils.Respond(w, resp)
}