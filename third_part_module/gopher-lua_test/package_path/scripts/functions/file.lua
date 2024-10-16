---
--- Created by walkskyer.
--- DateTime: 2024/10/16 10:22
---

m = {name="file"}
function m.read_file(path)
    local file = io.open(path, "r")
    local content = file:read("*all")
    file:close()
    return content
end
function m.write_file(path, content)
    local file = io.open(path, "w")
    file:write(content)
    file:close()
end


return m
