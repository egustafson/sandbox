#!/home/ericg/bin/lua

x = { "mtd8", "mtd9", "mtd10", "mtd0" }

for ii = 1, getn(x) do
   _, _, idx = strfind(x[ii], "[^0-9]*([0-9]*)")
   print("idx = "..idx)
end