package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

//Post ...
type Post struct {
	Title   string
	Content string
}

var totalSize int64 = 0

func main() {
	// Colour
	var Green = "\x1B[32m"
	var Bold = "\x1b[1m"
	var NC = "\x1b[0m"

	// Vars
	var filePaths []string

	// Flags
	filePath := flag.String("file", "", "Path to md file")
	dirPath := flag.String("dir", "", "Path to dir containing md files")
	flag.Parse()

	// Checks Values Of Flags
	isDirAndFilePathEmpty := len(*dirPath) == 0 && len(*filePath) == 0
	isDirAndFilePathFilled := len(*dirPath) > 0 && len(*filePath) > 0

	if isDirAndFilePathFilled {
		panic("Can't have both dir and file")
	} else if isDirAndFilePathEmpty {
		panic("dir and file can't both be empty")
	} else if *filePath != "" {
		filePaths = append(filePaths, *filePath)
	} else if *dirPath != "" {
		filePaths = findFilesInFolder(*dirPath)
	}

	start := time.Now()
	os.Mkdir("Generated/", 0755)
	for _, filePath := range filePaths {
		fileName := strings.Split(
			strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1], ".")
		if fileName[1] == "md" {
			file, _ := ioutil.ReadFile(filePath)
			createFile(fileName[0]+".html", file)
		}

	}
	elapsed := time.Since(start)

	fmt.Printf(Green+Bold+"Success! "+NC+"Generated "+Bold+"%d "+NC+"pages (%.2fkb total) in %.2f seconds\n",
		len(filePaths), float64(totalSize)*float64(0.001), elapsed.Seconds())
}

func createFile(name string, file []byte) {

	unsafe := blackfriday.Run(file)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	f, _ := os.Create(name)
	f.Write(html)
	fileStat, _ := os.Stat(name)
	totalSize += fileStat.Size()
	f.Close()
	os.Rename(name, "Generated/"+name)
}

func findFilesInFolder(dirPath string) []string {
	var filePaths []string
	filePath, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	for _, file := range filePath {
		if file.IsDir() {
			subFolderPath := findFilesInFolder(path.Join(dirPath, file.Name()))
			filePaths = append(filePaths, subFolderPath...)
			continue
		}
		filePaths = append(filePaths, path.Join(dirPath, file.Name()))
	}
	return filePaths

}
