package main

//GOROOT="/usr/local/go"
//GOPATH="/Users/g3n6i/go"
import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TextFile struct {
	Intro string
	Body  string
}

type ImageFile struct {
	Image string
	Intro string
	Body  string
}

func main() {
	file := flag.String("file", "", " filename of .txt file to be parsed")
	dir := flag.String("dir", "", " directory where we parse .txt files")
	flag.Parse()

	if *file != "" {
		directory := "."
		filename := strings.TrimSuffix(*file, filepath.Ext(*file))
		useImageTemplate := false
		if _, err := os.Stat(filename + ".png"); err == nil {
			//if image exists, use our HTML image template
			useImageTemplate = true
		}
		txtToHTML(directory, filename, useImageTemplate)
	} else if *dir != "" {
		files, err := ioutil.ReadDir(*dir)
		check(err)
		printedFiles := 0
		totalFileSize := 0.0
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".txt" {
				printedFiles++
				fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
				fmt.Println(fileName)
				useImageTemplate := false
				if _, err := os.Stat(fileName + ".png"); err == nil {
					//if image exists, use our HTML image template
					useImageTemplate = true
				}
				totalFileSize += txtToHTML(*dir, fileName, useImageTemplate)
			}
		}
		if printedFiles > 0 {
			FileSize_String := fmt.Sprintf("%.1f", totalFileSize)
			fmt.Printf("\033[32m\033[1mSuccess!\033[0m Generated \033[1m%s\033[0m pages \033[34m(\033[36m%skB total\033[34m)\033[0m.\n", strconv.Itoa(printedFiles), FileSize_String)
		}
	}
}

func txtToHTML(directory, fileName string, useImageTemplate bool) (fileSize float64) {
	var td interface{} // declares a variable of unknown type
	var templateSelector string
	text, err := ioutil.ReadFile(fileName + ".txt") //"first-post.txt"
	check(err)

	if useImageTemplate {
		templateSelector = "templateImg.tmpl"
		td = ImageFile{string(fileName + ".png"), "Our note reads as follows:", string(text)}
	} else {
		templateSelector = "template.tmpl"
		td = TextFile{"Our note reads as follows:", string(text)}
	}
	t := template.Must(template.New(templateSelector).ParseFiles(templateSelector))

	htmlfile := fileName + ".html"
	f, err := os.Create(directory + "/" + htmlfile)
	check(err)

	err = t.Execute(f, td)
	check(err)
	fileInfo, _ := os.Stat(htmlfile)
	fileSize = float64(fileInfo.Size() / 1000)
	f.Close()
	return fileSize
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
