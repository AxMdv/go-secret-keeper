package config

type Config struct {
	ConfigPath string `json:"-"`
	// DSN for acees to DB.
	DataBaseDSN string `json:"database_dsn"`
	// Enable HTTPS
	EnableHTTPS bool `json:"enable_https"`
	//
	GRPCRunAddr string `json:"grpc_server_address"`
}
