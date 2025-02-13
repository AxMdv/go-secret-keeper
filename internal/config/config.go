package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	grpcHostPortEnvName = "GRPC_HOST_PORT"
	dbDSNName           = "DB_DSN"

	logLvlName = "LOG_LEVEL"

	tokenTTLMinutesName = "TOKEN_TTL_MINUTES"
	secretJWTName       = "SECRET_JWT"
)

type Config struct {
	GRPCRunAddr string `json:"grpc_server_address"`
	// DSN for acees to DB.
	DataBaseDSN string `json:"database_dsn"`
	// Enable HTTPS
	EnableHTTPS bool `json:"enable_https"`
	//
	LoggingLevel string `json:"log_level"`
	//
	TokenTTLMinutes int `json:"tokenTTLInMinutes"`

	JWTSecretKey string `json:"-"`
}

func Parse() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	host := os.Getenv(grpcHostPortEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc server host and port not found")
	}
	dsn := os.Getenv(dbDSNName)
	if len(dsn) == 0 {
		return nil, errors.New("dsn not found")
	}
	logLvl := os.Getenv(logLvlName)
	if len(logLvl) == 0 {
		return nil, errors.New("log lvl not found")
	}
	jwtTTL := os.Getenv(tokenTTLMinutesName)
	if len(jwtTTL) == 0 {
		return nil, errors.New("jwtTTL not found")
	}
	ttlJWT, err := strconv.Atoi(jwtTTL)
	if err != nil {
		return nil, errors.New("fail to pars JWT ttl")
	}
	secretJWT := os.Getenv(secretJWTName)
	if len(secretJWT) == 0 {
		return nil, errors.New("secret jwt not found")
	}

	return &Config{
		GRPCRunAddr:     host,
		DataBaseDSN:     dsn,
		LoggingLevel:    logLvl,
		TokenTTLMinutes: ttlJWT,
		JWTSecretKey:    secretJWT,
	}, nil
}
