package templates

import "text/template"

// Changelog returns the CHANGELOG.md template
func Changelog() (t *template.Template) {
	const raw = `# Changelog
	All notable changes to this project will be documented in this file.
	
	The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
	and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
	
	## [Unreleased]
	- The entirety of ` + "`{{.Name}}`"

	t, err := template.New("changelog.md").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
