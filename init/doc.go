package init

//go:generate sh -c "echo \"package init\n\nimport (\" > init.go"
//go:generate sh -c "find ../components -name 'init.go' | sed 's#../#\t_ \"github.com/wzshiming/pipe/#' | sed 's#/init.go#\"#' >> init.go"
//go:generate sh -c "echo \")\" >> init.go"
//go:generate go fmt init.go
