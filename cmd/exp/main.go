package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Ahmed Ben Abdessalam",
		Bio:  `<script>alert'"test")</script>`,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
