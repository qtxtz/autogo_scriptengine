-- 验证 utf8.strReverse 的行为
local text = "中国人民"
print("原始字符串: " .. text)
print("字符序列:")
for i = 1, utf8.length(text) do
    local char = utf8.mid(text, i-1, 1)
    print("  索引 " .. (i-1) .. ": " .. char)
end

print()
print("反转结果: " .. utf8.strReverse(text))

-- 分析：
-- "中国人民" 的字符：中(0)、国(1)、人(2)、民(3)
-- 反转后应该是：民(3)、人(2)、国(1)、中(0)
-- 结果应该是："民人国中"

-- 但是懒人精灵的结果是："民国中人"
-- 这意味着：
-- 中(0) -> 民(3)
-- 国(1) -> 人(2)
-- 人(2) -> 国(1)
-- 民(3) -> 中(0)
-- 这就是反转！
