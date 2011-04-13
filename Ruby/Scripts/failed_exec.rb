#!/usr/bin/env ruby

def spawn(cmd)
  puts "Parent pid: #{Process.pid}"
  begin
    pipe = IO.popen(cmd)
    pid = pipe.pid
    puts "Child pid:  #{pid}"
    Process.waitpid(pid)
    puts "Child exited with [#{$?.exitstatus}]."
    pipe.close
  rescue => ex
    puts "ERROR - Child never spawned"
    puts "Caught #{ex.class}, (cause: #{ex})"
  end
end


cmds = ["echo successful command", "bogus-command param"]

cmds.each do |cmd|
  puts "--"
  spawn(cmd)
end
puts "-- done."
