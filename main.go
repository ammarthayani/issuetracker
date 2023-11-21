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

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db")
	}

	db.AutoMigrate(&models.Issue{})

	testIssue := models.Issue{Name: "test"}

	db.Create(&testIssue)

	// TODO: Create API Object to store mux router
	r := mux.NewRouter()
	r.HandleFunc("/", HomeController)

	http.Handle("/", r)
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
