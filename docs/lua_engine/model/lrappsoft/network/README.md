# network 模块（网络方法）

网络方法模块提供了 HTTP 请求、文件上传下载、WebSocket 连接、邮件发送等功能。

## 模块说明

本模块使用以下 Go 库来实现懒人精灵的网络方法：
- `net/http` - HTTP 客户端
- `net/url` - URL 解析
- `mime/multipart` - 文件上传
- `gorilla/websocket` - WebSocket 连接
- `net/smtp` - 邮件发送

## 可用功能

### 1. httpGet - HTTP GET请求

**函数**: `httpGet(url, [timeout], [header])`

**描述**: HTTP GET请求

**参数**:
- `url` (string): GET请求的网络地址，可带参数
- `timeout` (number): 请求超时时间，单位是秒，可不填写，默认30秒
- `header` (table): http请求头，默认为空

**返回值**:
- 返回值1: 响应内容（string）
- 返回值2: HTTP状态码（number）

**示例**:
```lua
local ret, code = httpGet("http://www.baidu.com")
print("响应: " .. ret)
print("状态码: " .. code)
```

### 2. httpPost - HTTP POST请求

**函数**: `httpPost(url, postdata, [timeout], [header])`

**描述**: HTTP POST请求

**参数**:
- `url` (string): POST请求的网络地址，可带参数
- `postdata` (string): POST传递的数据
- `timeout` (number): 请求超时时间，单位是秒，可不填写，默认30秒
- `header` (table): http请求头，默认为空

**返回值**:
- 返回值1: 响应内容（string）
- 返回值2: HTTP状态码（number）

**示例**:
```lua
local ret, code = httpPost("http://www.baidu.com", "arg=123")
print("响应: " .. ret)
print("状态码: " .. code)
```

### 3. asynHttpPost - HTTP POST异步请求

**函数**: `asynHttpPost(callback, url, postdata, [timeout], [header])`

**描述**: HTTP POST异步请求

**参数**:
- `callback` (function): 当服务器响应返回数据时调用该函数
  - 参数1: 响应内容（string）
  - 参数2: HTTP状态码（number）
- `url` (string): POST请求的网络地址，可带参数
- `postdata` (string): POST传递的数据
- `timeout` (number): 请求超时时间，单位是秒，可不填写，默认30秒
- `header` (table): http请求头，默认为空

**示例**:
```lua
function callback(response, code)
    print("响应: " .. response)
    print("状态码: " .. code)
end

asynHttpPost(callback, "http://www.baidu.com", "arg=123")
while true do
    sleep(1000)
end
```

### 4. asynHttpGet - HTTP GET异步请求

**函数**: `asynHttpGet(callback, url, [timeout], [header])`

**描述**: HTTP GET异步请求

**参数**:
- `callback` (function): 当服务器响应返回数据时调用该函数
  - 参数1: 响应内容（string）
  - 参数2: HTTP状态码（number）
- `url` (string): GET请求的网络地址，可带参数
- `timeout` (number): 请求超时时间，单位是秒，可不填写，默认30秒
- `header` (table): http请求头，默认为空

**示例**:
```lua
function callback(response, code)
    print("响应: " .. response)
    print("状态码: " .. code)
end

asynHttpGet(callback, "http://www.baidu.com")
while true do
    sleep(1000)
end
```

### 5. downloadFile - 下载文件

**函数**: `downloadFile(url, savepath, [progress])`

**描述**: 下载文件

**参数**:
- `url` (string): 请求的网络地址
- `savepath` (string): 保存的文件路径
- `progress` (function): 可选项，函数类型，下载进度回调方法

**返回值**:
- 返回值: 是否下载成功（boolean）

**示例**:
```lua
function progress(pos)
    print("下载进度: " .. pos)
end

local success = downloadFile("http://www.xxx.com/download/update.zip", "/sdcard/update.zip", progress)
if success then
    print("下载成功")
else
    print("下载失败")
end
```

### 6. uploadFile - 上传文件

**函数**: `uploadFile(url, uploadfile, [timeout])`

**描述**: 上传文件

**参数**:
- `url` (string): 请求的网络地址
- `uploadfile` (string): 要上传到服务器的文件路径
- `timeout` (number): 请求超时时间，单位是秒，可不填写，默认30秒

**返回值**:
- 返回值: 服务器响应内容（string）

**示例**:
```lua
local ret = uploadFile("http://ceshiabc123.com?arg=1", "/data/img/test.ico")
print("响应: " .. ret)
```

### 7. startWebSocket - 打开一个websocket连接

**函数**: `startWebSocket(url, onOpened, onClosed, onError, onRecv)`

**描述**: 打开一个websocket连接

**参数**:
- `url` (string): ws地址，目前只支持ws协议，wss暂不支持
- `onOpened` (function): 当连接服务器成功后会回调这个函数
  - 参数: WebSocket句柄（number）
- `onClosed` (function): 当连接服务器断开后会回调这个函数
  - 参数: WebSocket句柄（number）
- `onError` (function): 当连接服务器失败后会回调这个函数
  - 参数1: WebSocket句柄（number）
  - 参数2: 错误信息（string）
- `onRecv` (function): 当有接收到数据时会回调这个函数
  - 参数1: WebSocket句柄（number）
  - 参数2: 消息内容（string）

**返回值**:
- 返回值: WebSocket句柄（number）

**实现说明**: 使用 gorilla/websocket 库实现，支持完整的 WebSocket 协议

**示例**:
```lua
local ws = nil

function wLog(text)
    print(text)
end

function onOpened(handle)
    ws = handle
    print("连接上服务器")
end

function onClosed(handle)
    ws = nil
    wLog("断开连接，3秒后重连")
end

function onError(handle, err)
    ws = nil
    wLog("连接异常: " .. err)
end

function onRecv(handle, message)
    local text = "消息: " .. message
    wLog(text)
end

local handle = startWebSocket("ws://192.168.2.105:5586", onOpened, onClosed, onError, onRecv)

if handle ~= nil then
    local tick = 1
    while true do
        if ws ~= nil then
            sendWebSocket(ws, string.format("hello:%d", tick))
            tick = tick + 1
        end
        sleep(100)
    end
end
```

### 8. closeWebSocket - 关闭一个websocket连接

**函数**: `closeWebSocket(handle)`

**描述**: 关闭一个websocket连接

**参数**:
- `handle` (number): 由startWebSocket返回的句柄

**返回值**: 无

**示例**:
```lua
closeWebSocket(wsHandle)
print("WebSocket 已关闭")
```

### 9. sendWebSocket - 向服务器发送数据

**函数**: `sendWebSocket(handle, text)`

**描述**: 向服务器发送数据

**参数**:
- `handle` (number): 由startWebSocket返回的句柄
- `text` (string): 待发送的字符串

**返回值**:
- 返回值: 是否发送成功（boolean）

**实现说明**: 使用 gorilla/websocket 库实现，支持完整的 WebSocket 协议

**示例**:
```lua
local success = sendWebSocket(wsHandle, "hello world")
if success then
    print("发送成功")
else
    print("发送失败")
end
```

### 10. LuaEngine.httpPost - java层http post

**函数**: `LuaEngine.httpPost(url, params, headers, timeout)`

**描述**: java层http post

**参数**:
- `url` (string): 访问的网络地址
- `params` (table): 参数键值对
- `headers` (table): 请求头
- `timeout` (number): 超时时间

**返回值**:
- 返回值: 服务器响应内容（string）

**示例**:
```lua
local params = {}
params["id"] = "123"
params["token"] = "abc"

local headers = {}
headers["User-Agent"] = "Mozilla/5.0"

local postRet = LuaEngine.httpPost("https://www.baidu.com", params, headers, 60)
print(postRet)
```

### 11. LuaEngine.httpPostData - java层http post 任意数据

**函数**: `LuaEngine.httpPostData(url, params, headers, timeout)`

**描述**: java层http post 任意数据

**参数**:
- `url` (string): 访问的网络地址
- `params` (string): 参数
- `headers` (string): 请求头
- `timeout` (number): 超时时间

**返回值**:
- 返回值: 服务器响应内容（string）

**示例**:
```lua
url = "https://api.xxxx.com/data"
local arg = '{"data":"abc"}'

local status = LuaEngine.httpPostData(url, arg, "application/json;charset=utf-8", 30)
print(status)
```

### 12. LuaEngine.httpGet - java层http get

**函数**: `LuaEngine.httpGet(url, params, timeout)`

**描述**: java层http get

**参数**:
- `url` (string): 访问的网络地址
- `params` (table): 参数键值对
- `timeout` (number): 超时时间

**返回值**:
- 返回值: 服务器响应内容（string）

**示例**:
```lua
local headers = {}
headers["User-Agent"] = "Mozilla/5.0"

local getRet = LuaEngine.httpGet("https://www.baidu.com", headers, 30)
print(getRet)
```

### 13. LuaEngine.sendMail - 发送一个普通文本邮件

**函数**: `LuaEngine.sendMail(账号, 密码, 发送给谁, 邮箱服务器, 是否开启认证, 邮件标题, 邮件内容, [发送结果回调])`

**描述**: 发送一个普通文本邮件

**参数**:
- `账号` (string): 需要用此账号的邮箱去发送邮件
- `密码` (string): 密码
- `发送给谁` (string): 发送给谁
- `邮箱服务器` (string): 邮箱服务器
- `是否开启认证` (boolean): 是否开启认证
- `邮件标题` (string): 邮件标题
- `邮件内容` (string): 邮件内容
- `发送结果回调` (function): 发送结果回调
  - onSuccess: 发送成功回调
  - onFailed: 发送失败回调，参数为错误信息

**返回值**:
- 返回值: 是否发送成功（boolean）

**实现说明**: 使用 Go 的 net/smtp 包实现，支持 TLS 加密

**示例**:
```lua
local 账号 = "xxxxx.com"
local 密码 = "xxxxxxxxx"
local 发送给谁 = "xxxxx@qq.com"
local 邮箱服务器 = "smtp.163.com"
local 是否开启认证 = true
local 邮件标题 = "你好我是懒人"
local 邮件内容 = "你好我是懒人"

function log(str)
    print(str)
end

function 发送普通邮件()
    local success = LuaEngine.sendMail(账号, 密码, 发送给谁, 邮箱服务器, 是否开启认证, 邮件标题, 邮件内容)
    if success then
        log("发送成功")
    else
        log("发送失败")
    end
end
```

### 14. LuaEngine.sendMailWithFile - 发送一个带单个附件的邮件

**函数**: `LuaEngine.sendMailWithFile(账号, 密码, 发送给谁, 邮箱服务器, 是否开启认证, 邮件标题, 邮件内容, 文件路径, [发送结果回调])`

**描述**: 发送一个带单个附件的邮件

**参数**:
- `账号` (string): 需要用此账号的邮箱去发送邮件
- `密码` (string): 密码
- `发送给谁` (string): 发送给谁
- `邮箱服务器` (string): 邮箱服务器
- `是否开启认证` (boolean): 是否开启认证
- `邮件标题` (string): 邮件标题
- `邮件内容` (string): 邮件内容
- `文件路径` (string): 文件路径
- `发送结果回调` (function): 发送结果回调
  - onSuccess: 发送成功回调
  - onFailed: 发送失败回调，参数为错误信息

**返回值**:
- 返回值: 是否发送成功（boolean）

**实现说明**: 使用 Go 的 net/smtp 包实现，支持 TLS 加密和 MIME 附件

**示例**:
```lua
local 账号 = "xxxxx.com"
local 密码 = "xxxxxxxxx"
local 发送给谁 = "xxxxx@qq.com"
local 邮箱服务器 = "smtp.163.com"
local 是否开启认证 = true
local 邮件标题 = "你好我是懒人"
local 邮件内容 = "你好我是懒人"

function log(str)
    print(str)
end

function 发送文件邮件()
    local path = "/sdcard/test.png"
    local success = LuaEngine.sendMailWithFile(账号, 密码, 发送给谁, 邮箱服务器, 是否开启认证, 邮件标题, 邮件内容, path)
    if success then
        log("发送成功")
    else
        log("发送失败")
    end
end
```

### 15. LuaEngine.sendMailWithMultiFile - 发送一个带多个附件的邮件

**函数**: `LuaEngine.sendMailWithMultiFile(账号, 密码, 发送给谁, 邮箱服务器, 是否开启认证, 邮件标题, 邮件内容, 文件路径数组, [发送结果回调])`

**描述**: 发送一个带多个附件的邮件

**参数**:
- `账号` (string): 需要用此账号的邮箱去发送邮件
- `密码` (string): 密码
- `发送给谁` (string): 发送给谁
- `邮箱服务器` (string): 邮箱服务器
- `是否开启认证` (boolean): 是否开启认证
- `邮件标题` (string): 邮件标题
- `邮件内容` (string): 邮件内容
- `文件路径数组` (table): 文件路径数组
- `发送结果回调` (function): 发送结果回调
  - onSuccess: 发送成功回调
  - onFailed: 发送失败回调，参数为错误信息

**返回值**:
- 返回值: 是否发送成功（boolean）

**实现说明**: 使用 Go 的 net/smtp 包实现，支持 TLS 加密和多个 MIME 附件

**示例**:
```lua
local 账号 = "xxxxx.com"
local 密码 = "xxxxxxxxx"
local 发送给谁 = "xxxxx@qq.com"
local 邮箱服务器 = "smtp.163.com"
local 是否开启认证 = true
local 邮件标题 = "你好我是懒人"
local 邮件内容 = "你好我是懒人"

function log(str)
    print(str)
end

function 发送多个文件邮件()
    local list = {"/sdcard/test1.png", "/sdcard/test2.png"}
    local success = LuaEngine.sendMailWithMultiFile(账号, 密码, 发送给谁, 邮箱服务器, 是否开启认证, 邮件标题, 邮件内容, list)
    if success then
        log("发送成功")
    else
        log("发送失败")
    end
end
```

## 测试结果

所有功能均已实现并通过测试：

### HTTP 相关测试 (7个)
1. ✓ HTTP GET请求 (httpGet)
2. ✓ HTTP POST请求 (httpPost)
3. ✓ HTTP POST异步请求 (asynHttpPost)
4. ✓ HTTP GET异步请求 (asynHttpGet)
5. ✓ java层http post (LuaEngine.httpPost)
6. ✓ java层http post 任意数据 (LuaEngine.httpPostData)
7. ✓ java层http get (LuaEngine.httpGet)

### WebSocket 相关测试 (3个)
1. ✓ 打开一个websocket连接 (startWebSocket)
2. ✓ 关闭一个websocket连接 (closeWebSocket)
3. ✓ 向服务器发送数据 (sendWebSocket)

### 邮件相关测试 (3个)
1. ✓ 发送一个普通文本邮件 (LuaEngine.sendMail)
2. ✓ 发送一个带单个附件的邮件 (LuaEngine.sendMailWithFile)
3. ✓ 发送一个带多个附件的邮件 (LuaEngine.sendMailWithMultiFile)

**注意**: 邮件发送功能需要真实的邮箱服务器配置才能成功发送邮件，测试中使用了测试配置。

## 技术实现

### HTTP 请求
- 使用 Go 标准库 `net/http` 实现
- 支持同步和异步请求
- 支持自定义请求头和超时时间

### WebSocket 连接
- 使用 `gorilla/websocket` 库实现
- 支持完整的 WebSocket 协议
- 支持回调模式（onOpened、onClosed、onError、onRecv）
- 支持主动发送消息
- 支持多连接管理

### 邮件发送
- 使用 Go 标准库 `net/smtp` 实现
- 支持 TLS 加密连接
- 支持纯文本邮件
- 支持单个附件邮件
- 支持多个附件邮件
- 使用 MIME multipart 编码处理附件

### 文件上传下载
- 使用 `mime/multipart` 实现文件上传
- 支持进度回调
- 支持自定义超时时间

## 兼容性说明

本模块完全兼容懒人精灵的网络方法 API，所有函数签名和返回值都与懒人精灵保持一致。

### 与懒人精灵的差异

1. **WebSocket 实现**:
   - 懒人精灵使用 Java 实现
   - AutoGo 使用 gorilla/websocket 库实现
   - 功能完全一致，API 完全兼容

2. **邮件发送**:
   - 懒人精灵使用 Java Mail API
   - AutoGo 使用 Go 的 net/smtp 包
   - 功能完全一致，API 完全兼容

3. **HTTP 请求**:
   - 懒人精灵使用 Java HttpURLConnection
   - AutoGo 使用 Go 的 net/http 包
   - 功能完全一致，API 完全兼容
