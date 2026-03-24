# string 模块（字符串处理）

## 概述

string 模块扩展了 Lua 标准库的字符串处理功能，添加了懒人精灵特有的字符串分割函数和 UTF-8 字符串处理功能。

## 新增功能

### 1. splitStr(src, sep) - 字符串分割

**功能**：根据分隔符 sep 把 src 拆分为多个子串并返回结果（懒人扩展）

**参数**：
- `src` (string)：要分割的源字符串
- `sep` (string)：分隔符

**返回值**：
- `table`：分割后的子串数组

**示例**：
```lua
local parts = string.splitStr("ab#cd#ef#gh", "#")
-- parts = {"ab", "cd", "ef", "gh"}

-- 也可以直接调用
local parts = splitStr("ab#cd#ef#gh", "#")
```

## UTF-8 字符串处理

### 2. utf8.inStr(start, s, pattern) - UTF-8 字符串查找

**功能**：从指定位置开始查找 UTF-8 字符串中的子串

**参数**：
- `start` (number)：开始查找的位置（从 1 开始）
- `s` (string)：目标字符串
- `pattern` (string)：要查找的模式

**返回值**：
- `number`：找到的位置索引
- `nil`：找不到时返回 nil

**示例**：
```lua
local pos = utf8.inStr(1, "中国人民万岁", "人民")
-- pos = 3

local pos = utf8.inStr(4, "中国人民万岁", "万岁")
-- pos = 5
```

### 3. utf8.inStrRev(s, pattern, start, compare) - UTF-8 字符串反向查找

**功能**：从指定位置反向查找 UTF-8 字符串中的子串

**参数**：
- `s` (string)：目标字符串
- `pattern` (string)：要查找的模式
- `start` (number)：开始查找的位置（从 1 开始，默认从字符串末尾开始）
- `compare` (function)：比较函数（可选，暂未实现）

**返回值**：
- `number`：找到的位置索引
- `nil`：找不到时返回 nil

**示例**：
```lua
local pos = utf8.inStrRev("中国人民万岁", "人民", 6)
-- pos = 3

local pos = utf8.inStrRev("中国人民万岁", "人民")
-- pos = 3
```

### 4. utf8.strReverse(s) - UTF-8 字符串反转

**功能**：反转 UTF-8 字符串中的字符顺序

**参数**：
- `s` (string)：要反转的 UTF-8 字符串

**返回值**：
- `string`：反转后的 UTF-8 字符串

**示例**：
```lua
local result = utf8.strReverse("中国人民")
-- result = "民国中人"

local result = utf8.strReverse("Hello")
-- result = "olleH"
```

### 5. utf8.length(s) - UTF-8 字符串长度

**功能**：以字符为单位计算 UTF-8 字符串的长度（不是字节长度）

**参数**：
- `s` (string)：目标 UTF-8 字符串

**返回值**：
- `number`：字符串的字符长度

**示例**：
```lua
local len = utf8.length("中国人民万岁")
-- len = 6

local len = utf8.length("Hello")
-- len = 5

-- 与 string.len 的区别
local byteLen = string.len("中国人民万岁")  -- 字节长度，可能是 18
local charLen = utf8.length("中国人民万岁")  -- 字符长度，是 6
```

### 6. utf8.left(s, n) - UTF-8 字符串左侧截取

**功能**：截取 UTF-8 字符串左侧的 n 个字符

**参数**：
- `s` (string)：目标 UTF-8 字符串
- `n` (number)：截取的字符数

**返回值**：
- `string`：截取后的 UTF-8 字符串

**示例**：
```lua
local result = utf8.left("中国人民万岁", 2)
-- result = "中国"

local result = utf8.left("中国人民万岁", 10)
-- result = "中国人民万岁"（超过长度时返回整个字符串）
```

### 7. utf8.right(s, n) - UTF-8 字符串右侧截取

**功能**：截取 UTF-8 字符串右侧的 n 个字符

**参数**：
- `s` (string)：目标 UTF-8 字符串
- `n` (number)：截取的字符数

**返回值**：
- `string`：截取后的 UTF-8 字符串

**示例**：
```lua
local result = utf8.right("中国人民万岁", 2)
-- result = "万岁"

local result = utf8.right("中国人民万岁", 10)
-- result = "中国人民万岁"（超过长度时返回整个字符串）
```

### 8. utf8.mid(s, start, len) - UTF-8 字符串中间截取

**功能**：从指定位置开始截取 UTF-8 字符串的指定长度

**参数**：
- `s` (string)：目标 UTF-8 字符串
- `start` (number)：开始截取的位置（从 1 开始）
- `len` (number)：截取的字符数

**返回值**：
- `string`：截取后的 UTF-8 字符串

**示例**：
```lua
local result = utf8.mid("中国人民万岁", 2, 2)
-- result = "人民"

local result = utf8.mid("中国人民万岁", 3, 2)
-- result = "民万"

local result = utf8.mid("中国人民万岁", 5, 10)
-- result = "万岁"（超过长度时只截取到末尾）
```

### 9. utf8.strCut(s, start, len) - UTF-8 字符串裁剪

**功能**：移除 UTF-8 字符串中指定范围的字符并返回新字符串

**参数**：
- `s` (string)：目标 UTF-8 字符串
- `start` (number)：开始位置（从 1 开始）
- `len` (number)：要移除的字符数

**返回值**：
- `string`：裁剪后的 UTF-8 字符串

**示例**：
```lua
local result = utf8.strCut("中国人民万岁", 2, 2)
-- result = "中国万岁"（移除了"人民"）

local result = utf8.strCut("中国人民万岁", 3, 1)
-- result = "中国人万岁"（移除了"民"）

local result = utf8.strCut("中国人民万岁", 1, 2)
-- result = "民万岁"（移除了"中国"）
```

## Lua 标准库函数

以下函数是 Lua 标准库的一部分，AutoGo 完全支持：

### string.lower(s) / string.upper(s)
大小写转换

### string.find(s, pattern [, init [, plain]])
查找子串

### string.reverse(s)
字符串反转

### string.format(formatstring, ...)
字符串格式化

### string.char(...) / string.byte(s [, i [, j]])
字符编码转换

### string.len(s) / string.rep(s, n)
字符串长度与重复

## 使用示例

### 示例 1：处理 CSV 数据

```lua
-- 分割 CSV 行
local line = "张三,25,北京,工程师"
local fields = splitStr(line, ",")

for i, field in ipairs(fields) do
    print("字段 " .. i .. ": " .. field)
end
-- 输出：
-- 字段 1: 张三
-- 字段 2: 25
-- 字段 3: 北京
-- 字段 4: 工程师
```

### 示例 2：处理中文字符串

```lua
local text = "中国人民万岁"

-- 获取字符长度
print(utf8.length(text))  -- 6

-- 截取部分字符串
print(utf8.left(text, 2))   -- 中国
print(utf8.right(text, 2))  -- 万岁
print(utf8.mid(text, 3, 2)) -- 人民

-- 反转字符串
print(utf8.strReverse(text)) -- 岁万民人中

-- 裁剪字符串
print(utf8.strCut(text, 2, 2)) -- 中国万岁
```

### 示例 3：查找和替换

```lua
local text = "Hello 世界，Hello 中国"

-- 查找子串
local pos = utf8.inStr(1, text, "世界")
print(pos)  -- 7

-- 反向查找
local pos = utf8.inStrRev(text, "Hello", 20)
print(pos)  -- 9
```

## 注意事项

1. **UTF-8 支持**：所有 utf8.* 函数都正确处理多字节字符
2. **索引从 1 开始**：与 Lua 一致，所有位置索引都从 1 开始
3. **边界检查**：所有函数都会检查边界条件，不会越界
4. **全局函数**：`splitStr` 可以直接调用，也可以通过 `string.splitStr` 调用
5. **性能**：UTF-8 函数使用 Go 的 `unicode/utf8` 包，性能良好

## 与 Lua 标准库的区别

| 功能 | Lua 标准库 | AutoGo 扩展 |
|------|-----------|------------|
| 字符串分割 | ❌ 不支持 | ✅ splitStr |
| UTF-8 查找 | ❌ 不支持 | ✅ utf8.inStr/inStrRev |
| UTF-8 反转 | ❌ 不支持 | ✅ utf8.strReverse |
| UTF-8 长度 | ❌ 不支持 | ✅ utf8.length |
| UTF-8 截取 | ❌ 不支持 | ✅ utf8.left/right/mid |
| UTF-8 裁剪 | ❌ 不支持 | ✅ utf8.strCut |
| 字节长度 | ✅ string.len | ✅ 支持 |
| 字符长度 | ❌ 不支持 | ✅ utf8.length |

## 总结

string 模块为 AutoGo 提供了完整的字符串处理功能，包括：
- 懒人精灵特有的字符串分割功能
- 完整的 UTF-8 字符串处理支持
- 与 Lua 标准库的完全兼容

这些功能使得处理中文字符串和复杂文本变得更加简单和高效。
