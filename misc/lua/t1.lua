#!/home/ericg/bin/lua

tfile = "all.fm"

source = "tar -xO -f " .. tfile

if readfrom("|" .. source .. " index") == nil then
   print("can't read from fm upgrade file: " .. tfile)
   exit(1)
end

list = {}

l = read()
while l ~= nil do
   local t = {}

   _, _, t.type, t.file, t.meth, t.dest, t.md5, t.text =
      strfind(l, "^(.*):(.*):(.*):(.*):(.*):(.*)$")

   if t and t.type and t.file and t.meth and t.dest and t.md5 and t.text then
      tinsert(list, t)
   else
      print("invalid index entry: "..l)
   end

   _, _, _, ftype = strfind(t.file, "^([^\.]*)\.([^\.]*)$")
   print(l.." - ".. ftype)

   l = read()
end

readfrom()

exit(0)