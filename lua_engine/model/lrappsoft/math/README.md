# math 模块

math 模块提供了对 Lua 标准库 math 库的增强功能，兼容懒人精灵 API。

## 功能

该模块为 Lua 的 math 库添加了以下增强方法，同时保留原有的所有方法：

### math.tointeger

将值转换为整数。

**函数签名：**
```lua
math.tointeger(x)
```

**参数：**
- `x` (number/string): 要转换的值

**返回值：**
- `integer`: 转换后的整数，或 nil

**示例：**
```lua
print(math.tointeger(3.0))    -- 3
print(math.tointeger("666"))  -- 666
print(math.tointeger("abc"))  -- nil
```

### math.type

获取数字的类型。

**函数签名：**
```lua
math.type(x)
```

**参数：**
- `x` (number): 待检测的数

**返回值：**
- `string`: "integer" / "float" 或 nil

**示例：**
```lua
print(math.type(3))     -- "integer"
print(math.type(3.14))  -- "float"
print(math.type("abc")) -- nil
```

### math.ult

无符号比较两个整数。

**函数签名：**
```lua
math.ult(m, n)
```

**参数：**
- `m` (integer): 被比较的整数
- `n` (integer): 被比较的整数

**返回值：**
- `boolean`: 若 m < n（按无符号比较）则为 true

**示例：**
```lua
print(math.ult(1, 2))  -- true
print(math.ult(2, 1))  -- false
```

## 保留的原有方法

该模块会保留 Lua 标准库 math 的所有原有方法，包括但不限于：

- `math.abs` - 绝对值
- `math.acos` - 反余弦
- `math.asin` - 反正弦
- `math.atan` - 反正切
- `math.atan2` - 反正切（两个参数）
- `math.ceil` - 向上取整
- `math.cos` - 余弦
- `math.cosh` - 双曲余弦
- `math.deg` - 弧度转角度
- `math.exp` - 指数
- `math.floor` - 向下取整
- `math.fmod` - 取余数
- `math.frexp` - 分解浮点数
- `math.huge` - 最大数值
- `math.ldexp` - 重组浮点数
- `math.log` - 自然对数
- `math.log10` - 常用对数
- `math.max` - 取得参数中最大值
- `math.min` - 取得参数中最小值
- `math.modf` - 分离整数与小数部分
- `math.pi` - 圆周率
- `math.pow` - 幂运算
- `math.rad` - 角度转弧度
- `math.random` - 产生随机数
- `math.randomseed` - 设置随机种子
- `math.sin` - 正弦
- `math.sinh` - 双曲正弦
- `math.sqrt` - 平方根
- `math.tan` - 正切
- `math.tanh` - 双曲正弦

## 使用方法

该模块可以单独使用，也可以作为 SafeModules 的一部分使用：

```go
package main

import (
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft"
)

func main() {
    // 注册 math 模块
    lua_engine.RegisterModule(&lrappsoft.MathModule{})

    // 创建引擎
    engine := lua_engine.NewLuaEngine(&lua_engine.DefaultConfig())
    defer engine.Close()

    // 执行脚本
    err := engine.ExecuteString(`
        -- 使用增强的 math 函数
        print(math.tointeger(3.0))    -- 3
        print(math.tointeger("666"))  -- 666
        print(math.type(3))           -- "integer"
        print(math.type(3.14))        -- "float"
        print(math.ult(1, 2))         -- true

        -- 使用原有的 math 函数
        print(math.abs(-15))          -- 15
        print(math.ceil(3.1))         -- 4
        print(math.floor(3.9))        -- 3
        print(math.max(1, 10, 5))     -- 10
        print(math.min(1, 10, 5))     -- 1
    `)
    if err != nil {
        panic(err)
    }
}
```

## 注意事项

1. 该模块会保留 Lua 原有的 math 库内容，只添加增强方法
2. 这些方法兼容懒人精灵 API，可以无缝迁移
3. 如果 Lua 版本已经支持这些方法，该模块的实现会覆盖原有实现
