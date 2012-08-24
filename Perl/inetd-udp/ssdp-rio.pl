#!/usr/local/bin/perl

open(LOG, ">>/tmp/ssdp.log");
print LOG "SSDP: $$\n";

$ENV{PATH} = "/usr/bin:/bin";

sub mlog($) {
    my $msg = shift;
    system("/usr/bin/logger -t SSDP  \"$msg\"");
}

sub mdie($) {
    my $msg = shift;
    mlog $msg;
    exit 1;
}

sub sig_handle($) {
    mdie "SSDP server exit on signal";
}

# $mserve_ip must be a dotted-quad unless you modify the client
# NFS image to include /etc/resolv.conf.
#
my $mserve_ip   = "10.3.4.249"; 
my $mserve_port = "81";

#
# The box makes two different requests.  One comes from the kernel
# during initial booting, the second comes from the player application
# after the second boot when the player starts.
#
# The respones are different for Linux.  If there is a port number
# on the first "linux" request then the client box will use that port
# for portmapper lookups, which is generally bad when talking to 
# another linux box.
#
# The second "player" response includes a port number that indicates
# the port number to use for HTTP music related requests.  I use
# port 81 and setup a virtual server in Apache to respond to music
# requests, but you may want to do this differently.
#
my $player_request  = "^upnp:uuid:1D274DB0-F053-11d3-BF72-0050DA689B2F";
my $linux_request   = "^upnp:uuid:1D274DB1-F053-11d3-BF72-0050DA689B2F";

my $player_response = "http://$mserve_ip:$mserve_port/descriptor.xml\n";
my $linux_response  = "http://mserve_ip/descriptor.xml\n";

$datagram = <STDIN>;

print LOG "SSDP received: $datagram";
mlog "received: $datagram";

if ( $datagram =~ $linux_request ) {
    mlog "linux request from peer.";
#     print STDOUT $linux_response;
}

if ( $datagram =~ $player_request ) {
    mlog "player request from peer.";
#     print STDOUT $player_response;
}

exit(0);
