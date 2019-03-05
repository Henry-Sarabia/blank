# blank

[![GoDoc](https://godoc.org/github.com/Henry-Sarabia/blank?status.svg)](https://godoc.org/github.com/Henry-Sarabia/blank) [![Build Status](https://travis-ci.com/Henry-Sarabia/blank.svg?branch=master)](https://travis-ci.com/Henry-Sarabia/blank) [![Go Report Card](https://goreportcard.com/badge/github.com/Henry-Sarabia/blank)](https://goreportcard.com/report/github.com/Henry-Sarabia/blank) [![Coverage Status](https://coveralls.io/repos/github/Henry-Sarabia/blank/badge.svg?branch=master)](https://coveralls.io/github/Henry-Sarabia/blank?branch=master)

Blank implements blank checking and whitespace removal for strings. There are times when you
only need to make sure your string arguments are not empty. However, there are other times when you
need to not only make sure they aren't empty, but also not blank or made up of whitespace. This is
where the blank package can help.

## Installation 

If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To pull the most recent version of blank, use `go get`.

```
go get github.com/Henry-Sarabia/blank
```

Then import the package into your project.

```go
import "github.com/Henry-Sarabia/blank"
```

## Usage

### Blank checking

This package considers a string to be blank if it is solely made up of whitespace.

As an example, assume you are creating a search function that takes a string argument as the search
query. You want to avoid searching for any empty queries. This includes both empty strings and 
blank strings. The first case can be resolved by comparing the string against the empty string.
The second case is where the blank package and the `Is(string)` function is most useful.

```go
func search(qry string) error {
	if len(qry) == 0 {
		// return some error
	}
	
	if blank.Is(qry) {
		// return some other error
	}
	
	// rest of code
}
```

Similarly, the `Has([]string)` function can check an entire slice of strings for blanks.

### Whitespace Removal

This package considers whitespace to be any of these common characters: space, tab, newline, and
return. Other uncommon whitespace characters such as nonbreaking space and vertical tab are now supported
as well.

To remove the whitespace from a string, use the `Remove(string)` function.

```go
phrase := "this is a phrase"

str := blank.RemoveSpace(phrase)

fmt.Println(str)
// output: "thisisaphrase"
```

## Contributions

If you would like to contribute to this project, please adhere to the following
guidelines.

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