package main

import (
	"io/ioutil"
	"os"
	"html/template"
)


type post struct {
	content string
}

func main() {
	f, _ := os.Create("new.html")
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	firstPost, err := ioutil.ReadFile("first-post.txt")
	
	err = t.Execute(f, string(firstPost))
	if err != nil {
		panic(err)
	}
	
}
