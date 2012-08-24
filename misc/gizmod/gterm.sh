#!/bin/sh

# EG - there are problems with this script.  If the gizmod
#  server isn't up, the xterm dies before the user sees
#  the error message. (instantly dies)

exec xterm -geometry 80x24 +sb +ls -ut -T gterm -n gterm -e ./client 10.3.4.18
