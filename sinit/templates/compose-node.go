package templates

import "text/template"

// ComposeNode returns the Docker Compose file for a node project
func ComposeNode() (t *template.Template) {
	const raw = `version: '3'

	services:
		{{.Name}}:
			container_name: {{.Name}}
			image: node:{{.NodeVersion}}-stretch
			volumes:
				- ./:/code
			working_dir: /code
			command: sleep infinity`

	t, err := template.New("compose-node.yml").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
