module github.com/wzshiming/pipe

go 1.14

require (
	github.com/gorilla/handlers v1.4.2
	github.com/kubernetes-sigs/yaml v1.1.0
	github.com/spf13/pflag v1.0.5
	github.com/wzshiming/crun v0.3.3
	github.com/wzshiming/funcfg v0.1.1
	github.com/wzshiming/lockfile v0.0.5
	github.com/wzshiming/notify v0.0.5
	github.com/wzshiming/trie v0.0.1
	golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340
	golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sync => golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
)
