---
--- Created by walkskyer.
--- DateTime: 2024/10/16 10:05
---
--- Set Lua Path
package.path = package.path .. [[;]] .. GetLuaPath()
local my_math = require("my_math")

print("1+2=", my_math.add(1,2))
print("1-2=", my_math.sub(1,2))
print("1*2=", my_math.mul(1,2))
print("1/2=", my_math.div(1,2))
print("1%2=", my_math.mod(1,2))
print("1+2+3+4+5+6+7+9", my_math.Sum(1,2,3,4,5,6,7,9))


local my_file = require("file")
my_file.write_file("test.txt", "hello world")
print("read file:", my_file.read_file("test.txt"))