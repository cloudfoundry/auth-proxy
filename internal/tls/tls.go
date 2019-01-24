package tls

import (
	"crypto/tls"
)

type TLS struct {
	CAPath   string `env:"CA_PATH,   required, report"`
	CertPath string `env:"CERT_PATH, required, report"`
	KeyPath  string `env:"KEY_PATH,  required, report"`
}

func NewBaseTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
		CipherSuites:       supportedCipherSuites,
	}
}

var supportedCipherSuites = []uint16{
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
}
