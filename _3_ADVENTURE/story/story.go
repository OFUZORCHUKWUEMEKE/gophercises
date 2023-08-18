package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func init() {
	tpl := template.Must(template.New("").Parse(defaultHandlerTml))
}

var tpl *template.Template

var defaultHandlerTml = `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
	<style>
		.arcLink { margin-bottom: 10px; display: block; }
	</style>
</head>
<body>
	<h2>{{.Title}}</h2>
	<p>{{.Paragraph}}</p>
	{{range .Options}}
	<a class="arcLink" href="/?arc={{.Arc}}">{{.Text}}</a>
	{{end}}
</body>
</html>
`

type HandlerOption func(h *handler)

func NewHandler(s Story, t template.Template) http.Handler {
	h := handler{s, t}

	for _, opt := range opts {
		opt(&h)
	}
	return h

}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func (h handler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		log.Printf("%v", err)
		if err != nil {
			panic(err)
			http.Error(w, "Something went wrong", http.StatusBadRequest)
		}

	}
	http.Error(w, "Chapter not found", http.StatusNotExtended)

}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	fmt.Println(d)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
