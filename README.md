# Gogit

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Godoc](https://godoc.org/github.com/tbobm/gogit?status.svg)](https://godoc.org/github.com/tbobm/gogit)

Golang-based CLI wrapper for github.com

## Usage

- Create a repository

```bash
$ gogit repo create -n repository_name [-p] # -p makes the repository private
```

- List a user's repositories

```bash
$ gogit repo list -u username
```

- Display help

```bash
$ gogit --help
```

## Installation

Export your freshly created 
[personnal access token](https://github.com/settings/token), as `GITHUB_TOKEN` 

## Why ?

I'm tired of having to use my browser to create repositories, have 
informations about my issues, pull requests, etc.

I just want to have a CLI tool that'll allow me to avoid going to
github.com everytime.

Plus I like my terminal font more than my browser's, in addition to that.

