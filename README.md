disjoint
========

[![Go Report Card](https://goreportcard.com/badge/github.com/spakin/disjoint)](https://goreportcard.com/report/github.com/spakin/disjoint) [![Build Status](https://travis-ci.org/spakin/disjoint.svg?branch=master)](https://travis-ci.org/spakin/disjoint) [![Go Reference](https://pkg.go.dev/badge/github.com/spakin/disjoint.svg)](https://pkg.go.dev/github.com/spakin/disjoint)

Introduction
------------

`disjoint` is a package for the [Go programming language](http://golang.org/) that implements a [disjoint-set data structure](http://en.wikipedia.org/wiki/Disjoint-set_data_structure) (also known as a union-find data structure).  Disjoint sets are collections of unordered elements in which an element belongs to exactly one set at a time.  Sets can be merged destructively, meaning that the the original sets cease to exist once their union is taken.  And elements can be compared for belonging to the same set.  These operations run in amortized near-constant time.

Installation
------------

`disjoint` is a [Go module](https://golang.org/ref/mod) and therefore should be installed automatically from any program that imports it with
```Go
import "github.com/spakin/disjoint"
```
It can also be installed manually by issuing a
```
go get github.com/spakin/disjoint
```
command.

Usage
-----

See the [`disjoint` API documentation](https://pkg.go.dev/github.com/spakin/disjoint) for details and examples.

Author
------

[Scott Pakin](http://www.pakin.org/~scott/), *scott+disj@pakin.org*
