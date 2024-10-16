---
--- Created by walkskyer.
--- DateTime: 2024/10/16 10:05
---

m = { name = "functions" }

function m.add(a, b)
    return a + b
end

function m.sub(a, b)
    return a - b
end

function m.mul(a, b)
    return a * b
end

function m.div(a, b)
    return a / b
end

function m.Sum(...)
    local sum = 0
    for i, v in ipairs({ ... }) do
        sum = sum + v
    end
    return sum
end

function m.mod(a, b)
    return a % b
end

return m
