package controllers

import (
	"github.com/derekclevenger/accountapi/models"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


var GetExpenditure = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expenditure := &models.Expenditure{}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding request"))
		return
	}

	resp := expenditure.Get(uint(id))
	utils.Respond(w, resp)
}
