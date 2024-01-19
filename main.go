package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	for i := 1; i <= 5; i++ {
		go GetPost(i)
	}

	var input string
	fmt.Scanln(&input)
}

func GetPost(pageNumber int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", pageNumber)
	result, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer result.Body.Close()

	var post Post
	err = json.NewDecoder(result.Body).Decode(&post)
	if err != nil {
		panic(err.Error())
	}

	resultString, err := json.MarshalIndent(post, "", "  ")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(resultString))
}
