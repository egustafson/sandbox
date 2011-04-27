#!/usr/bin/env ruby

class Point
  attr_reader :x, :y
  
  def initialize(x, y)
    @x = x
    @y = y
  end
  
  def marshal
    {'x' => @x, 'y' => @y}
  end
  
  def to_s
    "[#{@x}, #{@y}]"
  end
end

class Point3d < Point
  attr_reader :z
  
  def initialize(x, y, z)
    super(x,y)
    @z = z
  end
  
  def marshal
    m = super   # 'super' ==> invoke super.method(args)
    m['z'] = @z
    m
  end
  
  def to_s
    # instance variables (@...) are not implicit in any class
    # definition, but are created and attached to the object
    # *instance* when first referenced, (initialize() in this
    # case).  Instance variables are visible INSIDE the object
    # at all levels of inheritance; they are not part of the
    # type (class) definition and as such are not really part
    # any notion of inheritance.
    "[#{@x}, #{@y}, #{@z}]"
  end
  
end

## Main

p2 = Point.new(1, 1)
p3 = Point3d.new(1, 1, 1);

puts "Point   = #{p2}"
puts "Point3d = #{p3}"

puts "Point   marshalled: #{p2.marshal}"
puts "Point3d marshalled: #{p3.marshal}"