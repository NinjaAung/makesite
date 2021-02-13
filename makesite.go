package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

//Post ...
type Post struct {
	Title   string
	Content string
}

func main() {
	start := time.Now()
	// Colour
	var Green = "\x1B[32m"
	var NC = "\x1b[0m"

	// Vars
	var filePaths []string
	var totalSize int64 = 0

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

		fileStat, _ := os.Stat(filePath)
		fileSize := fileStat.Size()
		totalSize += fileSize

		filePathSplit := strings.Split(filePath, "/")
		fileName := strings.Split(filePathSplit[len(filePathSplit)-1], ".")[0]

		file, _ := ioutil.ReadFile(filePath)
		fileSplit := strings.Split(string(file), "\n")
		content := strings.Join(fileSplit[1:], "\n")

		createFileFromTemplate(fileName+".html", Post{fileSplit[0], content})
	}
	elapsed := time.Since(start)

	fmt.Printf(Green+"Success! "+NC+"Generated %d pages (%.2fkb total) in %.3f seconds\n",
		len(filePaths), float64(totalSize)*float64(0.001), elapsed.Seconds())
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
