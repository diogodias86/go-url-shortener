package handlers

import (
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/diogodias86/go-url-shortener/urlgenerator"

	"github.com/diogodias86/go-url-shortener/db"
)

var templateEngine = template.Must(template.ParseGlob("templates/*.html"))

type data struct {
	ErrorMessage string
	NewURL       string
}

//FavIconHandler ...
func FavIconHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "templates/favicon.ico")

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
	path = strings.Replace(path, "/", "", -1)

	//Get the original URL from DB
	url := db.GetURL(path)

	if url == "" {
		//404 - Not Found
		http.NotFound(rw, r)
	} else {
		//Redirect
		http.Redirect(rw, r, url, http.StatusMovedPermanently)
	}
}

func postURL(rw http.ResponseWriter, r *http.Request) {
	_url := r.FormValue("url")
	_, err := url.ParseRequestURI(_url)

	data := &data{}

	//Basic validation
	if _url == "" || err != nil {
		data.ErrorMessage = "Enter a valid url."
	} else {
		//Generate the new url
		newURL := urlgenerator.Generate()

		//Save in DB
		db.Insert(_url, newURL)

		data.NewURL = r.Host + "/" + newURL
	}

	templateEngine.ExecuteTemplate(rw, "index", data)
}
