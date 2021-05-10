# Closable `chan`

This example shows how to construct a simple object that allows the
`chan` receiver to also signal it is done receiving on the `chan` (aka
close()).  Only thread that holds a writable end of a chan can invoke
`close()` on it, this method allows signaling in a thread safe,
non-blocking way.

Also worth noting:  if a buffered chan is closed with contents in the
buffer, and all references are dropped, then the chan and it's
buffered contents _should_ be garbage collected.
