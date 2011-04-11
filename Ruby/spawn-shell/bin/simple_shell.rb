#!/usr/bin/env ruby

def backtick(cmd)
  puts "Preparing to execute [#{cmd}]"
  result = `#{cmd}`
  puts "completed."
  puts "The results of [#{cmd}] are --"
  puts result
  puts "--"
end



backtick("echo foobar")
backtick("sleep 2")
print "done."
