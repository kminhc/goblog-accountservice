package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		GetAccount,
	},
}

func GetAccount(writer http.ResponseWriter, request *http.Request) {
	var accountId = mux.Vars(request)["accountId"]

	account, err := DBCLient.QueryAccount(accountId)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Content-Length", strconv.Itoa(len(data)))
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
}
