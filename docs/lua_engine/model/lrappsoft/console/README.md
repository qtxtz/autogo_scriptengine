# Console 模块

## 概述

Console 模块提供了控制台悬浮窗的控制接口，兼容懒人精灵的 `console` 接口。

## 模块信息

- **模块名称**: `console`
- **全局对象**: `console`
- **是否可用**: 是

## API 文档

### console.show

显示控制台悬浮窗。

**函数签名**: `console.show(...)`

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
console.show()
console.setPos(100, 100)
```

### console.showTitle

显示或者隐藏控制台标题栏。

**函数签名**: `console.showTitle(...)`

**参数**:
- `show` (boolean, 可选): 是否显示标题栏，默认为 true

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
console.show()
console.showTitle(false)
console.setPos(100, 100)
```

### console.lockConsole

锁定控制台窗口。

**函数签名**: `console.lockConsole()`

**返回值**:
- `any`: 无

**示例**:
```lua
console.show()
console.lockConsole()
```

### console.unlockConsole

解除锁定控制台窗口。

**函数签名**: `console.unlockConsole()`

**返回值**:
- `any`: 无

**示例**:
```lua
console.show()
console.setPos(100, 100)
console.lockConsole()
setStopCallBack(function(error)
    console.unlockConsole()
end)
local tick = 1
while true do
    console.println(3, "tick=>"..tick)
    tick = tick + 1
    sleep(1000)
end
```

### console.dismiss

关闭控制台窗口。

**函数签名**: `console.dismiss()`

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
console.show()
console.setPos(100, 100)
for i = 1, 100 do
    console.println(3, " 日志任务开始:"..i)
    sleep(500)
end
console.dismiss()
```

### console.setPos

设置控制台窗口的位置和大小。

**函数签名**: `console.setPos(x, y, ...)`

**参数**:
- `x` (number): 窗口左上角X坐标
- `y` (number): 窗口左上角Y坐标
- `width` (number, 可选): 窗口宽度
- `height` (number, 可选): 窗口高度

**返回值**:
- `any`: 无

**示例**:
```lua
console.show()
console.setPos(100, 100)
for i = 1, 100 do
    console.println(3, " 日志任务开始:"..i)
    sleep(500)
end
console.dismiss()
```

### console.println

打印日志到控制台窗口。

**函数签名**: `console.println(level, log)`

**参数**:
- `level` (number): 级别参数
- `log` (string): 日志内容

**返回值**:
- `any`: 无

**示例**:
```lua
console.show()
console.setPos(100, 100)
for i = 1, 100 do
    console.println(3, " 日志任务开始:"..i)
    sleep(500)
end
console.dismiss()
```

### console.clearLog

清除日志。

**函数签名**: `console.clearLog()`

**返回值**:
- `any`: 无

**示例**:
```lua
console.show()
console.clearLog()
console.setPos(100, 100)
for i = 1, 100 do
    console.println(3, " 日志任务开始:"..i)
    sleep(500)
end
console.dismiss()
```

### console.setTitle

设置控制台标题。

**函数签名**: `console.setTitle(title)`

**参数**:
- `title` (string): 标题字符串

**返回值**:
- `any`: 无

**示例**:
```lua
console.show()
console.setTitle("日志窗口")
console.clearLog()
console.setPos(100, 100)
for i = 1, 100 do
    console.println(3, " 日志任务开始:"..i)
    sleep(500)
end
console.dismiss()
```

## 实现说明

本模块基于 AutoGo 的 console 包实现，提供了懒人精灵兼容的接口。当前实现为模拟实现，实际的控制台功能需要调用 AutoGo 的 console API。

## 注意事项

1. 当前实现为模拟实现，主要用于接口兼容性测试
2. 实际的控制台显示和交互需要集成 AutoGo 的 console 包
3. 日志级别参数目前仅用于标识，不影响实际显示

## 测试

测试代码位于 `test/lrappsoft/console.go`，包含以下测试用例：

- 基本 show/dismiss 测试
- setPos 测试
- println 测试
- clearLog 测试
- setTitle 测试
- showTitle 测试
- lockConsole/unlockConsole 测试

## 更新日志

### v1.0.0 (2026-03-17)
- 初始版本
- 实现 console.show
- 实现 console.showTitle
- 实现 console.lockConsole
- 实现 console.unlockConsole
- 实现 console.dismiss
- 实现 console.setPos
- 实现 console.println
- 实现 console.clearLog
- 实现 console.setTitle
