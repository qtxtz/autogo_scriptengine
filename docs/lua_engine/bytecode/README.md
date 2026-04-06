# Lua 字节码功能

## 概述

AutoGo ScriptEngine 现已支持 Lua 字节码功能，允许预编译 Lua 脚本并重复执行，提高执行效率。

## 重要说明

⚠️ **gopher-lua 的字节码格式与标准 Lua 字节码不兼容**

- gopher-lua 使用自己的字节码格式，这是 gopher-lua 特有的实现
- 标准 Lua 编译器（luac）生成的 `.luac` 文件无法在 gopher-lua 中执行
- gopher-lua 的字节码仅能在 gopher-lua 虚拟机中执行
- 字节码格式可能会随 gopher-lua 版本变化，建议保留源码

## 功能特性

### 1. 编译 Lua 源码为字节码

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
)

func main() {
    // 初始化 Lua 引擎
    engine := lua_engine.NewLuaEngine(nil)
    defer engine.Close()
    
    // Lua 源码
    source := `
        local function add(a, b)
            return a + b
        end
        print("结果: " .. add(3, 5))
    `
    
    // 编译源码为字节码
    bytecode, err := engine.CompileString(source, "my_script")
    if err != nil {
        log.Fatalf("编译失败: %v", err)
    }
    
    fmt.Printf("字节码编译成功! 名称: %s\n", bytecode.GetName())
}
```

### 2. 执行预编译的字节码

```go
// 执行字节码
err = engine.ExecuteBytecode(bytecode)
if err != nil {
    log.Fatalf("执行失败: %v", err)
}
```

### 3. 从文件编译字节码

```go
// 从文件编译
bytecode, err := engine.CompileFile("scripts/main.lua")
if err != nil {
    log.Fatalf("编译文件失败: %v", err)
}

// 执行
err = engine.ExecuteBytecode(bytecode)
if err != nil {
    log.Fatalf("执行失败: %v", err)
}
```

### 4. 使用全局函数

```go
// 使用全局引擎
engine := lua_engine.GetLuaEngine()
defer engine.Close()

// 全局编译函数
bytecode, err := lua_engine.CompileString("print('Hello, World!')", "hello")
if err != nil {
    log.Fatal(err)
}

// 全局执行函数
err = lua_engine.ExecuteBytecode(bytecode)
if err != nil {
    log.Fatal(err)
}
```

### 5. 异步执行字节码

```go
// 异步执行字节码
err := engine.ExecuteBytecodeWithMode(bytecode, lua_engine.ExecuteModeAsync)
if err != nil {
    log.Fatal(err)
}
```

## require 函数支持字节码模块

`require` 函数现在支持加载 `.gluac` 字节码文件：

### 模块加载优先级

1. `module.gluac` - gopher-lua 字节码文件
2. `module.lua` - Lua 源码文件
3. `module/init.gluac` - 目录形式的字节码模块
4. `module/init.lua` - 目录形式的源码模块

### 使用示例

假设有以下文件结构：

```
scripts/
├── utils.gluac      # 预编译的字节码模块
├── utils.lua        # 源码模块（备用）
└── main.lua         # 主脚本
```

在 `main.lua` 中：

```lua
-- 加载字节码模块（优先加载 .gluac 文件）
local utils = require("utils")

-- 使用模块功能
utils.hello()
```

## 性能优势

预编译字节码的主要优势：

1. **减少解析开销**：避免每次执行时的词法分析和语法分析
2. **提高启动速度**：直接加载编译后的字节码，减少启动时间
3. **适合重复执行**：在循环或高频调用场景下性能提升明显

### 性能对比示例

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
)

func main() {
    engine := lua_engine.NewLuaEngine(nil)
    defer engine.Close()
    
    source := "local a = 1 + 1"
    
    // 测试直接执行
    start := time.Now()
    for i := 0; i < 10000; i++ {
        engine.ExecuteString(source)
    }
    directDuration := time.Since(start)
    
    // 测试字节码执行
    bytecode, _ := engine.CompileString(source)
    start = time.Now()
    for i := 0; i < 10000; i++ {
        engine.ExecuteBytecode(bytecode)
    }
    bytecodeDuration := time.Since(start)
    
    fmt.Printf("直接执行: %v\n", directDuration)
    fmt.Printf("字节码执行: %v\n", bytecodeDuration)
    fmt.Printf("性能提升: %.2fx\n", float64(directDuration)/float64(bytecodeDuration))
}
```

## API 参考

### Bytecode 类型

```go
type Bytecode struct {
    Proto *lua.FunctionProto  // 编译后的函数原型
    Name  string              // 字节码名称
}
```

### 方法列表

| 方法 | 说明 |
|------|------|
| `CompileString(source string, name ...string)` | 编译源码字符串为字节码 |
| `CompileFile(path string)` | 编译源码文件为字节码 |
| `ExecuteBytecode(bytecode *Bytecode)` | 执行字节码 |
| `ExecuteBytecodeWithMode(bytecode *Bytecode, mode ExecuteMode)` | 带模式的执行字节码 |
| `bytecode.GetName()` | 获取字节码名称 |
| `bytecode.GetFunctionProto()` | 获取函数原型 |

### 全局函数

| 函数 | 说明 |
|------|------|
| `lua_engine.CompileString(source, name)` | 全局编译函数 |
| `lua_engine.CompileFile(path)` | 全局编译文件函数 |
| `lua_engine.ExecuteBytecode(bytecode)` | 全局执行函数 |

## 最佳实践

### 1. 保留源码

由于 gopher-lua 字节码格式可能变化，建议始终保留原始源码：

```go
// 编译并保存字节码
bytecode, err := engine.CompileString(source, "script")
if err != nil {
    log.Fatal(err)
}

// 同时保存源码和字节码
// 源码用于备份和调试
// 字节码用于生产环境
```

### 2. 使用字节码缓存

对于频繁执行的脚本，使用字节码缓存：

```go
// 初始化时编译所有常用脚本
var scriptCache = make(map[string]*lua_engine.Bytecode)

func init() {
    engine := lua_engine.GetLuaEngine()
    
    // 预编译常用脚本
    scripts := map[string]string{
        "utils": "scripts/utils.lua",
        "math":  "scripts/math.lua",
    }
    
    for name, path := range scripts {
        bytecode, err := engine.CompileFile(path)
        if err != nil {
            log.Printf("编译 %s 失败: %v", name, err)
            continue
        }
        scriptCache[name] = bytecode
    }
}
```

### 3. 错误处理

```go
bytecode, err := engine.CompileString(source)
if err != nil {
    // 编译错误，检查源码语法
    log.Printf("编译错误: %v", err)
    return
}

err = engine.ExecuteBytecode(bytecode)
if err != nil {
    // 运行时错误
    log.Printf("执行错误: %v", err)
    return
}
```

## 限制和注意事项

1. **字节码格式不兼容**：gopher-lua 字节码与标准 Lua 字节码不兼容
2. **版本依赖**：字节码格式可能随 gopher-lua 版本变化
3. **调试困难**：字节码难以调试，建议开发阶段使用源码
4. **序列化限制**：gopher-lua 的 FunctionProto 不支持直接序列化

## 示例代码

完整的示例代码请参考：

- [examples/lua_engine/bytecode/main.go](../../examples/lua_engine/bytecode/main.go)

## 更新日志

### v1.0.0 (2026-04-06)

- 添加 `CompileString` 方法支持编译源码字符串
- 添加 `CompileFile` 方法支持编译源码文件
- 添加 `ExecuteBytecode` 方法支持执行字节码
- 添加 `ExecuteBytecodeWithMode` 方法支持异步执行
- 修改 `require` 函数支持加载 `.gluac` 字节码文件
- 添加全局编译和执行函数
