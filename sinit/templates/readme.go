package templates

import "text/template"

// Readme returns the readme template
func Readme() (t *template.Template) {
	const raw = `# {{.Name}}
	The ` + "`{{.Name}}` project is authored by `{{.Author}}`"

	t, err := template.New("readme.md").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
