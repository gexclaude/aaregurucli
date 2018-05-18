#!/bin/bash

gofmt -w .

find . -type d | xargs -L 1 golint