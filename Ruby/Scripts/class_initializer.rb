#!/usr/bin/env ruby

# The goal of this example is to demonstrate a Class level
# attribute that is initialized on the first _object_ construction
# of that Class.

class Foo
  @@class_var = nil
  
  def self.init
    puts "initialized class var"
    "Class Var - Value"
  end
  
  def initialize
    @@class_var ||= Foo.init
  end
  
  def to_s
    "#{@@class_var}"
  end
  
end


f1 = Foo.new
f2 = Foo.new

puts f1.to_s
puts f2.to_s

puts "done."