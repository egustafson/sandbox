## Makefile.am -- Process this file with automake to produce Makefile.in

ACLOCAL_AMFLAGS = -I m4

SUBDIRS = src

PKGDIR = /tmp/$(PACKAGE)-$(VERSION)
pkg-deb: all
	make install DESTDIR=$(PKGDIR)
	fpm -s dir -t deb -n $(PACKAGE) -v $(VERSION) -C $(PKGDIR) .
	-rm -rf $(PKGDIR)

