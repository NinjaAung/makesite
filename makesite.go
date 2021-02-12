package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"html/template"
	"strings"
)


type post struct {
	content string
}

func main() {
	fmt.Println("Starting")
	filePath := flag.String("file","post.txt","Path To html file")

	flag.Parse()

	filePathName := strings.Split(*filePath,".")[0]
	f, _ := os.Create(filePathName+".html")
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	firstPost, err := ioutil.ReadFile(string(*filePath))
	err = t.Execute(f, string(firstPost))
	if err != nil {
		panic(err)
	}
	
}
