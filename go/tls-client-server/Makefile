# Makefile - "I get by with a little help from my friends."

.PHONY: build
build:
	(cd cmd/mkdemocerts; go build -o ../../mkdemocerts)
	(cd cmd/tlsloadcert; go build -o ../../tlsloadcert)
	(cd cmd/tlsclient; go build -o ../../tlsclient)
	(cd cmd/tlsserver; go build -o ../../tlsserver)

.PHONY: certs
certs:
	( \
		cd cmd/mkdemocerts; \
		go build; \
		./mkdemocerts; \
	)
	ln -sf cmd/mkdemocerts/demo-ca-cert.pem .
	ln -sf cmd/mkdemocerts/demo-srv-cert.pem .
	ln -sf cmd/mkdemocerts/demo-srv-key.pem .

.PHONY: clean
clean:
	rm -f mkdemocerts tlsloadcert tlsclient tlsserver
	rm -f cmd/mkdemocerts/mkdemocerts
	rm -f cmd/tlsloadcert/tlsloadcert
	rm -f cmd/tlsclient/tlsclient
	rm -f cmd/tlsserver/tlsserver
	rm -f *.pem
	rm -f cmd/mkdemocerts/*.pem
