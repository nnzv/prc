#!/usr/bin/env bash

# TODO: document

GO=$(which go)

die() {
    echo "$@"
    exit 1
}

_test() {
    test -z $1 && die 'test what?'
    GOFLAGS=( 
              'test' 
                '-v' 
             ${1}.go
        ${1}_test.go
    )
    [[ $(dirname $1) = "." ]] && GOFLAGS+=( 'util.go' )
    ${GO} ${GOFLAGS[*]}
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
