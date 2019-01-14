package main

import (
	"github.com/derekclevenger/accountapi/app"
	"github.com/derekclevenger/accountapi/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	//User routes
	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT", "PATCH")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")

	//Account routes
	router.HandleFunc("/api/account", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/account/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/api/account/getall/{user_id}", controllers.GetAccounts).Methods("GET")
	router.HandleFunc("/api/account/{id}", controllers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/api/account/{id}", controllers.UpdateAccount).Methods("PUT", "PATCH")

	//Budget routes
	router.HandleFunc("/api/budget", controllers.CreateBudget).Methods("POST")
	router.HandleFunc("/api/budget/{id}", controllers.GetBudget).Methods("GET")
	router.HandleFunc("/api/budget/getall/{user_id}", controllers.GetBudgets).Methods("GET")
	router.HandleFunc("/api/budget/{id}", controllers.DeleteBudget).Methods("DELETE")
	router.HandleFunc("/api/budget/{id}", controllers.UpdateBudget).Methods("PUT", "PATCH")

	//Categories routes
	router.HandleFunc("/api/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category/{id}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/api/category/getall/{user_id}", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/category/{id}", controllers.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/api/category/{id}", controllers.UpdateCategory).Methods("PUT", "PATCH")

	//Expenditures routes
	router.HandleFunc("/api/expenditure/{id}", controllers.GetExpenditure).Methods("GET")

	//Transaction routes
	router.HandleFunc("/api/transaction", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/api/transaction/getbycategory", controllers.GetTransactionsByCategory).Methods("POST")
	router.HandleFunc("/api/transaction/{id}", controllers.GetTransaction).Methods("GET")
	router.HandleFunc("/api/transaction/getall/{user_id}", controllers.GetTransactions).Methods("GET")
	router.HandleFunc("/api/transaction/{id}", controllers.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/api/transaction/{id}", controllers.UpdateTransaction).Methods("PUT", "PATCH")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
