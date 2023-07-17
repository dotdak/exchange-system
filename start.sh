#!/bin/sh
USAGE="Usage: start.sh -s | -ls | -ut | -it | -t args"

set -e


usage()
{
    echo $USAGE
}

if [ $# == 0 ] ; then
    usage
    exit 1;
fi

local() {
  make local
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
        -ls | --local-start )       local
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
