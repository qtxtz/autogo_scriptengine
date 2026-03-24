-- 验证字符串位置和函数行为
local text = "中国人民万岁"
print("字符串: " .. text)
print("字符位置:")
for i = 1, utf8.length(text) do
    local char = utf8.mid(text, i, 1)
    print("  位置 " .. i .. ": " .. char)
end

print()
print("测试 utf8.mid:")
print("  utf8.mid(text, 2, 2) = " .. utf8.mid(text, 2, 2))
print("  说明: 从位置 2 开始，截取 2 个字符 -> '国人'")

print()
print("测试 utf8.strCut:")
print("  utf8.strCut(text, 2, 2) = " .. utf8.strCut(text, 2, 2))
print("  说明: 从位置 2 开始，移除 2 个字符 -> '中民万岁'")

print()
print("测试 utf8.strReverse:")
local text2 = "中国人民"
print("  字符串: " .. text2)
print("  字符位置:")
for i = 1, utf8.length(text2) do
    local char = utf8.mid(text2, i, 1)
    print("    位置 " .. i .. ": " .. char)
end
print("  utf8.strReverse(text2) = " .. utf8.strReverse(text2))
print("  说明: 反转所有字符 -> '民人国中'")
