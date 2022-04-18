package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {

	database.Initialize()

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))

}
