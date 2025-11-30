package htmltemplate

import (
	"fmt"
	"os"
	"path/filepath"
)

func (tr *HtmlTemplate) CreateSkeleton() error {
	// Directories to create
	layoutsDir := filepath.Join(tr.templatesDir, "layouts")
	pagesDir := filepath.Join(tr.templatesDir, "pages")

	dirs := []string{tr.templatesDir, layoutsDir, pagesDir}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", dir, err)
			}
		}
	}

	// Default templates
	files := map[string]string{
		filepath.Join(layoutsDir, "base.html"): `
{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
    {{ template "header" . }}
    {{ template "content" . }}
    {{ template "footer" . }}
</body>
</html>
{{ end }}
`,
		filepath.Join(layoutsDir, "header.html"): `
{{ define "header" }}
<header>
    <h1>Welcome, {{ .User }}</h1>
</header>
{{ end }}
`,
		filepath.Join(layoutsDir, "footer.html"): `
{{ define "footer" }}
<footer>
    <p>&copy; 2025 MySite</p>
</footer>
{{ end }}
`,
		filepath.Join(pagesDir, "home.html"): `
{{ define "content" }}
<h2>Items List</h2>
<ul>
    {{ range .Items }}
        <li>{{ . }}</li>
    {{ end }}
</ul>
{{ end }}
`,
	}

	// Create files if they do not exist
	for path, content := range files {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				return fmt.Errorf("failed to create file %s: %v", path, err)
			}
		}
	}

	return nil
}
