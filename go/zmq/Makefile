all: zecho zping

clean:
	-rm -f zecho zping

zecho: zecho.go
	go build -o $@ $<

zping: zping.go
	go build -o $@ $<
