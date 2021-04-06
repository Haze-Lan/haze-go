package discovery

type DiscoveryOption struct {
	DiscoveryType string
	server        *ServerOption
	config        *ConfigOption
}

type ServerOption struct {
	host string
	port uint8
}

type ConfigOption struct {
}
