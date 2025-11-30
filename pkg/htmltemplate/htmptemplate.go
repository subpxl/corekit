package htmltemplate

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// HtmlTemplate handles template rendering
type HtmlTemplate struct {
	templatesDir   string
	templates      map[string]*template.Template
	cacheTemplates bool
}

// NewHtmlTemplate creates a new renderer
func NewHtmlTemplate(templatesDir string, cache bool) *HtmlTemplate {
	return &HtmlTemplate{
		templatesDir:   templatesDir,
		templates:      make(map[string]*template.Template),
		cacheTemplates: cache,
	}
}

// Render renders the template and returns detailed wrapped errors
func (tr *HtmlTemplate) Render(w http.ResponseWriter, name string, data interface{}) error {
	var tmpl *template.Template
	var exists bool

	if tr.cacheTemplates {
		tmpl, exists = tr.templates[name]
	}

	if !exists {
		// Paths to template files
		layoutBase := filepath.Join(tr.templatesDir, "layouts", "base.html")
		layoutHeader := filepath.Join(tr.templatesDir, "layouts", "header.html")
		layoutFooter := filepath.Join(tr.templatesDir, "layouts", "footer.html")
		contentPath := filepath.Join(tr.templatesDir, name)

		var err error
		// Parse base template first
		tmpl, err = template.New("base.html").Funcs(templateFuncs()).ParseFiles(layoutBase)
		if err != nil {
			return &TemplateError{Op: "parse", File: layoutBase, Err: err}
		}

		// Parse other templates individually to know which file failed
		for _, path := range []string{layoutHeader, layoutFooter, contentPath} {
			_, err := tmpl.ParseFiles(path)
			if err != nil {
				return &TemplateError{Op: "parse", File: path, Err: err}
			}
		}

		if tr.cacheTemplates {
			tr.templates[name] = tmpl
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		return &TemplateError{Op: "execute", File: name, Err: err}
	}

	return nil
}
