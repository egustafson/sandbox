#!/usr/bin/env ruby

class Container
  def initialize(m)
    @msg = m
  end
  def to_s
    "Container('#{@msg}')"
  end
end

class FancyEx < StandardError
  def initialize(msg, container)
    super(msg)
    @c = container
  end
  
  def to_s
    "['#{super.to_str}' - #{@c}]"
  end
end

begin
  c = Container.new("container message")
  raise FancyEx.new("test raise", c)
rescue FancyEx => ex
  puts "Caught: #{ex}"
end