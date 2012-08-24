#!/usr/bin/perl

print "password: ";
$clearpass=<STDIN>;

print "salt: ";
$salt = <STDIN>;

chomp $clearpass;
chomp $salt;

$cryptpass = crypt( $clearpass, $salt );

print "crypt:  $cryptpass\n";


