#!/bin/sh

echo "// Code generated. DO NOT EDIT!" > init.go
echo "package init" >> init.go
echo "import (" >> init.go
find ../components -name 'init.go' | grep -v "/common/gen/" | sed "s#../#  _ \"$1/#" | sed 's#/init.go#\"#' >> init.go
echo ")" >> init.go
go fmt init.go

echo "// Code generated. DO NOT EDIT!" > init_gen.go
echo "package init" >> init_gen.go
echo "import (" >> init_gen.go
find ../components -name 'init.go' | grep "/common/gen/" | sed "s#../#  _ \"$1/#" | sed 's#/init.go#\"#' >> init_gen.go
echo ")" >> init_gen.go
go fmt init_gen.go