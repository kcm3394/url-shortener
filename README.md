# Simple URL shortener

My answer to this Gophercise: https://github.com/gophercises/urlshort

This URL shortener will look at the path of any incoming web request and redirect it to the appropriate page, or a
default "Hello, World!" page.

## How to run

```
$ go build main/main.go
$ go run main/main.go
```

With the default yaml file, go to localhost:8080/urlshort to be redirected to the original Gophercise repo.

## Flags
The yaml file is customizable via flags.

```
$ cd main
$ ./main -h

Usage of ./main:
  -yaml string
        a yaml file in the format of 
        '- path: /some-path
           url: https://www.some-url.com/demo' (default "default.yml")
```
