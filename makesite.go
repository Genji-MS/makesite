package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
)

type FileInfo struct {
	Intro string
	Body  string
}

func main() {
	file := flag.String("file", "", " filename of .txt file to be parsed")
	flag.Parse()
	if len(*file) > 0 {
		text, err := ioutil.ReadFile(*file) //"first-post.txt"
		check(err)
		td := FileInfo{"Our note reads as follows:", string(text)}
		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

		directory := "."
		filename := *file
		f, err := os.Create(directory + "/" + filename[:len(filename)-4] + ".html")
		check(err)
		err = t.Execute(f, td)
		check(err)
		f.Close()
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
