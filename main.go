package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Objects

// Article is a blog article
type Article struct {
	gorm.Model
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  `json:"author,omitempty"`
}

// Author is a author of blog articles
type Author struct {
	gorm.Model
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var articles []Article

func main() {
	router := mux.NewRouter()
	articles = append(articles, Article{ID: "1", Title: "My first article", Content: "This is the content of the first article", Author: Author{ID: "1", Name: "Jhon Doe"}})

	err := godotenv.Load(".env")
	CheckError(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db := OpenDb(host, port, user, password, dbname)

}

// OpenDb open the connexion to the psql database
func OpenDb(host string, port string, user string, password string, dbname string) *gorm.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := gorm.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	fmt.Println("Connected!")

	db.AutoMigrate(&Article{})

	db.AutoMigrate(&Author{})

	return db
}

// CheckError stop the program if an error is detected
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
