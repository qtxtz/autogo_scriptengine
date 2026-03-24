-- ffi 模块测试脚本
-- 注意：ffi 模块在 AutoGo 中不可用，这是 LuaJIT 的特性

print("=== ffi 模块测试 ===")
print("注意：ffi 是 LuaJIT 的特性，AutoGo 使用 gopher-lua，不支持 FFI")
print()

-- 测试 1: 尝试加载 ffi 模块
print("--- 测试 1: 加载 ffi 模块 ---")
local ffi = require("ffi")
if ffi then
    print("✓ ffi 模块加载成功")
else
    print("✗ ffi 模块加载失败")
end
print()

-- 测试 2: 尝试使用 ffi.cdef
print("--- 测试 2: 尝试使用 ffi.cdef ---")
local success, err = pcall(function()
    ffi.cdef[[
        int getpid(void);
    ]]
end)
if success then
    print("✓ ffi.cdef 调用成功（不应该成功）")
else
    print("✓ ffi.cdef 不可用（符合预期）")
    print("  错误信息: " .. tostring(err))
end
print()

-- 测试 3: 尝试使用 ffi.load
print("--- 测试 3: 尝试使用 ffi.load ---")
local success, err = pcall(function()
    local clib = ffi.load("libc")
end)
if success then
    print("✓ ffi.load 调用成功（不应该成功）")
else
    print("✓ ffi.load 不可用（符合预期）")
    print("  错误信息: " .. tostring(err))
end
print()

-- 测试 4: 尝试使用 ffi.new
print("--- 测试 4: 尝试使用 ffi.new ---")
local success, err = pcall(function()
    local ptr = ffi.new("int*", 42)
end)
if success then
    print("✓ ffi.new 调用成功（不应该成功）")
else
    print("✓ ffi.new 不可用（符合预期）")
    print("  错误信息: " .. tostring(err))
end
print()

-- 测试 5: 尝试使用 ffi.sizeof
print("--- 测试 5: 尝试使用 ffi.sizeof ---")
local success, err = pcall(function()
    local size = ffi.sizeof("int")
end)
if success then
    print("✓ ffi.sizeof 调用成功（不应该成功）")
else
    print("✓ ffi.sizeof 不可用（符合预期）")
    print("  错误信息: " .. tostring(err))
end
print()

print("=== ffi 模块测试完成 ===")
print("所有 ffi 功能都不可用，这是正常行为")
print("请使用 AutoGo 提供的 Go 模块或 Lua 标准库替代")
