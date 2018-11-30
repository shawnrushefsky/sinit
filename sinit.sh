#!/bin/bash

PROJECT_NAME=${1:-"new-project"}
LOWER_STACK=$(echo "$2" | tr '[:upper:]' '[:lower:]')
STACK=${LOWER_STACK:-"node"}

mkdir $PROJECT_NAME
cd $PROJECT_NAME
git init

touch README.md
echo "# $PROJECT_NAME
A description of your project
" > README.md

touch CHANGELOG.md
echo "# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
- The entirety of $PROJECT_NAME
" > CHANGELOG.md

if [ $STACK = node ];
then
    echo "Creating a Node.js Project..."
    npm init -y
    npm install --save-dev mocha chai eslint

    touch .eslintrc.js
    curl https://raw.githubusercontent.com/shawnrushefsky/sinit/master/resources/.eslintrc.js > .eslintrc.js

    touch .gitignore
    curl https://raw.githubusercontent.com/github/gitignore/master/Node.gitignore > .gitignore
fi