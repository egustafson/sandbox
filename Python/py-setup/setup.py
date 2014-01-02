#!/usr/bin/env python

from distutils.core import setup

setup( name='example',
       version='1.0',
       description='An example distribution for experimentation',
       author='Eric Gustafson',
       author_email='eg-git@elfwerks.org',
       url='http://www.elfwerks.org',
       py_modules=['ex_module'],
       packages=['ex_pkg1', 'ex_pkg2'],
       scripts=['ex_prog'],
     )

## Local Variables:
## mode: python
## End:

