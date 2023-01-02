package server

import (
	"html/template"
	"io"
	"main/res"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Ascii represents the ASCII art entity.
type Ascii struct {
	Word   string
	Art    string
	Banner string
}

var entity Ascii

// HomeHandler handles the home page request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &entity)
}

// GenerateHandler generates the ASCII art based on the given input.
func GenerateHandler(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("generate") != "" {
		word := r.FormValue("body")
		banner := r.FormValue("banner")
		if word == "" {
			http.Redirect(w, r, "https://http.cat/400", http.StatusSeeOther)
			return
		}

		// Validate that the input is ASCII characters only.
		for _, v := range word {
			if v > 127 {
				http.Redirect(w, r, "https://http.cat/400", http.StatusSeeOther)
				return
			}
		}

		entity = Ascii{Word: word, Art: res.AsciiArt(word, banner), Banner: banner}
	}

	if r.FormValue("download") != "" {
		source := strings.NewReader(entity.Art) //defining source
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Content-Disposition", "attachment; filename=file.txt")
		w.Header().Add("Content-Length", strconv.Itoa(len(entity.Art)))
		io.Copy(w, source) //w is the destination
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

var validPath = regexp.MustCompile("^/(ascii-art|$)")

// MakeHandler creates a handler function.
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Redirect(w, r, "https://http.cat/404", http.StatusSeeOther)
			return
		}
		fn(w, r)
	}
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

// renderTemplate renders the given template with the given data.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Ascii) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
