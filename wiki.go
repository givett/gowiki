// Creating a data structure with load and save methods.

package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"html/template"
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StausFound)
		return
	}

	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// t, _ := template.ParsFiles("view.html")
	// t.Execute(w, p)

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		// "<form action=\"/save/%s\" method=\"POST\">"+
		// "<textarea name=\"body\">%s</textarea><br>"+
		// "<input type=\"submit\" value=\"Save\">"+
		// "</form>",
		// p.Title, p.Title, p.Body)

	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)

	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p:= &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r. "/view/"+title, http.StatusFound)
}

func main() {

  // First main function:
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	fmt.Print("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
