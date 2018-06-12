#!/bin/bash

function pause()
{
    while true; do
        read -p "Continue? y/n: " yn
        case $yn in
            [Yy]* ) break;;
            [Nn]* ) exit;;
            * ) echo "Please answer yes or no.";;
        esac
    done
}

echo "go run Aareguru.go --help"
go run Aareguru.go --help

pause

echo "go run Aareguru.go cities"
go run Aareguru.go cities

pause

echo "go run Aareguru.go cities -f"
go run Aareguru.go cities -f

pause

echo "go run Aareguru.go"
go run Aareguru.go

pause

echo "go run Aareguru.go standard"
go run Aareguru.go standard

pause

echo "go run Aareguru.go standard -l"
go run Aareguru.go standard -l

pause

echo "go run Aareguru.go standard -f"
go run Aareguru.go standard -f

pause

echo "go run Aareguru.go schribmaschine"
go run Aareguru.go schribmaschine

pause

echo "go run Aareguru.go standard biel"
go run Aareguru.go standard biel

pause

echo "go run Aareguru.go schribmaschine interlaken"
go run Aareguru.go schribmaschine interlaken
