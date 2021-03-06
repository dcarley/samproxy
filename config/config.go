package config

// Config defines the interface the rest of the code uses to get items from the
// config. There are different implementations of the config using different
// backends to store the config. FileConfig is the default and uses a
// TOML-formatted config file. RedisPeerFileConfig uses a redis cluster to store
// the list of peers and then falls back to a filesystem config file for all
// other config elements.

type Config interface {

	// ReloadConfigs requests the configuration implementation to re-read
	// configs from whatever backs the source. Triggered by sending a USR1
	// signal to Samproxy
	ReloadConfig()

	// RegisterReloadCallback takes a name and a function that will be called
	// when the configuration is reloaded. This will happen infrequently. If
	// consumers of configuration set config values on startup, they should
	// check their values haven't changed and re-start anything that needs
	// restarting with the new values.
	RegisterReloadCallback(callback func())

	// GetListenAddr returns the address and port on which to listen for
	// incoming events
	GetListenAddr() (string, error)

	// GetPeerListenAddr returns the address and port on which to listen for
	// peer traffic
	GetPeerListenAddr() (string, error)

	// GetAPIKeys returns a list of Honeycomb API keys
	GetAPIKeys() ([]string, error)

	// GetPeers returns a list of other servers participating in this proxy cluster
	GetPeers() ([]string, error)

	// GetRedisHost returns the address of a Redis instance to use for peer
	// management. Only valid when the command line flag 'peer_type' is set to
	// 'redis'
	GetRedisHost() (string, error)

	// GetHoneycombAPI returns the base URL (protocol, hostname, and port) of
	// the upstream Honeycomb API server
	GetHoneycombAPI() (string, error)

	// GetLoggingLevel returns the verbosity with which we should log
	GetLoggingLevel() (string, error)

	// GetSendDelay returns the number of seconds to pause after a trace is
	// complete before sending it, to allow stragglers to arrive
	GetSendDelay() (int, error)

	// GetTraceTimeout is how long to wait before sending a trace even if it's
	// not complete. This should be longer than the longest expected trace
	// duration.
	GetTraceTimeout() (int, error)

	// GetSpanSeenDelay is a timer that bumps out sending the trace every time a
	// span is received. This one is used if you have traces of widely variable
	// duration, but don't want them to get sent until all spans arrive. Use with
	// care - if a trace continues to accumulate spans it may never get sent.
	GetSpanSeenDelay() (int, error)

	// GetOtherConfig attempts to fill the passed in struct with the contents of
	// a subsection of the config.   This is used by optional configurations to
	// allow different implementations of necessary interfaces configure
	// themselves
	GetOtherConfig(name string, configStruct interface{}) error

	// GetLoggerType returns the type of the logger to use. Valid types are in
	// the logger package
	GetLoggerType() (string, error)

	// GetCollectorType returns the type of the collector to use. Valid types
	// are in the collect package
	GetCollectorType() (string, error)

	// GetDefaultSamplerType returns the sampler type to use for all datasets
	// not explicitly defined
	GetDefaultSamplerType() (string, error)

	// GetSamplerTypeForDataset returns the sampler type to use for the given dataset
	GetSamplerTypeForDataset(string) (string, error)

	// GetMetricsType returns the type of metrics to use. Valid types are in the
	// metrics package
	GetMetricsType() (string, error)

	// GetUpstreamBufferSize returns the size of the libhoney buffer to use for the upstream
	// libhoney client
	GetUpstreamBufferSize() int
	// GetPeerBufferSize returns the size of the libhoney buffer to use for the peer forwarding
	// libhoney client
	GetPeerBufferSize() int
}
