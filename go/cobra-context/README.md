cobra passing a Context - Example
=================================

Goal: a multi-command level cobra application where the root Command
forms up an object (like a config, or some other resource) and
attaches it to the Command's context.  The leaf Command then extracts
and uses object.

