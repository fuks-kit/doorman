generate:
	openssl req -x509 -newkey rsa:4096 \
		-keyout doorman-key.pem -out doorman-cert.pem \
		-sha256 -days 3650 -nodes \
		-subj "/C=DE/L=Karlsruhe/O=fuks e.V./CN=doorman.local"
