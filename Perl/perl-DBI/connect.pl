#!/usr/local/bin/perl

use DBI;

my $dbh = DBI->connect( "dbi:mysql:perl_dbi", "root", "" ) 
    or die "Could not connect: $DBI::errstr\n";


my $sth = $dbh->prepare( "SELECT * FROM people" );

$sth->execute;

my @row;
while ( @row = $sth->fetchrow_array() ) {
    print "Row: @row\n";
}


$dbh->disconnect
    or warn "Disconnection failed: $DBI::errstr\n";
