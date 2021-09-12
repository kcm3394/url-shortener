# Simple URL shortener

My answer to this Gophercise: https://github.com/gophercises/urlshort

This URL shortener will look at the path of any incoming web request and redirect it to the appropriate page, or a
default "Hello, World!" page.

## How to run

```
$ go build main.go
$ go run main.go
```

With the default yaml file, go to localhost:8080/urlshort to be redirected to the original Gophercise repo.

## Flags
The file is customizable via flags.

```
$ ./main -h

Usage of ./main:
  -file string
        the file containing shortened paths to URLs in either yaml or json format (default "default.yml")
```
