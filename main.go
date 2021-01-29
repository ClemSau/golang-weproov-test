package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Objects

// Article is a blog article
type Article struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  `json:"author,omitempty"`
}

// Author is a author of blog articles
type Author struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var articles []Article

func main() {
	articles = append(articles, Article{ID: "1", Title: "My first article", Content: "This is the content of the first article", Author: Author{ID: "1", Name: "Jhon Doe"}})

	err := godotenv.Load(".env")
	CheckError(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}

// CheckError stop the program if an error is detected
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
