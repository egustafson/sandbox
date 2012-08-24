#!/usr/bin/perl

print "Configuration example.\n";

do 'main-config.rc';

print "CFG_NAME = $CFG_NAME\n";

foreach $key (keys %CFG_HASH) {

    print "CFG_HASH{$key} = $CFG_HASH{$key}\n";
}
