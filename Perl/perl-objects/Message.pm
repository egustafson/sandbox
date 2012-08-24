package Message;

sub new {
    my $invocant = shift;
    my $class    = ref($invocant) || $invocant;
    my $self     = { @_ };
    if ( !defined($self->{msg}) ) {
        return;
    }
    bless( $self, $class );
    return $self;
}

sub show {
    my $self = shift;
    print STDOUT "$self->{msg}\n";
}
1;
