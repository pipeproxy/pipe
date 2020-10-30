module github.com/pipeproxy/pipe

go 1.15

require (
	github.com/gogf/greuse v1.1.0
	github.com/gorilla/handlers v1.4.2
	github.com/lucas-clemente/quic-go v0.18.0
	github.com/mikioh/tcp v0.0.0-20190314235350-803a9b46060c
	github.com/mikioh/tcpinfo v0.0.0-20190314235526-30a79bb1804b // indirect
	github.com/mikioh/tcpopt v0.0.0-20190314235656-172688c1accc // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/spf13/pflag v1.0.5
	github.com/wzshiming/crun v0.3.3
	github.com/wzshiming/funcfg v0.1.2
	github.com/wzshiming/lockfile v0.0.5
	github.com/wzshiming/notify v0.0.5
	github.com/wzshiming/trie v0.0.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20200821140526-fda516888d29 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/golang/protobuf => github.com/golang/protobuf v1.4.2
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/net => golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/sync => golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200821140526-fda516888d29
	golang.org/x/text => golang.org/x/text v0.3.3
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.3.0
)
