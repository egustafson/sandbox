# -*- coding: utf-8 -*-

# This is a simple "system test case" that should only be run when
# system tests are explicitly run.  This test should not be run during
# normal unit testing.

import pytest
import requests


@pytest.fixture
def endpoint():
    return "http://localhost:5000/"


@pytest.mark.systest
def test_systest(endpoint):
    """A simple (stub) SYSTEM-TEST - NOT always run"""
    r = requests.get(endpoint)
    assert r.encoding == 'utf-8'
    j = r.json()
    assert j is not None
    print(j)
    assert j['status'] == "ok"
