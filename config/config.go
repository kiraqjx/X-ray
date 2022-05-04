package config

type ServerConfig struct {
	Grpc GrpcConfig
}

type GrpcConfig struct {
	Port uint32
}

type ClientConfig struct {
	GrpcServer GrpcServerConfig `yaml:"grpc-server"`
	ProxyHosts []ProxyHost      `yaml:"proxy-hosts"`
}

type GrpcServerConfig struct {
	Host string
}

type ProxyHost struct {
	Protocol int32
	Host     string
	Port     uint32
}
