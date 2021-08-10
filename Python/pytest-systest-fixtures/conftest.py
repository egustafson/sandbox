# -*- coding: utf-8 -*-

# pytest configuration

import pytest


def pytest_addoption(parser):
    parser.addoption(
        "--systest",
        action='store_true',
        help="run system tests, (requires service to be running)"
    )


def pytest_configure(config):
    config.addinivalue_line(
        "markers", "systest: run system tests"
    )


def pytest_runtest_setup(item):
    for m in item.iter_markers(name="systest"):
        if not item.config.getoption('--systest'):
            pytest.skip("skipping system tests")
            return
