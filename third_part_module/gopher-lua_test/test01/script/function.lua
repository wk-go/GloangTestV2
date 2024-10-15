module = {}
function fibonacci(n)
    if n == 0 then
        return 0
    elseif n == 1 then
        return 1
    end
    return fibonacci(n - 1) + fibonacci(n - 2)
end

function say_hello(name)
    if name == nil then
        name = "World"
    end
    print("Hello, " .. name)
end

-- 协程测试
function coro(target)
    if target == nil then
        target = 10
    end
    local i = 0
    while true do
        coroutine.yield(i)
        i = i + 1
        if i > target then
            break
        end
    end
    return i
end

module.fibonacci = fibonacci
module.say_hello = say_hello
module.coro = coro

return module
