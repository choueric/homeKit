package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/choueric/clog"
	"github.com/choueric/homeKit/homeKit"
)

var gBlob homeKit.IfaceInfoBlob

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view.html")

	var text string
	for _, v := range gBlob.InfoArray {
		text += fmt.Sprintf("%s: %v\n", v.Name, v.IP)
	}
	clog.Printf("view: %s\n", text)
	t.Execute(w, text)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	err := gBlob.FromJson(body)
	if err != nil {
		clog.Fatal(err)
	}
	for _, v := range gBlob.InfoArray {
		clog.Printf("%s: %v\n", v.Name, v.IP)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8088", nil)
}
