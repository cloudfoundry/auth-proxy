package main

import (
	envstruct "code.cloudfoundry.org/go-envstruct"
)

type CAPI struct {
	Addr       string `env:"CAPI_ADDR,        required, report"`
	CAPath     string `env:"CAPI_CA_PATH,     required, report"`
	CommonName string `env:"CAPI_COMMON_NAME, required, report"`
}

type UAA struct {
	ClientID     string `env:"UAA_CLIENT_ID,     required"`
	ClientSecret string `env:"UAA_CLIENT_SECRET, required"`
	Addr         string `env:"UAA_ADDR,          required, report"`
	CAPath       string `env:"UAA_CA_PATH,       required, report"`
}

type Config struct {
	GatewayAddr      string `env:"GATEWAY_ADDR, required, report"`
	Addr             string `env:"ADDR, required, report"`
	InternalIP       string `env:"INTERNAL_IP, report"`
	HealthAddr       string `env:"HEALTH_ADDR, report"`
	CertPath         string `env:"EXTERNAL_CERT, required, report"`
	KeyPath          string `env:"EXTERNAL_KEY, required, report"`
	SkipCertVerify   bool   `env:"SKIP_CERT_VERIFY, report"`
	SecurityEventLog string `env:"SECURITY_EVENT_LOG, report"`

	CAPI CAPI
	UAA  UAA
}

func LoadConfig() (*Config, error) {
	cfg := Config{
		SkipCertVerify: false,
		Addr:           ":8083",
		InternalIP:     "0.0.0.0",
		HealthAddr:     "localhost:6065",
		GatewayAddr:    "localhost:8081",
	}

	err := envstruct.Load(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
