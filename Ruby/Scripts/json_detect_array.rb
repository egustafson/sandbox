#!/usr/bin/env ruby

require 'json'

Obj = <<JSON
{ "obj" : { "k" : "v" } }
JSON

Arr = <<JSON
{ "obj" : [
  { "k1" : "v1" },
  { "k2" : "v2" }
  ]
}
JSON

MMap = <<JSON
{ "obj" : { "k1" : "v1" },
  "obj" : { "k2" : "v2" }
}
JSON

def resolve json
  o = JSON.parse json
  obj = o['obj']
  puts "obj is_a #{obj.class}"
  puts "obj is a Hash" if obj.is_a? Hash
  puts "obj is an Array" if obj.is_a? Array
end

def normalize json
  o = JSON.parse json
  obj = o['obj']
  if obj.is_a? Hash
    o_arr = [ obj ]
    o['obj'] = o_arr
  end
  JSON.generate o
end

mm = JSON.parse MMap  # this should fail
resolve MMap
puts "Normalized: #{normalize MMap}"


resolve Obj
resolve Arr

puts "Normalized: #{normalize Obj}"
puts "Normalized: #{normalize Arr}"



