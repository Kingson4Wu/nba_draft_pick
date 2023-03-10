package main

import (
	"embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"nba_draft_pick/draft"
	"net/http"
	"strings"
)

type ContactDetails struct {
	Email string
}

//go:embed forms.html
var f embed.FS

//go:embed photos
var photos embed.FS

//go:embed data
var data embed.FS

func main() {
	d := [14]int{140, 140, 140, 125, 105, 90, 60, 60, 60, 30, 20, 10, 10, 10}
	draft := draft.NewDraft(d)

	tmpl := template.Must(template.ParseFS(f, "forms.html"))

	t, _ := data.ReadFile("data/2019.txt")
	teams := strings.Split(string(t), "\n")

	var images string
	for i := 0; i < 14; i++ {
		b, _ := photos.ReadFile("photos/" + teams[i] + ".png")
		res := base64.StdEncoding.EncodeToString(b)
		images = images + "<img src=\"data:image/png;base64," + res + "\">"
		images = images + "<div>" + teams[i] + "(" + fmt.Sprintf("%.2f", float64(d[i])/float64(10)) + "%)</div>"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct {
			Success bool
			Images  template.HTML
		}{false, template.HTML(images)})
	})

	http.HandleFunc("/pick", func(w http.ResponseWriter, r *http.Request) {

		rt := draft.NewRoundResult()
		//fmt.Println(r)

		var results string
		for i := 0; i < 14; i++ {
			index := rt[i] - 1
			b, _ := photos.ReadFile("photos/" + teams[index] + ".png")
			res := base64.StdEncoding.EncodeToString(b)
			results = results + "<img src=\"data:image/png;base64," + res + "\">"
			results = results + "<div>" + teams[index] + "(" + fmt.Sprintf("%.2f", float64(d[index])/float64(10)) + "%)</div>"
		}

		tmpl.Execute(w, struct {
			Success bool
			Content string
			Result  template.HTML
		}{true,
			fmt.Sprintf("%v", rt), template.HTML(results)})
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct {
			Success bool
			Images  template.HTML
		}{false, template.HTML(images)})
	})

	http.ListenAndServe(":8988", nil)
}
