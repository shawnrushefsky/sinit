package templates

import "text/template"

// Codeowners returns the readme template
func Codeowners() (t *template.Template) {
	const raw = `# All files should be reviewed by {{.Author}}
	* {{.Email}}`

	t, err := template.New("codeowners").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
