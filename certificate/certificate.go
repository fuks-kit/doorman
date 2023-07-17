package certificate

import (
	"crypto/tls"
	_ "embed"
	"google.golang.org/grpc/credentials"
	"log"
)

//go:embed doorman-cert.pem
var Cert []byte

//go:embed doorman-key.pem
var Key []byte

func TLSCredentials() credentials.TransportCredentials {

	serverCert, err := tls.X509KeyPair(Cert, Key)
	if err != nil {
		log.Fatalln(err)
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config)
}
