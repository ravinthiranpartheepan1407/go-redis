package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Book struct {
	BookName   string `json:"bookname"`
	BookId     int    `json:"bookid"`
	BookAuthor string `json:"BookAuthor"`
}

func main() {
	fmt.Println("Welcome")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ping)
	}

	json, err := json.Marshal(Book{BookName: "Rosie", BookId: 1, BookAuthor: "Ravinthiran"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json)
	}

	bookNameErr := client.Set("BookName", json, 0).Err()
	bookIdErr := client.Set("BookId", json, 0).Err()
	bookAuthorErr := client.Set("BookAuthor", json, 0).Err()

	if bookNameErr != nil && bookIdErr != nil && bookAuthorErr != nil {
		fmt.Println(bookNameErr, bookIdErr, bookAuthorErr)
	}

	valueName, err := client.Get("BookName").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueName)
	}

	valueId, err := client.Get("BookId").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueId)
	}

	valueError, err := client.Get("BookAuthor").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueError)
	}
}
