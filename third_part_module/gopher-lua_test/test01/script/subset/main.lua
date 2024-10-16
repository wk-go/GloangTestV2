---
--- Created by walkskyer.
--- DateTime: 2024/10/16 14:19
---
a = 1 + 2
print("------------ os module must be work ------------")
print("time:", os.time())

--- table test
print("------------ table test ------------")
local t = {1,2,3,name="table01"}
print("table:", t)

--- io module must be not work
print("------------ io module must be not work ------------")
local f = io.open("subset-test.txt", "w")
f:write("hello world")
f:close()

print("read file:", io.open("subset-test.txt"):read("*all"))
