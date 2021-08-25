package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//func loadPage(title string) *Page {
//    filename := title + ".txt"
//    body, _ := ioutil.ReadFile(filename)
//    return &Page{Title: title, Body: body}
//}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	err := t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		log.Fatal(err)
		return
	}
	//_, err = fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	//if err != nil {
	//    log.Fatal(err)
	//    return
	//}

	//t, _ := template.ParseFiles("view.html")
	//t.Execute(w, p)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	//    "<form action=\"/save/%s\" method=\"POST\">"+
	//    "<textarea name=\"body\">%s</textarea><br>"+
	//    "<input type=\"submit\" value=\"Save\">"+
	//    "</form>", p.Title, p.Title, p.Body)

	//t, _ := template.ParseFiles("edit.html")
	//t.Execute(w, p)

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//renderTemplate(w, "save", p)
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//func main() {
//    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//    err := p1.save()
//    if err != nil {
//        return
//    }
//    p2, _ := loadPage("TestPage")
//    fmt.Println(string(p2.Body))
//}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
