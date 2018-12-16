package templates

import "text/template"

// CircleNode returns the circle ci job to run tests on a node project
func CircleNode() (t *template.Template) {
	const raw = `version: 2
jobs:
	build:
		docker:
			- image: circleci/node:{{.NodeVersion}}-stretch

		steps:
			- checkout

			# Download and cache dependencies
			- restore_cache:
					keys:
					- v1-dependencies-{{"{{"}} checksum "package.json" {{"}}"}}
					# fallback to using the latest cache if no exact match is found
					- v1-dependencies-

			- run: 
					name: Install Dependencies
					command: npm install

			- save_cache:
					paths:
						- node_modules
					key: v1-dependencies-{{"{{"}} checksum "package.json" {{"}}"}}
				
			# run tests!
			- run: 
					name: Run Tests
					command: npm test
`

	t, err := template.New("circle-node.yml").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
