all : pcap pcap_loop

pcap : pcap.c
	 gcc -o pcap -Wall -pedantic -g pcap.c -lpcap

pcap_loop : pcap_loop.c
	 gcc -o pcap_loop -Wall -pedantic -g pcap_loop.c -lpcap
