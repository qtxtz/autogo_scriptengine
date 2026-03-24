# dynamicui - 动态 UI 模块

## 概述

dynamicui 模块提供了懒人精灵动态 UI 方法的兼容层。**注意：此模块中的所有方法都是空实现，AutoGo 不支持懒人精灵的动态 UI 功能。**

## 重要说明

### ⚠️ 空实现警告

dynamicui 模块中的所有方法都是**空实现**，原因如下：

1. **架构差异**：AutoGo 和懒人精灵的 UI 系统架构完全不同
2. **不兼容性**：懒人精灵的动态 UI 系统无法在 AutoGo 中实现
3. **替代方案**：建议使用 AutoGo 的 imgui 模块或其他 UI 解决方案

### 配置选项

dynamicui 模块提供了配置选项来控制空方法的行为：

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
// 创建 dynamicui 模块（使用默认配置）
dynamicuiModule := dynamicui.New(nil)

// 创建 dynamicui 模块（自定义配置）
config := &dynamicui.EmptyMethodConfig{
    ThrowException: false, // 不抛出异常
    ShowWarning:    true,  // 显示警告信息
}
dynamicuiModule := dynamicui.New(config)

// 运行时修改配置
dynamicuiModule.SetConfig(&dynamicui.EmptyMethodConfig{
    ThrowException: true, // 抛出异常
    ShowWarning:    true, // 显示警告信息
})
```

## 功能说明

### 所有方法状态

| 方法名 | 状态 | 说明 |
|--------|------|------|
| ui.newLayout | ❌ 空实现 | 创建一个新的布局 |
| ui.show | ❌ 空实现 | 显示一个布局 |
| ui.dismiss | ❌ 空实现 | 关闭一个布局 |
| ui.newRow | ❌ 空实现 | 布局换行排列 |
| ui.addButton | ❌ 空实现 | 创建一个按钮 |
| ui.addTextView | ❌ 空实现 | 创建文字框控件 |
| ui.addEditText | ❌ 空实现 | 创建输入框控件 |
| ui.addCheckBox | ❌ 空实现 | 创建多选框控件 |
| ui.addRadioGroup | ❌ 空实现 | 创建单选框控件 |
| ui.addSpinner | ❌ 空实现 | 创建下拉框控件 |
| ui.addImageView | ❌ 空实现 | 创建图像控件 |
| ui.addLine | ❌ 空实现 | 创建线控件 |
| ui.addWebView | ❌ 空实现 | 创建一个浏览器控件 |
| ui.callJs | ❌ 空实现 | 调用webview打开的网页中的js函数 |
| ui.addTabView | ❌ 空实现 | 创建标签页控件 |
| ui.addTab | ❌ 空实现 | 创建子标签页控件 |
| ui.setLine | ❌ 空实现 | 重设线控件 |
| ui.setButton | ❌ 空实现 | 重设按钮控件 |
| ui.setEditText | ❌ 空实现 | 重设输入框控件 |
| ui.setEditHintText | ❌ 空实现 | 设置输入框默认提示字符串 |
| ui.setTextView | ❌ 空实现 | 重设文本框控件 |
| ui.setCheckBox | ❌ 空实现 | 重设多选框控件 |
| ui.setRadioGroup | ❌ 空实现 | 重设单选框控件 |
| ui.setSpinner | ❌ 空实现 | 重设下拉框控件 |
| ui.setWebView | ❌ 空实现 | 重设浏览器控件 |
| ui.setImageView | ❌ 空实现 | 重设图像控件 |
| ui.setImageViewEx | ❌ 空实现 | 重设图像控件（扩展） |
| ui.setText | ❌ 空实现 | 控件设置文字 |
| ui.setTitleText | ❌ 空实现 | 设置布局标题 |
| ui.setTextSize | ❌ 空实现 | 设置文字大小 |
| ui.setEnable | ❌ 空实现 | 设置控件可用状态 |
| ui.setVisiblity | ❌ 空实现 | 设置控件显示状态 |
| ui.setRowVisibleByGid | ❌ 空实现 | 批量设置行控件显示状态通过gid |
| ui.setBackground | ❌ 空实现 | 设置控件背景颜色 |
| ui.setTitleBackground | ❌ 空实现 | 设置标题栏背景颜色 |
| ui.setTextColor | ❌ 空实现 | 设置文字颜色 |
| ui.setInputType | ❌ 空实现 | 设置输入类型 |
| ui.getText | ❌ 空实现 | 获取文字 |
| ui.getEnable | ❌ 空实现 | 获取可用状态 |
| ui.getVisible | ❌ 空实现 | 获取显示状态 |
| ui.getTextColor | ❌ 空实现 | 获取文字颜色 |
| ui.setFullScreen | ❌ 空实现 | 设置控件宽度全屏 |
| ui.setPadding | ❌ 空实现 | 设置控件内边距 |
| ui.setGravity | ❌ 空实现 | 设置控件对齐方式 |
| ui.setOnClick | ❌ 空实现 | 设置控件单击事件 |
| ui.setOnBackEvent | ❌ 空实现 | 设置窗口监听返回按键消息 |
| ui.setOnClose | ❌ 空实现 | 设置窗口关闭事件 |
| ui.setOnChange | ❌ 空实现 | 设置控件改变事件 |
| ui.getValue | ❌ 空实现 | 获取控件值 |
| ui.getData | ❌ 空实现 | 获取当前界面所有控件的值 |
| ui.loadProfile | ❌ 空实现 | 读取设置 |
| ui.saveProfile | ❌ 空实现 | 保存配置 |
| ui.beginUiQueue | ❌ 空实现 | 创建一个快速ui命令队列 |
| ui.endUiQueue | ❌ 空实现 | 执行这个ui命令队列 |
| ui.addTableView | ❌ 空实现 | 添加表格控件 |
| ui.setTableViewAttrib | ❌ 空实现 | 设置表格属性 |
| ui.getTableViewAllData | ❌ 空实现 | 获取表格全部数据 |
| ui.getTableViewRowData | ❌ 空实现 | 获取表格指定行数据 |
| ui.addTableViewRow | ❌ 空实现 | 添加表格行 |
| ui.removeTableViewRow | ❌ 空实现 | 删除指定行 |
| ui.getTableViewRowCnt | ❌ 空实现 | 获取表格行数 |
| ui.getTableViewSelectIndex | ❌ 空实现 | 获取表格选中行索引 |

## 使用示例

### 示例 1：基本使用（默认配置）

```lua
-- 创建布局
ui.newLayout("layout1")

-- 添加控件
ui.addButton("layout1", "btn", "点击我")
ui.addTextView("layout1", "tv", "这是一个文本框")

-- 显示布局
ui.show("layout1")

-- 输出警告信息：
-- [警告] dynamicui.newLayout 方法未实现，此功能在 AutoGo 中不支持
-- [警告] dynamicui.addButton 方法未实现，此功能在 AutoGo 中不支持
-- [警告] dynamicui.addTextView 方法未实现，此功能在 AutoGo 中不支持
-- [警告] dynamicui.show 方法未实现，此功能在 AutoGo 中不支持
```

### 示例 2：配置为抛出异常

```go
// 在 Go 代码中配置
config := &dynamicui.EmptyMethodConfig{
    ThrowException: true, // 抛出异常
    ShowWarning:    true, // 显示警告信息
}
dynamicuiModule := dynamicui.New(config)
dynamicuiModule.Inject(state)
```

```lua
-- 在 Lua 代码中调用
ui.newLayout("layout1")

-- 输出：
-- [警告] dynamicui.newLayout 方法未实现，此功能在 AutoGo 中不支持
-- 抛出异常：dynamicui.newLayout 方法未实现，此功能在 AutoGo 中不支持
```

### 示例 3：配置为不显示警告

```go
// 在 Go 代码中配置
config := &dynamicui.EmptyMethodConfig{
    ThrowException: false, // 不抛出异常
    ShowWarning:    false, // 不显示警告信息
}
dynamicuiModule := dynamicui.New(config)
dynamicuiModule.Inject(state)
```

```lua
-- 在 Lua 代码中调用
ui.newLayout("layout1")
ui.addButton("layout1", "btn", "点击我")
ui.show("layout1")

-- 无任何输出，静默执行
```

## 迁移指南

### 从懒人精灵迁移到 AutoGo

由于 AutoGo 不支持懒人精灵的动态 UI 系统，建议使用以下替代方案：

#### 1. 使用 AutoGo 的 imgui 模块

AutoGo 提供了 imgui 模块，可以创建即时模式 GUI 界面。

```lua
-- 使用 imgui 创建界面
local imgui = require("imgui")

-- 创建窗口
if imgui.Begin("我的窗口", true) then
    -- 添加按钮
    if imgui.Button("点击我") then
        print("按钮被点击了")
    end

    -- 添加文本
    imgui.Text("这是一个文本框")

    -- 添加输入框
    local text, changed = imgui.InputText("输入框", "默认值")
    if changed then
        print("输入内容: " .. text)
    end

    imgui.End()
end
```

#### 2. 使用 AutoGo 的 hud 模块

AutoGo 的 hud 模块可以创建简单的悬浮窗口。

```lua
-- 使用 hud 显示信息
hud.create("my_hud", 100, 100, 200, 100)
hud.setText("my_hud", "这是一个悬浮窗口")
hud.show("my_hud")

-- 隐藏悬浮窗口
hud.hide("my_hud")
```

#### 3. 使用 Android 原生 UI

如果需要复杂的 UI 界面，可以直接使用 Android 原生 UI。

```lua
import('android.widget.*')
import('android.view.*')
import('android.graphics.*')

-- 创建 Android 原生布局
local layout = LinearLayout(activity)
layout.setOrientation(LinearLayout.VERTICAL)

-- 添加按钮
local button = Button(activity)
button.setText("点击我")
layout.addView(button)

-- 设置为内容视图
activity.setContentView(layout)
```

## 底层实现

### 空方法处理逻辑

```go
// handleEmptyMethod 处理空方法的通用逻辑
func (m *DynamicUIModule) handleEmptyMethod(methodName string, L *lua.LState) int {
    // 打印警告信息
    if m.config.ShowWarning {
        fmt.Printf("[警告] dynamicui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
    }

    // 根据配置决定是否抛出异常
    if m.config.ThrowException {
        L.RaiseError(fmt.Sprintf("dynamicui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
        return 0
    }

    // 默认返回 true（兼容懒人精灵的返回值）
    L.Push(lua.LBool(true))
    return 1
}
```

### 配置管理

```go
// SetConfig 设置空方法配置
func (m *DynamicUIModule) SetConfig(config *EmptyMethodConfig) {
    m.config = config
}

// GetConfig 获取空方法配置
func (m *DynamicUIModule) GetConfig() *EmptyMethodConfig {
    return m.config
}
```

## 注意事项

1. **所有方法都是空实现**：dynamicui 模块中的所有方法都不会产生任何实际效果
2. **警告信息**：默认情况下，调用任何方法都会打印警告信息
3. **配置灵活性**：可以通过配置开关控制是否抛出异常或显示警告
4. **兼容性**：空方法默认返回 `true`，以保持与懒人精灵 API 的兼容性
5. **替代方案**：建议使用 AutoGo 的 imgui、hud 或 Android 原生 UI

## 总结

dynamicui 模块提供了懒人精灵动态 UI 方法的兼容层，但所有方法都是空实现。通过配置选项，可以控制空方法的行为（抛出异常或空执行），并且所有空方法都会打印警告信息（除非明确禁用）。

建议使用 AutoGo 的 imgui、hud 或 Android 原生 UI 来替代懒人精灵的动态 UI 功能。
