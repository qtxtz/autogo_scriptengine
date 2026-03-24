-- 验证字符串位置
local text = "中国人民万岁"
print("字符串: " .. text)
print("字符位置:")
for i = 1, utf8.length(text) do
    local char = utf8.mid(text, i, 1)
    print("  位置 " .. i .. ": " .. char)
end

print()
print("utf8.mid(text, 2, 2) = " .. utf8.mid(text, 2, 2))
print("utf8.strCut(text, 2, 2) = " .. utf8.strCut(text, 2, 2))
print("utf8.strReverse('中国人民') = " .. utf8.strReverse("中国人民"))
