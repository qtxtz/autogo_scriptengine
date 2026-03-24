# node - 节点选择器模块

## 概述

node 模块提供了基于 AutoGo 的 uiacc 模块的节点选择器功能，包括各种选择器方法和节点操作方法。该模块是对懒人精灵 node 功能的封装，使用 AutoGo 的 uiacc 模块实现。

## 核心架构

### 缓存机制

node 模块使用线程安全的缓存机制来存储 UIObject 对象：

1. **缓存 UIObject**：当查找节点时，将 UIObject 对象缓存到 Go 的 map 中
2. **返回 ID**：返回包含 `__node_id__` 字段的 Lua 表
3. **通过 ID 访问**：节点方法通过 ID 从缓存中获取 UIObject 对象

### 优势

- **线程安全**：使用 `sync.RWMutex` 保证并发访问安全
- **高效访问**：通过 ID 快速查找 UIObject，无需重复查找
- **内存管理**：避免重复创建 UIObject 对象
- **完整功能**：支持所有节点操作方法

## 功能说明

### 选择器方法（42 个）

| 方法名 | 状态 | 说明 |
|--------|------|------|
| id | ✅ 已实现 | 根据节点的id属性匹配 |
| idContains | ✅ 已实现 | 根据id包含的部分字符串匹配 |
| idStartsWith | ✅ 已实现 | 根据id的前缀去匹配 |
| idEndsWith | ✅ 已实现 | 根据id的后缀去匹配 |
| idMatches | ✅ 已实现 | 正则匹配id |
| text | ✅ 已实现 | 根据节点的text属性匹配 |
| textContains | ✅ 已实现 | 根据text包含的部分字符串匹配 |
| textStartsWith | ✅ 已实现 | 根据text的前缀去匹配 |
| textEndsWith | ✅ 已实现 | 根据text的后缀去匹配 |
| textMatches | ✅ 已实现 | 正则匹配text |
| desc | ✅ 已实现 | 根据节点的desc属性匹配 |
| descContains | ✅ 已实现 | 根据desc包含的部分字符串匹配 |
| descStartsWith | ✅ 已实现 | 根据text的前缀去匹配 |
| descEndsWith | ✅ 已实现 | 根据text的后缀去匹配 |
| descMatches | ✅ 已实现 | 正则匹配text |
| className | ✅ 已实现 | 根据节点的className属性匹配 |
| classNameContains | ✅ 已实现 | 根据className属性所包含的字符串模糊匹配 |
| classNameStartsWith | ✅ 已实现 | 根据className属性前缀匹配 |
| classNameEndsWith | ✅ 已实现 | 根据className属性后缀匹配 |
| classNameMatches | ✅ 已实现 | 根据className属性正则匹配 |
| packageName | ✅ 已实现 | 根据packageName属性全字段匹配 |
| packageNameContains | ✅ 已实现 | 匹配包含指定字符串的的节点 |
| packageNameStartsWith | ✅ 已实现 | 匹配包名前缀为指定字符串的的节点 |
| packageNameEndsWith | ✅ 已实现 | 匹配包名后缀为指定字符串的的节点 |
| packageNameMatches | ✅ 已实现 | 包名正则匹配 |
| bounds | ✅ 已实现 | 根据节点的范围匹配 |
| boundsInside | ✅ 已实现 | 匹配该范围内的节点 |
| drawingOrder | ✅ 已实现 | 根据绘制顺序匹配 |
| depth | ✅ 已实现 | 根据深度匹配（通过遍历 UI 树实现） |
| index | ✅ 已实现 | 根据索引匹配 |
| visibleToUser | ✅ 已实现 | 根据是否可见匹配 |
| selected | ✅ 已实现 | 根据是否选中匹配 |
| clickable | ✅ 已实现 | 根据是否可点击匹配 |
| longClickable | ✅ 已实现 | 根据是否可长按点击匹配 |
| enabled | ✅ 已实现 | 根据是否可用匹配 |
| password | ✅ 已实现 | 根据是否是密码框匹配 |
| scrollable | ✅ 已实现 | 根据是否可以滚动来匹配 |
| checked | ✅ 已实现 | 是否被勾选来匹配 |
| checkable | ✅ 已实现 | 是否可以被勾选来匹配 |
| focusable | ✅ 已实现 | 根据是否允许抢占焦点来匹配 |
| focused | ✅ 已实现 | 根据是否抢占了焦点来匹配 |

### 选择器查找方法（3 个）

| 方法名 | 状态 | 说明 |
|--------|------|------|
| sel:findOne | ✅ 已实现 | 获取匹配到的第一个节点 |
| sel:findAll | ✅ 已实现 | 获取匹配到的所有节点 |
| sel:findOnce | ✅ 已实现 | 根据索引获取匹配到的节点 |

### 选择器操作方法（2 个）

| 方法名 | 状态 | 说明 |
|--------|------|------|
| sel:click | ✅ 已实现 | 点击所有匹配到的节点 |
| sel:longClick | ✅ 已实现 | 长按点击所有匹配到的节点 |

### 节点属性方法（8 个）

| 方法名 | 状态 | 说明 |
|--------|------|------|
| node:id | ✅ 已实现 | 获取节点的id属性值 |
| node:toJson | ✅ 已实现 | 获取节点的所有属性以json字符串的形式 |
| node:text | ✅ 已实现 | 获取节点的text属性值 |
| node:click | ✅ 已实现 | 点击该节点 |
| node:longClick | ✅ 已实现 | 长按点击该节点 |
| node:className | ✅ 已实现 | 获取节点的className属性值 |
| node:packageName | ✅ 已实现 | 获取节点的packageName属性值 |
| node:desc | ✅ 已实现 | 获取节点的desc属性值 |
| node:bounds | ✅ 已实现 | 获取节点的bounds属性值 |

## API 文档

### 选择器方法

#### id

根据节点的id属性匹配。

**语法：**
```lua
local sel = id(str)
```

**参数：**
- `str` (string) - 指定要查找的具体控件的id

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = id("com.nx.nxprojit:id/detail_tv_title")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### idContains

根据id包含的部分字符串匹配。

**语法：**
```lua
local sel = idContains(str)
```

**参数：**
- `str` (string) - 指定要查找的模糊的包含id某一段字符串的内容

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = idContains("detail_tv_title")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### idStartsWith

根据id的前缀去匹配。

**语法：**
```lua
local sel = idStartsWith(str)
```

**参数：**
- `str` (string) - 指定要查找的id的前缀

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = idStartsWith("com.nx.nxprojit")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:id())
end
```

---

#### idEndsWith

根据id的后缀去匹配。

**语法：**
```lua
local sel = idEndsWith(str)
```

**参数：**
- `str` (string) - 指定要查找的id的后缀

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = idEndsWith("detail_tv_title")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### idMatches

正则匹配id。

**语法：**
```lua
local sel = idMatches(str)
```

**参数：**
- `str` (string) - 正则表达式字符串

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = idMatches(".*detail_tv_title$")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### text

根据节点的text属性匹配。

**语法：**
```lua
local sel = text(str)
```

**参数：**
- `str` (string) - 指定要查找的具体控件的text

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = text("懒人高级版")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### textContains

根据text包含的部分字符串匹配。

**语法：**
```lua
local sel = textContains(str)
```

**参数：**
- `str` (string) - 指定要查找的模糊的包含text某一段字符串的内容

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = textContains("懒人")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### textStartsWith

根据text的前缀去匹配。

**语法：**
```lua
local sel = textStartsWith(str)
```

**参数：**
- `str` (string) - 指定要查找的text的前缀

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = textStartsWith("懒人")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### textEndsWith

根据text的后缀去匹配。

**语法：**
```lua
local sel = textEndsWith(str)
```

**参数：**
- `str` (string) - 指定要查找的text的后缀

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = textEndsWith("高级版")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### textMatches

正则匹配text。

**语法：**
```lua
local sel = textMatches(str)
```

**参数：**
- `str` (string) - 正则表达式字符串

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = textMatches(".*高级版$")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### desc

根据节点的desc属性匹配。

**语法：**
```lua
local sel = desc(str)
```

**参数：**
- `str` (string) - 指定要查找的具体控件的desc

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = desc("懒人高级版")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### descContains

根据desc包含的部分字符串匹配。

**语法：**
```lua
local sel = descContains(str)
```

**参数：**
- `str` (string) - 指定要查找的模糊的包含desc某一段字符串的内容

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = descContains("懒人")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### descStartsWith

根据text的前缀去匹配。

**语法：**
```lua
local sel = descStartsWith(str)
```

**参数：**
- `str` (string) - 指定要查找的desc的前缀

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = descStartsWith("懒人")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### descEndsWith

根据text的后缀去匹配。

**语法：**
```lua
local sel = descEndsWith(str)
```

**参数：**
- `str` (string) - 指定要查找的desc的后缀

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = descEndsWith("高级版")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### descMatches

正则匹配text。

**语法：**
```lua
local sel = descMatches(str)
```

**参数：**
- `str` (string) - 正则表达式字符串

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = descMatches(".*高级版$")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### className

根据节点的className属性匹配。

**语法：**
```lua
local sel = className(str)
```

**参数：**
- `str` (string) - 指定要查找的具体控件的类名

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### classNameContains

根据className属性所包含的字符串模糊匹配。

**语法：**
```lua
local sel = classNameContains(str)
```

**参数：**
- `str` (string) - 指定查找className包含该字符的节点

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = classNameContains("android.widget")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### classNameStartsWith

根据className属性前缀匹配。

**语法：**
```lua
local sel = classNameStartsWith(str)
```

**参数：**
- `str` (string) - 指定查找className前缀为该字符的节点

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = classNameStartsWith("android.widget")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### classNameEndsWith

根据className属性后缀匹配。

**语法：**
```lua
local sel = classNameEndsWith(str)
```

**参数：**
- `str` (string) - 指定查找className后缀为该字符的节点

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = classNameEndsWith("TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### classNameMatches

根据className属性正则匹配。

**语法：**
```lua
local sel = classNameMatches(str)
```

**参数：**
- `str` (string) - 指定要匹配className的正则表达式

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = classNameMatches("^android.*TextView$")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### packageName

根据packageName属性全字段匹配。

**语法：**
```lua
local sel = packageName(str)
```

**参数：**
- `str` (string) - 指定要匹配节点所属的包名

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = packageName("com.nx.nxprojit"):className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### packageNameContains

匹配包含指定字符串的的节点。

**语法：**
```lua
local sel = packageNameContains(str)
```

**参数：**
- `str` (string) - 包含该字符串的包名的节点将被匹配到

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = packageNameContains("nxprojit"):className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### packageNameStartsWith

匹配包名前缀为指定字符串的的节点。

**语法：**
```lua
local sel = packageNameStartsWith(str)
```

**参数：**
- `str` (string) - 包名为该字符的前缀的节点将被匹配到

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = packageNameStartsWith("com.nx"):className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### packageNameEndsWith

匹配包名后缀为指定字符串的的节点。

**语法：**
```lua
local sel = packageNameEndsWith(str)
```

**参数：**
- `str` (string) - 包名为该字符的后缀的节点将被匹配到

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = packageNameEndsWith("nx.nxprojit"):className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### packageNameMatches

包名正则匹配。

**语法：**
```lua
local sel = packageNameMatches(str)
```

**参数：**
- `str` (string) - 根据这个正则字符串匹配所有符合规则包名的节点

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = packageNameMatches("com.nx.*"):className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### bounds

根据节点的范围匹配。

**语法：**
```lua
local sel = bounds(l, t, r, b)
```

**参数：**
- `l` (number) - 左范围
- `t` (number) - 上范围
- `r` (number) - 右范围
- `b` (number) - 下范围

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = bounds(30, 212, 236, 250)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### boundsInside

匹配该范围内的节点。

**语法：**
```lua
local sel = boundsInside(l, t, r, b)
```

**参数：**
- `l` (number) - 左范围
- `t` (number) - 上范围
- `r` (number) - 右范围
- `b` (number) - 下范围

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = boundsInside(30, 212, 236, 250)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### drawingOrder

根据绘制顺序匹配。

**语法：**
```lua
local sel = drawingOrder(level)
```

**参数：**
- `level` (number) - 绘制顺序级别

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = drawingOrder(1)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### depth

根据深度匹配。

**语法：**
```lua
local sel = depth(level)
```

**参数：**
- `level` (number) - 深度级别（从 0 开始，0 表示根节点）

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**实现原理：**
- 通过遍历 UI 树计算每个节点的深度
- 使用 `GetParent()` 方法向上遍历到根节点，计算深度
- 在查找时过滤出指定深度的节点

**示例：**
```lua
-- 查找深度为 2 的所有 TextView
local sel = className("android.widget.TextView"):depth(2)
local nodes = sel:findAll()
if nodes ~= nil then
    for i = 1, #nodes do
        print(nodes[i]:text())
    end
end

-- 查找深度为 1 的按钮
local sel = depth(1):clickable(true)
local node = sel:findOne(10000)
if node ~= nil then
    node:click()
end
```

---

#### index

根据索引匹配。

**语法：**
```lua
local sel = index(idx)
```

**参数：**
- `idx` (number) - 索引值

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = index(0)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### visibleToUser

根据是否可见匹配。

**语法：**
```lua
local sel = visibleToUser(b)
```

**参数：**
- `b` (boolean) - 是否可见

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = visibleToUser(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### selected

根据是否选中匹配。

**语法：**
```lua
local sel = selected(b)
```

**参数：**
- `b` (boolean) - 是否选中

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = selected(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### clickable

根据是否可点击匹配。

**语法：**
```lua
local sel = clickable(b)
```

**参数：**
- `b` (boolean) - 是否可点击

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = clickable(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### longClickable

根据是否可长按点击匹配。

**语法：**
```lua
local sel = longClickable(b)
```

**参数：**
- `b` (boolean) - 是否可长按点击

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = longClickable(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### enabled

根据是否可用匹配。

**语法：**
```lua
local sel = enabled(b)
```

**参数：**
- `b` (boolean) - 是否可用

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = enabled(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### password

根据是否是密码框匹配。

**语法：**
```lua
local sel = password(b)
```

**参数：**
- `b` (boolean) - 是否是密码框

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = password(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### scrollable

根据是否可以滚动来匹配。

**语法：**
```lua
local sel = scrollable(b)
```

**参数：**
- `b` (boolean) - 是否可以滚动

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = scrollable(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### checked

是否被勾选来匹配。

**语法：**
```lua
local sel = checked(b)
```

**参数：**
- `b` (boolean) - 是否被勾选

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = checked(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### checkable

是否可以被勾选来匹配。

**语法：**
```lua
local sel = checkable(b)
```

**参数：**
- `b` (boolean) - 是否可以被勾选

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = checkable(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### focusable

根据是否允许抢占焦点来匹配。

**语法：**
```lua
local sel = focusable(b)
```

**参数：**
- `b` (boolean) - 是否允许抢占焦点

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = focusable(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### focused

根据是否抢占了焦点来匹配。

**语法：**
```lua
local sel = focused(b)
```

**参数：**
- `b` (boolean) - 是否抢占了焦点

**返回值：**
- `selector` - 返回一个选择器对象，该对象支持级联选择

**示例：**
```lua
local sel = focused(true)
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

### 选择器查找方法

#### sel:findOne

获取匹配到的第一个节点。

**语法：**
```lua
local node = sel:findOne(timeout)
```

**参数：**
- `timeout` (number, 可选) - 匹配超时时间（毫秒），默认 10000

**返回值：**
- `node` - 返回匹配到的节点对象（包含 `__node_id__` 的 Lua 表）

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### sel:findAll

获取匹配到的所有节点。

**语法：**
```lua
local nodes = sel:findAll(timeout)
```

**参数：**
- `timeout` (number, 可选) - 匹配超时时间（毫秒），默认 10000

**返回值：**
- `nodes` - 返回匹配到的节点对象数组（每个节点都是包含 `__node_id__` 的 Lua 表）

**示例：**
```lua
local sel = className("android.widget.TextView")
local nodes = sel:findAll(10000)
if nodes ~= nil then
    for i = 1, #nodes do
        print(nodes[i]:text())
    end
end
```

---

#### sel:findOnce

根据索引获取匹配到的节点。

**语法：**
```lua
local node = sel:findOnce(index)
```

**参数：**
- `index` (number, 可选) - 节点索引，默认 0

**返回值：**
- `node` - 返回匹配到的节点对象（包含 `__node_id__` 的 Lua 表）

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOnce(1) -- 匹配查找到的第二个节点
if node ~= nil then
    print(node:text())
end
```

---

### 选择器操作方法

#### sel:click

点击所有匹配到的节点。

**语法：**
```lua
local success = sel:click()
```

**返回值：**
- `boolean` - true表示成功，false表示失败

**示例：**
```lua
local sel = className("android.widget.TextView"):text("雷电游戏中心")
sel:click()
```

**实现细节：**
- 使用 AutoGo 的 touch 模块点击节点中心
- 支持点击多个匹配的节点
- 每次点击后等待 100ms

---

#### sel:longClick

长按点击所有匹配到的节点。

**语法：**
```lua
local success = sel:longClick()
```

**返回值：**
- `boolean` - true表示成功，false表示失败

**示例：**
```lua
local sel = className("android.widget.TextView"):text("雷电游戏中心")
sel:longClick()
```

**实现细节：**
- 使用 AutoGo 的 touch 模块长按节点中心
- 长按时间为 500ms
- 支持长按多个匹配的节点
- 每次长按后等待 100ms

---

### 节点属性方法

#### node:id

获取节点的id属性值。

**语法：**
```lua
local id = node:id()
```

**返回值：**
- `string` - 返回id

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:id())
end
```

---

#### node:toJson

获取节点的所有属性以json字符串的形式。

**语法：**
```lua
local jsonStr = node:toJson()
```

**返回值：**
- `string` - 返回json字符串

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:toJson())
end
```

**返回的 JSON 格式：**
```json
{
    "id": "com.nx.nxprojit:id/detail_tv_title",
    "text": "懒人高级版",
    "desc": "描述",
    "className": "android.widget.TextView",
    "packageName": "com.nx.nxprojit"
}
```

---

#### node:text

获取节点的text属性值。

**语法：**
```lua
local text = node:text()
```

**返回值：**
- `string` - 返回text

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:text())
end
```

---

#### node:click

点击该节点。

**语法：**
```lua
local success = node:click()
```

**返回值：**
- `boolean` - true表示成功，false表示失败

**示例：**
```lua
local sel = className("android.widget.TextView"):text("确定")
local node = sel:findOne(10000)
if node ~= nil then
    node:click()
end
```

**实现细节：**
- 使用 AutoGo 的 touch 模块点击节点中心
- 从缓存中获取 UIObject 对象

---

#### node:longClick

长按点击该节点。

**语法：**
```lua
local success = node:longClick()
```

**返回值：**
- `boolean` - true表示成功，false表示失败

**示例：**
```lua
local sel = className("android.widget.TextView"):text("确定")
local node = sel:findOne(10000)
if node ~= nil then
    node:longClick()
end
```

**实现细节：**
- 使用 AutoGo 的 touch 模块长按节点中心
- 长按时间为 500ms
- 从缓存中获取 UIObject 对象

---

#### node:className

获取节点的className属性值。

**语法：**
```lua
local className = node:className()
```

**返回值：**
- `string` - 返回className

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:className())
end
```

---

#### node:packageName

获取节点的packageName属性值。

**语法：**
```lua
local packageName = node:packageName()
```

**返回值：**
- `string` - 返回packageName

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:packageName())
end
```

---

#### node:desc

获取节点的desc属性值。

**语法：**
```lua
local desc = node:desc()
```

**返回值：**
- `string` - 返回desc

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print(node:desc())
end
```

---

#### node:bounds

获取节点的bounds属性值。

**语法：**
```lua
local bounds = node:bounds()
```

**返回值：**
- `table` - 返回bounds表，包含以下字段：
  - `left` (number) - 左边界
  - `top` (number) - 上边界
  - `right` (number) - 右边界
  - `bottom` (number) - 下边界
  - `width` (number) - 宽度
  - `height` (number) - 高度
  - `centerX` (number) - 中心点 X 坐标
  - `centerY` (number) - 中心点 Y 坐标

**示例：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    local bounds = node:bounds()
    print("左:", bounds.left)
    print("上:", bounds.top)
    print("右:", bounds.right)
    print("下:", bounds.bottom)
    print("宽度:", bounds.width)
    print("高度:", bounds.height)
    print("中心X:", bounds.centerX)
    print("中心Y:", bounds.centerY)
end
```

---

## 使用示例

### 示例 1：基本选择器使用

```lua
-- 使用 id 选择器
local sel = id("com.nx.nxprojit:id/detail_tv_title")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
end

-- 使用 text 选择器
local sel = text("确定")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点 ID:", node:id())
end

-- 使用 className 选择器
local sel = className("android.widget.TextView")
local nodes = sel:findAll(10000)
if nodes ~= nil then
    print("找到", #nodes, "个 TextView")
end
```

### 示例 2：级联选择器

```lua
-- 级联选择器
local sel = packageName("com.nx.nxprojit"):className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
end

-- 多条件级联
local sel = packageName("com.nx.nxprojit"):className("android.widget.TextView"):textContains("懒人")
local nodes = sel:findAll(10000)
if nodes ~= nil then
    for i = 1, #nodes do
        print("节点", i, ":", nodes[i]:text())
    end
end
```

### 示例 3：模糊匹配

```lua
-- textContains 模糊匹配
local sel = textContains("懒人")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
end

-- idContains 模糊匹配
local sel = idContains("detail_tv_title")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点 ID:", node:id())
end

-- textMatches 正则匹配
local sel = textMatches(".*高级版$")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
end
```

### 示例 4：使用 bounds 选择器

```lua
-- bounds 选择器
local sel = bounds(30, 212, 236, 250)
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
end

-- boundsInside 选择器
local sel = boundsInside(30, 212, 236, 250)
local nodes = sel:findAll(10000)
if nodes ~= nil then
    print("找到", #nodes, "个节点")
end
```

### 示例 5：使用属性选择器

```lua
-- clickable 选择器
local sel = clickable(true)
local nodes = sel:findAll(10000)
if nodes ~= nil then
    print("找到", #nodes, "个可点击节点")
end

-- enabled 选择器
local sel = enabled(true)
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
end

-- scrollable 选择器
local sel = scrollable(true)
local nodes = sel:findAll(10000)
if nodes ~= nil then
    print("找到", #nodes, "个可滚动节点")
end
```

### 示例 6：获取节点信息

```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    -- 获取节点 ID
    print("节点 ID:", node:id())
    
    -- 获取节点文本
    print("节点文本:", node:text())
    
    -- 获取节点类名
    print("节点类名:", node:className())
    
    -- 获取节点包名
    print("节点包名:", node:packageName())
    
    -- 获取节点描述
    print("节点描述:", node:desc())
    
    -- 获取节点所有属性（JSON 格式）
    print("节点 JSON:", node:toJson())
    
    -- 获取节点边界
    local bounds = node:bounds()
    print("节点边界:", bounds.left, bounds.top, bounds.right, bounds.bottom)
end
```

### 示例 7：节点操作

```lua
-- 点击节点
local sel = text("确定")
local node = sel:findOne(10000)
if node ~= nil then
    node:click()
end

-- 长按节点
local sel = text("设置")
local node = sel:findOne(10000)
if node ~= nil then
    node:longClick()
end

-- 选择器点击所有匹配的节点
local sel = className("android.widget.Button"):clickable(true)
sel:click()

-- 选择器长按所有匹配的节点
local sel = className("android.widget.TextView"):longClickable(true)
sel:longClick()
```

### 示例 8：遍历所有匹配的节点

```lua
local sel = className("android.widget.TextView")
local nodes = sel:findAll(10000)
if nodes ~= nil then
    for i = 1, #nodes do
        local node = nodes[i]
        print("节点", i)
        print("  ID:", node:id())
        print("  文本:", node:text())
        print("  类名:", node:className())
        print("  包名:", node:packageName())
        print("  描述:", node:desc())
        print("  JSON:", node:toJson())
        
        local bounds = node:bounds()
        print("  边界:", bounds.left, bounds.top, bounds.right, bounds.bottom)
    end
end
```

### 示例 9：使用 index 选择器

```lua
-- 获取第一个匹配的节点
local sel = className("android.widget.TextView"):index(0)
local node = sel:findOne(10000)
if node ~= nil then
    print("第一个节点:", node:text())
end

-- 获取第二个匹配的节点
local sel = className("android.widget.TextView"):index(1)
local node = sel:findOne(10000)
if node ~= nil then
    print("第二个节点:", node:text())
end
```

---

## 与懒人精灵的差异

### 主要差异

| 特性 | 懒人精灵 | AutoGo |
|------|----------|---------|
| 返回值 | UiObject 对象 | 包含 `__node_id__` 的 Lua 表 |
| 访问方式 | 使用冒号语法（:） | 使用冒号语法（:） |
| 节点方法 | 支持多种节点操作 | 支持所有节点操作方法 |
| 选择器方法 | 完全支持 | 完全支持（depth 除外） |
| 查找方法 | 完全支持 | 完全支持 |
| 操作方法 | 完全支持 | 完全支持 |
| 缓存机制 | 无 | 使用线程安全的缓存 |

### 代码迁移示例

**懒人精灵代码：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
    print("节点 ID:", node:id())
    node:click()
end
```

**AutoGo 代码：**
```lua
local sel = className("android.widget.TextView")
local node = sel:findOne(10000)
if node ~= nil then
    print("节点文本:", node:text())
    print("节点 ID:", node:id())
    node:click()
end
```

**注意：** 代码完全兼容，无需修改！

---

## 底层实现

node 模块基于 AutoGo 的 uiacc 模块实现，提供了基于辅助功能服务的控件定位、交互操作等功能。

### AutoGo uiacc 模块

- 提供基于辅助功能服务的控件定位、交互操作等功能
- 支持多种选择器条件（text、id、class、package 等）
- 支持模糊匹配和精确匹配
- 支持正则表达式匹配

### AutoGo touch 模块

- 提供触摸操作功能
- 支持点击（Tap）和长按（LongPress）
- 用于实现节点的点击和长按操作

### 缓存机制

- 使用 `map[int64]*uiacc.UiObject` 缓存 UIObject 对象
- 使用 `sync.RWMutex` 保证线程安全
- 通过唯一 ID 快速查找 UIObject 对象

### 相关文档

- [AutoGo 官方文档](../../AutoGo/docs)
- [uiacc 模块 API](../../AutoGo/docs/API/uiacc.md)
- [touch 模块 API](../../AutoGo/docs/API/touch.md)
- [accessibility 模块](../accessibility/README.md)

---

## 使用建议

1. **选择器使用**：优先使用 text、id、className 等常用选择器
2. **级联选择**：支持级联选择，可以组合多个条件
3. **模糊匹配**：使用 Contains、StartsWith、EndsWith 进行模糊匹配
4. **正则匹配**：使用 Matches 方法进行正则表达式匹配
5. **节点操作**：支持所有节点操作方法（click、longClick 等）
6. **缓存机制**：自动缓存 UIObject 对象，无需手动管理
7. **线程安全**：支持多线程并发访问

---

## 更新日志

### v2.0.0 (2026-03-16)
- 重构节点对象实现，使用缓存机制
- 实现 sel:click 和 sel:longClick 方法
- 实现 node:click 和 node:longClick 方法
- 实现 index 选择器方法
- 添加更多节点属性方法（className、packageName、desc、bounds）
- 优化性能，避免重复查找 UIObject
- 保证线程安全，支持并发访问

### v1.0.0 (2026-03-16)
- 初始版本
- 实现了 41 个选择器方法
- 实现了 3 个查找方法（findOne, findAll, findOnce）
- 实现了 3 个节点属性方法（id, toJson, text）
- 提供了完整的 API 文档和使用示例
