#!/usr/bin/env bash

# TODO: document

GO=$(which go)

_test() {
    FILES=( 
             "${1}.go"
        "${1}_test.go"
    )
    test -z ${1} && FILES=()
    if [[ $(dirname ${1}) = . ]]; then
        FILES+=("util.go")
    fi
    ${GO} test -v ${FILES[*]}
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
