disjoint
========

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

Copyright and License
---------------------

Copyright &copy; 2015 Scott Pakin

The `disjoint` package is released under the [BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause).  Informally, this means you can do whatever you want with the code as long as you give me (Scott Pakin) credit for writing the original version.

Author
------

[Scott Pakin](http://www.pakin.org/~scott/), *scott+npbm@pakin.org*
