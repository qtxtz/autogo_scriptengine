# http 模块

http 模块提供了 HTTP/HTTPS 网络请求功能，兼容懒人精灵 API。该模块整合了 `http`、`ssl.https` 和 `ltn12` 功能。

## 功能

该模块提供了以下模块：

### http 模块

发起 HTTP/HTTPS 网络请求。

#### http.request

发起 HTTP/HTTPS 请求，支持多种参数配置。

**函数签名：**
```lua
http.request(arg)
```

**参数：**

当 arg 为字符串时：
- `url` (string): 请求地址

当 arg 为表时：
- `url` (string, 必填): 请求地址
- `method` (string, 可选): HTTP 方法（GET/POST/PUT/DELETE 等），默认为 "GET"
- `headers` (table, 可选): 请求头表
- `body` (string, 可选): 请求体内容
- `timeout` (number, 可选): 超时时间（秒），默认为 30
- `sslverify` (boolean, 可选): 是否验证 SSL 证书，默认为 true

**返回值：**
- `number`: HTTP 响应状态码（成功时，如 200、404 等）
- `string`: HTTP 响应状态码对应的文本描述
- `table`: 服务器返回的 HTTP 头信息，是一个包含各种头字段的表
- `string`: 服务器的响应主体内容
- `nil, string`: 失败时返回 nil 和错误信息

**示例：**

```lua
local http = require("http")
local statusCode, statusText, headers, body = http.request("https://httpbin.org/get")

if statusCode then
    print("请求成功")
    print("状态码:", statusCode)
    print("响应体:", body)
else
    print("请求失败:", statusText)
end
```

### ssl.https 模块

发起 HTTPS 网络请求，功能与 http.request 相同。

#### ssl.https.request

发起 HTTPS 请求，支持多种参数配置。

**函数签名：**
```lua
ssl.https.request(arg)
```

**参数：**

当 arg 为字符串时：
- `url` (string): 请求地址

当 arg 为表时：
- `url` (string, 必填): 请求地址
- `method` (string, 可选): HTTP 方法（GET/POST/PUT/DELETE 等），默认为 "GET"
- `headers` (table, 可选): 请求头表
- `body` (string, 可选): 请求体内容
- `timeout` (number, 可选): 超时时间（秒），默认为 30
- `sslverify` (boolean, 可选): 是否验证 SSL 证书，默认为 true

**返回值：**
- `number`: HTTP 响应状态码（成功时，如 200、404 等）
- `string`: HTTP 响应状态码对应的文本描述
- `table`: 服务器返回的 HTTP 头信息，是一个包含各种头字段的表
- `string`: 服务器的响应主体内容
- `nil, string`: 失败时返回 nil 和错误信息

**示例：**

```lua
local https = require("ssl.https")
local statusCode, statusText, headers, body = https.request{
    url = "https://httpbin.org/headers",
    method = "GET",
    headers = {
        ["User-Agent"] = "LuaScript/1.0",
        ["Accept"] = "application/json"
    },
    timeout = 30
}

if statusCode == 200 then
    print("请求成功:", body)
else
    print("请求失败:", statusText)
end
```

### ltn12 模块

提供了 Lua 的过滤器和源（filters and sources）功能，用于数据流处理。

#### ltn12.sink.table

将数据收集到表中。

**函数签名：**
```lua
ltn12.sink.table(tbl)
```

**参数：**
- `tbl` (table): 用于收集数据的表

**返回值：**
- `function`: sink 函数，每次调用时将数据添加到表中

**示例：**
```lua
local ltn12 = require("ltn12")
local responseBody = {}
local sink = ltn12.sink.table(responseBody)

sink("chunk1")
sink("chunk2")
sink("chunk3")

for i, v in ipairs(responseBody) do
    print("[" .. i .. "] " .. v)
end
```

#### ltn12.sink.file

将数据写入文件。

**函数签名：**
```lua
ltn12.sink.file(file)
```

**参数：**
- `file` (file): 文件句柄

**返回值：**
- `function`: sink 函数，每次调用时将数据写入文件

#### ltn12.sink.chain

链接多个 sink。

**函数签名：**
```lua
ltn12.sink.chain(sink1, sink2, ...)
```

**参数：**
- `sink1, sink2, ...` (function): 多个 sink 函数

**返回值：**
- `function`: 组合的 sink 函数

#### ltn12.sink.null

丢弃所有数据。

**函数签名：**
```lua
ltn12.sink.null()
```

**参数：**
- 无

**返回值：**
- `function`: sink 函数，丢弃所有传入的数据

**示例：**
```lua
local ltn12 = require("ltn12")
local sink = ltn12.sink.null()

sink("chunk1")
sink("chunk2")
sink("chunk3")
```

#### ltn12.sink.error

抛出错误的 sink。

**函数签名：**
```lua
ltn12.sink.error(errorMsg)
```

**参数：**
- `errorMsg` (string): 错误信息

**返回值：**
- `function`: sink 函数，调用时抛出指定错误

#### ltn12.source.file

从文件读取数据。

**函数签名：**
```lua
ltn12.source.file(file)
```

**参数：**
- `file` (file): 文件句柄

**返回值：**
- `function`: source 函数，每次调用时从文件读取数据

#### ltn12.source.string

从字符串读取数据。

**函数签名：**
```lua
ltn12.source.string(str)
```

**参数：**
- `str` (string): 要读取的字符串

**返回值：**
- `function`: source 函数，返回字符串数据

**示例：**
```lua
local ltn12 = require("ltn12")
local source = ltn12.source.string("test data")

local data = source()
print(data)  -- 输出：test data
```

#### ltn12.source.chain

链接多个 source。

**函数签名：**
```lua
ltn12.source.chain(source1, source2, ...)
```

**参数：**
- `source1, source2, ...` (function): 多个 source 函数

**返回值：**
- `function`: 组合的 source 函数

## 使用方法

```go
package main

import (
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/http"
)

func main() {
    // 注册 http 模块（包含 ssl.https 和 ltn12）
    lua_engine.RegisterModule(&http.HttpModule{})

    // 创建引擎
    engine := lua_engine.NewLuaEngine(&lua_engine.DefaultConfig())
    defer engine.Close()

    // 执行脚本
    err := engine.ExecuteString(`
        -- 使用 http 模块
        local http = require("http")
        local statusCode, statusText, headers, body = http.request("https://httpbin.org/get")

        if statusCode then
            print("http 请求成功")
            print("状态码:", statusCode)
        end

        -- 使用 ssl.https 模块
        local https = require("ssl.https")
        local statusCode, statusText, headers, body = https.request{
            url = "https://httpbin.org/post",
            method = "POST",
            headers = {
                ["Content-Type"] = "application/json"
            },
            body = '{"key": "value"}',
            timeout = 30
        }

        if statusCode == 200 then
            print("ssl.https 请求成功")
        end

        -- 使用 ltn12 模块
        local ltn12 = require("ltn12")
        local responseBody = {}
        local sink = ltn12.sink.table(responseBody)

        sink("chunk1")
        sink("chunk2")

        print("ltn12 收集数据:", #responseBody)
    `)
    if err != nil {
        panic(err)
    }
}
```

## 与 ssl.https 和 ltn12 组合使用

```lua
local https = require("ssl.https")
local ltn12 = require("ltn12")
local responseBody = {}

-- 发起 HTTPS 请求
local statusCode, statusText, headers, body = https.request{
    url = "https://httpbin.org/get",
    method = "GET",
    timeout = 30
}

if statusCode then
    print("请求成功")
    print("状态码:", statusCode)
    
    -- 使用 ltn12.sink.table 收集响应
    local sink = ltn12.sink.table(responseBody)
    sink(body)

    print("响应体长度:", #body)
else
    print("请求失败:", statusText)
end
```

## 注意事项

1. 该模块兼容懒人精灵 API，可以无缝迁移
2. 使用 `require("http")` 来加载 http 模块
3. 使用 `require("ssl.https")` 来加载 ssl.https 模块
4. 使用 `require("ltn12")` 来加载 ltn12 模块
5. 支持 HTTP/HTTPS 协议
6. 默认超时时间为 30 秒
7. 默认验证 SSL 证书，可以通过 `sslverify = false` 禁用
8. 支持自定义请求头
9. 支持 GET、POST、PUT、DELETE 等 HTTP 方法
10. 响应头以表的形式返回，可以通过 `pairs` 遍历
11. 在网络请求失败时，第一个返回值为 nil，第二个返回值为错误信息
12. 建议在实际应用中添加适当的错误处理和重试机制
13. ltn12 模块提供了灵活的数据流处理机制，可以组合使用多个 sink 和 source
14. 所有模块都通过 `package.preload` 机制加载，不影响文件系统模块的正常加载
15. 支持模块缓存，多次 require 同一个模块会返回同一个实例
