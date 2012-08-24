#!/usr/local/bin/perl

use DBI;
use Finance::QuoteHist;

## ############################################################
my $dbh = DBI->connect( "dbi:mysql:stock_history", "root", "" ) 
    or die "Could not connect: $DBI::errstr\n";
my $sth = $dbh->prepare( "INSERT INTO basic_history ( ticker, date, open, close, high, low, volume ) VALUES ( ?,?,?,?,?,?,? )" );

$q = new Finance::QuoteHist
    (
#      symbols    => [qw(IBM UPS AMZN)],
     symbols    => 'IBM',
     start_date => '01/01/1980',
     end_date   => '12/31/1989',      
     );

# Adjusted values
foreach $row ($q->quote_get()) {
    ($symbol, $date, $open, $high, $low, $close, $volume) = @$row;

    print "@$row\n";
    $date =~ /^(\d+)\/(\d+)\/(\d+)$/;
    $sql_date = $1 . $2 . $3;
        
    $sth->execute( $symbol, 
                   $sql_date, 
                   $open, 
                   $high, 
                   $low, 
                   $close, 
                   $volume );
}

$dbh->disconnect
    or warn "Disconnection failed: $DBI::errstr\n";
