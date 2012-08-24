#!/usr/bin/perl

($port) = @ARGV;
$port = 2345 unless $port;

$AF_INET = 2;
$SOCK_STREAM = 1;

$sockaddr = 'S n a4 x8';

($name, $aliases, $proto) = getprotobyname('tcp');
if ($port !~ /^\d+$/) {
    ($name, $aliases, $port) = getservbyport($port, 'tcp');
}

print "Port = $port\n";

$this = pack($sockaddr, $AF_INET, $port, "\0\0\0\0");

select(NS); $| = 1; select(stdout);

socket(S, $AF_INET, $SOCK_STREAM, $proto) || die "socket: $!";
bind(S,$this) || die "bind: $!";
listen(S,5) || die "connect: $!";

select(S); $| = 1; select(stdout);

for ($con = 1; ; $con++) {
    printf("Listening for connection %d...\n", $con);
    ($addr = accept(NS,S)) || die $!;
    
    if (($child = fork()) == 0) {
        print "accept ok\n";

        ($af,$port,$inetaddr) = unpack($sockaddr,$addr);
        @inetaddr = unpack('C4',$inetaddr);
        print "$con:  $af $port @inetaddr\n";
        
        while (<NS>) {
            print "$con: $_";
            if ( /^quit/ ) {
                print NS "Good bye.\n";
                close(NS);
                exit;
            }
            if ( /^monitor/ ) {
                
            }
            if ( /^hello +(.*)$/ ) {
                chop($name = $1);
                print NS "Pleased to meet you $name.\n";
            }
        }
        close(NS);
        exit;
    }
    close(NS);
}
