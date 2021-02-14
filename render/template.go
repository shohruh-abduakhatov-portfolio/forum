package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	model "forum.com/model"
)

var (
	templates     *template.Template
	layout        *template.Template
	layoutFuncs   map[string]interface{}
	errorTemplate string
)

func init() {
	layoutFuncs = template.FuncMap{
		"yield": func() (string, error) {
			return "", fmt.Errorf("yield called inappropriately")
		},
	}
	templates = template.Must(template.New("t").ParseGlob(filepath.Join(".", "templates", "**", "*.html")))
	layout = template.Must(
		template.
			New("base.html").
			Funcs(layoutFuncs).
			ParseFiles("templates/base.html"),
	)
	errorTemplate = `
	<html>
		<body>
			<h1>Error rendering template %s</h1>
			<p>%s</p>
		</body>
	</html>
	`
}

func Template(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	data["CurrentUser"] = model.RequestUser(r)
	data["Flash"] = r.URL.Query().Get("flash")

	log.Println(">>> CurrentUser", data["CurrentUser"])

	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}
	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)

	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}

func Basic(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	data["CurrentUser"] = model.RequestUser(r)
	data["Flash"] = r.URL.Query().Get("flash")

	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}
