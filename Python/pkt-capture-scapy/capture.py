#!/usr/bin/env python

# Note: the following program almost certainly needs to be run with
# root privileges.  The scapy library needs to place interfaces into
# promiscuous mode.

from scapy.all import sniff
from scapy.layers.inet import IP

# action to take when a packet is matched
def monitor_callback(pkt):
    if IP in pkt:
        return pkt.sprintf("%IP.src% -> %IP.dst%")


# BPF filter to match packets
#
#   the filter below catpures only the SA packet for a session on port 22
#
filter = "tcp[tcpflags] & (tcp-syn|tcp-ack) == tcp-syn|tcp-ack and port 22"

# number of packets to capture before exiting (0 = unlimited)
count = 2

# start sniffing
sniff(prn=monitor_callback, filter=filter, count=count)


