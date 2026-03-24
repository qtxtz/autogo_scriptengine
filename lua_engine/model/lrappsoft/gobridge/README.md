# gobridge 模块

gobridge 模块提供了 Lua 和 Go 语言之间的桥接功能，支持：

1. **注册 Lua 回调函数** - 让 Go 代码可以调用 Lua 函数
2. **调用 Go 动态库** - 从 Lua 调用编译为 .so 文件的 Go 函数
3. **类型转换** - 在 Lua 值和 Go 值之间进行转换

## 模块信息

- **模块名称**: `gobridge`
- **可用性**: 仅在 Android 平台可用（需要 cgo 支持）
- **依赖**: cgo, dlfcn

## 方法列表

### 1. gobridge.register(funcname, function)

注册 Lua 回调函数，让 Go 代码可以通过 `CallLuaCallback` 方法调用。

**参数**:
- `funcname` (string) - 注册的函数名称
- `function` (function) - Lua 回调函数

**返回值**: 无

**示例**:
```lua
-- 注册一个加法函数
gobridge.register("add", function(a, b)
    return a + b
end)

-- 注册一个问候函数
gobridge.register("greet", function(name)
    return "Hello, " .. name .. "!"
end)
```

**Go 端调用示例**:
```go
// 从 Go 代码调用已注册的 Lua 回调
results, err := gobridgeModule.CallLuaCallback("add", 1, 2)
if err == nil {
    fmt.Printf("结果: %v\n", results[0]) // 输出: 结果: 3
}

results, err = gobridgeModule.CallLuaCallback("greet", "World")
if err == nil {
    fmt.Printf("结果: %v\n", results[0]) // 输出: 结果: Hello, World!
}
```

---

### 2. gobridge.call(libpath, funcname, ...)

调用 Go 编译的动态库（.so 文件）中的函数。

**参数**:
- `libpath` (string) - Go 编译的动态库路径（如 "libgo.so"）
- `funcname` (string) - Go 导出的函数名
- `...` (any) - 传递给 Go 函数的参数（可变参数）

**返回值**:
- `any` - 返回 Go 函数的执行结果（多种类型）
- `string` - 错误信息（如果调用失败）

**注意事项**:
- 支持传递多种类型的参数（int, string, []byte）
- 支持返回多种类型的返回值
- 可以传递二进制数据
- 动态库会自动缓存，重复调用同一库时不会重复加载

**示例**:
```lua
-- 调用加法函数
local sum = gobridge.call("libgo.so", "Add", 8, 3)
print(sum)  -- 输出: 11

-- 调用字符串处理函数
local msg = gobridge.call("libgo.so", "Greet", "World")
print(msg)  -- 输出: Hello, World!

-- 调用 MD5 函数
local data = gobridge.tobytes("123")
local md5 = gobridge.call("libgo.so", "GetMD5", data)
print(md5)  -- 输出: 202cb962ac59075b964b07152d234b70
```

**对应的 Go 注册函数**:
```go
package main

import (
    "sync"
    "crypto/md5"
    "encoding/hex"
    "github.com/LrGo/LibGo/LibGo/bridge"
)

// Add 加法函数
func Add(a, b int) int {
    return a + b
}

// Greet 问候函数
func Greet(name string) string {
    return "Hello, " + name + "!"
}

// GetMD5 获取数据的 MD5
func GetMD5(data []byte) string {
    hash := md5.Sum(data)
    return hex.EncodeToString(hash[:])
}

var registerOnce sync.Once

func init() {
    registerOnce.Do(func() {
        funcs := []struct {
            name string
            fn   interface{}
        }{
            {"Add", Add},     // 注册加法函数
            {"Greet", Greet}, // 注册问候函数
            {"GetMD5", GetMD5}, // 注册 MD5 函数
        }

        for _, f := range funcs {
            if !bridge.Register(f.name, f.fn) {
                bridge.PrintLog("函数 %s 注册失败", f.name)
            }
        }
    })
}

func main() {}
```

---

### 3. gobridge.tobytes(str)

将字符串转换为字节数组（十六进制表示）。

**参数**:
- `str` (string) - 要转换的字符串

**返回值**:
- `string` - 字节数组的十六进制表示

**示例**:
```lua
local hex = gobridge.tobytes("Hello")
print(hex)  -- 输出: 48656c6c6f
```

---

### 4. gobridge.tostring(hexStr)

将字节数组（十六进制表示）转换为字符串。

**参数**:
- `hexStr` (string) - 字节数组的十六进制表示

**返回值**:
- `string` - 转换后的字符串

**示例**:
```lua
local str = gobridge.tostring("48656c6c6f")
print(str)  -- 输出: Hello
```

---

## 类型转换

### Lua 到 Go 的类型转换

| Lua 类型 | Go 类型 | 说明 |
|---------|---------|------|
| nil | nil | 空值 |
| boolean | bool | 布尔值 |
| number | float64 | 数字（整数和浮点数都转换为 float64） |
| string | string | 字符串 |
| - | []byte | 字节数组（通过 gobridge.tobytes 转换） |

### Go 到 Lua 的类型转换

| Go 类型 | Lua 类型 | 说明 |
|---------|---------|------|
| nil | nil | 空值 |
| bool | boolean | 布尔值 |
| int | number | 整数 |
| int64 | number | 长整数 |
| float64 | number | 浮点数 |
| string | string | 字符串 |
| []byte | string | 字节数组（转换为字符串） |

---

## 使用场景

### 场景 1: Go 调用 Lua 回调

```lua
-- Lua 脚本
gobridge.register("onEvent", function(eventType, data)
    print("收到事件:", eventType, data)
    return "处理完成"
end)
```

```go
// Go 代码
results, err := gobridgeModule.CallLuaCallback("onEvent", "click", 123)
if err == nil {
    fmt.Println("Lua 返回:", results[0]) // 输出: Lua 返回: 处理完成
}
```

### 场景 2: Lua 调用 Go 动态库

```lua
-- Lua 脚本
local result = gobridge.call("libutils.so", "Calculate", 10, 20)
print("计算结果:", result)
```

```go
// Go 动态库代码
package main

import "C"

//export Calculate
func Calculate(a, b int) int {
    return a * b
}

func main() {}
```

### 场景 3: 二进制数据传递

```lua
-- Lua 脚本
local data = gobridge.tobytes("Hello World")
local hash = gobridge.call("libcrypto.so", "SHA256", data)
print("SHA256:", hash)
```

---

## 注意事项

1. **平台限制**
   - gobridge 模块仅在 Android 平台可用
   - 需要编译时启用 cgo 支持

2. **动态库要求**
   - .so 文件必须使用相同的编译器和构建参数编译
   - 函数必须使用 `//export` 注释导出
   - 函数签名必须与调用时使用的类型匹配

3. **类型安全**
   - 类型转换时会进行基本的类型检查
   - 不匹配的类型可能导致运行时错误
   - 建议在调用前验证参数类型

4. **内存管理**
   - 动态库会自动缓存，不会重复加载
   - 字符串和字节数组会自动进行内存管理
   - 不需要手动释放资源

5. **错误处理**
   - 调用失败时会返回 nil 和错误信息
   - 建议始终检查返回的错误信息
   - 错误信息会包含详细的失败原因

---

## 完整示例

### Lua 脚本示例

```lua
-- 注册回调函数
gobridge.register("add", function(a, b)
    print("Lua: 计算加法", a, "+", b)
    return a + b
end)

-- 调用 Go 动态库函数
local sum = gobridge.call("libmath.so", "Add", 100, 200)
print("100 + 200 =", sum)

-- 处理二进制数据
local data = gobridge.tobytes("Hello")
local md5 = gobridge.call("libcrypto.so", "MD5", data)
print("MD5:", md5)

local original = gobridge.tostring(md5)
print("原始数据:", original)
```

### Go 动态库示例

```go
package main

import "C"

//export Add
func Add(a, b int) int {
    return a + b
}

//export MD5
func MD5(data []byte) string {
    // 实现 MD5 计算
    return "md5_hash"
}

func main() {}
```

### 编译动态库

```bash
# 编译为 Android ARM64 动态库
GOOS=android GOARCH=arm64 go build -buildmode=c-shared -o libmath.so

# 编译为 Android ARMv7 动态库
GOOS=android GOARCH=arm go build -buildmode=c-shared -o libmath.so
```

---

## 参考文档

- [lr_document - gobridge](../../lr_document)
- [gopher-lua 文档](https://github.com/yuin/gopher-lua)
- [Go cgo 文档](https://pkg.go.dev/cmd/cgo)
