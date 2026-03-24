# lfs 模块

lfs 模块提供了 LuaFileSystem 库的文件系统相关函数，兼容懒人精灵 API。

## 功能

该模块提供了以下文件系统方法：

### lfs.attributes

获取文件或目录的属性信息。

**函数签名：**
```lua
lfs.attributes(path)
```

**参数：**
- `path` (string): 要获取属性的文件或目录的绝对路径

**返回值：**
- `table`: 包含文件或目录属性信息的表
- `string`: 错误信息（如果失败）

**属性表包含的字段：**
- `mode` (string): 文件类型（"file", "directory", "link" 等）
- `size` (number): 文件大小（字节）
- `modification` (number): 修改时间（Unix 时间戳）
- `access` (number): 访问时间（Unix 时间戳）

**示例：**
```lua
local path = "/mnt/sdcard/test.png"
local ret = lfs.attributes(path)
if ret then
    print("文件类型:", ret.mode)
    print("文件大小:", ret.size)
    print("修改时间:", ret.modification)
else
    print("获取属性失败")
end
```

### lfs.chdir

改变当前工作目录。

**函数签名：**
```lua
lfs.chdir(path)
```

**参数：**
- `path` (string): 要切换到的新目录的绝对路径

**返回值：**
- `boolean`: 是否成功切换目录
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard"
local success, err = lfs.chdir(path)
if success then
    print("成功切换到目录")
else
    print("切换目录失败:", err)
end
```

### lfs.currentdir

获取当前工作目录。

**函数签名：**
```lua
lfs.currentdir()
```

**参数：**
- 无

**返回值：**
- `string`: 包含当前工作目录绝对路径的字符串
- `string`: 错误信息（如果失败）

**示例：**
```lua
local currentDir = lfs.currentdir()
print("当前工作目录:", currentDir)
```

### lfs.dir

遍历目录中的文件和子目录。

**函数签名：**
```lua
lfs.dir(path)
```

**参数：**
- `path` (string): 要遍历的目录的绝对路径

**返回值：**
- `function`: 迭代器函数，用于遍历目录内容
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard"
for file in lfs.dir(path) do
    print(file)
end
```

### lfs.link

创建硬链接。

**函数签名：**
```lua
lfs.link(src, dest)
```

**参数：**
- `src` (string): 原始文件的路径
- `dest` (string): 新链接的路径

**返回值：**
- `boolean`: 是否成功创建硬链接
- `string`: 错误信息（如果失败）

**示例：**
```lua
local src = getWorkPath() .. "/test.png"
local dest = getWorkPath() .. "/test1.png"
local success, err = lfs.link(src, dest)
if success then
    print("硬链接创建成功")
else
    print("硬链接创建失败:", err)
end
```

### lfs.lock

锁定文件。

**函数签名：**
```lua
lfs.lock(file, mode)
```

**参数：**
- `file` (file): 要锁定的文件句柄
- `mode` (string): 锁定模式："r"(读锁定), "w"(写锁定), "u"(解锁)

**返回值：**
- `boolean`: 是否成功锁定文件
- `string`: 错误信息（如果失败）

**示例：**
```lua
local filePath = "/mnt/sdcard/test.png"
local file, err = io.open(filePath, "w")
if file then
    local lock, lockErr = lfs.lock(file, "w")
    if lock then
        print("文件锁定成功")
        -- 进行文件操作
        lfs.unlock(file)
        print("文件已解锁")
    else
        print("文件锁定失败:", lockErr)
    end
    file:close()
else
    print("打开文件失败:", err)
end
```

### lfs.mkdir

创建目录。

**函数签名：**
```lua
lfs.mkdir(path)
```

**参数：**
- `path` (string): 要创建的目录的路径

**返回值：**
- `boolean`: 是否成功创建目录
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard/new_folder"
local created = lfs.mkdir(path)
if created then
    print("目录创建成功")
else
    print("目录创建失败")
end
```

### lfs.rmdir

删除空目录。

**函数签名：**
```lua
lfs.rmdir(path)
```

**参数：**
- `path` (string): 要删除的空目录的路径

**返回值：**
- `boolean`: 是否成功删除目录
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard/testdir"
lfs.mkdir(path)
local deleted = lfs.rmdir(path)
if deleted then
    print("目录删除成功")
else
    print("目录删除失败")
end
```

### lfs.symlinkattributes

获取符号链接文件的属性。

**函数签名：**
```lua
lfs.symlinkattributes(path)
```

**参数：**
- `path` (string): 要获取属性的符号链接文件的路径

**返回值：**
- `table`: 包含符号链接文件属性信息的表
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard/test.txt"
local attributes = lfs.symlinkattributes(path)
if attributes then
    print("符号链接属性:")
    for key, value in pairs(attributes) do
        print(key .. ": " .. value)
    end
else
    print("获取符号链接属性失败")
end
```

### lfs.touch

更新文件的访问和修改时间。

**函数签名：**
```lua
lfs.touch(path)
```

**参数：**
- `path` (string): 要更新时间的文件的路径

**返回值：**
- `boolean`: 是否成功更新文件时间
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard/old_file"
local updated = lfs.touch(path)
if updated then
    print("文件时间更新成功")
else
    print("文件时间更新失败")
end
```

### lfs.unlock

解锁文件。

**函数签名：**
```lua
lfs.unlock(file)
```

**参数：**
- `file` (file): 要解锁的文件句柄

**返回值：**
- `boolean`: 是否成功解锁文件
- `string`: 错误信息（如果失败）

**示例：**
```lua
local filePath = "/mnt/sdcard/test.png"
local file, err = io.open(filePath, "w")
if file then
    local lock, lockErr = lfs.lock(file, "w")
    if lock then
        print("文件锁定成功")
        -- 进行文件操作
        lfs.unlock(file)
        print("文件已解锁")
    else
        print("文件锁定失败:", lockErr)
    end
    file:close()
else
    print("打开文件失败:", err)
end
```

### lfs.lock_dir

锁定目录。

**函数签名：**
```lua
lfs.lock_dir(path)
```

**参数：**
- `path` (string): 要锁定的目录的路径

**返回值：**
- `boolean`: 是否成功锁定目录
- `string`: 错误信息（如果失败）

**示例：**
```lua
local path = "/mnt/sdcard/locked_folder"
local locked = lfs.lock_dir(path)
if locked then
    print("目录锁定成功")
else
    print("目录锁定失败")
end
```

## 使用方法

```go
package main

import (
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/lfs"
)

func main() {
    // 注册 lfs 模块
    lua_engine.RegisterModule(&lfs.LfsModule{})

    // 创建引擎
    engine := lua_engine.NewLuaEngine(&lua_engine.DefaultConfig())
    defer engine.Close()

    // 执行脚本
    err := engine.ExecuteString(`
        -- 获取当前工作目录
        local currentDir = lfs.currentdir()
        print("当前工作目录:", currentDir)

        -- 创建目录
        local testDir = "/tmp/test_folder"
        lfs.mkdir(testDir)
        print("目录创建成功")

        -- 获取文件属性
        local attrs = lfs.attributes(testDir)
        print("目录类型:", attrs.mode)

        -- 遍历目录
        print("目录内容:")
        for file in lfs.dir("/tmp") do
            print("  ", file)
        end

        -- 删除目录
        lfs.rmdir(testDir)
        print("目录删除成功")
    `)
    if err != nil {
        panic(err)
    }
}
```

## 注意事项

1. 该模块兼容懒人精灵 API，可以无缝迁移
2. 路径参数建议使用绝对路径
3. 删除目录时，目录必须为空
4. 文件锁定功能在不同操作系统上可能有不同的实现
5. 符号链接操作需要相应的系统权限
6. 在使用文件操作时，建议添加错误处理
