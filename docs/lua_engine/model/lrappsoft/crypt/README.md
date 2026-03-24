# cryptLib 模块

cryptLib 模块提供了常用的加解密相关函数，兼容懒人精灵 API。

## 功能

该模块提供了以下加解密方法：

### cryptLib.aes_crypt

对数据执行 AES 加密或解密操作，支持多种模式（ECB/CBC/CFB/OFB/CTR）和可选填充。

**函数签名：**
```lua
cryptLib.aes_crypt(data, key, operation, mode, [iv], [padding])
```

**参数：**
- `data` (string): 要加密或解密的原始数据
- `key` (string): 加密密钥（16/24/32 字节，对应 AES-128/192/256）
- `operation` (string): 操作类型，"encrypt" 或 "decrypt"
- `mode` (string): 加密模式："ecb", "cbc", "cfb", "ofb", "ctr"
- `iv` (string, 可选): 初始化向量（16 字节），ECB 模式不需要
- `padding` (boolean, 可选): 是否启用 PKCS#7 填充，默认为 true

**返回值：**
- `string`: 加密或解密后的数据
- `string`: 错误信息（如果失败）

**注意事项：**
- ECB 模式不需要 iv 参数
- 密钥长度必须为 16/24/32 字节
- 当模式不是 ECB 时，iv 长度必须为 16 字节

**示例：**
```lua
local data = "Hello LazyScript!"
local key = "1234567890123456"  -- 16 字节 AES-128
local iv = "0000000000000000"   -- 16 字节 IV

-- AES-CBC 加密
local encrypted = cryptLib.aes_crypt(data, key, "encrypt", "cbc", iv, true)
print("加密结果:", encrypted)

-- AES-CBC 解密
local decrypted = cryptLib.aes_crypt(encrypted, key, "decrypt", "cbc", iv, true)
print("解密结果:", decrypted)
```

### cryptLib.aes_keygen

随机生成指定长度的 AES 密钥（16/24/32）。

**函数签名：**
```lua
cryptLib.aes_keygen(key_length)
```

**参数：**
- `key_length` (number): 密钥长度（16/24/32）

**返回值：**
- `string`: 随机生成的密钥
- `string`: 错误信息（如果失败）

**示例：**
```lua
local key = cryptLib.aes_keygen(32) -- 生成 32 字节密钥 (AES-256)
print("生成的密钥:", key)
```

### cryptLib.aes_ivgen

生成一个 16 字节的随机初始化向量 (IV)，用于 CBC/CFB/OFB 等模式。

**函数签名：**
```lua
cryptLib.aes_ivgen()
```

**返回值：**
- `string`: 16 字节随机 IV
- `string`: 错误信息（如果失败）

**示例：**
```lua
local iv = cryptLib.aes_ivgen()
print("生成的IV:", iv)
```

### cryptLib.rsa_generate_key

生成 RSA 公钥/私钥对，返回 PEM 格式的公钥和私钥。

**函数签名：**
```lua
cryptLib.rsa_generate_key([key_bits])
```

**参数：**
- `key_bits` (number, 可选): 密钥长度，默认 2048（常见 1024/2048/4096）

**返回值：**
- `string`: PEM 格式的公钥
- `string`: PEM 格式的私钥
- `string`: 错误信息（如果失败）

**注意事项：**
- 密钥长度越长越安全，但性能越低；一般建议至少使用 2048 位

**示例：**
```lua
local pubkey, privkey = cryptLib.rsa_generate_key(2048)
print("公钥:\n", pubkey)
print("私钥:\n", privkey)
```

### cryptLib.rsa_encrypt

使用 RSA 公钥（或私钥）对数据进行加密；当使用私钥加密时可用于签名场景。

**函数签名：**
```lua
cryptLib.rsa_encrypt(data, key, is_public_key)
```

**参数：**
- `data` (string): 要加密的数据
- `key` (string): 公钥或私钥（PEM 格式）
- `is_public_key` (boolean): true 表示使用公钥加密，false 表示使用私钥加密

**返回值：**
- `string`: 加密后的数据（原始字节）
- `string`: 错误信息（如果失败）

**注意事项：**
- RSA 加密的数据长度受限（通常不能超过密钥长度减填充开销，例如 PKCS#1 填充为 key_len-11），大数据推荐使用混合加密（对称加密数据，RSA 加密对称密钥）

**示例：**
```lua
-- 假设已生成 pubkey, privkey
local original_data = "需要加密的敏感数据"
local encrypted_data = cryptLib.rsa_encrypt(original_data, pubkey, true)
print("加密后长度:", #encrypted_data)

-- 使用私钥解密
local decrypted = cryptLib.rsa_decrypt(encrypted_data, privkey, false)
print("解密结果:", decrypted)
```

### cryptLib.rsa_decrypt

使用 RSA 私钥（或公钥）对数据进行解密；当使用公钥解密私钥"加密"的数据时可用于验证签名。

**函数签名：**
```lua
cryptLib.rsa_decrypt(data, key, is_public_key)
```

**参数：**
- `data` (string): 要解密的数据
- `key` (string): 私钥或公钥（PEM 格式）
- `is_public_key` (boolean): true 表示使用公钥解密，false 表示使用私钥解密

**返回值：**
- `string`: 解密后的原始数据
- `string`: 错误信息（如果失败）

**示例：**
```lua
local decrypted = cryptLib.rsa_decrypt(encrypted_data, privkey, false)
print("解密结果:", decrypted)
```

## 使用方法

```go
package main

import (
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft"
)

func main() {
    // 注册 cryptLib 模块
    lua_engine.RegisterModule(&lrappsoft.CryptModule{})

    // 创建引擎
    engine := lua_engine.NewLuaEngine(&lua_engine.DefaultConfig())
    defer engine.Close()

    // 执行脚本
    err := engine.ExecuteString(`
        -- AES 加密示例
        local data = "Hello LazyScript!"
        local key = "1234567890123456"
        local iv = cryptLib.aes_ivgen()
        
        local encrypted = cryptLib.aes_crypt(data, key, "encrypt", "cbc", iv, true)
        print("加密结果:", encrypted)
        
        local decrypted = cryptLib.aes_crypt(encrypted, key, "decrypt", "cbc", iv, true)
        print("解密结果:", decrypted)
        
        -- RSA 密钥生成示例
        local pubkey, privkey = cryptLib.rsa_generate_key(2048)
        print("RSA 密钥对生成成功")
        
        -- RSA 加密示例
        local rsa_data = "需要加密的数据"
        local rsa_encrypted = cryptLib.rsa_encrypt(rsa_data, pubkey, true)
        print("RSA 加密成功")
        
        -- RSA 解密示例
        local rsa_decrypted = cryptLib.rsa_decrypt(rsa_encrypted, privkey, false)
        print("RSA 解密结果:", rsa_decrypted)
    `)
    if err != nil {
        panic(err)
    }
}
```

## 注意事项

1. 该模块兼容懒人精灵 API，可以无缝迁移
2. AES 加密支持多种模式，但 ECB 模式不安全，建议使用 CBC 模式
3. RSA 加密的数据长度受限，大数据推荐使用混合加密
4. 密钥和 IV 应该安全存储，不要硬编码在代码中
5. 使用随机生成的密钥和 IV 可以提高安全性
