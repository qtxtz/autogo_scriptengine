# extension 模块（扩展方法）

## 概述

extension 模块提供了懒人精灵的扩展方法，包括加密编码、文件操作、图片处理、OCR、Plugin、节点查找、MySQL 数据库、线程管理等功能。

## 配置选项

extension 模块支持配置未实现方法的行为：

```go
type EmptyMethodConfig struct {
    ThrowException bool // 是否抛出异常（默认 false）
    ShowWarning    bool // 是否显示警告信息（默认 true）
}
```

### 使用示例

```go
// 创建模块实例并设置配置
extensionModule := extension.New(&extension.EmptyMethodConfig{
    ThrowException: false, // 不抛出异常，返回默认值
    ShowWarning:    true,  // 显示警告信息
})

// 运行时修改配置
extensionModule.SetConfig(&extension.EmptyMethodConfig{
    ThrowException: true, // 抛出异常
    ShowWarning:    true,
})

// 获取当前配置
config := extensionModule.GetConfig()
```

### 配置说明

- **ThrowException**: 当设置为 true 时，调用未实现的方法会抛出异常；当设置为 false 时，返回默认值
- **ShowWarning**: 当设置为 true 时，调用未实现的方法会打印警告信息到控制台

### 默认行为

默认情况下，未实现的方法会：

- 打印警告信息（ShowWarning: true）
- 不抛出异常（ThrowException: false）
- 返回默认值（nil 或 0）

## 重要说明

### ✅ 完全可用的功能（28 个）

以下功能在 AutoGo 中**完全可用**：

**加密编码类（7 个）**：

1. **MD5** - 字符串的 MD5
2. **fileMD5** - 获取文件的 MD5 码
3. **encodeBase64** - base64 编码
4. **decodeBase64** - base64 解码
5. **encodeUrl** - url 格式编码
6. **decodeUrl** - url 格式解码
7. **getFileBase64** - 获取文件 base64 编码

**文件操作类（6 个）**：
8\. **fileExist** - 文件或文件夹是否存在
9\. **mkdir** - 创建文件夹
10\. **delfile** - 删除文件或文件夹
11\. **readFile** - 读取文件所有内容
12\. **writeFile** - 写字符串到文件
13\. **fileSize** - 获取文件大小

**图片处理类（5 个）**：
14\. **getImage** - 获取本地图片数据
15\. **rotateImage** - 旋转本地图片
16\. **scaleImage** - 缩放图片
17\. **binaryImage** - 二值化本地图片
18\. **binaryRect** - 区域截图并二值化

### ⚠️ 部分可用的功能（11 个）

以下功能在 AutoGo 中**部分可用**，但需要注意使用限制：

**Plugin 相关（2 个）**：
3\. **LuaEngine.loadApk** - 加载一个 apk 插件

- **实现状态**：使用 AutoGo 的 plugin.LoadApk
- **限制**：返回的是包装的 Lua 表，不是原始的 ClassLoader 对象
- **说明**：可以加载 APK 插件，但 API 与懒人精灵不同

1. **LuaEngine.getContext** - 获取 android 上下文对象
   - **实现状态**：使用 AutoGo 的 plugin.Context
   - **限制**：返回的是包装的 Lua 表，不是原始的 Context 对象
   - **说明**：可以获取 Context，但 API 与懒人精灵不同

**OCR 相关（2 个）**：
5\. **qrDecode** - 二维码解析

- **实现状态**：已实现
- **说明**：使用 github.com/makiuchi-d/gozxing 库解析二维码
- **参数**：
  - `filePath` (string) - 二维码图片文件路径
- **返回值**：
  - `content` (string) - 二维码内容，解析失败返回 nil
- **示例**：
  ```lua
  local content = qrDecode("/sdcard/qrcode.png")
  if content then
      print("二维码内容:", content)
  else
      print("解析失败")
  end
  ```

1. **bdOcr** - 百度 ocr
   - **实现状态**：已实现
   - **说明**：调用百度 OCR API 进行文字识别
   - **参数**：
     - `filePath` (string) - 图片文件路径
     - `apiKey` (string) - 百度 OCR API Key（从百度 AI 开放平台获取）
     - `secretKey` (string) - 百度 OCR Secret Key（从百度 AI 开放平台获取）
     - `type` (number) - OCR 识别类型
       - `0`: 通用文字识别（标准版）
       - `1`: 通用文字识别（含位置信息版）
       - `2`: 通用文字识别（高精度版）
       - `3`: 通用文字识别（高精度含位置版）
   - **返回值**：
     - `result` (string) - OCR 识别结果（JSON 格式），失败返回 nil
   - **示例**：
     ```lua
     local result = bdOcr("/sdcard/test.png", "your_api_key", "your_secret_key", 0)
     if result then
         print("OCR 结果:", result)
         -- 解析 JSON 结果
         local json = require("json")
         local data = json.decode(result)
         if data.words_result then
             for i, word in ipairs(data.words_result) do
                 print("文字", i, ":", word.words)
             end
         end
     else
         print("OCR 识别失败")
     end
     ```
   - **注意事项**：
     - 需要先在百度 AI 开放平台（https://ai.baidu.com/）注册账号并创建应用
     - 获取 API Key 和 Secret Key
     - 开通文字识别服务
     - API 调用需要网络连接
     - 不同识别类型的费用和限制不同，请参考百度官方文档

### ❌ 不可用的功能（1 个）

以下功能在 AutoGo 中**不可用**，但提供了错误提示：

**Java 相关（1 个）**：
1. **import** - 加载 java 类
   - **实现状态**：抛出异常
   - **错误信息**：`import('xxx') 不可用：AutoGo 不支持直接调用 Java 类。请使用 AutoGo 的 plugin 模块或 rhino 模块。`
   - **替代方案**：
     - 使用 AutoGo 的 `plugin` 模块加载 APK 插件
     - 使用 AutoGo 的 `rhino` 模块执行 JavaScript 代码

### ✅ 已实现的功能（36 个）

以下功能在 AutoGo 中**已实现**：

**MySQL 相关（4 个）**：
1. **mysql.connectSQL** - 连接 mysql 数据库
2. **mysql.closeSQL** - 关闭数据库连接
3. **mysql.executeQuerySQL** - 执行 sql 语句并返回结果
4. **mysql.executeSQL** - 执行 sql 语句
   - **实现状态**：已实现，使用 `github.com/go-sql-driver/mysql` 驱动
   - **说明**：支持连接池、事务、查询和执行操作
   - **注意事项**：
     - 使用 Go 的 MySQL 驱动库
     - 返回的是 Lua 表对象，不是原始的连接对象
     - 支持连接池管理（最大 10 个连接，空闲 5 个）
     - 连接超时时间根据参数设置

**线程管理相关（6 个）**：
5. **beginThread** - 启动一个线程
6. **Thread.newThread** / **thread:stopThread** - 创建和停止线程
7. **setCache** / **getCache** / **delCache** - 线程安全的缓存访问
   - **实现状态**：使用 Go 的 goroutine 实现
   - **限制**：与懒人精灵的线程模型完全不同
   - **说明**：
     - 使用 Go 的 goroutine 和 channel 机制
  - 不支持懒人精灵的线程限制（最多 10 个）
  - **支持线程安全的缓存访问**（setCache, getCache, delCache）
  - **不支持懒人精灵的线程共享变量机制**
- **替代方案**：
  - 使用 Go 的 goroutine 和 channel
  - 使用 Lua 的协程（coroutine）
  - **使用 setCache/getCache/delCache 实现线程间数据共享**

## 总结

extension 模块提供了懒人精灵的扩展方法，实现情况如下：

- **完全可用**：34 个（加密编码、文件操作、图片处理、MySQL、qrDecode、bdOcr）
- **部分可用**：2 个（Plugin）
- **不可用**：1 个（import）
- **总计**：37 个方法

**注意**：nodeLib 相关方法已移至 accessibility 模块，请参考 [accessibility 模块文档](../accessibility/README.md)。

### 功能分类

| 类别      | 数量  | 说明                                                                             |
| ------- | --- | ------------------------------------------------------------------------------ |
| 加密编码    | 7 个 | MD5, fileMD5, encodeBase64, decodeBase64, encodeUrl, decodeUrl, getFileBase64  |
| 文件操作    | 6 个 | fileExist, mkdir, delfile, readFile, writeFile, fileSize                       |
| 图片处理    | 5 个 | getImage, rotateImage, scaleImage, binaryImage, binaryRect                     |
| OCR     | 2 个 | qrDecode（已实现）, bdOcr（已实现）                                                      |
| Plugin  | 2 个 | LuaEngine.loadApk, LuaEngine.getContext                                        |
| MySQL   | 4 个 | mysql.connectSQL, mysql.closeSQL, mysql.executeQuerySQL, mysql.executeSQL      |
| 线程管理    | 6 个 | beginThread, Thread.newThread, thread:stopThread, setCache, getCache, delCache |
| Java    | 1 个 | import（抛出异常）                                                                   |
| nodeLib | 9 个 | 已移至 accessibility 模块，详见 [accessibility 模块文档](../accessibility/README.md)       |

## 使用示例

### 示例 1：加密编码操作

```lua
-- 字符串 MD5
local md5 = MD5("123")
print("MD5:", md5)

-- 文件 MD5
local fileMd5 = fileMD5("/sdcard/test.png")
print("文件 MD5:", fileMd5)

-- Base64 编码
local encoded = encodeBase64("欢迎使用懒人精灵")
print("编码:", encoded)

-- Base64 解码
local decoded = decodeBase64(encoded)
print("解码:", decoded)

-- URL 编码
local urlEncoded = encodeUrl("欢迎使用懒人精灵")
print("URL 编码:", urlEncoded)

-- URL 解码
local urlDecoded = decodeUrl(urlEncoded)
print("URL 解码:", urlDecoded)

-- 文件 Base64
local fileBase64 = getFileBase64("/sdcard/test.png")
print("文件 Base64:", string.sub(fileBase64, 1, 50) .. "...")
```

### 示例 2：文件操作

```lua
-- 检查文件是否存在
local exists = fileExist("/sdcard/test.txt")
print("文件存在:", exists)

-- 创建文件夹
local success = mkdir("/sdcard/testdir")
print("创建文件夹:", success)

-- 写入文件
writeFile("/sdcard/test.txt", "Hello, World!")
print("写入文件成功")

-- 读取文件
local content = readFile("/sdcard/test.txt")
print("文件内容:", content)

-- 追加写入文件
writeFile("/sdcard/test.txt", "\n追加内容", true)
print("追加内容成功")

-- 获取文件大小
local size = fileSize("/sdcard/test.txt")
print("文件大小:", size)

-- 删除文件
delfile("/sdcard/test.txt")
print("删除文件成功")
```

### 示例 3：图片处理

```lua
-- 获取图片信息
local w, h, t = getImage("/sdcard/test.png")
print("图片宽度:", w)
print("图片高度:", h)
print("图片类型:", t)

-- 旋转图片
local success = rotateImage("/sdcard/test.png", 90)
print("旋转图片:", success)

-- 缩放图片
local success = scaleImage("/sdcard/src.png", "/sdcard/dst.jpg", 200, 300)
print("缩放图片:", success)

-- 二值化图片
local success = binaryImage("/sdcard/test.png", "/sdcard/binary.png", 150)
print("二值化图片:", success)

-- 区域截图并二值化
local success = binaryRect("/sdcard/rect.png", 100, 100, 300, 300, 150)
print("区域截图二值化:", success)
```

### 示例 4：线程管理（使用线程安全的缓存）

```lua
-- 使用 beginThread 启动线程
function thread_func(arg)
    while true do
        -- 使用 getCache 读取共享数据
        local exit = getCache("exit")
        local data = getCache("data")
        
        if exit == false then
            break
        end
        
        print("我是子线程:" .. arg .. " 共享数据:" .. data)
        sleep(100)
    end
end

-- 使用 setCache 设置初始数据
setCache("exit", true)
setCache("data", 0)

for i = 1, 5 do
    beginThread(thread_func, i)
end

for i = 1, 10 do
    local tick = 10 - i
    -- 使用 setCache 更新共享数据
    setCache("data", tick)
    print("倒计时【" .. tick .. "】秒后结束线程")
    sleep(1000)
end

-- 使用 setCache 设置退出标志
setCache("exit", false)
print("线程结束")

-- 使用 delCache 清理缓存
delCache("exit")
delCache("data")
```

### 示例 5：使用 Thread.newThread

```lua
function thread_func(arg)
    while true do
        local exit = getCache("exit")
        if exit == false then
            break
        end
        print("我是子线程")
        sleep(1000)
    end
end

setCache("exit", true)
local thread = Thread.newThread(thread_func, 11)
sleep(3000)
setCache("exit", false)
thread:stopThread()
delCache("exit")
```

### 示例 6：MySQL 数据库操作

```lua
-- 连接 MySQL 数据库
local con, err = mysql.connectSQL("192.168.1.24", 3306, "mydatabase", "root", "mtest@123", 20000)
if con == nil then
    print("连接失败:", err)
    return
end
print("连接成功")

-- 执行查询 SQL
local result, err = mysql.executeQuerySQL(con, "SELECT * FROM users WHERE age > 18")
if result == nil then
    print("查询失败:", err)
else
    -- 遍历结果
    for i, row in pairs(result) do
        print("第" .. i .. "行:")
        for key, value in pairs(row) do
            print("  " .. key .. ": " .. value)
        end
    end
end

-- 执行 SQL（INSERT/UPDATE/DELETE）
local rows, err = mysql.executeSQL(con, "INSERT INTO users (name, age) VALUES ('张三', 25)")
if rows == nil then
    print("执行失败:", err)
else
    print("影响的行数:", rows)
end

-- 关闭数据库连接
mysql.closeSQL(con)
print("连接已关闭")
```

### 示例 7：Plugin 相关

```lua
-- 加载 APK 插件
local loader = LuaEngine.loadApk("app-release.apk")
if loader ~= nil then
    print("插件加载成功")
    -- 注意：返回的是包装的 Lua 表，不是原始的 ClassLoader 对象
end

-- 获取 Context
local ctx = LuaEngine.getContext()
if ctx ~= nil then
    print("Context 获取成功")
    -- 注意：返回的是包装的 Lua 表，不是原始的 Context 对象
end
```

### 示例 7：import 方法（会抛出异常）

```lua
-- 尝试 import Java 类（会抛出异常）
import('java.lang.*')
-- 错误信息：import('java.lang.*') 不可用：AutoGo 不支持直接调用 Java 类。请使用 AutoGo 的 plugin 模块或 rhino 模块。
```

### 示例 8：import 方法（会抛出异常）

```lua
-- 尝试 import Java 类（会抛出异常）
import('java.lang.*')
-- 错误信息：import('java.lang.*') 不可用：AutoGo 不支持直接调用 Java 类。请使用 AutoGo 的 plugin 模块或 rhino 模块。
```

## 懒人共享变量代码修改方案

### 问题说明

在懒人精灵中，线程间共享变量是通过直接访问全局变量或表来实现的：

```lua
-- 懒人精灵的线程间共享变量方式
local var = {
    exit = true,
    data = 0,
}

function thread_func(arg)
    while var.exit do
        print("我是子线程:" .. arg .. " 共享数据:" .. var.data)
        sleep(100)
    end
end

for i = 1, 5 do
    beginThread(thread_func, i)
end

for i = 1, 10 do
    local tick = 10 - i
    var.data = tick
    print("倒计时【" .. tick .. "】秒后结束线程")
    sleep(1000)
end

var.exit = false
print("线程结束")
```

**问题**：在 AutoGo 中，由于使用 Go 的 goroutine 机制，懒人精灵的线程共享变量方式**不适用**。直接访问全局变量会导致线程安全问题。

### 解决方案

使用 extension 模块提供的**线程安全的缓存访问方法**：`setCache`、`getCache`、`delCache`。

### 修改步骤

#### 步骤 1：将全局变量改为缓存

**原始代码（懒人精灵）**：

```lua
local var = {
    exit = true,
    data = 0,
}
```

**修改后（AutoGo）**：

```lua
-- 使用 setCache 初始化缓存
setCache("exit", true)
setCache("data", 0)
```

#### 步骤 2：在线程函数中使用 getCache 读取数据

**原始代码（懒人精灵）**：

```lua
function thread_func(arg)
    while var.exit do
        print("我是子线程:" .. arg .. " 共享数据:" .. var.data)
        sleep(100)
    end
end
```

**修改后（AutoGo）**：

```lua
function thread_func(arg)
    while true do
        -- 使用 getCache 读取共享数据
        local exit = getCache("exit")
        local data = getCache("data")
        
        if exit == false then
            break
        end
        
        print("我是子线程:" .. arg .. " 共享数据:" .. data)
        sleep(100)
    end
end
```

#### 步骤 3：在主线程中使用 setCache 更新数据

**原始代码（懒人精灵）**：

```lua
for i = 1, 10 do
    local tick = 10 - i
    var.data = tick
    print("倒计时【" .. tick .. "】秒后结束线程")
    sleep(1000)
end

var.exit = false
```

**修改后（AutoGo）**：

```lua
for i = 1, 10 do
    local tick = 10 - i
    -- 使用 setCache 更新共享数据
    setCache("data", tick)
    print("倒计时【" .. tick .. "】秒后结束线程")
    sleep(1000)
end

-- 使用 setCache 设置退出标志
setCache("exit", false)
```

#### 步骤 4：使用 delCache 清理缓存（可选）

```lua
-- 使用 delCache 清理缓存
delCache("exit")
delCache("data")
```

### 完整示例对比

#### 懒人精灵原始代码

```lua
-- 懒人精灵的线程间共享变量方式
local var = {
    exit = true,
    data = 0,
}

function thread_func(arg)
    while var.exit do
        print("我是子线程:" .. arg .. " 共享数据:" .. var.data)
        sleep(100)
    end
end

for i = 1, 5 do
    beginThread(thread_func, i)
end

for i = 1, 10 do
    local tick = 10 - i
    var.data = tick
    print("倒计时【" .. tick .. "】秒后结束线程")
    sleep(1000)
end

var.exit = false
print("线程结束")
```

#### AutoGo 修改后代码

```lua
-- AutoGo 的线程安全缓存方式
-- 使用 setCache 初始化缓存
setCache("exit", true)
setCache("data", 0)

function thread_func(arg)
    while true do
        -- 使用 getCache 读取共享数据
        local exit = getCache("exit")
        local data = getCache("data")
        
        if exit == false then
            break
        end
        
        print("我是子线程:" .. arg .. " 共享数据:" .. data)
        sleep(100)
    end
end

for i = 1, 5 do
    beginThread(thread_func, i)
end

for i = 1, 10 do
    local tick = 10 - i
    -- 使用 setCache 更新共享数据
    setCache("data", tick)
    print("倒计时【" .. tick .. "】秒后结束线程")
    sleep(1000)
end

-- 使用 setCache 设置退出标志
setCache("exit", false)
print("线程结束")

-- 使用 delCache 清理缓存
delCache("exit")
delCache("data")
```

### API 说明

#### setCache(key, value)

设置缓存（线程安全）

**参数**：

- `key` (string) - 缓存键
- `value` (any) - 缓存值（支持 string, number, boolean）

**返回值**：无

**示例**：

```lua
setCache("exit", true)
setCache("data", 0)
setCache("name", "张三")
setCache("age", 25)
```

#### getCache(key)

获取缓存（线程安全）

**参数**：

- `key` (string) - 缓存键

**返回值**：

- `value` (any) - 缓存值，如果不存在则返回 nil

**示例**：

```lua
local exit = getCache("exit")
local data = getCache("data")
local name = getCache("name")
local age = getCache("age")
```

#### delCache(key)

删除缓存（线程安全）

**参数**：

- `key` (string) - 缓存键

**返回值**：无

**示例**：

```lua
delCache("exit")
delCache("data")
delCache("name")
delCache("age")
```

### 注意事项

1. **线程安全**：
   - `setCache`、`getCache`、`delCache` 都是线程安全的
   - 使用 Go 的 `sync.RWMutex` 实现
   - 读写锁机制，支持并发读取
2. **数据类型**：
   - 支持的数据类型：string, number, boolean
   - 不支持 table 类型（会转换为 nil）
   - 不支持 function 类型（会转换为 nil）
3. **性能考虑**：
   - 缓存操作使用互斥锁，会有一定的性能开销
   - 对于高频访问的数据，建议在本地缓存
   - 对于低频访问的数据，直接使用缓存即可
4. **内存管理**：
   - 缓存数据会一直保存在内存中
   - 使用 `delCache` 清理不再需要的缓存
   - 避免缓存大量数据导致内存泄漏
5. **键名冲突**：
   - 使用有意义的键名，避免冲突
   - 建议使用前缀区分不同模块的缓存
   - 例如：`thread_exit`, `thread_data`, `user_info`

### 常见问题

#### Q1：为什么不能直接使用全局变量？

**A**：在 AutoGo 中，使用 Go 的 goroutine 机制，懒人精灵的线程共享变量方式不适用。直接访问全局变量会导致线程安全问题，可能引发数据竞争和不可预期的行为。

#### Q2：setCache/getCache 的性能如何？

**A**：setCache/getCache 使用 Go 的 `sync.RWMutex` 实现，性能良好。对于大多数应用场景，性能开销可以忽略不计。如果对性能有极高要求，可以考虑在本地缓存数据。

#### Q3：支持哪些数据类型？

**A**：支持 string, number, boolean 类型。不支持 table 和 function 类型。如果需要存储复杂数据，可以考虑使用 JSON 序列化。

#### Q4：如何清理缓存？

**A**：使用 `delCache(key)` 清理指定的缓存。建议在不再需要缓存数据时及时清理，避免内存泄漏。

#### Q5：缓存数据会持久化吗？

**A**：不会。缓存数据只保存在内存中，程序重启后会丢失。如果需要持久化，建议使用文件或数据库。

## 迁移指南

### 从懒人精灵迁移到 AutoGo

1. **保留可用的功能**：
   - 所有加密编码操作（MD5, fileMD5, encodeBase64, decodeBase64, encodeUrl, decodeUrl, getFileBase64）
   - 所有文件操作（fileExist, mkdir, delfile, readFile, writeFile, fileSize）
   - 所有图片处理（getImage, rotateImage, scaleImage, binaryImage, binaryRect）
2. **部分使用的功能**：
   - Plugin 相关：可以使用，但 API 不同
   - OCR 相关：等待后续实现或使用 AutoGo 的 ppocr/dotocr 模块
   - nodeLib 相关：已移至 accessibility 模块，详见 [accessibility 模块文档](../accessibility/README.md)
3. **移除不可用的功能**：
   - import：会抛出异常，提示用户使用 plugin 或 rhino 模块
   - 线程管理：使用 Go 的 goroutine 替代，使用 setCache/getCache/delCache 实现线程间数据共享
4. **使用替代方案**：
   - **Java 调用**：使用 AutoGo 的 plugin 模块或 rhino 模块
   - **数据库**：使用已实现的 MySQL 方法（mysql.connectSQL, mysql.closeSQL, mysql.executeQuerySQL, mysql.executeSQL）
   - **线程管理**：使用 Go 的 goroutine 和 channel，使用 setCache/getCache/delCache 实现线程间数据共享
   - **节点查找**：使用 AutoGo 的 uiacc 模块
   - **OCR**：使用 AutoGo 的 ppocr 或 dotocr 模块

## 注意事项

1. **加密编码操作**：
   - MD5 返回 32 位十六进制字符串
   - Base64 编码使用标准 Base64 格式
   - URL 编码使用 QueryEscape 方法
2. **文件操作**：
   - writeFile 的第三个参数为 true 时表示追加模式
   - delfile 可以删除文件或文件夹
   - fileSize 返回文件大小（字节数）
3. **图片处理**：
   - rotateImage 支持的角度：90, 180, 270
   - scaleImage 会创建新文件，不会修改原文件
   - binaryImage 和 binaryRect 默认阈值为 150
4. **线程管理**：
   - 使用 Go 的 goroutine 实现，与懒人精灵的线程模型不同
   - 不支持懒人精灵的线程限制（最多 10 个）
   - 不支持懒人精灵的线程共享变量机制
   - 建议使用 Go 的 goroutine 和 channel 机制
5. **Plugin 相关**：
   - 返回的是包装的 Lua 表，不是原始的对象
   - API 与懒人精灵不同，需要参考 AutoGo 的 plugin 模块文档
6. **不可用功能**：
   - import 会抛出异常，提示用户使用替代方案
   - MySQL 会返回错误，提示用户使用第三方库
   - OCR 相关方法暂未实现，等待后续更新
   - nodeLib 相关方法已移至 accessibility 模块，详见 [accessibility 模块文档](../accessibility/README.md)

## API 映射

### AutoGo API 映射

extension 模块通过以下 AutoGo API 实现功能：

| extension 方法          | AutoGo API                                                    | 说明           |
| --------------------- | ------------------------------------------------------------- | ------------ |
| MD5                   | crypto/md5                                                    | 字符串 MD5      |
| fileMD5               | files.GetMd5                                                  | 文件 MD5       |
| encodeBase64          | encoding/base64                                               | Base64 编码    |
| decodeBase64          | encoding/base64                                               | Base64 解码    |
| encodeUrl             | net/url.QueryEscape                                           | URL 编码       |
| decodeUrl             | net/url.QueryUnescape                                         | URL 解码       |
| getFileBase64         | files.ReadBytes + encoding/base64                             | 文件 Base64    |
| fileExist             | files.Exists                                                  | 文件是否存在       |
| mkdir                 | files.Create                                                  | 创建文件夹        |
| delfile               | files.Remove                                                  | 删除文件或文件夹     |
| readFile              | files.Read                                                    | 读取文件         |
| writeFile             | files.Write / files.Append                                    | 写入文件         |
| fileSize              | files.ReadBytes                                               | 获取文件大小       |
| getImage              | images.ReadFromPath                                           | 读取图片         |
| rotateImage           | images.Rotate + images.Save                                   | 旋转图片         |
| scaleImage            | images.Resize + images.Save                                   | 缩放图片         |
| binaryImage           | images.ReadFromPath + images.ApplyBinarization + images.Save  | 二值化图片        |
| binaryRect            | images.CaptureScreen + images.ApplyBinarization + images.Save | 区域截图二值化      |
| LuaEngine.loadApk     | plugin.LoadApk                                                | 加载 APK 插件    |
| LuaEngine.getContext  | plugin.NewContext                                             | 获取 Context   |
| mysql.connectSQL      | database/sql + github.com/go-sql-driver/mysql                 | 连接 MySQL 数据库 |
| mysql.closeSQL        | database/sql                                                  | 关闭数据库连接      |
| mysql.executeQuerySQL | database/sql                                                  | 执行查询 SQL     |
| mysql.executeSQL      | database/sql                                                  | 执行 SQL       |
| setCache              | sync.RWMutex + map\[string]interface{}                        | 设置缓存（线程安全）   |
| getCache              | sync.RWMutex + map\[string]interface{}                        | 获取缓存（线程安全）   |
| delCache              | sync.RWMutex + map\[string]interface{}                        | 删除缓存（线程安全）   |
| beginThread           | go func() + lua.LState.PCall                                  | 启动 goroutine |
| Thread.newThread      | go func() + lua.LState.PCall                                  | 创建 goroutine |
| thread:stopThread     | close(channel)                                                | 停止 goroutine |

## 待实现功能

以下功能计划在后续版本中实现：

1. **qrDecode** - 二维码解析
   - 需要集成第三方二维码库
   - 支持多种二维码格式
2. **bdOcr** - 百度 OCR
   - 需要集成百度 OCR API
   - 支持多种 OCR 类型
3. **nodeLib 相关方法** - 节点查找
   - 已移至 accessibility 模块
   - 详见 [accessibility 模块文档](../accessibility/README.md)

## 总结

extension 模块提供了懒人精灵的大部分扩展方法，通过 AutoGo API 实现了核心功能：

- **完全可用**：32 个（加密编码、文件操作、图片处理、MySQL、线程安全缓存）
- **部分可用**：11 个（OCR、Plugin、nodeLib）
- **不可用**：3 个（import、线程管理）
- **总计**：49 个方法

对于不可用的功能，提供了清晰的错误提示和替代方案建议。建议用户根据实际需求选择合适的替代方案。

**重要更新**：

- ✅ MySQL 相关方法已完全实现，使用 `github.com/go-sql-driver/mysql` 驱动
- ✅ 线程安全的缓存访问方法已实现（setCache, getCache, delCache）
- ✅ 提供了详细的懒人共享变量代码修改方案
- ⚠️ 线程管理使用 Go 的 goroutine 机制，与懒人精灵的线程模型不同
- ⚠️ 不支持懒人精灵的线程共享变量机制，需要使用 setCache/getCache/delCache 替代

