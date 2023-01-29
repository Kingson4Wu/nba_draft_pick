package main

import (
	"embed"
	"fmt"
	"html/template"
	"nba_draft_pick/draft"
	"net/http"
)

type ContactDetails struct {
	Email string
}

//go:embed forms.html
var f embed.FS

func main() {
	d2019 := draft.NewDraft([14]int{140, 140, 140, 125, 105, 90, 60, 60, 60, 30, 20, 10, 10, 10})

	tmpl := template.Must(template.ParseFS(f, "forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
		/*details := ContactDetails{
			Email: r.FormValue("email"),
		}*/
	})

	http.HandleFunc("/pick", func(w http.ResponseWriter, r *http.Request) {

		r2019 := d2019.NewRoundResult()
		fmt.Println(r2019)
		tmpl.Execute(w, struct {
			Success bool
			Content string
		}{true,
			fmt.Sprintf("%v", r2019)})
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
