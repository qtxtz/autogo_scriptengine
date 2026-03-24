# ui 模块（懒人精灵 UI 方法）

## 概述

ui 模块提供了懒人精灵的 UI 相关方法，包括窗口管理、WebView 控件、UI 控件操作、HUD 显示等功能。

## ⚠️ 重要说明

### ❌ 所有方法均为空操作

**重要**：ui 模块中的所有方法都是**空操作**，不会执行任何实际功能。

**原因**：
- 懒人精灵的自定义 UI 系统（包括 showUI、窗口管理、WebView 控件、UI 控件操作等）是懒人精灵特有的功能
- AutoGo 使用完全不同的架构和 UI 系统
- 这些功能无法通过 AutoGo API 或 adb 命令实现
- 实现这些功能需要重写整个 UI 系统，工作量巨大且与 AutoGo 的设计理念不符

**说明**：
- 所有方法可以正常调用，不会报错
- 方法会返回默认值以保持 API 兼容性
- 这些方法的存在只是为了保持与懒人精灵 API 的兼容性
- **不建议**在 AutoGo 中使用这些方法

### 🔧 配置选项

ui 模块提供了配置选项来控制空方法的行为：

```go
type EmptyMethodConfig struct {
    ThrowException bool // 是否抛出异常
    ShowWarning    bool // 是否显示警告信息
}
```

#### 配置说明

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `ThrowException` | bool | false | 是否在调用空方法时抛出异常 |
| `ShowWarning` | bool | true | 是否在调用空方法时打印警告信息 |

#### 配置示例

```go
// 创建 ui 模块（使用默认配置）
uiModule := ui.New(nil)

// 创建 ui 模块（自定义配置）
config := &ui.EmptyMethodConfig{
    ThrowException: false, // 不抛出异常
    ShowWarning:    true,  // 显示警告信息
}
uiModule := ui.New(config)

// 运行时修改配置
uiModule.SetConfig(&ui.EmptyMethodConfig{
    ThrowException: true, // 抛出异常
    ShowWarning:    true, // 显示警告信息
})
```

#### 使用示例

```lua
-- 默认配置（显示警告，不抛出异常）
closeWindow()
-- 输出：[警告] ui.closeWindow 方法未实现，此功能在 AutoGo 中不支持

-- 配置为抛出异常
-- 在 Go 代码中设置：config.ThrowException = true
closeWindow()
-- 输出：
-- [警告] ui.closeWindow 方法未实现，此功能在 AutoGo 中不支持
-- 抛出异常：ui.closeWindow 方法未实现，此功能在 AutoGo 中不支持

-- 配置为不显示警告
-- 在 Go 代码中设置：config.ShowWarning = false
closeWindow()
-- 无任何输出，静默执行
```

## 功能列表

### 窗口管理功能（1 个）

1. **closeWindow** - 关闭窗口
   - **实现状态**：空操作
   - **返回值**：无

### WebView 功能（2 个）

2. **getUIWebViewUrl** - 获取当前浏览器的地址
   - **实现状态**：空操作
   - **返回值**：空字符串 ""

3. **setUIWebViewUrl** - 设置浏览器地址
   - **实现状态**：空操作
   - **返回值**：无

### UI 控件操作功能（15 个）

4. **getUISelected** - 获取单选框或者下拉框当前选中项
   - **实现状态**：空操作
   - **返回值**：0

5. **getUISelectText** - 获取单选框或者下拉框当前选中项的文本内容
   - **实现状态**：空操作
   - **返回值**：空字符串 ""

6. **getUIText** - 获取控件显示的文本内容
   - **实现状态**：空操作
   - **返回值**：空字符串 ""

7. **setUIText** - 设置控件文本
   - **实现状态**：空操作
   - **返回值**：无

8. **setUISelect** - 设置单选框或者下拉框选项被选中
   - **实现状态**：空操作
   - **返回值**：无

9. **setUICheck** - 设置多选框被选中或者反选
   - **实现状态**：空操作
   - **返回值**：无

10. **getUICheck** - 获取多选框状态
    - **实现状态**：空操作
    - **返回值**：false

11. **getUIEnable** - 获取控件是否可用
    - **实现状态**：空操作
    - **返回值**：false

12. **setUIEnable** - 设置控件是否可用
    - **实现状态**：空操作
    - **返回值**：无

13. **getUIVisible** - 获取当前控件可见的值
    - **实现状态**：空操作
    - **返回值**：0

14. **setUIVisible** - 设置控件隐藏或可见
    - **实现状态**：空操作
    - **返回值**：无

15. **setUITextColor** - 修改控件字体颜色
    - **实现状态**：空操作
    - **返回值**：无

16. **setUIBackground** - 设置窗口或者控件的背景
    - **实现状态**：空操作
    - **返回值**：无

17. **setUIConfig** - 给指定窗口ui加载一个配置
    - **实现状态**：空操作
    - **返回值**：无

### 消息提示功能（2 个）

18. **toast** - 弹窗显示信息
    - **实现状态**：空操作
    - **返回值**：无

19. **hideToast** - 关闭消息显示
    - **实现状态**：空操作
    - **返回值**：无

### 界面显示功能（2 个）

20. **showUI** - 显示一个自定义的界面
    - **实现状态**：空操作
    - **返回值**：空字符串 ""

21. **showUIEx** - 显示一个自定义的界面（扩展版本）
    - **实现状态**：空操作
    - **返回值**：空字符串 ""

### HUD 显示功能（3 个）

22. **createHUD** - 创建一个HUD用于显示
    - **实现状态**：空操作
    - **返回值**：0

23. **showHUD** - 显示HUD或者刷新
    - **实现状态**：空操作
    - **返回值**：无

24. **hideHUD** - 隐藏并销毁HUD
    - **实现状态**：空操作
    - **返回值**：false

## 总结

ui 模块提供了懒人精灵的 UI 相关方法，但所有方法都是空操作：

- **空操作方法**：24 个（所有方法）
- **总计**：24 个方法

### 功能分类

| 类别 | 数量 | 说明 |
|------|------|------|
| 窗口管理 | 1 个 | closeWindow |
| WebView 功能 | 2 个 | getUIWebViewUrl, setUIWebViewUrl |
| UI 控件操作 | 15 个 | getUISelected, getUISelectText, getUIText, setUIText, setUISelect, setUICheck, getUICheck, getUIEnable, setUIEnable, getUIVisible, setUIVisible, setUITextColor, setUIBackground, setUIConfig |
| 消息提示 | 2 个 | toast, hideToast |
| 界面显示 | 2 个 | showUI, showUIEx |
| HUD 显示 | 3 个 | createHUD, showHUD, hideHUD |

## 迁移指南

### 从懒人精灵迁移到 AutoGo

**重要**：懒人精灵的 UI 系统与 AutoGo 完全不同，以下功能**不会被实现**：

1. **懒人精灵 UI 系统**：
   - 使用自定义的 UI 系统
   - 支持 .ui 文件定义界面
   - 支持窗口管理、WebView 控件、UI 控件操作等
   - 支持事件回调机制（onload, onclick, onchecked, onselected, onclose 等）

2. **AutoGo UI 系统**：
   - 使用 imgui 模块创建即时模式 GUI
   - 使用 uiacc 模块进行应用内控件交互
   - 不支持懒人精灵的 .ui 文件格式
   - 不支持懒人精灵的窗口管理系统

### 替代方案

#### 1. 自定义界面

**懒人精灵**：
```lua
function onUIEvent(handle, event, arg1, arg2, arg3)
    if event == "onclick" then
        print("按钮点击:", arg1, arg2)
    end
end

local ret = showUI("myui.ui", 600, 600, onUIEvent)
```

**AutoGo（使用 imgui）**：
```lua
-- 使用 AutoGo 的 imgui 模块创建自定义界面
imgui.Begin("我的界面")
if imgui.Button("点击我") then
    print("按钮被点击")
end
imgui.End()
```

#### 2. 应用内控件交互

**懒人精灵**：
```lua
-- 获取控件文本
local text = getUIText(handle, 0, "idEdit1")
print(text)

-- 设置控件文本
setUIText(handle, 0, "idEdit1", "新文本")
```

**AutoGo（使用 uiacc）**：
```lua
-- 使用 AutoGo 的 uiacc 模块查找和操作应用内控件
local node = uiacc.findNode("text=搜索")
if node then
    local text = node.text
    print(text)
    
    -- 设置文本（如果控件支持）
    node:setText("新文本")
end
```

#### 3. 消息提示

**懒人精灵**：
```lua
toast("这是一条消息", 0, 0, 24)
sleep(5000)
hideToast()
```

**AutoGo（使用 toast 模块）**：
```lua
-- 使用 AutoGo 的 toast 模块显示消息
toast.show("这是一条消息")
sleep(5000)
toast.hide()
```

#### 4. HUD 显示

**懒人精灵**：
```lua
local id = createHUD()
showHUD(id, "HelloWorld!", 12, "0xffff0000", "0xffffffff", 0, 100, 100, 0, 0)
sleep(5000)
hideHUD(id)
```

**AutoGo（使用 imgui 或 overlay）**：
```lua
-- 使用 AutoGo 的 imgui 模块创建悬浮窗口
imgui.SetNextWindowPos(100, 100)
imgui.Begin("HUD", false, imgui.WindowFlags_NoDecoration)
imgui.Text("HelloWorld!")
imgui.End()
```

#### 5. WebView 功能

**懒人精灵**：
```lua
-- 获取浏览器地址
local url = getUIWebViewUrl(handle, 0, "idWebView")
print(url)

-- 设置浏览器地址
setUIWebViewUrl(handle, 0, "idWebView", "http://www.example.com")
```

**AutoGo（使用 http 模块）**：
```lua
-- 使用 AutoGo 的 http 模块获取网页内容
local response = http.get("http://www.example.com")
if response then
    print(response.body)
end
```

## 注意事项

1. **不要使用这些方法**：
   - 所有 ui 模块的方法都是空操作
   - 调用这些方法不会产生任何实际效果
   - 建议使用 AutoGo 提供的替代方案

2. **使用 AutoGo 的替代模块**：
   - **自定义界面**：使用 imgui 模块
   - **应用内控件交互**：使用 uiacc 模块
   - **消息提示**：使用 toast 模块
   - **网络请求**：使用 http 模块

3. **API 兼容性**：
   - ui 模块的方法可以正常调用，不会报错
   - 方法会返回默认值以保持 API 兼容性
   - 这些方法的存在只是为了保持与懒人精灵 API 的兼容性

## 为什么不实现这些功能

### 技术原因

1. **架构差异**：
   - 懒人精灵使用自定义的 UI 系统，基于 Android 的 View 系统
   - AutoGo 使用即时模式 GUI（imgui），两者架构完全不同

2. **文件格式**：
   - 懒人精灵使用 .ui 文件定义界面（XML 格式）
   - AutoGo 使用代码直接创建界面
   - 解析 .ui 文件需要完整的解析器和渲染引擎

3. **事件系统**：
   - 懒人精灵使用事件回调机制（onload, onclick, onchecked, onselected, onclose 等）
   - AutoGo 使用即时模式，每次渲染时检查状态
   - 两者的事件处理方式完全不同

4. **WebView 控件**：
   - 懒人精灵的 WebView 控件是自定义的，支持 JavaScript 交互
   - AutoGo 不支持 WebView 控件
   - 无法通过 adb 或 AutoGo API 实现

5. **窗口管理**：
   - 懒人精灵支持多窗口、窗口层级、窗口动画等
   - AutoGo 不支持窗口管理系统
   - 实现这些功能需要重写整个窗口管理器

### 工作量考虑

1. **实现复杂度**：
   - 需要实现完整的 UI 渲染引擎
   - 需要实现 .ui 文件解析器
   - 需要实现事件系统
   - 需要实现窗口管理系统
   - 需要实现 WebView 控件

2. **维护成本**：
   - 需要持续维护和更新 UI 系统
   - 需要处理不同 Android 版本的兼容性问题
   - 需要处理不同设备的适配问题

3. **设计理念**：
   - AutoGo 的设计理念是轻量级、高性能
   - 懒人精灵的 UI 系统过于复杂，不符合 AutoGo 的设计理念
   - AutoGo 已经提供了 imgui 和 uiacc 模块，可以满足大部分需求

## 总结

ui 模块提供了懒人精灵的 UI 相关方法，但所有方法都是空操作。这是由于：

1. **架构差异**：懒人精灵和 AutoGo 使用完全不同的 UI 系统
2. **技术限制**：无法通过 AutoGo API 或 adb 命令实现这些功能
3. **工作量巨大**：实现这些功能需要重写整个 UI 系统
4. **设计理念**：不符合 AutoGo 轻量级、高性能的设计理念

**建议**：使用 AutoGo 提供的替代方案（imgui、uiacc、toast、http 等模块）来实现类似功能。
