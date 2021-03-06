package templates

// CircleHugo returns the build job for a hugo site
func CircleHugo() string {
	return `version: 2

jobs:
	build:
		docker:
			- image: cibuilds/hugo:latest
		working_directory: ~/hugo
		steps:
			# checkout the repository
			- checkout

			# install git submodules for managing third-party dependencies
			- run: git submodule sync && git submodule update --init

			# build with Hugo
			- run: HUGO_ENV=production hugo -v

			# The built site will be in <workspace>/hugo/public
			- persist_to_workspace:
					root: ~/
					paths:
						- hugo
`
}
