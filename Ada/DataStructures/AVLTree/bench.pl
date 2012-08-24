#!/usr/local/bin/perl

$low  = 1;
$high = 20;

for ( $ii = $low; $ii < $high+1; $ii++ ) {

    $sum = 0;
    for ( $jj = 0; $jj < 5; $jj++ ) {
        $start = (times)[2];
        system("./main $ii 0");
        $end   = (times)[2];
        $sum = $sum + ( $end - $start );
    }
    $bench_time[$ii] = ( $sum / 5 );
    printf("2^%2d inserts -> %6.2fs\n", $ii, $bench_time[$ii]);
}

