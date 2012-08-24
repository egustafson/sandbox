#!/usr/local/bin/perl

use Thread;

my $thread_id = Thread->self->tid();

print "my thread id: $thread_id\n";
