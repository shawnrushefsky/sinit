package templates

import "text/template"

// PackageJSON returns the readme template
func PackageJSON() (t *template.Template) {
	const raw = `{
		"name": "{{.Name}}",
		"version": "0.0.1",
		"description": "A description of ` + "`{{.Name}}`" + `",
		"main": "src/index.js",
		"directories": {
			"test": "test"
		},
		"scripts": {
			"test": "mocha --recursive",
			"coverage": "nyc mocha --recursive"
		},
		"keywords": [],
		"author": "{{.Author}}",
		"license": "UNLICENSED",
		"devDependencies": {
			"chai": "^{{.ChaiVersion}}",
			"eslint": "^{{.ESLintVersion}}",
			"mocha": "^{{.MochaVersion}}",
			"nyc": "^{{.NYCVersion}}"
		},
		"engines": {
			"node": "{{.NodeVersion}}"
		}
	}`

	t, err := template.New("package.json").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
