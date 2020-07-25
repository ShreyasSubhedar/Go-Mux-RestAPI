package main

// Initialize the mock data (book)
var books []Book
var book1 = Book{
	ID:    "1",
	Ispn:  "23423",
	Title: "One night at call center",
	Author: &Author{
		Firstname: "Chetan",
		Lastname:  "Bhagat",
		Mobile:    "9988776655",
	},
}
var book2 = Book{
	ID:    "2",
	Ispn:  "23223",
	Title: "Revolution 2020",
	Author: &Author{
		Firstname: "Chetan",
		Lastname:  "Bhagat",
		Mobile:    "9988776655",
	},
}
var book3 = Book{
	ID:    "3",
	Ispn:  "23223",
	Title: "Five point someone",
	Author: &Author{
		Firstname: "Chetan",
		Lastname:  "Bhagat",
		Mobile:    "9988776655",
	},
}
