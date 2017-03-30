package main
import (
	"io/ioutil"
	"net/http"
	"regexp"
	"fmt"
	"os/exec"
	"os"
	"time"
)

type Script struct {
	Title string
	Body  []byte
}

/*
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}*/

func compileHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Printf("Received %s\n", title)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	s := &Script{Title: title, Body: []byte(body)}
  filename := s.Title + ".mq4";
	err = ioutil.WriteFile("MQL4/Experts/" + filename, s.Body, 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  exec.Command("./compile.sh", filename).Run()
	file, err := os.Open("MQL4/Experts/" + s.Title + ".ex4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  http.ServeContent(w, r, "result.ex4", time.Now(), file)
}


var validPath = regexp.MustCompile("^/(compile)/([a-zA-Z0-9]+)$")
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/compile/", makeHandler(compileHandler))
	http.ListenAndServe(":8080", nil)
}
