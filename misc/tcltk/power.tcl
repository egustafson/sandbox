#!/usr/local/bin/wish -f

proc power {base p} {
    set result 1
    while {$p>0} {
        set result [expr $result*$base]
        set p [expr $p-1]
    }
    return $result
}

set base 0
set power 0

entry .base -width 6 -relief sunken -textvariable base
label .label1 -text "to the power"
entry .power -width 6 -relief sunken -textvariable power
label .label2 -text "is"
label .result -textvariable result
button .exit -text "Exit" -command exit
pack .base .label1 .power .label2 .result .exit -side left -padx 1m -pady 2m
bind .base <Return> {set result [power $base $power]}
bind .power <Return> {set result [power $base $power]}
