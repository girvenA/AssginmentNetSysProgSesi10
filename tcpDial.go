package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func main() {
	addr := "www.google.com:443"
	dial := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	configTLS := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.DialWithDialer(dial, "tcp", addr, configTLS)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	fmt.Printf("TLS Version: %s\n", tellVer(state.Version))
	fmt.Printf("Cipher Suite: %s\n", tls.CipherSuiteName(state.CipherSuite))

	for _, cert := range state.PeerCertificates {
		fmt.Printf("Issuer Organization: %s\n", cert.Issuer.Organization)
	}
}

func tellVer(ver uint16) string {
	switch ver {
	case tls.VersionTLS10:
		return "TLS Version: TLS 1.0"
	case tls.VersionTLS11:
		return "TLS Version: TLS 1.1"
	case tls.VersionTLS12:
		return "TLS Version: TLS 1.2"
	case tls.VersionTLS13:
		return "TLS Version: TLS 1.3"
	default:
		return "Unknown TLS version"
	}
}
