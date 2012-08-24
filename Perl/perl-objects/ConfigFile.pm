package ConfigFile;

@EXPORT = qw(import get has);

# Construct the config file object.
#
#  import( <invocant>, <cfg file name> )
#
#  Currently only the user's home directory is
#  searched.  
#
#  Enhancements:  - ENV var for search path
#                 - third param for search path
#
sub load {
    my $invocant      = shift;
    my $class         = ref($invocant) || $invocant;
    my $cfg_file_name = shift;
    unless ($cfg_file_name) {
        warn "no config file defined,";
        return;
    }
    unless (open(CFG_FH, "$ENV{HOME}/$cfg_file_name")) {
        warn "failed to open $ENV{HOME}/$cfg_file_name";
        warn "failed to open config file ($cfg_file_name)";
        return;
    }
    my $self = { };
    while ( <CFG_FH> ) {
        if ( /^\s*(\w+)\s*=\s*"([^"]*)"/ ) {
            $self->{$1} = $2;
        } elsif ( /^\s*(\w+)\s*=\s*([^\s]+)/ ) {
            $self->{$1} = $2;
        }
    }

    bless($self, $self);

    close CFG_FH;
}

sub has {
    
}

sub get {

}

