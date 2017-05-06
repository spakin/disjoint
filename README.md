disjoint
========

[![Go Report Card](https://goreportcard.com/badge/github.com/spakin/disjoint)](https://goreportcard.com/report/github.com/spakin/disjoint) [![GoDoc](https://godoc.org/github.com/spakin/disjoint?status.svg)](https://godoc.org/github.com/spakin/disjoint)

Introduction
------------

`disjoint` is a package for the [Go programming language](http://www.golang.org/) that implements a [disjoint-set data structure](http://en.wikipedia.org/wiki/Disjoint-set_data_structure) (also known as a union-find data structure).  Disjoint sets are collections of unordered elements in which an element belongs to exactly one set at a time.  Sets can be merged destructively, meaning that the the original sets cease to exist once their union is taken.  And elements can be compared for belonging to the same set.  These operations run in amortized near-constant time.

Installation
------------

Instead of manually downloading and installing `disjoint` from GitHub, the recommended approach is to ensure your `GOPATH` environment variable is set properly then issue a

    go get github.com/spakin/disjoint

command.

Usage
-----

See the [`disjoint` API documentation](http://godoc.org/github.com/spakin/disjoint) for details.

Author
------

[Scott Pakin](http://www.pakin.org/~scott/), *scott+disj@pakin.org*
