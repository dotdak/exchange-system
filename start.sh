#!/bin/sh
set -e

usage()
{
    echo "usage: start.sh [[-h]]"
}

start() {
  make build
  export $(cat .env | xargs)
  chmod +x ./app
  ./app
}

unittest() {
    make unit-test
}

integrationtest () {
    make integration-test-setup
    make integration-test-start
    make integration-test-teardown
}

test() {
    unittest;
    integrationtest
}

while [ "$1" != "" ]; do
    case $1 in
        -s | --start )              start
                                    ;;
        -ut | --unittest )          unittest
                                    ;;
        -it | --integrationtest )   integrationtest
                                    ;;
        -t | --test )               test
                                    ;;
        -h | --help )               usage
                                    exit
                                    ;;
        * )                         usage
                                    exit 1
    esac
    shift
done
