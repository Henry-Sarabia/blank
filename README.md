# blank

[![GoDoc](https://godoc.org/github.com/Henry-Sarabia/blank?status.svg)](https://godoc.org/github.com/Henry-Sarabia/blank)
[![Build Status](https://travis-ci.com/Henry-Sarabia/blank.svg?branch=master)](https://travis-ci.com/Henry-Sarabia/blank)
[![Go Report Card](https://goreportcard.com/badge/github.com/Henry-Sarabia/blank)](https://goreportcard.com/report/github.com/Henry-Sarabia/blank)
[![Coverage Status](https://coveralls.io/repos/github/Henry-Sarabia/blank/badge.svg?branch=master)](https://coveralls.io/github/Henry-Sarabia/blank?branch=master)

The **Blank** package offers two main functionalities.

**Blank** can remove whitespace from a string.
The package defines whitepsace as a character that is not typically visible.
These characters range anywhere from the ordinary space to a less common vertical tab.

**Blank** can check if a string is blank.
The package considers a string to be blank if it is comprised solely of whitespace.


## Installation 

If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To pull the most recent version of **Blank**, use `go get`.

```
go get -u github.com/Henry-Sarabia/blank
```

Then import the package into your project.

```go
import "github.com/Henry-Sarabia/blank"
```

## Usage

### Whitespace Removal

The package considers whitespace to be any character that is not typically visible.
The most common of these characters are: space, tab, newline, return, formfeed, nonbreaking space, and vertical tab.
For more information, visit the [unicode package](https://golang.org/pkg/unicode/#IsSpace) and the [unicode seperator category](http://www.fileformat.info/info/unicode/category/Zs/list.htm).

To remove the whitespace from a string, use the `Remove` function.

```go
phrase := "this is a phrase"

str := blank.Remove(phrase)

fmt.Println(str)
// output: "thisisaphrase"
```

### Blank Detection

The package considers a string to be blank if it is comprised solely of whitespace.

For example, assume we are creating a search function that takes a string as a search query.
We want to avoid searching for blank queries.
Blank queries can be detected using the `Is` function.

```go
func search(qry string) error {
	if blank.Is(qry) {
		// return error
	}
	
	// rest of code
}
```

Similarly, the `Has` function can process an entire slice of strings; it will check if any of the strings are blank.

Let's slightly alter our example.
Assume the search function takes a slice of strings as a list of queries.
We still want to avoid seraching for blank queries.
Blank queries can be detected using the `Has` function.

```go
func search(qrs []string) error {
	if blank.Has(qrs) {
		// return error
	}
	
	// rest of code
}
```

## Contributions

If you would like to contribute to this project, please adhere to the following guidelines.

* Submit an issue describing the problem.
* Fork the repo and add your contribution.
* Add appropriate tests.
* Run go fmt, go vet, and golint.
* Prefer idiomatic Go over non-idiomatic code.
* Follow the basic Go conventions found [here](https://github.com/golang/go/wiki/CodeReviewComments).
* If in doubt, try to match your code to the current codebase.
* Create a pull request with a description of your changes.

I'll review pull requests as they come in and merge them if everything checks out.

Any and all contributions are greatly appreciated. Thank you!