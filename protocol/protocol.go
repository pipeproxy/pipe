package protocol

type Protocol interface {
	WriteHeader(Map) error
	WriteBody(interface{}) error
	ReadHeader(Map) error
	ReadBody(interface{}) error
	Close() error
}

type Map interface {
	Add(key string, value string)
	Set(key string, values []string)
	Get(key string) []string
	Del(key string)
	Range(func(key string, values []string) bool)
	Clone() Map
}
