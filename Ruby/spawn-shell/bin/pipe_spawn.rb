#!/usr/bin/env ruby

def spawn_pipe(cmd)
  pipe = IO.popen(cmd)
  pid = pipe.pid
  puts "Parent[#{Process.pid}] waiting for child[#{pid}]"
  Process.waitpid(pid)
  success = ($?.exited? and $?.exitstatus == 0)
  results = pipe.read
  pipe.close
  unless success
    puts "Something FAILED."
  end
  results
end


pipe_output = spawn_pipe("sleep 3; echo foobar")

puts "Received --"
puts pipe_output
puts "-- from the shell command"

