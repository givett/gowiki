// Creating a data structure with load and save methods.

package main

import (
	"fmt"
	"io/ioutil"
)

// Page - define Page as a struct with two fields representing the title and body.
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage - loadPage constructs the file name from the title parameter,
// reads the file's contents into a new variable body, and returns a pointer
// to a Page literal constructed with the proper title and body values.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
