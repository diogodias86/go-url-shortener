package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

var templateEngine = template.Must(template.ParseGlob("templates/*.html"))

type data struct {
	ErrorMessage string
	NewURL       string
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
	fmt.Println("Redirecting...")

	//TODO: REMOVE
	if strings.Contains(path, "favicon") {
		return
	}

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
