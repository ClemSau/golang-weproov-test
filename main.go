package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

var db *gorm.DB
var err error

// OpenDb open the connexion to the psql database
func OpenDb(host string, port string, user string, password string, dbname string) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := gorm.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	fmt.Println("Connected!")

	db.AutoMigrate(&Article{})

	db.AutoMigrate(&Author{})
}

// CheckError stop the program if an error is detected
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Endpoints

// GetArticles return the list of all the articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	db.Find(&articles)
	json.NewEncoder(w).Encode(&articles)
}

// GetArticle return an article by id
func GetArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article Article
	db.First(&article, params["id"])
	json.NewEncoder(w).Encode(&article)
}

// CreateArticle create an article with the given arguments
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	db.Create(&article)
}

// DeleteArticle delete an article by id
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article Article
	db.First(&article, params["id"])
	db.Delete(&article)
	var articles []Article
	db.Find(&articles)
	json.NewEncoder(w).Encode(&articles)
}

// UpdateArticle update a given article
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article Article
	db.First(&article, params["id"])
	var newArticle Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.Title = newArticle.Title
	article.Content = newArticle.Content
	article.Author = newArticle.Author
	db.Save(&article)
}

// GetAuthors return the list of all the authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []Author
	db.Find(&authors)
	json.NewEncoder(w).Encode(&authors)
}

// GetAuthor return an author by id
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author Author
	db.First(&author, params["id"])
	json.NewEncoder(w).Encode(&author)
}

// CreateAuthor create an author with the given arguments
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	db.Create(&author)
}

// DeleteAuthor delete an author by id
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author Article
	db.First(&author, params["id"])
	db.Delete(&author)
	var authors []Author
	db.Find(&authors)
	json.NewEncoder(w).Encode(&authors)
}

func main() {
	router := mux.NewRouter()

	err := godotenv.Load(".env")
	CheckError(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	OpenDb(host, port, user, password, dbname)

	router.HandleFunc("/articles", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", CreateArticle).Methods("CREATE")
	router.HandleFunc("/articles/{id}", DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", UpdateArticle).Methods("UPDATE")
	router.HandleFunc("/authors", GetAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", GetAuthor).Methods("GET")
	router.HandleFunc("/authors/{id}", CreateAuthor).Methods("CREATE")
	router.HandleFunc("/authors/{id}", DeleteAuthor).Methods("DELETE")
	router.HandleFunc("/authors/{id}", UpdateAuthor).Methods("UPDATE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
