#!/usr/bin/env python

import requests

if __name__ == '__main__':
    msg = "PUTVAL hostname.domain/disk-sda/disk_octets interval=10 1251533299:197141504:175136768"
    headers = {'Content-Type': 'text/plain'}
    r = requests.post('http://localhost:8000/path', headers=headers, data=msg)
    print("Response: {}".format(r.status_code))
    r = requests.post('http://localhost:8000/path', headers=headers, data=msg)
    print("Response: {}".format(r.status_code))
