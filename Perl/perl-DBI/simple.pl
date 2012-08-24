#!/usr/local/bin/perl

use DBI;

my @drivers = DBI->available_drivers();

my @mysqlSources = DBI->data_sources( "mysql" );
foreach my $mysqlSource ( @mysqlSources ) {
    print "Data Source for mysql is $mysqlSource\n";
}

foreach my $driver ( @drivers ) {
    print "Driver: $driver\n";
    my @dataSources = DBI->data_sources( $driver );
    foreach my $dataSource ( @dataSources ) {
        print "\tData Source is $dataSource\n";
    }
}
