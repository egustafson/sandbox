#!/usr/bin/env ruby

require 'json'

json_txt = <<JSON
{ 
  "obj" : { "k" : "v" },
  "arr" : ["a", "b", "c"]
}
JSON

o = JSON.parse json_txt

raise "JSON did NOT parse to a Hash" unless o.is_a? Hash 

puts "Element 'foo' is nil (does not exist in the hash)" if o['foo'].nil?

puts "success."