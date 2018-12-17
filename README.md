# sinit
`sinit` is a project initiation tool that creates code projects in a style I personally like. This is only intended to automate common work patterns for me personally, but if you get use out it, too, awesome! It is written in Go, because I wanted to learn Go.

# install

clone down the repo, and run `go build -o /usr/local/bin/sinit`

# use
```
Usage of sinit:
  -n -name string
    	The name of your new project (default "new-project")
  -s -stack string
    	The project stack you want to create. e.g. node (default "node")
  -t -theme string
    	The name of the theme you want to use (for static sites, etc.) (default "mediumish")
  -u -theme-repo string
    	The URL for the repo of the theme you want to use (for hugo) (default "https://github.com/lgaida/mediumish-gohugo-theme.git")
  -d -deploy string
    	The deployment target, e.g. s3
```

# Stacks

## node
This will create a new node project with mocha, chai, eslint, and nyc. It also explicitly notates your node version in the package.json, .nvmrc, and in the docker-compose.yml. It also sets up a circleci pipeline to run tests on commit-to-master.

## hugo
This creates a hugo static site, and installs a theme for you. It will automatically provide a circleci job to build the site and persist it to the workspace. You may optionally specify a deployment target.

Currently supported deployment targets:
- `s3`
- `gcs`
- `firebase`
