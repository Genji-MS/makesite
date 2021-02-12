package main

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

type FileInfo struct {
	Intro string
	Body  string
}

func main() {
	file := flag.String("file", "", " filename of .txt file to be parsed")
	dir := flag.String("dir", "", " directory where we parse .txt files")
	flag.Parse()
	if len(*file) > 0 {
		text, err := ioutil.ReadFile(*file) //"first-post.txt"
		check(err)
		td := FileInfo{"Our note reads as follows:", string(text)}
		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

		directory := "."
		filename := *file
		f, err := os.Create(directory + "/" + strings.TrimSuffix(filename, filepath.Ext(filename)) + ".html")
		//works, however it's assuming the suffix of a file is 4 characters in length. So we use the above
		//f, err := os.Create(directory + "/" + filename[:len(filename)-4] + ".html")
		check(err)
		err = t.Execute(f, td)
		check(err)
		f.Close()
	} else if len(*dir) > 0 {
		// Check in directory, if files have proper extension display them
		files, err := ioutil.ReadDir(*dir)
		check(err)
		printedFiles := 0
		totalFileSize := 0.0
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".txt" {
				printedFiles++
				fmt.Println(file.Name())
				totalFileSize += txtToHTML(*dir, file.Name())
				//removes the file extension from a file
				//fileTitle := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
				//fmt.Println(*dir + "/" + fileTitle)
			}
		}
		if printedFiles > 0 {
			//format to float <> string https://yourbasic.org/golang/convert-int-to-string/
			FileSize_String := fmt.Sprintf("%.1f", totalFileSize)
			// Terminal commands for color/font https://stackoverflow.com/questions/2924697/how-does-one-output-bold-text-in-bash
			fmt.Printf("\033[32m\033[1mSuccess!\033[0m Generated \033[1m%s\033[0m pages (%skB total).\n", strconv.Itoa(printedFiles), FileSize_String)
		}
		return
	}
	// t, err := template.New("Note").Parse(" \"{{.Intro}}\"  \"{{.Body}}\"")
	// if err != nil {
	// 	panic(err)
	// }

	//.new creates the instance .ParseFiles parses the document and does a 'find&replace' of var
	// t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// err = t.ExecuteTemplate(os.Stdout, "template.tmpl", td)
	// check(err)

	//ls of directory
	// directory := "."
	// files, err := ioutil.ReadDir(directory)
	// check(err)

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	//write to a file
	//https://stackoverflow.com/questions/32551811/read-file-as-template-execute-it-and-write-it-back
	// directory := "."
	// f, err := os.Create(directory + "/first-post.html")
	// check(err)
	// err = t.Execute(f, td)
	// check(err)
	// f.Close()

	// f, err := os.Create("/tmp/first-post2")
	// check(err)
	// defer f.Close()
}

func txtToHTML(directory string, textfile string) (fileSize float64) {
	text, err := ioutil.ReadFile(textfile) //"first-post.txt"
	check(err)
	td := FileInfo{"Our note reads as follows:", string(text)}
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	htmlfile := strings.TrimSuffix(textfile, filepath.Ext(textfile)) + ".html"
	f, err := os.Create(directory + "/" + htmlfile)
	//works, however it's assuming the suffix of a file is 4 characters in length. So we use the above
	//f, err := os.Create(directory + "/" + filename[:len(filename)-4] + ".html")
	check(err)
	err = t.Execute(f, td)
	check(err)
	//we cannot call .Stat().Size() so we store multiple variables to pull statistics of a file, then the file size
	fileInfo, _ := os.Stat(htmlfile)
	fileSize = float64(fileInfo.Size() / 1000)
	//fmt.Println(fileSize) // /1000 because we want kB not Bytes
	f.Close()
	return fileSize
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
