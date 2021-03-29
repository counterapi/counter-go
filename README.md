# Counter API Go Library

<p align="center">
  <a href="https://counterapi.dev/" target="_blank">
    <img width="180" src="https://raw.githubusercontent.com/counterapi/docs/master/src/.vuepress/public/favicons/apple-icon-180x180.png" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/counterapi/counter/Code%20Check" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/counterapi/counter/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/counterapi/counter" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/counterapi/counter"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/counterapi/counter/blob/master/LICENSE"><img src="https://img.shields.io/github/license/counterapi/counter" alt="License"></a>
</p>

Go Library for Counter API.

```shell
CLI command to counter things

Usage:
  counter [command]

Available Commands:
  counts      Fetches counts of counter
  down        Count down for given name
  get         Fetches counter information
  help        Help about any command
  set         Sets counter
  up          Count up for given name
  version     Print the version/build number

Flags:
  -h, --help   help for counter

Use "counter [command] --help" for more information about a command.

```

## Requirements

* Go installed.

## What does it do?

Go library for Counter API. It does provide both CLI and SDK.

## How to use it

1. Install Go package.

```shell
go get -u github.com/counterapi/counter
```

```shell
counter up --name MyCounter
```

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

## License

[MIT](https://github.com/counterapi/counter/blob/master/LICENSE)
