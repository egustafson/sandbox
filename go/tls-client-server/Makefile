#

build:
	(cd cmd/mkdemocerts; go build -o ../../mkdemocerts)
	(cd cmd/tlsloadcert; go build -o ../../tlsloadcert)
	(cd cmd/tlsclient; go build -o ../../tlsclient)
	(cd cmd/tlsserver; go build -o ../../tlsserver)

clean:
	rm -f mkdemocerts tlsloadcert tlsclient tlsserver
	rm -f cmd/mkdemocerts/mkdemocerts
	rm -f cmd/tlsloadcert/tlsloadcert
	rm -f cmd/tlsclient/tlsclient
	rm -f cmd/tlsserver/tlsserver
	rm -f *.pem
	rm -f cmd/mkdemocerts/*.pem
