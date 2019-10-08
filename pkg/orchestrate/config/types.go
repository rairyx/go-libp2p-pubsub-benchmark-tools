package config

const (
	defaultsLoc       string = "./defaults"
	defaultConfigName string = "config.default.json"
)

// Config is a struct to hold the config options
type Config struct {
	Orchestra Orchestra `json:"orchestra,omitempty"`
	Subnet    Subnet    `json:"subnet,omitempty"`
	Host      Host      `json:"host,omitempty"`
	General   General   `json:"general,omitempty"`
}

// Orchestra holds the orchestra configs
type Orchestra struct {
	// OmitSubnet omits building the subnet
	OmitSubnet bool `json:"omitSubnet,omitempty"`
	// HostsIfOmitSubnet are the host addresses if the subnet is omitted
	HostsIfOmitSubnet []string `json:"hostsIfOmitSubnet,omitempty"`
	// MessageNanoSecondInterval is the amount of nano seconds between each message send
	MessageNanoSecondInterval uint `json:"messageNanoSecondInterval,omitempty"`
	// ClientTimeoutSeconds is the time to wait to receive a response on the rpc channel
	ClientTimeoutSeconds int `json:"clientTimeoutSeconds,omitempty"`
	// MessageLocation is the location of the message json file
	MessageLocation string `json:"messageLocation,omitempty"`
	// MessageByteSize is the size of the message in bytes
	MessageByteSize uint `json:"messageByteSize,omitempty"`
	// TODO: implement
	//// MessageIntervalStdDev is the standard deviation of the message send interval
	//MessageIntervalStdDev uint
	//// MessageIntervalRampUpSeconds is the amount of time in seconds to ramp up to full send
	//MessageIntervalRampUpSeconds uint
}

// Subnet holds the configs for the subnet
type Subnet struct {
	// NumHosts is the number of hosts to spin up
	NumHosts int `json:"numHosts,omitempty"`
	// PubsubCIDR is the range of ip addrs for the pubsub to listen. Ports are incremented before IP
	PubsubCIDR string `json:"pubsubCIDR,omitempty"`
	// PubsubPortRange is the range of ports of pubsub to listen. Range is inclusive. Ports are incremented before IP
	PubsubPortRange [2]int `json:"pubsubPortRange,omitempty"`
	// RPCCIDR is the range of ip addrs for the rpc host to listen. Ports are incremented before IP
	RPCCIDR string `json:"rpcCIDR,omitempty"`
	// RPCPortRange is the range of ports for the rpc to listen. Range is inclusive. Ports are incremented before IP
	RPCPortRange [2]int `json:"rpcPortRange,omitempty"`
	// PeerTopology is the named peering topology
	PeerTopology string `json:"peerTopology,omitempty"`
}

// Host holds the configs for the hosts
type Host struct {
	// Transports are the transport protocols which the host is to use (e.g. "tcp", "ws", etc)
	Transports []string `json:"transports,omitempty"`
	// Muxers are the transport muxers (e.g. yamux, mplex, etc.)
	Muxers [][]string `json:"muxers,omitempty"`
	// Security specifies the security to use
	Security string `json:"security,omitempty"`
	// OmitRelay disables the relay
	OmitRelay bool `json:"omitRelay,omitempty"`
	// OmitConnectionManager disables the connection manager
	OmitConnectionManager bool `json:"omitConnectionManager,omitempty"`
	// OmitNatPortMap disables the nat port map
	OmitNATPortMap bool `json:"omitNATPortMap,omitempty"`
	// OmitRPCServer disables the rpc server
	OmitRPCServer bool `json:"omitRPCServer,omitempty"`
	// OmitDiscoveryService disables the discovery service
	OmitDiscoveryService bool `json:"omitDiscoveryService,omitempty"`
	// OmitBootstrapPeers disables bootstrapping of peers
	OmitBootstrapPeers bool `json:"omitBootstrapPeers,omitempty"`
	// OmitRouting disables ipfs routing (e.g. dht);
	// note: DHT is the only router supported, for now...
	OmitRouting bool `json:"omitRouting,omitempty"`
}

// General store general config directives
type General struct {
	// LogerLocation points to the log file. One will be create if not exists. Default is std out.
	LoggerLocation string `json:"loggerLocation,omitempty"`
}
