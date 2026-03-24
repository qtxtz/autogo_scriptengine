-- string 模块测试脚本
print("=== string 模块测试 ===")
print()

-- 测试 1: splitStr - 字符串分割
print("--- 测试 1: splitStr - 字符串分割 ---")
local parts = splitStr("ab#cd#ef#gh", "#")
if #parts == 4 and parts[1] == "ab" and parts[2] == "cd" and parts[3] == "ef" and parts[4] == "gh" then
    print("✓ splitStr 测试通过")
    print("  分割结果: " .. table.concat(parts, ", "))
else
    print("✗ splitStr 测试失败")
end
print()

-- 测试 2: utf8.inStr - UTF-8 字符串查找
print("--- 测试 2: utf8.inStr - UTF-8 字符串查找 ---")
local pos = utf8.inStr(1, "中国人民万岁", "人民")
if pos == 3 then
    print("✓ utf8.inStr 测试通过")
    print("  查找位置: " .. pos)
else
    print("✗ utf8.inStr 测试失败")
    print("  期望: 3, 实际: " .. (pos or "nil"))
end
print()

-- 测试 3: utf8.inStrRev - UTF-8 字符串反向查找
print("--- 测试 3: utf8.inStrRev - UTF-8 字符串反向查找 ---")
local pos = utf8.inStrRev("中国人民万岁", "人民", 6)
if pos == 3 then
    print("✓ utf8.inStrRev 测试通过")
    print("  查找位置: " .. pos)
else
    print("✗ utf8.inStrRev 测试失败")
    print("  期望: 3, 实际: " .. (pos or "nil"))
end
print()

-- 测试 4: utf8.strReverse - UTF-8 字符串反转
print("--- 测试 4: utf8.strReverse - UTF-8 字符串反转 ---")
local result = utf8.strReverse("中国人民")
if result == "民国中人" then
    print("✓ utf8.strReverse 测试通过")
    print("  反转结果: " .. result)
else
    print("✗ utf8.strReverse 测试失败")
    print("  期望: 民国人中, 实际: " .. result)
end
print()

-- 测试 5: utf8.length - UTF-8 字符串长度
print("--- 测试 5: utf8.length - UTF-8 字符串长度 ---")
local len = utf8.length("中国人民万岁")
if len == 6 then
    print("✓ utf8.length 测试通过")
    print("  字符长度: " .. len)
else
    print("✗ utf8.length 测试失败")
    print("  期望: 6, 实际: " .. (len or "nil"))
end
print()

-- 测试 6: utf8.left - UTF-8 字符串左侧截取
print("--- 测试 6: utf8.left - UTF-8 字符串左侧截取 ---")
local result = utf8.left("中国人民万岁", 2)
if result == "中国" then
    print("✓ utf8.left 测试通过")
    print("  截取结果: " .. result)
else
    print("✗ utf8.left 测试失败")
    print("  期望: 中国, 实际: " .. result)
end
print()

-- 测试 7: utf8.right - UTF-8 字符串右侧截取
print("--- 测试 7: utf8.right - UTF-8 字符串右侧截取 ---")
local result = utf8.right("中国人民万岁", 2)
if result == "万岁" then
    print("✓ utf8.right 测试通过")
    print("  截取结果: " .. result)
else
    print("✗ utf8.right 测试失败")
    print("  期望: 万岁, 实际: " .. result)
end
print()

-- 测试 8: utf8.mid - UTF-8 字符串中间截取
print("--- 测试 8: utf8.mid - UTF-8 字符串中间截取 ---")
local result = utf8.mid("中国人民万岁", 2, 2)
if result == "人民" then
    print("✓ utf8.mid 测试通过")
    print("  截取结果: " .. result)
else
    print("✗ utf8.mid 测试失败")
    print("  期望: 人民, 实际: " .. result)
end
print()

-- 测试 9: utf8.strCut - UTF-8 字符串裁剪
print("--- 测试 9: utf8.strCut - UTF-8 字符串裁剪 ---")
local result = utf8.strCut("中国人民万岁", 2, 2)
if result == "中国万岁" then
    print("✓ utf8.strCut 测试通过")
    print("  裁剪结果: " .. result)
else
    print("✗ utf8.strCut 测试失败")
    print("  期望: 中国万岁, 实际: " .. result)
end
print()

-- 测试 10: 综合测试 - 处理 CSV 数据
print("--- 测试 10: 综合测试 - 处理 CSV 数据 ---")
local line = "张三,25,北京,工程师"
local fields = splitStr(line, ",")
if #fields == 4 and fields[1] == "张三" and fields[2] == "25" and fields[3] == "北京" and fields[4] == "工程师" then
    print("✓ CSV 数据处理测试通过")
    for i, field in ipairs(fields) do
        print("  字段 " .. i .. ": " .. field)
    end
else
    print("✗ CSV 数据处理测试失败")
end
print()

-- 测试 11: 边界测试 - 超出长度
print("--- 测试 11: 边界测试 - 超出长度 ---")
local result = utf8.left("中国人民万岁", 10)
if result == "中国人民万岁" then
    print("✓ 边界测试通过（超出长度返回整个字符串）")
    print("  结果: " .. result)
else
    print("✗ 边界测试失败")
    print("  期望: 中国人民万岁, 实际: " .. result)
end
print()

-- 测试 12: 边界测试 - 空字符串
print("--- 测试 12: 边界测试 - 空字符串 ---")
local result = utf8.left("", 2)
if result == "" then
    print("✓ 空字符串测试通过")
else
    print("✗ 空字符串测试失败")
    print("  期望: (空), 实际: " .. result)
end
print()

print("=== string 模块测试完成 ===")
