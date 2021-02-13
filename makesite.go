package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//Post ...
type Post struct {
	Title   string
	Content string
}

func main() {
	fmt.Println("Starting")
	// Vars
	var filePaths []string
	var posts []Post

	// Flags
	filePath := flag.String("file", "", "Path to html file")
	dirPath := flag.String("dir", "", "Path to dir containinf html files")
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
			filePaths = append(filePaths, path.Join(*dirPath, file.Name()))
		}
	}

	for _, filePath := range filePaths {
		file, _ := ioutil.ReadFile(filePath)
		fileSplit := strings.Split(string(file), "\n")
		content := strings.Join(fileSplit[1:], "\n")
		posts = append(posts, Post{fileSplit[0], content})
	}

	f, _ := os.Create("first-post.html")
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	err := t.Execute(f, posts)
	if err != nil {
		panic(err)
	}

}
