# touch 模块（触控方法）

## 概述

touch 模块提供了懒人精灵的触控相关方法，包括触摸模拟、输入法控制、按键操作和 UI 控制等功能。由于 AutoGo 运行在 Android 环境，大部分方法已经通过 AutoGo API 实现。

## 配置选项

touch 模块支持配置未实现方法的行为：

```go
type EmptyMethodConfig struct {
    ThrowException bool // 是否抛出异常（默认 false）
    ShowWarning    bool // 是否显示警告信息（默认 true）
}
```

### 使用示例

```go
// 创建模块实例并设置配置
touchModule := touch.New(&touch.EmptyMethodConfig{
    ThrowException: false, // 不抛出异常，返回默认值
    ShowWarning:    true,  // 显示警告信息
})

// 运行时修改配置
touchModule.SetConfig(&touch.EmptyMethodConfig{
    ThrowException: true, // 抛出异常
    ShowWarning:    true,
})

// 获取当前配置
config := touchModule.GetConfig()
```

### 配置说明

- **ThrowException**: 当设置为 true 时，调用未实现的方法会抛出异常；当设置为 false 时，返回默认值
- **ShowWarning**: 当设置为 true 时，调用未实现的方法会打印警告信息到控制台

### 默认行为

默认情况下，未实现的方法会：
- 打印警告信息（ShowWarning: true）
- 不抛出异常（ThrowException: false）
- 返回默认值（nil、0 或空字符串）

## 重要说明

### ✅ 可用功能

以下功能在 AutoGo 中**可用**（通过 AutoGo API 实现）：

1. **inputText** - 模拟输入文字
2. **imeLib.lock** - 锁定使用懒人输入法
3. **imeLib.unlock** - 解锁懒人输入法
4. **imeLib.setText** - 输入法模拟输入文字
5. **imeLib.deleteChar** - 输入法删除一个字符
6. **imeLib.finishInput** - 输入法模拟完成输入
7. **imeLib.keyEvent** - 输入法输入字符
8. **tap** - 点击
9. **longTap** - 长点击
10. **touchDown** - 按下手指
11. **touchUp** - 弹起手指
12. **touchMove** - 模拟滑动
13. **touchMoveEx** - 模拟滑动增强版
14. **swipe** - 划动
15. **keyPress** - 按键
16. **keyDown** - 按键按下
17. **keyUp** - 按键弹起

### ⚠️ 部分可用功能

以下方法在 AutoGo 中**部分可用**，但需要注意使用限制：

1. **setOnTouchListener** - 获取用户触摸屏幕坐标
   - **实现状态**：检查 root 权限，返回 true/false
   - **限制**：虽然会检查 root 权限，但实际的触摸事件监听功能暂不实现
   - **原因**：实现完整的触摸监听需要：
     - 找到正确的触摸输入设备文件（/dev/input/eventX）
     - 解析复杂的 getevent 输出格式（十六进制数据）
     - 在 goroutine 中安全地调用 Lua 回调函数
     - 处理多指触摸和事件同步
     - 涉及 Lua 状态的并发安全问题
   - **建议**：如果需要触摸监听功能，建议使用 AutoGo 的 uiacc 模块进行控件交互

### ❌ 不会实现的功能（懒人精灵 UI）

以下方法是懒人精灵的**自定义 UI 功能**，AutoGo **不会实现**这些功能：

**重要说明**：懒人精灵的自定义 UI 系统（包括窗口管理、WebView 控件、UI 控件操作等）是懒人精灵特有的功能，AutoGo 使用不同的架构，因此这些功能不会被实现。

#### 窗口管理功能（2 个）

2. **closeWindow** - 关闭窗口
   - **原因**：AutoGo 不支持懒人精灵的自定义 UI 窗口系统
   - **说明**：懒人精灵支持自定义 UI 窗口，AutoGo 使用 imgui 模块替代
   - **建议**：使用 AutoGo 的 imgui 模块创建自定义界面

3. **getUIWebViewUrl** - 获取当前浏览器的地址
4. **setUIWebViewUrl** - 设置浏览器地址
   - **原因**：AutoGo 不支持懒人精灵的 WebView 控件
   - **说明**：这些方法是针对懒人精灵自定义 WebView 控件的，无法通过 adb 或 AutoGo API 实现
   - **建议**：使用 AutoGo 的 http 模块获取网页内容

#### UI 控件操作功能（15 个）

5. **getUISelected** - 获取单选框或者下拉框当前选中项
6. **getUISelectText** - 获取单选框或者下拉框当前选中项的文本内容
7. **getUIText** - 获取控件显示的文本内容
8. **setUIText** - 设置控件文本
9. **setUISelect** - 设置单选框或者下拉框选项被选中
10. **setUICheck** - 设置多选框被选中或者反选
11. **getUICheck** - 获取单选框状态
12. **getUIEnable** - 获取控件是否可用
13. **setUIEnable** - 设置控件是否可用
14. **getUIVisible** - 获取当前控件可见的值
15. **setUIVisible** - 设置控件隐藏或可见
16. **setUITextColor** - 修改控件字体颜色
17. **setUIBackground** - 设置窗口或者控件的背景
18. **setUIConfig** - 给指定窗口ui加载一个配置
   - **原因**：AutoGo 不支持懒人精灵的自定义 UI 控件系统
   - **说明**：这些方法是针对懒人精灵自定义 UI 控件的，AutoGo 使用不同的 UI 系统
   - **建议**：使用 AutoGo 的 uiacc 模块进行应用内控件交互，使用 imgui 模块创建自定义界面

## 总结

touch 模块提供了懒人精灵的触控相关方法，通过 AutoGo API 实现了核心功能：

- **完全可用方法**：17 个（通过 AutoGo API 实现）
- **部分可用方法**：1 个（setOnTouchListener，检查 root 权限但实际监听功能暂不实现）
- **不会实现的方法**：17 个（懒人精灵自定义 UI 功能）
- **总计**：35 个方法

### 功能分类

| 类别 | 数量 | 说明 |
|------|------|------|
| 触摸操作 | 7 个 | tap, longTap, touchDown, touchUp, touchMove, touchMoveEx, swipe |
| 输入法操作 | 7 个 | inputText, imeLib.lock, imeLib.unlock, imeLib.setText, imeLib.deleteChar, imeLib.finishInput, imeLib.keyEvent |
| 按键操作 | 3 个 | keyPress, keyDown, keyUp |
| 触摸监听 | 1 个 | setOnTouchListener（部分实现） |
| 懒人精灵 UI | 17 个 | 窗口管理、WebView、UI 控件操作（不会实现） |

## 使用示例

### 示例 1：基本触摸操作

```lua
-- 点击屏幕上的点
tap(540, 960)

-- 长按屏幕上的点
longTap(540, 960)

-- 划动操作
swipe(100, 1000, 100, 100, 500)
```

### 示例 2：多指触摸操作

```lua
-- 按下第一根手指
touchDown(0, 300, 500)

-- 按下第二根手指
touchDown(1, 600, 500)

-- 移动第一根手指
for i = 1, 100 do
    touchMove(0, 300 + i, 500)
    sleep(10)
end

-- 弹起第一根手指
touchUp(0)

-- 弹起第二根手指
touchUp(1)
```

### 示例 3：输入文字

```lua
-- 直接输入文字
inputText("Hello, AutoGo!")

-- 使用输入法输入文字
imeLib.lock()
imeLib.setText("测试文本")
imeLib.finishInput()
imeLib.unlock()

-- 删除一个字符
imeLib.deleteChar()
```

### 示例 4：按键操作

```lua
-- 按下 Home 键
keyPress("home")

-- 按下返回键
keyPress("back")

-- 按下最近任务键
keyPress("recent")

-- 使用数字按键码
keyPress(3)  -- Home 键的按键码
```

### 示例 5：综合操作流程

```lua
-- 打开应用
runApp("com.tencent.mm")

-- 等待应用启动
sleep(2000)

-- 点击搜索框
tap(540, 150)

-- 输入搜索内容
inputText("AutoGo")

-- 点击搜索按钮
tap(800, 150)

-- 等待搜索结果
sleep(1000)

-- 滑动查看更多结果
swipe(540, 800, 540, 400, 500)

-- 返回上一页
keyPress("back")
```

### 示例 6：模拟滑动操作

```lua
-- 使用 touchMove 模拟滑动
print("开始从下往上滑动")
touchDown(0, 297, 327)
for i = 1, 300 do
    touchMove(0, 297, 327 - i)
    sleep(1)
end
touchUp(0)
print("滑动结束")

-- 使用 touchMoveEx 模拟滑动（带时间控制）
touchDown(0, 100, 100)
sleep(50)
touchMoveEx(0, 300, 300, 800)
touchUp(0)
```

### 示例 7：输入法高级操作

```lua
-- 锁定输入法
imeLib.lock()

-- 设置停止回调
setStopCallBack(function(error)
    imeLib.unlock()
end)

-- 输入文字
imeLib.setText("Hello, World!")

-- 模拟按键事件
imeLib.keyEvent(0, 7)  -- 按下字符 "0"
imeLib.keyEvent(1, 7)  -- 弹起字符 "0"

-- 完成输入
imeLib.finishInput()

-- 解锁输入法
imeLib.unlock()
```

## 按键标识符说明

支持以下按键标识符：

| 标识符 | 说明 | 按键码 |
|--------|------|--------|
| home | Home 键 | 3 |
| back | 返回键 | 4 |
| recent | 最近任务键 | 187 |
| power | 电源键 | 26 |
| menu | 菜单键 | 82 |
| volume_up | 音量加键 | 24 |
| volume_down | 音量减键 | 25 |
| camera | 相机键 | 27 |
| enter | 回车键 | 66 |
| delete | 删除键 | 67 |

也可以直接使用数字按键码，参考 Android KeyEvent.KEYCODE_* 常量。

## 迁移指南

### 从懒人精灵迁移到 AutoGo

1. **保留可用的功能**：
   - 所有触摸操作（tap, longTap, touchDown, touchUp, touchMove, touchMoveEx, swipe）
   - 所有输入法操作（inputText, imeLib.*）
   - 所有按键操作（keyPress, keyDown, keyUp）

2. **部分使用的功能**：
   - setOnTouchListener：会检查 root 权限并返回 true/false，但实际的触摸事件监听功能暂不实现

3. **移除不会实现的功能**：
   - 所有懒人精灵自定义 UI 功能（17 个）
   - 包括窗口管理、WebView 控件、UI 控件操作等
   - 这些方法可以正常调用但不执行任何功能

4. **使用替代方案**：
   - **UI 界面**：使用 AutoGo 的 imgui 模块创建自定义界面
   - **应用内控件交互**：使用 AutoGo 的 uiacc 模块进行控件查找和操作
   - **浏览器功能**：使用 AutoGo 的 http 模块获取网页内容
   - **触摸监听**：setOnTouchListener 只检查 root 权限，如需完整功能建议使用 uiacc 模块

### 懒人精灵 UI 功能说明

**重要**：懒人精灵的自定义 UI 系统（包括 showUI、窗口管理、WebView 控件、UI 控件操作等）是懒人精灵特有的功能。AutoGo 使用不同的架构和 UI 系统：

- **懒人精灵**：使用自定义的 UI 系统，支持 .ui 文件、WebView 控件、窗口管理等
- **AutoGo**：使用 imgui 模块创建即时模式 GUI，使用 uiacc 模块进行应用内控件交互

因此，以下 17 个懒人精灵 UI 相关的功能**不会被实现**：

1. closeWindow - 关闭窗口
2. getUIWebViewUrl - 获取浏览器地址
3. setUIWebViewUrl - 设置浏览器地址
4. getUISelected - 获取下拉框选中项
5. getUISelectText - 获取下拉框选中项文本
6. getUIText - 获取控件文本
7. setUIText - 设置控件文本
8. setUISelect - 设置下拉框选中项
9. setUICheck - 设置复选框状态
10. getUICheck - 获取复选框状态
11. getUIEnable - 获取控件可用状态
12. setUIEnable - 设置控件可用状态
13. getUIVisible - 获取控件可见状态
14. setUIVisible - 设置控件可见状态
15. setUITextColor - 设置控件字体颜色
16. setUIBackground - 设置控件背景
17. setUIConfig - 加载 UI 配置

这些方法在 touch 模块中保持为空操作，以保持 API 兼容性，但不会执行任何实际功能。

## 注意事项

1. **触摸操作**：
   - touchDown, touchMove, touchUp 支持多指触摸（0-4）
   - swipe 操作在所有模式下都可用
   - touchDown/touchUp/touchMove 只在 root 或激活模式下可用

2. **按键操作**：
   - root 或激活模式下所有按键都有效
   - 无障碍模式下只支持 home、back、recent、power 四个按键

3. **输入法操作**：
   - 需要安装懒人精灵输入法或 AutoGo 输入法
   - 使用前需要调用 imeLib.lock() 锁定输入法
   - 使用后建议调用 imeLib.unlock() 解锁输入法

4. **UI 控制方法**：
   - 所有 UI 控制方法都是空操作
   - 可以正常调用但不执行任何功能
   - 建议使用 AutoGo 的 imgui 模块替代

## API 映射

### AutoGo API 映射

touch 模块通过以下 AutoGo API 实现功能：

| touch 方法 | AutoGo API | 说明 |
|------------|------------|------|
| tap | motion.Click | 单击操作 |
| longTap | motion.LongClick | 长按操作 |
| touchDown | motion.TouchDown | 触摸按下 |
| touchUp | motion.TouchUp | 触摸抬起 |
| touchMove | motion.TouchMove | 触摸移动 |
| swipe | motion.Swipe | 滑动操作 |
| keyPress/keyDown/keyUp | motion.KeyAction | 按键操作 |
| inputText | ime.InputText | 输入文字 |
| imeLib.deleteChar | motion.KeyAction(67) | 删除字符 |
| imeLib.finishInput | motion.KeyAction(66) | 完成输入 |
