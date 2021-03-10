package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type MyElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]MyElement)
var DATAFILE = "/tmp/datafile.gob"

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	if err = encoder.Encode(DATA); err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}
	return err
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	if err != nil {
		fmt.Println("Empty key/value stpre!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}

func ADD(k string, n MyElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func LOOKUP(k string) *MyElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}
}

func CHANGE(k string, n MyElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, d := range DATA {
		fmt.Printf("keq: %s value:%v\n", k, d)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("home.gohtml"))
	myT.ExecuteTemplate(w, "home.gohtml", nil)
}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store")

	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\"> List all elements!</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\"> Change an element!</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\"> Insert new element!</a>")

	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")

	for k, v := range DATA {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV Store")
	tmpl := template.Must(template.ParseFiles("update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	n := MyElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	if CHANGE(key, n) {
		fmt.Println("Update operations failed")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct {
			Success bool
		}{true})
	}
}

func insertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV store!")
	tmpl := template.Must(template.ParseFiles("insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := MyElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	if !ADD(key, n) {
		fmt.Println("Add operation failed!")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func main() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}

	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Using port number: ", PORT)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
