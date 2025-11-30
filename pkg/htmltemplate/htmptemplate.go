package htmltemplate

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type HtmlTemplate struct {
	templatesDir   string
	templates      map[string]*template.Template
	cacheTemplates bool
}

func NewHtmlTemplate(templatesDir string, cache bool) *HtmlTemplate {
	return &HtmlTemplate{
		templatesDir:   templatesDir,
		templates:      make(map[string]*template.Template),
		cacheTemplates: cache,
	}
}

func (tr *HtmlTemplate) Render(w http.ResponseWriter, name string, data interface{}) error {
	var tmpl *template.Template
	var exists bool

	if tr.cacheTemplates {
		tmpl, exists = tr.templates[name]
	}

	if !exists {
		// Load templates as before
		layoutBase := filepath.Join(tr.templatesDir, "layouts", "base.html")
		layoutHeader := filepath.Join(tr.templatesDir, "layouts", "header.html")
		layoutFooter := filepath.Join(tr.templatesDir, "layouts", "footer.html")
		contentPath := filepath.Join(tr.templatesDir, name)

		var err error
		tmpl, err = template.New("base.html").Funcs(templateFuncs()).ParseFiles(
			layoutBase,
			layoutHeader,
			layoutFooter,
			contentPath,
		)
		if err != nil {
			log.Printf("Template parsing error: %v", err)
			return err
		}

		if tr.cacheTemplates {
			tr.templates[name] = tmpl
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, "base", data)
}
