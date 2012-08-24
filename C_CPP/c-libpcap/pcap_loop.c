#include <stdio.h>
#include <pcap.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <netinet/in_systm.h>
#include <net/if.h>
#include <netinet/if_ether.h>
#include <netinet/ip.h>
#include <netinet/tcp.h>


void log_pkt_long( const u_char* packet, 
                   const struct pcap_pkthdr* pkt_header );
void print_eth_addr( const u_char* addr );

void pkt_callback( u_char* args, 
                   const struct pcap_pkthdr* header, 
                   const u_char* packet );

static int pkt_count;

int main(int argc, char* argv[]) {

    pcap_t* handle;
    char dev[] = "fxp0";
    char errbuf[PCAP_ERRBUF_SIZE];
    struct pcap_pkthdr header;
    const u_char* packet;

    int ms_time = 1000;

    int lcv;

    pkt_count = 0;

    printf("Device:  %s\n", dev);
    printf("BUFSIZ:  %d\n", BUFSIZ);

    handle = pcap_open_live(dev, BUFSIZ, 1, ms_time, errbuf);
    if ( NULL == handle ) {
        printf("ERROR:  pcap_open_live returned: %s\n", errbuf);
        return 1;
    }
    printf("pcap_open_live() - success\n");

/*     packet = pcap_next(handle, &header); */
/*     log_pkt_long( packet, &header ); */

    lcv = 0; 
    while ( lcv < 1000 ) {
        pkt_count = 0;
        pcap_dispatch(handle, -1, pkt_callback, NULL);
        printf("Processed %d packets in %d ms\n", pkt_count, ms_time);
        lcv ++;
    }

    pcap_close(handle);
    return 0;
}


void pkt_callback( u_char* args, 
                   const struct pcap_pkthdr* header, 
                   const u_char* packet ) 
{
/*     printf("----------------------------------------\n"); */
/*     log_pkt_long( packet, header ); */
    pkt_count++;
}



void log_pkt_long( const u_char* packet, 
                   const struct pcap_pkthdr* pkt_header )
{
    const struct ether_header*  eth;
    const struct ip*            ip;
    const struct tcphdr*        tcp;
    const u_char*               payload;

    int size_ether = sizeof(struct ether_header);
    int size_ip = sizeof(struct ip);
    int size_tcp = sizeof(struct tcphdr);

    printf("Jacked a packet with length of [%d]\n", pkt_header->len);

    eth     = (struct ether_header*)packet;
    ip      = (struct ip*)(packet + size_ether);
    tcp     = (struct tcphdr*)(packet + size_ether + size_ip);
    payload = (u_char*)(packet + size_ether + size_ip + size_tcp);


    printf("ethernet type: %d\n", eth->ether_type);
    printf("ethernet dest: ");
    print_eth_addr( eth->ether_dhost );
    printf("\nethernet src:  ");
    print_eth_addr( eth->ether_shost );
    printf("\n");

    printf("ip version:  %d\n", ip->ip_v);
    printf("ip tos:      %d\n", ip->ip_tos);
    printf("ip length:   %d\n", ip->ip_len);
    printf("ip id:       %d\n", ip->ip_id);
    printf("ip frag off  %d\n", ip->ip_off);
    printf("ip ttl:      %d\n", ip->ip_ttl);
    printf("ip prot:     %d\n", ip->ip_p);
    printf("ip ck sum:   %d\n", ip->ip_sum);
    printf("ip src addr: %s\n", inet_ntoa(ip->ip_src));
    printf("ip dst addr: %s\n", inet_ntoa(ip->ip_dst));
}

void print_eth_addr( const u_char* addr ) {

    int ii;
    for ( ii = 0; ii < 6; ii++ ) {
        printf("%02x", addr[ii]);
        if ( ii < 5 ) {
            printf(":");
        }
    }
}
