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

	// Flags
	filePath := flag.String("file", "", "Path to html file")
	dirPath := flag.String("dir", "", "Path to dir containing html files")
	flag.Parse()

	// Panic when both or none of flags are filled
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

		filePathSplit := strings.Split(filePath, "/")
		fileName := strings.Split(filePathSplit[len(filePathSplit)-1], ".")[0]

		file, _ := ioutil.ReadFile(filePath)
		fileSplit := strings.Split(string(file), "\n")
		content := strings.Join(fileSplit[1:], "\n")

		createFileFromTemplate(fileName+".html", Post{fileSplit[0], content})
	}

}

func createFileFromTemplate(name string, post Post) {
	f, _ := os.Create(name)
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	fmt.Println("Creating ", name)
	err := t.Execute(f, post)
	if err != nil {
		panic(err)
	}
}
