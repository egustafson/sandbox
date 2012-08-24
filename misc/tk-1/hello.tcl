#!/usr/local/bin/wish

text .text 
button .button -text "Exit" -command exit

pack .text
pack .button

.text insert end "hello\n"
.text insert end "good bye"