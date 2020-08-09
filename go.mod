module github.com/wzshiming/pipe

go 1.14

require (
	github.com/gorilla/handlers v1.4.2
	github.com/lucas-clemente/quic-go v0.17.3
	github.com/prometheus/client_golang v0.8.0
	github.com/spf13/pflag v1.0.5
	github.com/wzshiming/crun v0.3.3
	github.com/wzshiming/funcfg v0.1.2
	github.com/wzshiming/lockfile v0.0.5
	github.com/wzshiming/notify v0.0.5
	github.com/wzshiming/trie v0.0.1
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20200808120158-1030fc2bf1d9 // indirect
	golang.org/x/text v0.3.3 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/golang/protobuf => github.com/golang/protobuf v1.4.2
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	golang.org/x/net => golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sync => golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200808120158-1030fc2bf1d9
	golang.org/x/text => golang.org/x/text v0.3.3
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.3.0
)
