package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"text/template"
)

var templateEngine = template.Must(template.ParseGlob("templates/*.html"))

type data struct {
	ErrorMessage string
	NewURL       string
}

//FavIconHandler ...
func FavIconHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "image/x-icon")
	// w.Header().Set("Cache-Control", "public, max-age=7776000")
	// fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=")
}

//IndexHandler ...
func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleURLRedirect(r.URL.Path, rw, r)
		return
	}

	if r.Method != http.MethodPost {
		templateEngine.ExecuteTemplate(rw, "index", nil)
	} else {
		postURL(rw, r)
	}
}

func handleURLRedirect(path string, rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("path: %s\n", path)

	//Get the original URL from DB

	//Redirect
	http.Redirect(rw, r, "", http.StatusMovedPermanently)
}

func postURL(rw http.ResponseWriter, r *http.Request) {
	_url := r.FormValue("url")
	_, err := url.ParseRequestURI(_url)

	data := &data{}

	//Basic validation
	if _url == "" || err != nil {
		data.ErrorMessage = "Enter a valid url."
	} else {
		//Save in DB

		data.NewURL = "new.url/e1wa3"
	}

	templateEngine.ExecuteTemplate(rw, "index", data)
}
