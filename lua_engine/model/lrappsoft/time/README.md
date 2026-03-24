# time 模块（懒人精灵兼容）

time 模块提供了时间相关的工具函数，兼容懒人精灵的时间 API。

## ⚠️ 重要限制说明

在使用本模块前，请务必了解以下与懒人精灵原版 API 的差异：

### 1. getNetWorkTime 的实现差异
- **懒人精灵**：从网络服务器获取精确的网络时间
- **AutoGo**：在网络请求失败时，会**回退到本地时间**
- **影响**：在网络不可用时，返回的是本地时间而非网络时间

### 2. 网络时间精度
- **懒人精灵**：可能使用专门的时间服务器协议（如 NTP）
- **AutoGo**：使用 HTTP 请求获取时间，精度较低
- **影响**：网络时间可能存在数秒的误差

### 3. tickCount 的精度
- **懒人精灵**：使用高精度计时器
- **AutoGo**：使用 Go 的 time.Since()，精度为毫秒级
- **影响**：在高精度计时场景下可能不够精确

## 模块加载

```lua
local time = require("time")
```

## 全局方法

除了通过 `require("time")` 加载模块外，以下三个方法也可以直接作为全局方法调用，无需 require：

- `systemTime()` - 返回系统当前时间戳（毫秒）
- `getNetWorkTime()` - 获取网络时间字符串
- `tickCount()` - 返回脚本自启动以来的运行时长（毫秒）

### 使用全局方法

```lua
-- 直接调用，不需要 require
local timestamp = systemTime()
local networkTime = getNetWorkTime()
local duration = tickCount()
```

### 使用模块方法

```lua
-- 通过模块调用
local time = require("time")
local timestamp = time.systemTime()
local networkTime = time.getNetWorkTime()
local duration = time.tickCount()
```

两种方式的结果完全一致，可以根据需要选择使用。

## API 参考

### 1. time.systemTime - 系统时间戳（毫秒）

**函数签名**：
```lua
timestamp = time.systemTime()
```

**参数**：
- 无

**返回值**：
- `timestamp` (number)：系统当前时间戳（毫秒）

**示例**：
```lua
local time = require("time")

-- 获取当前时间戳（毫秒）
local tm = time.systemTime()
print("当前时间戳（毫秒）: " .. tm)
-- 输出示例: 当前时间戳（毫秒）: 1672531200000
```

**说明**：
- 返回从 1970-01-01 00:00:00 UTC 到现在的毫秒数
- 精度为毫秒级
- 不依赖网络，直接获取系统时间

---

### 2. time.getNetWorkTime - 获取网络时间字符串

**函数签名**：
```lua
timeString = time.getNetWorkTime()
```

**参数**：
- 无

**返回值**：
- `timeString` (string)：网络时间字符串，格式为 `年-月-日_时-分-秒`

**示例**：
```lua
local time = require("time")

-- 获取网络时间
local tm = time.getNetWorkTime()
print("网络时间: " .. tm)
-- 输出示例: 网络时间: 2023-01-01_12-00-00
```

**⚠️ 重要说明**：
- 在网络不可用时，会回退到本地时间
- 时间格式固定为 `YYYY-MM-DD_HH-mm-ss`
- 使用的是 HTTP 请求，精度可能较低

---

### 3. time.tickCount - 脚本运行时间（毫秒）

**函数签名**：
```lua
duration = time.tickCount()
```

**参数**：
- 无

**返回值**：
- `duration` (number)：脚本自启动以来的运行时长（毫秒）

**示例**：
```lua
local time = require("time")

-- 获取脚本运行时间
local t = time.tickCount()
print("脚本已运行: " .. t .. " 毫秒")

-- 延迟 1 秒后再次测量
-- mSleep(1000)  -- 假设有延迟函数
local t2 = time.tickCount()
print("脚本已运行: " .. t2 .. " 毫秒")
```

**说明**：
- 从脚本启动时开始计时
- 精度为毫秒级
- 不会因为系统时间调整而受影响

---

## 完整示例

```lua
local time = require("time")

-- 示例 1: 获取系统时间戳
print("=== 示例 1: 获取系统时间戳 ===")
local timestamp = time.systemTime()
print("当前时间戳（毫秒）: " .. timestamp)

-- 示例 2: 获取网络时间
print("\n=== 示例 2: 获取网络时间 ===")
local networkTime = time.getNetWorkTime()
print("网络时间: " .. networkTime)

-- 示例 3: 测量脚本运行时间
print("\n=== 示例 3: 测量脚本运行时间 ===")
local startTime = time.tickCount()
print("脚本已运行: " .. startTime .. " 毫秒")

-- 模拟一些工作
local sum = 0
for i = 1, 1000000 do
    sum = sum + i
end

local endTime = time.tickCount()
local duration = endTime - startTime
print("计算完成，耗时: " .. duration .. " 毫秒")

-- 示例 4: 结合 Lua 标准库使用
print("\n=== 示例 4: 结合 Lua 标准库使用 ===")
local now = os.date("%Y-%m-%d %H:%M:%S")
print("当前时间（os.date）: " .. now)

local secs = os.time()
print("当前时间戳（秒）: " .. secs)

-- 示例 5: 时间转换
print("\n=== 示例 5: 时间转换 ===")
-- 将毫秒时间戳转换为秒时间戳
local timestampMs = time.systemTime()
local timestampSec = math.floor(timestampMs / 1000)
print("毫秒时间戳: " .. timestampMs)
print("秒时间戳: " .. timestampSec)

-- 将秒时间戳转换为可读时间
local readableTime = os.date("%Y-%m-%d %H:%M:%S", timestampSec)
print("可读时间: " .. readableTime)
```

## 与 Lua 标准库的配合使用

time 模块与 Lua 标准库的 `os` 模块可以很好地配合使用：

### os.date - 日期格式化

```lua
local time = require("time")

-- 获取当前时间戳（秒）
local secs = os.time()

-- 格式化为可读时间
local now = os.date("%Y-%m-%d %H:%M:%S")
print("当前时间: " .. now)

-- 获取详细时间信息
local timeTable = os.date("*t", secs)
print("年份: " .. timeTable.year)
print("月份: " .. timeTable.month)
print("日期: " .. timeTable.day)
print("小时: " .. timeTable.hour)
print("分钟: " .. timeTable.min)
print("秒: " .. timeTable.sec)
print("星期: " .. timeTable.wday)
```

### os.time - 获取时间戳

```lua
local time = require("time")

-- 获取当前时间戳（秒）
local timestamp = os.time()
print("当前时间戳（秒）: " .. timestamp)

-- 获取指定时间的时间戳
local ta = {
    year = 2023,
    month = 1,
    day = 1,
    hour = 12,
    min = 0,
    sec = 0
}
local specifiedTimestamp = os.time(ta)
print("指定时间的时间戳: " .. specifiedTimestamp)

-- 获取 7 天前的时间
local now = os.date("*t")
local sevenDaysAgo = {
    year = now.year,
    month = now.month,
    day = now.day - 7,
    hour = now.hour,
    min = now.min,
    sec = now.sec
}
local sevenDaysAgoTimestamp = os.time(sevenDaysAgo)
local sevenDaysAgoTime = os.date("%Y-%m-%d %H:%M:%S", sevenDaysAgoTimestamp)
print("7 天前的时间: " .. sevenDaysAgoTime)
```

## 时间格式化参考

os.date 支持的格式化字符：

| 字符 | 说明 | 示例 |
|------|------|------|
| %a | 星期缩写 | Wed |
| %A | 星期全称 | Wednesday |
| %b | 月份缩写 | Sep |
| %B | 月份全称 | September |
| %c | 日期和时间 | 09/16/98 23:48:10 |
| %d | 一个月中的几号 | 01~31 |
| %H | 24 小时制小时 | 00~23 |
| %I | 12 小时制小时 | 01~12 |
| %j | 一年中的第几天 | 001~366 |
| %M | 分钟 | 00~59 |
| %m | 月份 | 01~12 |
| %p | AM/PM | am/pm |
| %S | 秒 | 00~59 |
| %w | 星期几（0=周日） | 0~6 |
| %x | 日期 | 09/16/98 |
| %X | 时间 | 23:48:10 |
| %y | 两位年份 | 00~99 |
| %Y | 四位年份 | 1998 |
| %% | 百分号 | % |

## 常用时间操作示例

### 计算两个时间之间的差值

```lua
local time = require("time")

-- 记录开始时间
local start = time.tickCount()

-- 执行一些操作
for i = 1, 1000000 do
    -- 模拟工作
end

-- 记录结束时间
local end_time = time.tickCount()

-- 计算耗时
local elapsed = end_time - start
print("操作耗时: " .. elapsed .. " 毫秒")
```

### 判断是否超时

```lua
local time = require("time")

-- 设置超时时间（毫秒）
local timeout = 5000
local startTime = time.tickCount()

-- 执行循环
while true do
    -- 检查是否超时
    if time.tickCount() - startTime > timeout then
        print("操作超时！")
        break
    end
    
    -- 执行其他操作
    -- ...
end
```

### 获取特定日期的信息

```lua
-- 获取今年 2 月份的天数
local now = os.date("*t")
local febLastDay = {
    year = now.year,
    month = 3,
    day = 0
}
local daysInFeb = os.date("%d", os.time(febLastDay))
print("今年 2 月份有 " .. daysInFeb .. " 天")

-- 获取本月有多少天
local currentMonth = os.date("*t")
local nextMonth = {
    year = currentMonth.year,
    month = currentMonth.month + 1,
    day = 0
}
local daysInCurrentMonth = os.date("%d", os.time(nextMonth))
print("本月有 " .. daysInCurrentMonth .. " 天")
```

## 兼容性说明

本模块旨在与懒人精灵的时间 API 保持兼容，但由于底层实现差异，存在以下限制：

1. **网络时间精度**：使用 HTTP 请求而非专门的时间协议，精度较低
2. **网络回退机制**：在网络不可用时回退到本地时间
3. **计时精度**：tickCount 使用 Go 的 time.Since()，精度为毫秒级

如果需要更高精度的时间功能，建议使用懒人精灵原版环境或专门的 NTP 客户端。
