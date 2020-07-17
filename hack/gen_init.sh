#!/bin/sh

echo "package init" > init.go
echo "import (" >> init.go
find ../components -name 'init.go' | sed "s#../#  _ \"$1/#" | sed 's#/init.go#\"#' >> init.go
echo ")" >> init.go
go fmt init.go