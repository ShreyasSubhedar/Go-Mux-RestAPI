package main

// Book Structs (Model)
type Book struct {
	ID     string  `json:"id"`
	Ispn   string  `json:"ispn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Mobile    string `json:"mobile"`
}
