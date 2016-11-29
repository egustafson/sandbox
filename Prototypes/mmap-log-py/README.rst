==========================================
Prototype mmap backed Log (Journal) Module
==========================================

...




Development
===========

Install locally using:

.. code::

    > python setup.py develop --user

The above command deploys the project in "Development Mode" using the
develop_ verb.  The project is deployed into the user's home directory
with the '--user' flag.

.. _develop:
https://setuptools.pypa.io/en/latest/setuptools.html#automatic-script-creation

The files will be deployed into a tree under ``~/.local``.  Adding
``~/.local/bin`` to your path is sufficient to execute the local
copies.


Testing
-------

Unit tests use `Python unittest` and can be run using Python Nose_ (v1)

.. _Python unittest:  https://docs.python.org/3/library/unittest.html
.. _Nose: https://nose.readthedocs.org/en/latest/

.. code::

   > nosetests



.. Local Variables:
.. mode: rst
.. End:
