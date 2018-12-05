# Gogit

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Godoc](https://godoc.org/github.com/tbobm/gogit?status.svg)](https://godoc.org/github.com/tbobm/gogit)

Golang-based CLI wrapper for github.com

## Usage

### Beforehand

Before anything else, export your freshly created 
[personnal access token](https://github.com/settings/token), as `GITHUB_TOKEN` 
You can do so by including the `export GITHUB_TOKEN=xxxx...` in your `~/.bashrc`
or `~/.profile`.

### Commands

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
$ gogit --help # Or simply 'gogit', which defaults to displaying the help
```

## Installation

### From a release

Download the corresponding archive from [the release page](https://github.com/tbobm/gogit/releases),
extract its content, and place the `gogit` executable in your `$PATH` variable.

### From source

Clone the master branch, and run `make install`.
This will run the `go install` command, placing the `gogit` executable in your `$GOPATH`

[dep](https://github.com/golang/dep) is used to manage packages.


```bash
$ git clone https://github.com/tbobm/gogit
$ cd gogit
$ make install
```

## Why ?

I'm tired of having to use my browser to create repositories, have 
informations about my issues, pull requests, etc.

I just want to have a CLI tool that'll allow me to avoid going to
github.com everytime.

Plus I like my terminal font more than my browser's, in addition to that.

## License

[Apache License](https://github.com/tbobm/gogit/blob/master/LICENSE)

Copyright (c) 2017 Theo Massard
