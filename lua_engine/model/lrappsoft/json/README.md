# JSON 模块

## 概述

JSON 模块提供了 JSON 编码和解码功能，兼容懒人精灵的 `jsonLib` 接口。

## 模块信息

- **模块名称**: `json`
- **全局对象**: `jsonLib`
- **是否可用**: 是

## API 文档

### jsonLib.encode

把 Lua 表格编码成 JSON 字符串。

**函数签名**: `jsonLib.encode(tb)`

**参数**:
- `tb` (any): 要编码的 Lua 值（表格、字符串、数字、布尔值等）

**返回值**:
- `string`: JSON 字符串

**示例**:
```lua
local tb = {
    code = 1,
    data = {
        ret = "hello ok",
        status = 123
    }
}
print(jsonLib.encode(tb))
-- 输出: {"code":1,"data":{"ret":"hello ok","status":123}}
```

### jsonLib.decode

把 JSON 字符串转换成 Lua 表格。

**函数签名**: `jsonLib.decode(json)`

**参数**:
- `json` (string): JSON 字符串

**返回值**:
- `any`: Lua 表格或值

**示例**:
```lua
local json = "{\"data\":{\"ret\":\"hello ok\",\"status\":123},\"code\":1}"
local result = jsonLib.decode(json)
print(result.code)        -- 输出: 1
print(result.data.ret)    -- 输出: hello ok
print(result.data.status) -- 输出: 123
```

## 实现说明

### Lua 表格到 JSON 的转换规则

1. **数组判断**: 如果表格只包含从 1 开始的连续数字键，则转换为 JSON 数组
2. **对象判断**: 如果表格包含字符串键或非连续数字键，则转换为 JSON 对象
3. **空表格**: 如果表格为空，根据上下文判断为数组或对象

### 类型映射

| Lua 类型 | JSON 类型 |
|---------|----------|
| nil | null |
| boolean | boolean |
| number | number |
| string | string |
| table (数组) | array |
| table (对象) | object |

## 注意事项

1. Lua 不区分空数组和空对象，都表示为空表格 `{}`
2. `encode` 方法会自动判断表格类型并生成对应的 JSON
3. 对于复杂的嵌套结构，建议使用 `encode` 和 `decode` 进行往返转换

## 测试

测试代码位于 `test/lrappsoft/json.go`，包含以下测试用例：

- 基本 encode/decode 测试
- 嵌套表格测试
- 数组和对象测试
- 错误处理测试

## 更新日志

### v1.0.0 (2026-03-17)
- 初始版本
- 实现 jsonLib.encode
- 实现 jsonLib.decode
