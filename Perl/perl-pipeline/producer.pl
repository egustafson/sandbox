#!/usr/local/bin/perl

$| = 1;

@data = ( 1 .. 10 );

foreach $datum ( @data ) {
    sleep ( 1 );
    printf "\%05d\n", $datum;
}
