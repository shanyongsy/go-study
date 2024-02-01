-- sayHello 函数
function sayHello(name)
    print("Hello World! " .. name)
end

-- add 函数
function add(a, b)
    return a + b
end

-- 构造一个配置表
config = {
    name = "John",
    age = 30,
    address = {
        city = "New York",
        postal_code = "10001",
    },
    hobbies = {"reading", "coding", "traveling"},
}

-- 除法运算，返回余数和商
function div(a, b)
    return a % b, math.floor(a / b)
end

