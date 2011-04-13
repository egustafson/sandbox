#!/usr/bin/env ruby

require 'timeout'

def spawn_timeout(cmd, timeout)
  success = false
  pid = nil
  Timeout.timeout(timeout) do
    pid = fork { exec(cmd) }
    Process.waitpid(pid)
    success = ($?.exited? and $?.exitstatus == 0)
  end
  unless success
    puts "Something FAILED."
  end
rescue Timeout::Error
  Process.kill("SIGTERM", pid)
  Process.detach(pid)
  puts "TIMEOUT - BAD, FAILURE."
end

spawn_timeout("echo foobar", 2)
print "done."
