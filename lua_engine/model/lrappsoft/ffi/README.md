# ffi 模块（Foreign Function Interface）

## 重要说明

**ffi 模块是 LuaJIT 的特性，AutoGo 使用的是 gopher-lua（Lua 5.1），两者之间存在根本性差异。**

### 主要限制

1. **ffi.cdef 不可用**：无法定义 C 语言类型和函数
2. **ffi.load 不可用**：无法加载动态库（.so 文件）
3. **ffi.new 不可用**：无法创建 cdata 对象
4. **ffi.sizeof 不可用**：无法获取类型大小

### 为什么不可用？

- **LuaJIT**：使用 JIT 编译器，支持 FFI（Foreign Function Interface），可以直接调用 C 函数
- **gopher-lua**：纯 Go 实现的 Lua 5.1 解释器，不支持 FFI 特性

### 替代方案

如果需要调用 C 函数或系统功能，请使用以下替代方案：

1. **使用 AutoGo 提供的 Go 模块**：
   - `cv` 模块：图像处理
   - `time` 模块：时间功能
   - `http` 模块：网络请求
   - `crypt` 模块：加密解密
   - `lfs` 模块：文件系统操作

2. **使用 Go 语言扩展**：
   - 在 Go 代码中实现需要的功能
   - 通过 Lua 模块接口暴露给 Lua 脚本

3. **使用 Lua 标准库**：
   - `os` 模块：操作系统功能
   - `io` 模块：输入输出
   - `package` 模块：包管理

## API 参考

### ffi.cdef(str)

**功能**：定义 C 语言的类型和函数（不可用）

**参数**：
- `str` (string)：C 语言的类型声明和函数原型

**返回值**：
- 返回 `nil` 和错误信息

**示例**：
```lua
local ffi = require("ffi")

-- 这个功能在 AutoGo 中不可用
ffi.cdef[[
    int getpid(void);
]]

-- 会返回错误：ffi.cdef 在 AutoGo 中不可用
```

### ffi.load(str)

**功能**：加载一个动态库（不可用）

**参数**：
- `str` (string)：要加载的库的路径和名称

**返回值**：
- 返回 `nil` 和错误信息

**示例**：
```lua
local ffi = require("ffi")

-- 这个功能在 AutoGo 中不可用
ffi.cdef[[
    int add(int a, int b);
]]

local mylib = ffi.load("/data/local/tmp/liblrapi.so")

-- 会返回错误：ffi.load 在 AutoGo 中不可用
```

### ffi.new(ctype, init)

**功能**：创建 cdata 对象（不可用）

**参数**：
- `ctype` (string)：C 类型
- `init` (any)：初始值

**返回值**：
- 返回 `nil`

**示例**：
```lua
local ffi = require("ffi")

-- 这个功能在 AutoGo 中不可用
local ptr = ffi.new("int*", 42)

-- 会返回 nil
```

### ffi.sizeof(ctype)

**功能**：获取类型大小（不可用）

**参数**：
- `ctype` (string)：C 类型

**返回值**：
- 返回 `0`

**示例**：
```lua
local ffi = require("ffi")

-- 这个功能在 AutoGo 中不可用
local size = ffi.sizeof("int")

-- 会返回 0
```

## 懒人精灵代码迁移指南

如果你的懒人精灵代码使用了 ffi，需要进行以下修改：

### 示例 1：调用 C 函数

**懒人精灵代码**：
```lua
local ffi = require("ffi")
ffi.cdef[[
    int getpid(void);
]]
local clib = ffi.load("libc")
local pid = clib.getpid()
print("当前进程pid:" .. pid)
```

**AutoGo 替代方案**：
```lua
-- 使用 os 模块获取进程信息
local pid = os.getpid()
print("当前进程pid:" .. pid)
```

### 示例 2：加载自定义库

**懒人精灵代码**：
```lua
local ffi = require("ffi")
ffi.cdef[[
    int add(int a, int b);
]]
local mylib = ffi.load("/data/local/tmp/liblrapi.so")
local result = mylib.add(3, 4)
print(result)
```

**AutoGo 替代方案**：
```lua
-- 在 Go 代码中实现 add 函数
-- 通过 Lua 模块暴露给 Lua 脚本

-- 或者使用 Lua 直接实现
function add(a, b)
    return a + b
end

local result = add(3, 4)
print(result)
```

### 示例 3：使用 C 结构体

**懒人精灵代码**：
```lua
local ffi = require("ffi")
ffi.cdef[[
    typedef struct {
        int x;
        int y;
    } Point;
]]
local p = ffi.new("Point", 100, 200)
print(p.x, p.y)
```

**AutoGo 替代方案**：
```lua
-- 使用 Lua 表
local p = { x = 100, y = 200 }
print(p.x, p.y)

-- 或者使用 cv 模块的 Point 功能
local cv = require("cv")
local p = cv.newPoint(100, 200)
local coords = cv.getPoint(p)
print(coords.x, coords.y)
```

## 总结

ffi 模块在 AutoGo 中**不可用**，这是由于技术架构的根本差异导致的。如果你需要调用 C 函数或系统功能，请：

1. 使用 AutoGo 提供的 Go 模块
2. 在 Go 代码中实现需要的功能
3. 使用 Lua 标准库
4. 使用 Lua 表代替 C 结构体

这样可以保持代码的可移植性和跨平台性。
