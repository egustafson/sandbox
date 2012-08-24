#!/usr/bin/perl

($them,$port) = @ARGV;
$port = 2345 unless $port;
$them = 'localhost' unless $them;

$AF_INET = 2;
$SOCK_STREAM = 1;

$SIG{'INT'} = 'dokill';
sub dokill {
    kill 9,$child if $child;
}

$sockaddr = 'S n a4 x8';

chop($hostname = `hostname`);

($name,$aliases,$proto) = getprotobyname('tcp');
($name,$aliases,$port)  = getservbyname($port,'tcp')
    unless $port =~ /^\d+$/;;
($name,$aliases,$type,$len,$thisaddr) =
    gethostbyname($hostname);
($name,$aliases,$type,$len,$thataddr) = gethostbyname($them);

$this = pack($sockaddr, $AF_INET, 0, $thisaddr);
$that = pack($sockaddr, $AF_INET, $port, $thataddr);

# Make the socket filehandle.

if (socket(S, $AF_INET, $SOCK_STREAM, $proto)) {
    print "socket ok\n";
} else {
    die $!;
}

# Give the socket an address.

if (bind(S, $this)) {
    print "bind ok\n";
} else {
    die $!;
}

# Call up the server.

if (connect(S,$that)) {
    print "connect ok\n";
} else {
    die $!;
}

# Set socket to be command buffered.

select(S); $| = 1; select(STDOUT);

# Avoid deadlock by forking

if ($child = fork) {
    print S "hello `hostname`\n";
    print S "quit\n";
    close(S);
    sleep 3;
    do dokill();
} else {
    while(<S>) {
        print;
    }
    close(S);
}
