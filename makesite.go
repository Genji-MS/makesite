package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type FileInfo struct {
	Intro string
	Body  string
}

func main() {
	text, err := ioutil.ReadFile("first-post.txt")

	td := FileInfo{"Our note reads as follows:", string(text)}

	// t, err := template.New("Note").Parse(" \"{{.Intro}}\"  \"{{.Body}}\"")
	// if err != nil {
	// 	panic(err)
	// }

	//.new creates the instance .ParseFiles parses the document and does a 'find&replace' of var
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.ExecuteTemplate(os.Stdout, "template.tmpl", td)
	check(err)

	//ls of directory
	// directory := "."
	// files, err := ioutil.ReadDir(directory)
	// check(err)

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	//write to a file
	directory := "."
	f, err := os.Create(directory + "/first-post.html")
	check(err)
	err = t.Execute(f, td)
	check(err)
	f.Close()

	// f, err := os.Create("/tmp/first-post2")
	// check(err)
	// defer f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
