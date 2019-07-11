# Gopher-cutter

[![CircleCI](https://circleci.com/gh/vlsidlyarevich/gopher-cutter/tree/master.svg?style=svg)](https://circleci.com/gh/vlsidlyarevich/gopher-cutter/tree/master)

Simple link shortener with MongoDB persist and Base62 link shorting algorithm. 

## Api usage

* GET http://host:port/cut?url=testurl - Shorten link and get short version
* GET http://host:port/shorturl - Proceed to full version of short link

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

You need to install Golang to run this project.

``` bash 
✗ go version
go version go1.12.1 darwin/amd64
```

Also you need MongoDB instance. You should add your MongoDB installation url and database name to [config.toml](https://github.com/vlsidlyarevich/gopher-cutter/blob/master/configs/config.toml#L9-L10)

### Installing

* Clone the project and open in your favourite IDE.
* Install all needed dependencies via `go get -d ./...` or via IDE tools.

## Running

Run [main.go](https://github.com/vlsidlyarevich/gopher-cutter/blob/master/cmd/gopher-cutter/main.go)

## Authors

**Vladislav Sidlyarevich** - [Github profile](https://github.com/vlsidlyarevich)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details