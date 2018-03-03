#!/bin/bash

POSITIONAL=()
while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -v|--version)
    RELEASEVERSION="$2"
    shift # past argument
    shift # past value
    ;;
    *)    # unknown option
    POSITIONAL+=("$1") # save it in an array for later
    shift # past argument
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

if [ -z "${RELEASEVERSION}" ] ; then
    echo "releaseversion must be set"
    exit -1
fi

echo RELEASE VERSION  = "${RELEASEVERSION}"

while true; do
    read -p "Do you wish to release with this version? y/n: " yn
    case $yn in
        [Yy]* ) break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

git tag -a "${RELEASEVERSION}" -m "${RELEASEVERSION} release tag"
if [ $? -ne 0 ]; then { echo "Failed, aborting." ; exit 1; }; fi

git push origin "${RELEASEVERSION}"
if [ $? -ne 0 ]; then { echo "Failed, aborting." ; exit 1; }; fi

rm -rf dist
if [ $? -ne 0 ]; then { echo "Failed, aborting." ; exit 1; }; fi

goreleaser
if [ $? -ne 0 ]; then { echo "Failed, aborting." ; exit 1; }; fi

echo "done"
