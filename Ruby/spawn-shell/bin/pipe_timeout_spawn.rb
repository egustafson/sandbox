#!/usr/bin/env ruby

require 'timeout'

def pipe_timeout_spawn(cmd, timeout)
  success = false
  pid = nil
  pipe = nil
  Timeout.timeout(timeout) do
    pipe = IO.popen(cmd)
    pid = pipe.pid
    Process.waitpid(pid)
    success = ($?.exited? and $?.exitstatus == 0)
    if success
      pipe.read
    else
      nil
    end
  end
rescue Timeout::Error
  Process.kill("SIGTERM", pid)
  Process.detach(pid)
  puts "TIMEOUT"
  nil
ensure
  pipe.close if pipe
end


pipe_results = pipe_timeout_spawn("sleep 3; echo Successful sub-shell command.", 5)
puts "Received --[#{pipe_results}]--"

pipe_results = pipe_timeout_spawn("sleep 3; echo Failure - timeout", 2)
if pipe_results
  puts "FAILURE - results: --[#{pipe_results}]--"
else
  puts "Success, timeout killed process."
end

pipe_results = pipe_timeout_spawn("sleep 1", 2)
if pipe_results and pipe_results.length < 1
  puts "Success, received a zero length string, spawned command had no output."
elsif pipe_results 
  puts "FAILURE, received --[#{pipe_results}]--, should have received an empty string."
  puts 
else
  puts "FAILURE, received a FALSE result"
end
