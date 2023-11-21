package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ammarthayani/issuetracker/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type API struct {
	r  *mux.Router
	db *gorm.DB
}

func createApi() *API {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db")
	}
	r := mux.NewRouter()
	return &API{
		r:  r,
		db: db,
	}
}

func main() {
	api := createApi()

	api.db.AutoMigrate(&models.Issue{})

	api.r.HandleFunc("/", HomeController)

	http.Handle("/", api.r)
	http.ListenAndServe(":8000", nil)
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/home/index.html")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db")
	}

	var issue models.Issue
	db.First(&issue)
	tmpl.Execute(w, issue)
}
