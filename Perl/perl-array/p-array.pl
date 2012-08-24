#!/usr/local/bin/perl

@pool = ();

if ( @pool ) {
    print "true\n";
} else {
    print "false\n";
}


push(@pool, 1);


if ( @pool ) {
    print "true\n";
} else {
    print "false\n";
}



pop(@pool);




if ( @pool ) {
    print "true\n";
} else {
    print "false\n";
}
