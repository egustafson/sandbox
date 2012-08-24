#include <pcap.h>
#include <stdio.h>
int main()
{
    pcap_t *handle=0;                        /* Session handle */
    char* dev;                                /* The device to sniff on */
    char errbuf[PCAP_ERRBUF_SIZE]; /* Error string */
    struct bpf_program filter;            /* The compiled filter */
    char filter_app[] = "port 23";       /* The filter expression */
    bpf_u_int32 mask;                     /* Our netmask */
    bpf_u_int32 net;                        /* Our IP */
    struct pcap_pkthdr header;          /* The header that pcap gives us */
    const u_char *packet;                 /* The actual packet */
    int result;
 
    /* Define the device */
    dev = pcap_lookupdev(errbuf);

    printf("Listening on %s.\n", dev);

    /* Find the properties for the device */
    result = pcap_lookupnet(dev, &net, &mask, errbuf);
    printf("pcap_lookupnet() -> %d\n", result);
    printf("net  = 0x%08x\n", net);
    printf("mask = 0x%08x\n", mask);

    /* Open the session in promiscuous mode */
    handle = pcap_open_live(dev, BUFSIZ, 1, -1, errbuf);
    printf("handle = 0x%08x\n", handle);

    printf("BUFSIZ = %d\n", BUFSIZ);

    /* Compile and apply the filter */
/*     pcap_compile(handle, &filter, filter_app, 0, net); */
    pcap_compile(handle, &filter, filter_app, 0, mask);

    result = pcap_setfilter(handle, &filter); 
    printf("result = %d\n", result);

    sleep(10);

    printf("Grabbing a packet ...\n");
    /* Grab a packet */
    packet = pcap_next(handle, &header);

    /* Print its length */
    if ( packet ) {
        printf("Jacked a packet with length of [%d]\n", header.len);
    } else {
        printf("did not capture a packet.\n");
    }

    /* And close the session */
    pcap_close(handle);
    return(0);
}
