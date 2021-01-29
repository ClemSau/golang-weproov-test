package main

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
}
