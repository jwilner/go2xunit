# go2xunit 0.1.3

Converts `go test -v` output to xunit compatible XML output (used in
[Jenkins][jenkins]/[Hudson][hudson]).


# Install

    go get bitbucket.org/tebeka/go2xunit


# Usage
By default `go2xunit` reads data from standard input and emits XML to standard
output. However you can use `-input` and `-output` flags to change this.

The `-fail` switch will cause `go2xunit` to exit with non zero status if there
are failed tests.

    go test -v | go2xunit -output tests.xml

Here's an example script (`run-tests.sh`) that can be used with [Jenkins][jenkins]/[Hudson][hudson].

    #!/bin/bash

    outfile=gotest.out

    go test -v | tee $outfile
    go2xunit -fail -input $outfile -output tests.xml


[jenkins]: http://jenkins-ci.org/
[hudson]: http://hudson-ci.org/

Contact
=======
Miki Tebeka <miki.tebeka@gmail.com>

Bug reports go [here][bugs].

[bugs]: https://bitbucket.org/tebeka/go2xunit/issues

