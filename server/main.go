package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/choueric/clog"
	"github.com/choueric/homeKit/homeKit"
)

/*
 * URL, METHOD, HANDLER
 */

var gBlob homeKit.IfaceInfoBlob

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view.html")

	t.Execute(w, gBlob)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		clog.Printf("Path: %s\n", r.URL.Path)
		body, _ := ioutil.ReadAll(r.Body)
		err := gBlob.FromJson(body)
		if err != nil {
			clog.Fatal(err)
		}
		for _, v := range gBlob.InfoArray {
			clog.Printf("%s: %v\n", v.Name, v.IP)
		}
	} else {
		clog.Warn("unsupported method: %s", r.Method)
		w.Write([]byte("unsupported method"))
	}
}

func main() {
	//http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/save/", saveHandler)

	port := ":8088"
	clog.Printf("start server at %s\n", port)
	http.ListenAndServe(port, nil)
}
