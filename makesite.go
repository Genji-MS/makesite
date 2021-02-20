package main

//GOROOT="/usr/local/go"
//GOPATH="/Users/g3n6i/go"
import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kortschak/zalgo"
)

type FileInfo struct {
	Intro string
	Body  string
}

func main() {
	file := flag.String("file", "", " filename of .txt file to be parsed")
	dir := flag.String("dir", "", " directory where we parse .txt files")
	flag.Parse()

	if *file != "" {
		text, err := ioutil.ReadFile(*file)
		check(err)
		td := FileInfo{"Our note reads as follows:", string(text)}
		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		directory := "."
		filename := *file
		f, err := os.Create(directory + "/" + strings.TrimSuffix(filename, filepath.Ext(filename)) + ".html")
		check(err)
		err = t.Execute(f, td)
		check(err)
		f.Close()
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
				totalFileSize += txtToHTML(*dir, fileName)
			}
		}
		if printedFiles > 0 {
			FileSize_String := fmt.Sprintf("%.1f", totalFileSize)
			fmt.Printf("\033[32m\033[1mSuccess!\033[0m Generated \033[1m%s\033[0m pages \033[34m(\033[36m%skB total\033[34m)\033[0m.\n", strconv.Itoa(printedFiles), FileSize_String)
		}
	}
}

func txtToHTML(directory, fileName string) (fileSize float64) {
	text, err := ioutil.ReadFile(fileName + ".txt") //"first-post.txt"
	check(err)

	g := new(bytes.Buffer)
	glitch := zalgo.NewCorrupter(g)
	glitch.Zalgo = func(n int, r rune, z *zalgo.Corrupter) bool {
		z.Up += 0.001
		z.Middle += complex(0.001, 0.001)
		z.Down += complex(real(z.Down)*0.001, 0)
		return false
	}

	fmt.Fprintln(glitch, string(text))
	//fmt.Println(g.String())

	td := FileInfo{"Our note reads as follows:", g.String()}
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
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
