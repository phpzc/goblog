package main

import (
	"database/sql"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var router *mux.Router

var db *sql.DB

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	//router := http.NewServeMux()
	//router := mux.NewRouter()

	//中间件 强制内容为HTML
	router.Use(forceHTMLMiddleware)

	// homeURL, _ := router.Get("home").URL()
	// fmt.Println("homeURL: ", homeURL)
	// articleURL, _ := router.Get("articles.show").URL("id", "23")
	// fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))

}
