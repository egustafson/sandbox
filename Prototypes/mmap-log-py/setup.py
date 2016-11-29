#!/usr/bin/env python

from setuptools import setup, find_packages

setup(
    name = "mmap-persistent-log-py",
    version = "0.1",
    author = "Eric Gustafson",
    author_email = "eg-git@elfwerks.org",
    license = "Apache Software License",
    packages = find_packages(exclude=['tests']),
    classifiers = [
        'Programming Language :: Python',
        'Development Status :: 1 - Planning',
        'License :: OSI Approved :: Apache Software License',
    ],
)
