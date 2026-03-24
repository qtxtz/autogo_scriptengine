# VirtualScreen 模块

## 概述

VirtualScreen 模块提供了虚拟屏幕的创建和管理功能，兼容懒人精灵的 `virtualDisplay` 接口。该模块基于 AutoGo 的 vdisplay 和 motion 包实现。

## 模块信息

- **模块名称**: `virtualscreen`
- **全局对象**: `virtualDisplay`
- **是否可用**: 是

## API 文档

### virtualDisplay.createVirtualDisplay

创建虚拟屏。

**函数签名**: `virtualDisplay.createVirtualDisplay(width, height, dpi)`

**参数**:
- `width` (number): 屏幕宽度
- `height` (number): 屏幕高度
- `dpi` (number, 可选): 屏幕的DPI，默认320

**返回值**:
- `number`: 创建成功返回 displayId，失败返回 -1

**示例**:
```lua
displayId = virtualDisplay.createVirtualDisplay(720, 1280, 320)
if displayId ~= -1 then
    print("创建虚拟屏幕成功:" , displayId)
else
    print("创建虚拟屏幕失败")
end
```

### virtualDisplay.showVirtualDisplay

显示虚拟屏。

**函数签名**: `virtualDisplay.showVirtualDisplay(displayId, config)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `config` (table): 配置表（可选），包含以下可选字段：
  - `x`: 窗口X坐标（整数，默认0）
  - `y`: 窗口Y坐标（整数，默认0）
  - `width`: 窗口宽度（整数，默认600）
  - `height`: 窗口高度（整数，默认600）
  - `hastitle`: 是否显示标题栏（布尔值，默认true）
  - `title`: 窗口标题文本（字符串）
  - `titlecolor`: 标题文字颜色（ARGB格式，默认0xFFFFFFFF白色）
  - `titlebgcolor`: 标题栏背景色（ARGB格式，默认0xFF87CEFA天蓝色）
  - `hasclose`: 是否显示关闭按钮（布尔值，默认true）
  - `closecolor`: 关闭按钮颜色（ARGB格式，默认0xFFFFFFFF白色）
  - `hasresize`: 是否显示设置窗口大小按钮（布尔值，默认true）
  - `resizecolor`: 设置窗口大小按钮颜色（ARGB格式，默认0xFFFFFFFF白色）
  - `hastoggle`: 是否显示窗口收缩按钮（布尔值，默认true）
  - `togglecolor`: 设置窗口收缩按钮颜色（ARGB格式，默认0xFFFFFFFF白色）
  - `titlesize`: 设置标题栏字体大小

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
config = {
    x = 100,
    y = 100,
    width = 480,
    height = 800,
    hastitle = true,
    title = "虚拟屏测试",
    titlesize = 13,
    titlecolor = 0xffffffff,
    titlebgcolor = 0xdd294a7a
}
local result = virtualDisplay.showVirtualDisplay(displayId, config)
```

### virtualDisplay.closeVirtualDisplay

关闭虚拟屏。

**函数签名**: `virtualDisplay.closeVirtualDisplay(displayId)`

**参数**:
- `displayId` (number): 创建返回的 displayId

**返回值**:
- 无

**示例**:
```lua
virtualDisplay.closeVirtualDisplay(displayId)
```

### virtualDisplay.snapToBitmap

虚拟屏截图。

**函数签名**: `virtualDisplay.snapToBitmap(displayId, x, y, w, h)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `x` (number): 窗口X坐标（整数）
- `y` (number): 窗口Y坐标（整数）
- `w` (number): 窗口宽度（整数）
- `h` (number): 窗口高度（整数）

**返回值**:
- `object`: 返回一个 Java 的 Bitmap 对象

**注意**: 如果 x,y,w,h 都写 0 表示全屏截图

**示例**:
```lua
local bitmap = virtualDisplay.snapToBitmap(displayId, 0, 0, 0, 0)
```

### virtualDisplay.switchToVirtualDisplay

切换虚拟屏幕。

**函数签名**: `virtualDisplay.switchToVirtualDisplay(displayId)`

**参数**:
- `displayId` (number): 创建返回的 displayId

**返回值**:
- `boolean`: true表示成功，false表示失败

**说明**: 当虚拟屏幕创建好后，如果想在这个虚拟屏里面实现图色功能，就需要用这个方法指定此虚拟屏幕的 id。如果想切换回系统主屏只需要将 id 指定为 0 即可。

**示例**:
```lua
virtualDisplay.switchToVirtualDisplay(displayId)
```

### virtualDisplay.runAppWithVirtualDisplay

在虚拟屏里面运行指定应用。

**函数签名**: `virtualDisplay.runAppWithVirtualDisplay(displayId, pkg, forceStopApp)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `pkg` (string): 已安装的 app 包名
- `forceStopApp` (boolean): 每次启动是否强制先关闭之前已运行的实例

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
local result = virtualDisplay.runAppWithVirtualDisplay(displayId, "com.example.app", true)
```

### virtualDisplay.tap

虚拟屏里面点击。

**函数签名**: `virtualDisplay.tap(displayId, x, y)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `x` (number): 要点击的 X 坐标
- `y` (number): 要点击的 Y 坐标

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
virtualDisplay.tap(displayId, 100, 200)
```

### virtualDisplay.touchDown

虚拟屏里面模拟手指按下。

**函数签名**: `virtualDisplay.touchDown(displayId, id, x, y)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `id` (number): 手指 id（0-5）
- `x` (number): 要点击的 X 坐标
- `y` (number): 要点击的 Y 坐标

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
virtualDisplay.touchDown(displayId, 1, 100, 200)
```

### virtualDisplay.touchUp

虚拟屏里面模拟手指弹起。

**函数签名**: `virtualDisplay.touchUp(displayId, id)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `id` (number): 手指 id（0-5）

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
virtualDisplay.touchUp(displayId, 1)
```

### virtualDisplay.touchMove

虚拟屏里面模拟手指滑动。

**函数签名**: `virtualDisplay.touchMove(displayId, id, x, y)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `id` (number): 手指 id（0-5）
- `x` (number): 要移动到的 X 坐标
- `y` (number): 要移动到的 Y 坐标

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
virtualDisplay.touchMove(displayId, 1, 150, 250)
```

### virtualDisplay.touchMoveEx

虚拟屏里面模拟手指增强滑动。

**函数签名**: `virtualDisplay.touchMoveEx(displayId, id, x, y, time)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `id` (number): 模拟手指的索引号 0-5 之间
- `x` (number): 整数型，当前屏幕横坐标
- `y` (number): 整数型，当前屏幕纵坐标
- `time` (number): 整数型，滑动到 x,y 坐标点所需要的时间单位是毫秒

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
virtualDisplay.touchMoveEx(displayId, 1, 100, 300, 1000)
```

### virtualDisplay.swipe

虚拟屏里面模拟滑动。

**函数签名**: `virtualDisplay.swipe(displayId, x1, y1, x2, y2, time)`

**参数**:
- `displayId` (number): 创建返回的 displayId
- `x1` (number): 整数型，划动的起点 x 坐标
- `y1` (number): 整数型，划动的起点 y 坐标
- `x2` (number): 整数型，划动的终点 x 坐标
- `y2` (number): 整数型，划动的终点 y 坐标
- `time` (number): 整数型，滑动到 x,y 坐标点所需要的时间单位是毫秒

**返回值**:
- `boolean`: true表示成功，false表示失败

**示例**:
```lua
virtualDisplay.swipe(displayId, 100, 1000, 100, 300, 1000)
```

## 实现说明

本模块基于 AutoGo 的 vdisplay 和 motion 包实现，提供了懒人精灵兼容的接口。当前实现为模拟实现，实际的控制台功能需要调用 AutoGo 的 vdisplay 和 motion API。

## 注意事项

1. 虚拟屏功能必须 root 或者激活模式才能使用
2. 虚拟屏需要 Android 10 及以上系统（API 29+）
3. 当前实现为模拟实现，主要用于接口兼容性测试
4. 实际的虚拟屏显示和交互需要集成 AutoGo 的 vdisplay 和 motion 包

## 测试

测试代码位于 `test/lrappsoft/virtualscreen.go`，包含以下测试用例：

- 创建虚拟屏测试
- 显示虚拟屏测试
- 关闭虚拟屏测试
- 虚拟屏截图测试
- 切换虚拟屏幕测试
- 在虚拟屏运行应用测试
- 虚拟屏点击测试
- 虚拟屏触摸事件测试
- 虚拟屏滑动测试

## 更新日志

### v1.0.0 (2026-03-17)
- 初始版本
- 实现 virtualDisplay.createVirtualDisplay
- 实现 virtualDisplay.showVirtualDisplay
- 实现 virtualDisplay.closeVirtualDisplay
- 实现 virtualDisplay.snapToBitmap
- 实现 virtualDisplay.switchToVirtualDisplay
- 实现 virtualDisplay.runAppWithVirtualDisplay
- 实现 virtualDisplay.tap
- 实现 virtualDisplay.touchDown
- 实现 virtualDisplay.touchUp
- 实现 virtualDisplay.touchMove
- 实现 virtualDisplay.touchMoveEx
- 实现 virtualDisplay.swipe
