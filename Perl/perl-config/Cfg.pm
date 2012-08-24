package Cfg;

sub load_config {

    if ( -f 'main-config.rc' ) {
        do 'main-config.rc';
        print "config file loaded\n";
    }

}

1;
