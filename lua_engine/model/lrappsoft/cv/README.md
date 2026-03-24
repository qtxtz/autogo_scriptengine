# cv 模块（懒人精灵兼容）

cv 模块提供了基于 AutoGo 的 OpenCV 图像处理功能，兼容懒人精灵的 cv 扩展 API。

## ⚠️ 重要限制说明

在使用本模块前，请务必了解以下与懒人精灵原版 API 的差异：

### 1. Mat 对象的内存管理
- **懒人精灵**：需要手动调用 `image:release()` 释放 Mat 对象
- **AutoGo**：Mat 对象**没有 Release 方法**，Go 的垃圾回收会自动管理内存
- **影响**：不需要手动释放 Mat 对象，但需要注意不要创建过多未使用的 Mat 对象

### 2. 指针操作的真实性
- **懒人精灵**：newInt/newDouble 等函数创建真正的 C 语言指针
- **AutoGo**：这些函数**返回 Lua 数值而非真正的指针**
- **影响**：
  - 不能用于需要真正指针的 OpenCV 函数
  - 仅用于传递数值参数
  - `cv.deletePtr()` 实际上不执行任何操作

### 3. Point 和 Point2f 的实现差异
- **懒人精灵**：newPoint 和 newPoint2f 返回真正的 C 语言指针
- **AutoGo**：使用 ID 和 map 来管理 Point 对象
- **影响**：
  - 返回的是内部 ID，不是真正的指针
  - 但功能完全兼容，可以正常使用 getPoint/setPoint/getPoint2f/setPoint2f
  - deletePtr 会从内部 map 中删除对象

### 4. Point2f 的精度问题
- **懒人精灵**：Point2f 支持浮点坐标
- **AutoGo**：使用 `image.Point` 实现，坐标会被**转换为整数**
- **影响**：浮点坐标会丢失小数部分

### 5. OpenCV 功能范围
- **懒人精灵**：集成 OpenCV 4.5.3 的完整函数集合
- **AutoGo**：仅实现了本文档列出的基础函数
- **影响**：高级 OpenCV 功能可能不可用

## 模块加载

```lua
local cv = require("cv")
```

## API 参考

### 1. cv.snapShot - 区域截图并返回 Mat 矩阵

**函数签名**：
```lua
mat = cv.snapShot(l, t, r, b)
```

**参数**：
- `l` (number, 必填)：截图的左边界（像素）
- `t` (number, 必填)：截图的上边界（像素）
- `r` (number, 必填)：截图的右边界（像素）
- `b` (number, 必填)：截图的下边界（像素）

**返回值**：
- `mat` (userdata)：截图对应的 Mat 对象，失败时返回 nil

**示例**：
```lua
local cv = require("cv")

-- 截取屏幕左上角 500x500 区域
local image = cv.snapShot(0, 0, 500, 500)
if image then
    print("截图成功")
    -- 注意：AutoGo 的 Mat 对象没有 release() 方法
    -- Go 的垃圾回收会自动管理内存
else
    print("截图失败")
end
```

### 2. cv.newPoint - 创建 Point 指针

**函数签名**：
```lua
ptr = cv.newPoint(x, y)
```

**参数**：
- `x` (number, 必填)：点的 x 坐标
- `y` (number, 必填)：点的 y 坐标

**返回值**：
- `ptr` (userdata)：Point 指针

**示例**：
```lua
local cv = require("cv")

-- 创建一个点
local point = cv.newPoint(100, 200)
print("创建点成功")

-- 使用完毕后删除
cv.deletePtr(point)
```

### 3. cv.getPoint - 获取 Point 坐标

**函数签名**：
```lua
coords = cv.getPoint(ptr)
```

**参数**：
- `ptr` (userdata, 必填)：Point 指针

**返回值**：
- `coords` (table)：包含 x 和 y 坐标的表，格式为 `{x = number, y = number}`

**示例**：
```lua
local cv = require("cv")

local point = cv.newPoint(100, 200)
local coords = cv.getPoint(point)
print(string.format("坐标: x=%d, y=%d", coords.x, coords.y))
-- 输出: 坐标: x=100, y=200
```

### 4. cv.setPoint - 设置 Point 坐标

**函数签名**：
```lua
cv.setPoint(ptr, x, y)
```

**参数**：
- `ptr` (userdata, 必填)：Point 指针
- `x` (number, 必填)：新的 x 坐标
- `y` (number, 必填)：新的 y 坐标

**返回值**：
- 无

**示例**：
```lua
local cv = require("cv")

local point = cv.newPoint(100, 200)
cv.setPoint(point, 500, 500)
local coords = cv.getPoint(point)
print(string.format("新坐标: x=%d, y=%d", coords.x, coords.y))
-- 输出: 新坐标: x=500, y=500
```

### 5. cv.deletePtr - 删除并回收指针

**函数签名**：
```lua
cv.deletePtr(ptr)
```

**参数**：
- `ptr` (userdata, 必填)：要回收的指针

**返回值**：
- 无

**⚠️ 重要说明**：
- 在 AutoGo 中，此函数**不执行任何操作**
- Go 的垃圾回收会自动管理内存
- 保留此函数是为了与懒人精灵 API 兼容

**示例**：
```lua
local cv = require("cv")

local point = cv.newPoint(100, 200)
-- 使用 point...
cv.deletePtr(point)  -- 实际上不执行任何操作，但保留用于兼容性
```

### 6. cv.newPoint2f - 创建 Point2f 指针

**函数签名**：
```lua
ptr = cv.newPoint2f(x, y)
```

**参数**：
- `x` (number, 必填)：点的 x 坐标（浮点数）
- `y` (number, 必填)：点的 y 坐标（浮点数）

**返回值**：
- `ptr` (userdata)：Point2f 指针

**⚠️ 重要说明**：
- 实际使用 `image.Point` 实现
- 浮点坐标会被**转换为整数**
- 例如：`cv.newPoint2f(100.5, 200.7)` 实际创建的是 `{x=100, y=200}`

**示例**：
```lua
local cv = require("cv")

-- 创建浮点坐标点（实际会被转换为整数）
local point = cv.newPoint2f(100.5, 200.7)
local coords = cv.getPoint2f(point)
print(string.format("坐标: x=%d, y=%d", coords.x, coords.y))
-- 输出: 坐标: x=100, y=200 (小数部分丢失)
```

### 7. cv.getPoint2f - 获取 Point2f 坐标

**函数签名**：
```lua
coords = cv.getPoint2f(ptr)
```

**参数**：
- `ptr` (userdata, 必填)：Point2f 指针

**返回值**：
- `coords` (table)：包含 x 和 y 坐标的表，格式为 `{x = number, y = number}`

**示例**：
```lua
local cv = require("cv")

local point = cv.newPoint2f(100.0, 200.0)
local coords = cv.getPoint2f(point)
print(string.format("坐标: x=%.1f, y=%.1f", coords.x, coords.y))
-- 输出: 坐标: x=100.0, y=200.0
```

### 8. cv.setPoint2f - 设置 Point2f 坐标

**函数签名**：
```lua
cv.setPoint2f(ptr, x, y)
```

**参数**：
- `ptr` (userdata, 必填)：Point2f 指针
- `x` (number, 必填)：新的 x 坐标（浮点数）
- `y` (number, 必填)：新的 y 坐标（浮点数）

**返回值**：
- 无

**⚠️ 重要说明**：
- 浮点坐标会被**转换为整数**

**示例**：
```lua
local cv = require("cv")

local point = cv.newPoint2f(100.0, 200.0)
cv.setPoint2f(point, 500.5, 500.7)
local coords = cv.getPoint2f(point)
print(string.format("新坐标: x=%d, y=%d", coords.x, coords.y))
-- 输出: 新坐标: x=500, y=500 (小数部分丢失)
```

### 9. cv.newInt - 创建整型指针

**函数签名**：
```lua
ptr = cv.newInt(val)
```

**参数**：
- `val` (number, 必填)：初始整数值

**返回值**：
- `ptr` (number)：返回数值本身（非真正的指针）

**⚠️ 重要说明**：
- **不创建真正的指针**，直接返回 Lua 数值
- 仅用于传递数值参数
- 不能用于需要真正指针的 OpenCV 函数

**示例**：
```lua
local cv = require("cv")

-- 创建整型值（实际返回数值）
local val = cv.newInt(5)
print(val)  -- 输出: 5
```

### 10. cv.getInt - 获取整型值

**函数签名**：
```lua
val = cv.getInt(ptr)
```

**参数**：
- `ptr` (number, 必填)：数值

**返回值**：
- `val` (number)：返回传入的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newInt(5)
print(cv.getInt(val))  -- 输出: 5
```

### 11. cv.setInt - 设置整型值

**函数签名**：
```lua
newPtr = cv.setInt(ptr, val)
```

**参数**：
- `ptr` (number, 必填)：原数值（未使用）
- `val` (number, 必填)：新的整数值

**返回值**：
- `newPtr` (number)：返回新设置的数值

**⚠️ 重要说明**：
- 不修改原指针，直接返回新值
- 与懒人精灵行为不同

**示例**：
```lua
local cv = require("cv")

local val = cv.newInt(5)
local newVal = cv.setInt(val, 8)
print(newVal)  -- 输出: 8
```

### 12. cv.newDouble - 创建双精度浮点型指针

**函数签名**：
```lua
ptr = cv.newDouble(val)
```

**参数**：
- `val` (number, 必填)：初始双精度浮点值

**返回值**：
- `ptr` (number)：返回数值本身（非真正的指针）

**⚠️ 重要说明**：
- **不创建真正的指针**，直接返回 Lua 数值

**示例**：
```lua
local cv = require("cv")

-- 创建双精度浮点值
local val = cv.newDouble(5.0)
print(val)  -- 输出: 5.0
```

### 13. cv.getDouble - 获取双精度浮点值

**函数签名**：
```lua
val = cv.getDouble(ptr)
```

**参数**：
- `ptr` (number, 必填)：数值

**返回值**：
- `val` (number)：返回传入的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newDouble(5.0)
print(cv.getDouble(val))  -- 输出: 5.0
```

### 14. cv.setDouble - 设置双精度浮点值

**函数签名**：
```lua
newPtr = cv.setDouble(ptr, val)
```

**参数**：
- `ptr` (number, 必填)：原数值（未使用）
- `val` (number, 必填)：新的双精度浮点值

**返回值**：
- `newPtr` (number)：返回新设置的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newDouble(5.0)
local newVal = cv.setDouble(val, 8.0)
print(newVal)  -- 输出: 8.0
```

### 15. cv.newFloat - 创建单精度浮点型指针

**函数签名**：
```lua
ptr = cv.newFloat(val)
```

**参数**：
- `val` (number, 必填)：初始单精度浮点值

**返回值**：
- `ptr` (number)：返回数值本身（非真正的指针）

**⚠️ 重要说明**：
- **不创建真正的指针**，直接返回 Lua 数值

**示例**：
```lua
local cv = require("cv")

-- 创建单精度浮点值
local val = cv.newFloat(5.0)
print(val)  -- 输出: 5.0
```

### 16. cv.getFloat - 获取单精度浮点值

**函数签名**：
```lua
val = cv.getFloat(ptr)
```

**参数**：
- `ptr` (number, 必填)：数值

**返回值**：
- `val` (number)：返回传入的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newFloat(5.0)
print(cv.getFloat(val))  -- 输出: 5.0
```

### 17. cv.setFloat - 设置单精度浮点值

**函数签名**：
```lua
newPtr = cv.setFloat(ptr, val)
```

**参数**：
- `ptr` (number, 必填)：原数值（未使用）
- `val` (number, 必填)：新的单精度浮点值

**返回值**：
- `newPtr` (number)：返回新设置的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newFloat(5.0)
local newVal = cv.setFloat(val, 8.0)
print(newVal)  -- 输出: 8.0
```

### 18. cv.newLong - 创建长整型指针

**函数签名**：
```lua
ptr = cv.newLong(val)
```

**参数**：
- `val` (number, 必填)：初始长整数值

**返回值**：
- `ptr` (number)：返回数值本身（非真正的指针）

**⚠️ 重要说明**：
- **不创建真正的指针**，直接返回 Lua 数值

**示例**：
```lua
local cv = require("cv")

-- 创建长整型值
local val = cv.newLong(1000000000)
print(val)  -- 输出: 1000000000
```

### 19. cv.getLong - 获取长整型值

**函数签名**：
```lua
val = cv.getLong(ptr)
```

**参数**：
- `ptr` (number, 必填)：数值

**返回值**：
- `val` (number)：返回传入的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newLong(1000000000)
print(cv.getLong(val))  -- 输出: 1000000000
```

### 20. cv.setLong - 设置长整型值

**函数签名**：
```lua
newPtr = cv.setLong(ptr, val)
```

**参数**：
- `ptr` (number, 必填)：原数值（未使用）
- `val` (number, 必填)：新的长整数值

**返回值**：
- `newPtr` (number)：返回新设置的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newLong(1000000000)
local newVal = cv.setLong(val, 1000)
print(newVal)  -- 输出: 1000
```

### 21. cv.newByte - 创建字节型指针

**函数签名**：
```lua
ptr = cv.newByte(val)
```

**参数**：
- `val` (number, 必填)：初始字节值（0-255）

**返回值**：
- `ptr` (number)：返回数值本身（非真正的指针）

**⚠️ 重要说明**：
- **不创建真正的指针**，直接返回 Lua 数值

**示例**：
```lua
local cv = require("cv")

-- 创建字节值
local val = cv.newByte(190)
print(val)  -- 输出: 190
```

### 22. cv.getByte - 获取字节值

**函数签名**：
```lua
val = cv.getByte(ptr)
```

**参数**：
- `ptr` (number, 必填)：数值

**返回值**：
- `val` (number)：返回传入的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newByte(190)
print(cv.getByte(val))  -- 输出: 190
```

### 23. cv.setByte - 设置字节值

**函数签名**：
```lua
newPtr = cv.setByte(ptr, val)
```

**参数**：
- `ptr` (number, 必填)：原数值（未使用）
- `val` (number, 必填)：新的字节值（0-255）

**返回值**：
- `newPtr` (number)：返回新设置的数值

**示例**：
```lua
local cv = require("cv")

local val = cv.newByte(190)
local newVal = cv.setByte(val, 100)
print(newVal)  -- 输出: 100
```

## 完整示例

```lua
local cv = require("cv")

-- 示例 1: 截图
print("=== 示例 1: 截图 ===")
local image = cv.snapShot(0, 0, 500, 500)
if image then
    print("截图成功")
else
    print("截图失败")
end

-- 示例 2: Point 操作
print("\n=== 示例 2: Point 操作 ===")
local point = cv.newPoint(100, 200)
local coords = cv.getPoint(point)
print(string.format("初始坐标: x=%d, y=%d", coords.x, coords.y))

cv.setPoint(point, 500, 500)
coords = cv.getPoint(point)
print(string.format("修改后坐标: x=%d, y=%d", coords.x, coords.y))

cv.deletePtr(point)

-- 示例 3: Point2f 操作（注意精度问题）
print("\n=== 示例 3: Point2f 操作 ===")
local point2f = cv.newPoint2f(100.5, 200.7)
coords = cv.getPoint2f(point2f)
print(string.format("坐标（小数部分丢失）: x=%d, y=%d", coords.x, coords.y))

-- 示例 4: 数值操作
print("\n=== 示例 4: 数值操作 ===")
local intVal = cv.newInt(5)
print(string.format("整型值: %d", cv.getInt(intVal)))

local doubleVal = cv.newDouble(5.5)
print(string.format("双精度值: %.1f", cv.getDouble(doubleVal)))

local byteVal = cv.newByte(190)
print(string.format("字节值: %d", cv.getByte(byteVal)))
```

## 兼容性说明

本模块旨在与懒人精灵的 cv 扩展 API 保持兼容，但由于底层实现差异，存在以下限制：

1. **Mat 对象管理**：AutoGo 的 Mat 对象没有 `release()` 方法，依赖 Go 的垃圾回收
2. **指针真实性**：newInt/newDouble 等函数不创建真正的 C 语言指针
3. **浮点精度**：Point2f 的浮点坐标会被转换为整数
4. **功能范围**：仅实现了基础函数，高级 OpenCV 功能可能不可用

如果需要完整的 OpenCV 4.5.3 功能，建议使用懒人精灵原版环境。
