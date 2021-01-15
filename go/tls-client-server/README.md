Go Example - basic TLS client / server (TCP)
============================================

This directory (and sub-dirs) contains a complete and functional,
Golang example of TLS client and server using basic TCP.  The bulk of
the example is directly applicable to net/http as well.  Also included
is an example of building PEM encoded certificates and keys.  The demo
is intended to work with "bring your own certs", but has only been
tested with the certs generated from the 'mkdemocert'.

Usage
-----

Setup:
```
> make clean
> make build  # places all 4 exec's in the top directory.
> make cert   # uses mkdemocert and leaves the certs in ./
```

Show that the certs can be read (parsed):
```
> tlsloadcert
loading CA from:    demo-ca-cert.pem
loading cert from:  demo-srv-cert.pem
loading key from:   demo-srv-key.pem
CA block is: CERTIFICATE
initial block is:  CERTIFICATE
key block is:  RSA PRIVATE KEY
ok.
>
```

### Client / Server

| Server (window) | Client (window) |
| --------------- | --------------- |
| `> tlsserver`   | `> tlsclient`   |

Server Window:
```
> tlsserver
2021/01/14 17:03:16 loading CA from:    demo-ca-cert.pem
2021/01/14 17:03:16 loading cert from:  demo-srv-cert.pem
2021/01/14 17:03:16 loading key from:   demo-srv-key.pem
2021/01/14 17:03:16 listening ...
2021/01/14 17:03:21 connected
2021/01/14 17:03:21 successfully wrote 21 bytes
2021/01/14 17:03:21 connection closed
2021/01/14 17:03:21 done.
>
```

Client Window:
```
> tlsclient
2021/01/14 17:03:21 loading CA from:    demo-ca-cert.pem
2021/01/14 17:03:21 loading cert from:  demo-srv-cert.pem
2021/01/14 17:03:21 loading key from:   demo-srv-key.pem
2021/01/14 17:03:21 msg recv: response from server
2021/01/14 17:03:21 done.
>
```


References
----------

The following citations from the web were used in creating this example.

* [Creating a Certificate Authority + Signing Certificates in Go](https://shaneutt.com/blog/golang-ca-and-signed-cert-go/)



