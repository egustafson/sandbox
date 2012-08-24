LIBS = -L/usr/local/lib -lomnithread -lomniORB4 -lomniCodeSets4


all : eg1

echo.hh echoSK.cc:
	omniidl -bcxx echo.idl


eg1 : eg1.cc echoSK.cc echo.hh
	gcc -c -I/usr/local/include echoSK.cc
	gcc -c -I/usr/local/include eg1.cc
	gcc -o eg1 echoSK.o eg1.o ${LIBS}

clean:
	rm -f *.o eg1 echoSK.* echo.hh
