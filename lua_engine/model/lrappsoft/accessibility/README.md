# accessibility - 无障碍服务模块

## 概述

accessibility 模块提供了基于 AutoGo 的 uiacc 模块的无障碍服务相关功能，包括节点查找、服务检测等。该模块是对懒人精灵 nodeLib 功能的封装，使用 AutoGo 的 uiacc 模块实现。

## 配置选项

accessibility 模块支持配置未实现方法的行为：

```go
type EmptyMethodConfig struct {
    ThrowException bool // 是否抛出异常（默认 false）
    ShowWarning    bool // 是否显示警告信息（默认 true）
}
```

### 使用示例

```go
// 创建模块实例并设置配置
accessibilityModule := accessibility.New(&accessibility.EmptyMethodConfig{
    ThrowException: false, // 不抛出异常，返回默认值
    ShowWarning:    true,  // 显示警告信息
})

// 运行时修改配置
accessibilityModule.SetConfig(&accessibility.EmptyMethodConfig{
    ThrowException: true, // 抛出异常
    ShowWarning:    true,
})

// 获取当前配置
config := accessibilityModule.GetConfig()
```

### 配置说明

- **ThrowException**: 当设置为 true 时，调用未实现的方法会抛出异常；当设置为 false 时，返回默认值
- **ShowWarning**: 当设置为 true 时，调用未实现的方法会打印警告信息到控制台

### 默认行为

默认情况下，未实现的方法会：
- 打印警告信息（ShowWarning: true）
- 不抛出异常（ThrowException: false）
- 返回默认值（nil）

## 功能说明

### 已实现的功能（6 个）

| 方法名 | 状态 | 说明 |
|--------|------|------|
| nodeLib.isAccServiceOk | ✅ 已实现 | 检测无障碍服务是否开启 |
| nodeLib.openSnapService | ✅ 已实现 | 打开截图服务 |
| nodeLib.isSnapServiceOk | ✅ 已实现 | 检测截图服务是否开启 |
| nodeLib.findOne | ✅ 已实现 | 查找一个节点 |
| nodeLib.findAll | ✅ 已实现 | 查找所有满足要求的节点 |
| nodeLib.findChildNodes | ✅ 已实现 | 查找一个节点的所有子节点 |

### 未实现的功能（3 个）

| 方法名 | 状态 | 原因 |
|--------|------|------|
| nodeLib.findNextNode | ❌ 未实现 | AutoGo 的 uiacc 模块不直接支持兄弟节点查找 |
| nodeLib.findPreNode | ❌ 未实现 | AutoGo 的 uiacc 模块不直接支持兄弟节点查找 |
| nodeLib.findInNode | ❌ 未实现 | AutoGo 的 uiacc 模块不直接支持在指定节点中查找 |

## API 文档

### nodeLib.isAccServiceOk

检测无障碍服务是否开启。

**语法：**
```lua
local ok = nodeLib.isAccServiceOk()
```

**返回值：**
- `boolean` - 始终返回 `true`（AutoGo 的 uiacc 模块无需开启无障碍服务）

**示例：**
```lua
local r = nodeLib.isAccServiceOk()
if r then
    print("无障碍服务已经开启")
else
    print("无障碍服务没有开启")
end
```

**注意事项：**
- AutoGo 的 nodeLib.isAccServiceOk 始终返回 true
- 不需要手动开启无障碍服务
- AutoGo 使用不同的技术实现节点查找

---

### nodeLib.openSnapService

打开截图服务。

**语法：**
```lua
local ok = nodeLib.openSnapService()
```

**返回值：**
- `boolean` - 始终返回 `true`（AutoGo 使用 images.CaptureScreen 截图，无需开启截图服务）

**示例：**
```lua
local r = nodeLib.openSnapService()
if r then
    if nodeLib.isAccServiceOk() == false then
        r = false
    end
end
if r then
    print("截图服务开启成功")
else
    print("截图服务开启失败")
end
```

**注意事项：**
- AutoGo 的 nodeLib.openSnapService 始终返回 true
- 不需要手动开启截图服务
- AutoGo 使用 images.CaptureScreen 截图

---

### nodeLib.isSnapServiceOk

检测截图服务是否开启。

**语法：**
```lua
local ok = nodeLib.isSnapServiceOk()
```

**返回值：**
- `boolean` - 始终返回 `true`（AutoGo 使用 images.CaptureScreen 截图，无需检测截图服务）

**示例：**
```lua
local r = nodeLib.isSnapServiceOk()
if r then
    print("截图服务已经开启")
else
    print("截图服务没有开启")
end
```

**注意事项：**
- AutoGo 的 nodeLib.isSnapServiceOk 始终返回 true
- 不需要手动开启截图服务
- AutoGo 使用 images.CaptureScreen 截图

---

### nodeLib.findOne

查找出一个节点。

**语法：**
```lua
local node = nodeLib.findOne(selector, fuzzy_match)
```

**参数：**
- `selector` (table) - 节点选择器，是一个表格，里面的键值对通过节点工具勾选直接生成
- `fuzzy_match` (boolean) - `true` 表示使用模糊匹配（包含匹配），`false` 表示完全匹配

**返回值：**
- `table` - 找到的节点对象，包含以下字段：
  - `class` (string) - 节点类名
  - `id` (string) - 节点 ID
  - `package` (string) - 节点包名
  - `text` (string) - 节点文本
  - `desc` (string) - 节点描述
  - `bounds` (table) - 节点位置信息，包含以下字段：
    - `left` (number) - 左边界
    - `top` (number) - 上边界
    - `right` (number) - 右边界
    - `bottom` (number) - 下边界
    - `width` (number) - 宽度
    - `height` (number) - 高度
- `nil` - 未找到节点

**示例：**
```lua
-- 完全匹配
local node = nodeLib.findOne({text = "确定"}, false)
if node ~= nil then
    print("节点文本:", node.text)
    print("节点 ID:", node.id)
    print("节点类名:", node.class)
    print("节点包名:", node.package)
    print("节点描述:", node.desc)
    print("节点位置:", node.bounds.left, node.bounds.top, node.bounds.right, node.bounds.bottom)
end

-- 模糊匹配
local node = nodeLib.findOne({text = "确定"}, true)
if node ~= nil then
    print("找到节点:", node.text)
end
```

**注意事项：**
- 返回的是包装的 Lua 表，不是原始的 UiObject 对象
- 使用点语法访问节点属性（如 `node.text`）
- 不支持节点操作方法（如点击、长按等）
- 模糊匹配使用包含匹配，不是正则表达式

---

### nodeLib.findAll

查找所有满足要求的节点。

**语法：**
```lua
local nodes = nodeLib.findAll(selector, fuzzy_match)
```

**参数：**
- `selector` (table) - 节点选择器，是一个表格，里面的键值对通过节点工具勾选直接生成
- `fuzzy_match` (boolean) - `true` 表示使用模糊匹配（包含匹配），`false` 表示完全匹配

**返回值：**
- `table` - 找到的节点对象数组，每个节点对象包含以下字段：
  - `class` (string) - 节点类名
  - `id` (string) - 节点 ID
  - `package` (string) - 节点包名
  - `text` (string) - 节点文本
  - `desc` (string) - 节点描述
  - `bounds` (table) - 节点位置信息
- `nil` - 未找到节点

**示例：**
```lua
-- 完全匹配
local nodes = nodeLib.findAll({text = "确定"}, false)
if nodes ~= nil then
    print("找到节点数量:", #nodes)
    for i, node in ipairs(nodes) do
        print("节点", i, "文本:", node.text)
        print("节点", i, "ID:", node.id)
        print("节点", i, "类名:", node.class)
        print("节点", i, "包名:", node.package)
        print("节点", i, "描述:", node.desc)
        print("节点", i, "位置:", node.bounds.left, node.bounds.top, node.bounds.right, node.bounds.bottom)
    end
end

-- 模糊匹配
local nodes = nodeLib.findAll({class = "android.widget.TextView"}, true)
if nodes ~= nil then
    print("找到", #nodes, "个 TextView")
end
```

**注意事项：**
- 返回的是包装的 Lua 表数组，不是原始的 UiObject 对象数组
- 使用点语法访问节点属性（如 `node.text`）
- 不支持节点操作方法（如点击、长按等）
- 模糊匹配使用包含匹配，不是正则表达式

---

### nodeLib.findChildNodes

查找一个节点的所有子节点。

**语法：**
```lua
local children = nodeLib.findChildNodes(selector, fuzzy_match)
```

**参数：**
- `selector` (table) - 节点选择器，是一个表格，里面的键值对通过节点工具勾选直接生成
- `fuzzy_match` (boolean) - `true` 表示使用模糊匹配（包含匹配），`false` 表示完全匹配

**返回值：**
- `table` - 找到的子节点对象数组，每个节点对象包含以下字段：
  - `class` (string) - 节点类名
  - `id` (string) - 节点 ID
  - `package` (string) - 节点包名
  - `text` (string) - 节点文本
  - `desc` (string) - 节点描述
  - `bounds` (table) - 节点位置信息
- `nil` - 未找到节点

**示例：**
```lua
-- 查找父节点的所有子节点
local children = nodeLib.findChildNodes({class = "android.widget.LinearLayout"}, true)
if children ~= nil then
    print("找到子节点数量:", #children)
    for i, node in ipairs(children) do
        print("子节点", i, "文本:", node.text)
        print("子节点", i, "ID:", node.id)
        print("子节点", i, "类名:", node.class)
        print("子节点", i, "包名:", node.package)
        print("子节点", i, "描述:", node.desc)
        print("子节点", i, "位置:", node.bounds.left, node.bounds.top, node.bounds.right, node.bounds.bottom)
    end
end
```

**注意事项：**
- 返回的是包装的 Lua 表数组，不是原始的 UiObject 对象数组
- 使用点语法访问节点属性（如 `node.text`）
- 不支持节点操作方法（如点击、长按等）
- 模糊匹配使用包含匹配，不是正则表达式

---

### nodeLib.findNextNode

获取指定节点的下一个兄弟节点。

**语法：**
```lua
local nextNode = nodeLib.findNextNode(selector, fuzzy_match)
```

**参数：**
- `selector` (table) - 节点选择器
- `fuzzy_match` (boolean) - 是否使用模糊匹配

**返回值：**
- `nil` - AutoGo 的 uiacc 模块不直接支持兄弟节点查找

**示例：**
```lua
local nextNode = nodeLib.findNextNode({class="android.widget.TextView",id="",package="com.nx.nxproj",text="名称: tmp.lr"},true)
if nextNode ~= nil then
    print(nextNode)
else
    print("未找到下一个兄弟节点")
end
```

**注意事项：**
- AutoGo 的 uiacc 模块不直接支持兄弟节点查找
- 返回 nil 表示不可用
- 可以使用 uiacc 模块的 children 方法遍历子节点来实现类似功能

---

### nodeLib.findPreNode

获取指定节点的上一个兄弟节点。

**语法：**
```lua
local preNode = nodeLib.findPreNode(selector, fuzzy_match)
```

**参数：**
- `selector` (table) - 节点选择器
- `fuzzy_match` (boolean) - 是否使用模糊匹配

**返回值：**
- `nil` - AutoGo 的 uiacc 模块不直接支持兄弟节点查找

**示例：**
```lua
local preNode = nodeLib.findPreNode({class="android.widget.TextView",id="",package="com.nx.nxproj",text="名称: tmp.lr"},false)
if preNode ~= nil then
    print(preNode)
else
    print("未找到上一个兄弟节点")
end
```

**注意事项：**
- AutoGo 的 uiacc 模块不直接支持兄弟节点查找
- 返回 nil 表示不可用
- 可以使用 uiacc 模块的 children 方法遍历子节点来实现类似功能

---

### nodeLib.findInNode

在指定节点中查找符合要求的子节点。

**语法：**
```lua
local childNodes = nodeLib.findInNode(src, dst, isfindall, fuzzy_match)
```

**参数：**
- `src` (table) - 要被查找的节点对象
- `dst` (table) - 要查找的节点对象
- `isfindall` (boolean) - 是否查找并返回所有符合要求的节点，`true` 返回所有，`false` 找到一个就返回
- `fuzzy_match` (boolean) - 是否使用模糊匹配

**返回值：**
- `nil` - AutoGo 的 uiacc 模块不直接支持在指定节点中查找

**示例：**
```lua
local ret = nodeLib.findOne(0,0,0,0,{class="android.widget.FrameLayout",id="com.nx.nxproj.assist:id/title"},true)
if ret ~= nil then
    local r = nodeLib.findInNode(ret,{text=".*新建.*"},true,true)
    print(ret)
    if r~= nil then
        for i=1,#r do
            print(r[i])
        end
    end
end
```

**注意事项：**
- AutoGo 的 uiacc 模块不直接支持在指定节点中查找
- 返回 nil 表示不可用
- 可以使用 uiacc 模块的 children 方法获取所有子节点，然后手动过滤

---

## 与懒人精灵的差异

### 主要差异

| 特性 | 懒人精灵 | AutoGo |
|------|----------|---------|
| 返回值 | UiObject 对象 | Lua 表 |
| 访问方式 | 使用冒号语法（:） | 使用点语法（.） |
| 节点方法 | 支持点击、长按等操作 | 不支持节点操作 |
| 模糊匹配 | 支持正则表达式 | 支持包含匹配 |
| 无障碍服务 | 需要手动开启 | 无需开启 |
| 截图服务 | 需要手动开启 | 无需开启 |

### 代码迁移示例

**懒人精灵代码：**
```lua
-- 查找节点并点击
local node = nodeLib.findOne({text = "确定"}, true)
if node ~= nil then
    print("节点文本:", node:text())
    node:click()
end
```

**AutoGo 代码：**
```lua
-- 查找节点（注意使用点语法访问属性）
local node = nodeLib.findOne({text = "确定"}, true)
if node ~= nil then
    print("节点文本:", node.text)
    -- 不支持节点操作方法，需要使用其他方式点击
    -- 例如使用 touch 模块点击节点位置
    touch.tap(node.bounds.left + node.bounds.width/2, node.bounds.top + node.bounds.height/2)
end
```

---

## 底层实现

accessibility 模块基于 AutoGo 的 uiacc 模块实现，提供了基于辅助功能服务的控件定位、交互操作等功能。无需开启 APP 的无障碍服务。

### AutoGo uiacc 模块

- 提供基于辅助功能服务的控件定位、交互操作等功能
- 无需开启 APP 的无障碍服务
- 支持多种选择器条件（text、id、class、package 等）
- 支持模糊匹配和精确匹配

### 相关文档

- [AutoGo 官方文档](../../AutoGo/docs)
- [uiacc 模块 API](../../AutoGo/docs/API/uiacc.md)
- [lrappsoft 扩展方法](../extension/README.md)

---

## 使用建议

1. **节点查找**：优先使用 nodeLib.findOne 和 nodeLib.findAll 方法
2. **属性访问**：使用点语法访问节点属性（如 `node.text`）
3. **节点操作**：不支持节点操作方法，需要使用其他模块（如 touch 模块）
4. **模糊匹配**：使用包含匹配，不是正则表达式
5. **服务检测**：无需手动检测无障碍服务和截图服务

---

## 更新日志

### v1.0.0 (2026-03-16)
- 初始版本
- 实现了 6 个 nodeLib 方法
- 使用 AutoGo 的 uiacc 模块作为底层实现
- 提供了完整的 API 文档和使用示例
