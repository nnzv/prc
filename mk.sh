#!/bin/sh

# TODO: document

GO=`which go`

_test() {
    test -z $1 && {
        echo "test what?"
        exit 1
    }
    ${GO} test -v $1.go $1_test.go
}

_http() {
    echo "visit: http://localhost:6060/pkg/git.sr.ht/~nzv/prc"
    echo "Press Ctrl+C to stop"
    godoc
}

case $1 in 
    test|t) _test $2 ;;
     doc|d) _http    ;;
esac
