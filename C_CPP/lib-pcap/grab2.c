#include <pcap.h>
#include <stdio.h>
int main()
{
    pcap_t *handle;                        /* Session handle */
    char *dev;                                /* The device to sniff on */
    char errbuf[PCAP_ERRBUF_SIZE]; /* Error string */
    struct bpf_program filter;            /* The compiled filter */
    char filter_app[] = "port 22";       /* The filter expression */
    bpf_u_int32 mask;                     /* Our netmask */
    bpf_u_int32 net;                        /* Our IP */
    struct pcap_pkthdr header;          /* The header that pcap gives us */
    const u_char *packet;                 /* The actual packet */
    /* Define the device */
    dev = pcap_lookupdev(errbuf);
    /* Find the properties for the device */
    pcap_lookupnet(dev, &net, &mask, errbuf);
    /* Open the session in promiscuous mode */
    handle = pcap_open_live(dev, BUFSIZ, 1, 0, errbuf);
    /* Compile and apply the filter */
    pcap_compile(handle, &filter, filter_app, 0, net);
    pcap_setfilter(handle, &filter);
    /* Grab a packet */
    packet = pcap_next(handle, &header);
    /* Print its length */
    printf("Jacked a packet with length of [%d]\n", header.len);
    /* And close the session */
    pcap_close(handle);
    return(0);
}
