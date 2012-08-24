IDLDIR  = .

IDLS    = exservant
MODULES = ex

all: idls

idls:
	omniidl -I${IDLDIR} -bpython ${IDLDIR}/${IDLS}.idl

clean:
	rm -rf *.pyc ${MODULES} ${MODULES}__POA ${IDLS}_idl.py
