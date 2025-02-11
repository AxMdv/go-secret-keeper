package config

import "flag"

type Config struct {
	ConfigPath string `json:"-"`
	// DSN for acees to DB.
	DataBaseDSN string `json:"database_dsn"`
	// Enable HTTPS
	EnableHTTPS bool `json:"enable_https"`
	//
	GRPCRunAddr string `json:"grpc_server_address"`
}

func Parse() *Config {
	cfg := Config{}
	flag.StringVar(&cfg.GRPCRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&cfg.DataBaseDSN, "d", "", "dsn for acees to DB")
	return &cfg
}
