package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"path"
	"os"
	"html/template"
)

type Post struct {
	Content string
}

func main() {
	fmt.Println("Starting")
	// Vars
	var filePaths []string 
	var posts []Post

	// Flags
	filePath := flag.String("file","","Path to html file")
	dirPath  := flag.String("dir","","Path to dir containinf html files")
	flag.Parse()

	if len(*dirPath) > 0 && len(*filePath) > 0 {
		panic("Can't have both dir and file")
	} else if len(*dirPath) == 0 && len(*filePath) == 0 {
		panic("dir and file can't both be empty")
	}

	if *filePath != "" {
		filePaths = append(filePaths, *filePath)
	}
	if *dirPath != "" {
		files, _ := ioutil.ReadDir(*dirPath)
		for _, file := range files {
			filePaths = append(filePaths,path.Join(*dirPath,file.Name()))
		}
	}
	
	for _, filePath := range filePaths {
		content, _ := ioutil.ReadFile(filePath)
		posts = append(posts, Post{string(content)})
	}


	f, _ := os.Create("first-post.html")
	t := template.Must(template.New("POSTS").ParseFiles("template.tmpl"))
	
	err := t.Execute(f, posts)
	if err != nil {
		panic(err)
	}
	

}

// func createSite() {
// 	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
// 	firstPost, err := ioutil.ReadFile(string(*filePath))
// 	err = t.Execute(f, string(firstPost))
// 	if err != nil {
// 		panic(err)
// 	}

// }
