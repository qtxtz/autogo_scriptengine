# image 模块（图色方法）

图色方法模块提供了图像捕获、处理、识别等相关功能，包括截图、OCR文本识别、目标检测等。

## 模块说明

本模块使用 AutoGo 的 `images`、`ppocr` 和 `yolo` 包来实现懒人精灵的图色方法。

### ⚠️ 功能限制

以下功能在 AutoGo 中**不可用**，使用空操作实现：

1. **PaddleOcr.loadOnnxModel** - 加载自定义 ONNX 模型
   - AutoGo 不支持自定义 ONNX 模型
   - 返回空字符串

2. **PaddleOcr.loadNnccModel** - 加载自定义 NCNN 模型
   - AutoGo 不支持自定义 NCNN 模型
   - 返回空字符串

3. **PaddleOcr.detectWithPadding** - 带填充的 OCR 识别
   - AutoGo 不支持 padding 参数
   - 使用标准 OCR 识别

## 可用功能

### 1. LuaEngine.snapShot - Java 层截图方法

**函数**: `LuaEngine.snapShot(l, t, r, b)`

**描述**: 截取屏幕指定区域

**参数**:
- `l` (number): 左范围
- `t` (number): 上范围
- `r` (number): 右范围
- `b` (number): 下范围

**返回值**:
- 返回一个 image.NRGBA 对象

**示例**:
```lua
local bitmap = LuaEngine.snapShot(100, 100, 200, 200)
-- 使用 bitmap 进行后续操作
```

### 2. LuaEngine.releaseBmp - 释放 bitmap 资源

**函数**: `LuaEngine.releaseBmp(bmp)`

**描述**: 释放 bitmap 资源（Go 的垃圾回收会自动处理，此函数为兼容性提供）

**参数**:
- `bmp` (userdata): bitmap 对象

**返回值**: 无

**示例**:
```lua
LuaEngine.releaseBmp(bitmap)
```

### 3. PaddleOcr.loadModel - 加载自带的模型

**函数**: `PaddleOcr.loadModel(isUseOnnxModel)`

**描述**: 加载 OCR 模型（AutoGo 默认使用 v2 模型）

**参数**:
- `isUseOnnxModel` (boolean): 是否使用 ONNX 模型（AutoGo 不支持此参数）

**返回值**:
- `true`: 加载成功

**示例**:
```lua
if PaddleOcr.loadModel(true) then
    print("OCR 模型加载成功")
end
```

### 4. PaddleOcr.detect - OCR 识别

**函数**: `PaddleOcr.detect(bmp)`

**描述**: 对 bitmap 进行 OCR 文字识别

**参数**:
- `bmp` (userdata): bitmap 对象

**返回值**:
- 返回 JSON 字符串，包含识别结果

**示例**:
```lua
local bitmap = LuaEngine.snapShot(0, 0, 500, 500)
local jsonStr = PaddleOcr.detect(bitmap)
if jsonStr then
    local results = json.decode(jsonStr)
    for i, item in ipairs(results) do
        print(string.format("文字: %s, 坐标: (%d, %d), 置信度: %.2f",
            item.label, item.X, item.Y, item.Score))
    end
end
```

### 5. YoloV5.loadModel - 加载 YOLO 模型

**函数**: `YoloV5.loadModel(paramPath, binPath, labels)`

**描述**: 加载 YOLO 模型（支持 v5 和 v8 版本）

**参数**:
- `paramPath` (string): 模型参数文件路径
- `binPath` (string): 模型二进制文件路径
- `labels` (string): 标签文件路径

**返回值**:
- `true`: 加载成功
- `false`: 加载失败

**说明**:
- AutoGo 支持 YoloV5 和 YoloV8
- 默认使用 v5 版本，CPU 线程数设为 4

**示例**:
```lua
if YoloV5.loadModel("/sdcard/yolo.param", "/sdcard/yolo.bin", "/sdcard/labels.txt") then
    print("YOLO 模型加载成功")
else
    print("YOLO 模型加载失败")
end
```

### 6. YoloV5.detect - YOLO 目标检测

**函数**: `YoloV5.detect(bmp)`

**描述**: 对 bitmap 进行 YOLO 目标检测

**参数**:
- `bmp` (userdata): bitmap 对象

**返回值**:
- 返回 JSON 字符串，包含检测结果

**检测结果格式**:
```json
[
  {
    "X": 100,
    "Y": 200,
    "宽": 50,
    "高": 80,
    "标签": "person",
    "精度": 0.95
  }
]
```

**示例**:
```lua
local bitmap = LuaEngine.snapShot(0, 0, 500, 500)
local jsonStr = YoloV5.detect(bitmap)
if jsonStr then
    local results = json.decode(jsonStr)
    for i, item in ipairs(results) do
        print(string.format("标签: %s, 坐标: (%d, %d), 大小: %dx%d, 置信度: %.2f",
            item.标签, item.X, item.Y, item.宽, item.高, item.精度))
    end
end
```

### 7. YoloV5.close - 关闭 YOLO 模型

**函数**: `YoloV5.close()`

**描述**: 关闭 YOLO 模型实例，释放相关资源

**返回值**: 无

**示例**:
```lua
YoloV5.close()
print("YOLO 模型已关闭")
```

### 8. CaptureScreen - 截取屏幕指定区域

**函数**: `CaptureScreen(x1, y1, x2, y2)`

**描述**: 截取屏幕的指定区域

**参数**:
- `x1` (number): 区域左上角 x 坐标
- `y1` (number): 区域左上角 y 坐标
- `x2` (number): 区域右下角 x 坐标
- `y2` (number): 区域右下角 y 坐标

**返回值**:
- 返回 image.NRGBA 对象

**示例**:
```lua
local img = CaptureScreen(0, 0, 500, 500)
-- 使用 img 进行后续操作
```

### 9. Pixel - 获取指定坐标点的颜色值

**函数**: `Pixel(x, y)`

**描述**: 获取指定坐标点的颜色值

**参数**:
- `x` (number): x 坐标
- `y` (number): y 坐标

**返回值**:
- 返回颜色字符串，格式为 "RRGGBB"

**示例**:
```lua
local color = Pixel(100, 100)
print(string.format("颜色: %s", color))
```

### 10. CmpColor - 比较指定坐标点的颜色

**函数**: `CmpColor(x, y, colorStr, sim)`

**描述**: 比较指定坐标点 (x, y) 的颜色

**参数**:
- `x` (number): x 坐标
- `y` (number): y 坐标
- `colorStr` (string): 颜色字符串，格式如 "FFFFFF|CCCCCC-101010"
- `sim` (number): 相似度，取值范围 0.1-1.0

**返回值**:
- `true`: 颜色匹配
- `false`: 颜色不匹配

**示例**:
```lua
if CmpColor(100, 100, "FF0000", 0.9) then
    print("颜色匹配")
end
```

### 11. FindColor - 在指定区域内查找目标颜色

**函数**: `FindColor(x1, y1, x2, y2, colorStr, sim, dir)`

**描述**: 在指定区域内查找目标颜色

**参数**:
- `x1` (number): 区域左上角 x 坐标
- `y1` (number): 区域左上角 y 坐标
- `x2` (number): 区域右下角 x 坐标
- `y2` (number): 区域右下角 y 坐标
- `colorStr` (string): 颜色字符串，格式如 "FFFFFF|CCCCCC-101010"
- `sim` (number): 相似度，取值范围 0.1-1.0
- `dir` (number): 查找方向
  - 0: 从左到右，从上到下
  - 1: 从右到左，从上到下
  - 2: 从左到右，从下到上
  - 3: 从右到左，从下到上

**返回值**:
- 返回找到颜色的坐标 (x, y)，如果未找到则返回 (-1, -1)

**示例**:
```lua
local x, y = FindColor(0, 0, 500, 500, "FF0000", 0.9, 0)
if x >= 0 and y >= 0 then
    print(string.format("找到颜色在: (%d, %d)", x, y))
end
```

### 12. GetColorCountInRegion - 计算指定区域内符合颜色条件的像素数量

**函数**: `GetColorCountInRegion(x1, y1, x2, y2, colorStr, sim)`

**描述**: 计算指定区域内符合颜色条件的像素数量

**参数**:
- `x1` (number): 区域左上角 x 坐标
- `y1` (number): 区域左上角 y 坐标
- `x2` (number): 区域右下角 x 坐标
- `y2` (number): 区域右下角 y 坐标
- `colorStr` (string): 颜色字符串，格式如 "FFFFFF|CCCCCC-101010"
- `sim` (number): 相似度，取值范围 0.1-1.0

**返回值**:
- 返回符合条件的颜色像素数量

**示例**:
```lua
local count = GetColorCountInRegion(0, 0, 500, 500, "FF0000", 0.9)
print(string.format("找到 %d 个匹配的像素", count))
```

### 13. DetectsMultiColors - 多点颜色比对

**函数**: `DetectsMultiColors(colors, sim)`

**描述**: 根据指定的颜色串信息在屏幕进行多点颜色比对

**参数**:
- `colors` (string): 颜色模板字符串，例如 "369,1220,ffab2d-101010,370,1221,24b1ff-101010"
- `sim` (number): 相似度，取值范围 0.1-1.0

**返回值**:
- `true`: 比对成功
- `false`: 比对失败

**示例**:
```lua
if DetectsMultiColors("100,100,FF0000-101010,200,200,00FF00-101010", 0.9) then
    print("多点颜色匹配")
end
```

### 14. FindMultiColors - 多点找色

**函数**: `FindMultiColors(x1, y1, x2, y2, colors, sim, dir)`

**描述**: 在指定区域内查找匹配的多点颜色序列

**参数**:
- `x1` (number): 区域左上角 x 坐标
- `y1` (number): 区域左上角 y 坐标
- `x2` (number): 区域右下角 x 坐标
- `y2` (number): 区域右下角 y 坐标
- `colors` (string): 颜色模板字符串
- `sim` (number): 相似度，取值范围 0.1-1.0
- `dir` (number): 查找方向（同 FindColor）

**返回值**:
- 返回匹配的首个颜色点的屏幕坐标 (x, y)，如果未找到则返回 (-1, -1)

**示例**:
```lua
local x, y = FindMultiColors(0, 0, 500, 500, "FF0000-101010,10,10,00FF00-101010", 0.9, 0)
if x >= 0 and y >= 0 then
    print(string.format("找到多点颜色在: (%d, %d)", x, y))
end
```

### 15. ReadFromPath - 读取图片文件

**函数**: `ReadFromPath(path)`

**描述**: 读取指定路径的图片文件

**参数**:
- `path` (string): 图片文件路径

**返回值**:
- 返回 image.NRGBA 对象，失败则返回 nil

**示例**:
```lua
local img = ReadFromPath("/sdcard/test.png")
if img then
    print("图片读取成功")
end
```

### 16. ReadFromUrl - 加载网络图片

**函数**: `ReadFromUrl(url)`

**描述**: 从 URL 加载图片

**参数**:
- `url` (string): 图片 URL 地址

**返回值**:
- 返回 image.NRGBA 对象，失败则返回 nil

**示例**:
```lua
local img = ReadFromUrl("https://example.com/test.png")
if img then
    print("图片加载成功")
end
```

### 17. ReadFromBase64 - 解码 Base64 图片

**函数**: `ReadFromBase64(base64Str)`

**描述**: 解码 Base64 数据并返回图片

**参数**:
- `base64Str` (string): Base64 编码的图片数据

**返回值**:
- 返回 image.NRGBA 对象，失败则返回 nil

**示例**:
```lua
local img = ReadFromBase64("iVBORw0KGgoAAAANS...")
if img then
    print("图片解码成功")
end
```

### 18. EncodeToBase64 - 编码图片为 Base64

**函数**: `EncodeToBase64(img, format, quality)`

**描述**: 将图片编码为 Base64 字符串

**参数**:
- `img` (userdata): image.NRGBA 对象
- `format` (string): 编码格式（"png", "jpg" 等），默认 "png"
- `quality` (number): 图片质量，默认 90

**返回值**:
- 返回 Base64 编码的字符串

**示例**:
```lua
local base64Str = EncodeToBase64(img, "png", 90)
print(string.format("Base64 长度: %d", #base64Str))
```

### 19. Save - 保存图片

**函数**: `Save(img, path, quality)`

**描述**: 将图片保存到指定路径

**参数**:
- `img` (userdata): image.NRGBA 对象
- `path` (string): 保存路径
- `quality` (number): 图片质量，默认 90

**返回值**:
- `true`: 保存成功
- `false`: 保存失败

**示例**:
```lua
if Save(img, "/sdcard/screenshot.png", 90) then
    print("图片保存成功")
end
```

### 20. Clip - 剪切图片

**函数**: `Clip(img, x1, y1, x2, y2)`

**描述**: 从图片指定位置剪切区域

**参数**:
- `img` (userdata): image.NRGBA 对象
- `x1` (number): 剪切区域左上角 x 坐标
- `y1` (number): 剪切区域左上角 y 坐标
- `x2` (number): 剪切区域右下角 x 坐标
- `y2` (number): 剪切区域右下角 y 坐标

**返回值**:
- 返回剪切后的 image.NRGBA 对象

**示例**:
```lua
local clipped = Clip(img, 100, 100, 200, 200)
if clipped then
    print("图片剪切成功")
end
```

### 21. Resize - 调整图片大小

**函数**: `Resize(img, width, height)`

**描述**: 调整图片大小

**参数**:
- `img` (userdata): image.NRGBA 对象
- `width` (number): 目标宽度
- `height` (number): 目标高度

**返回值**:
- 返回调整大小后的 image.NRGBA 对象

**示例**:
```lua
local resized = Resize(img, 800, 600)
if resized then
    print("图片调整大小成功")
end
```

### 22. Rotate - 旋转图片

**函数**: `Rotate(img, degree)`

**描述**: 将图片顺时针旋转指定角度

**参数**:
- `img` (userdata): image.NRGBA 对象
- `degree` (number): 旋转角度（顺时针）

**返回值**:
- 返回旋转后的 image.NRGBA 对象

**示例**:
```lua
local rotated = Rotate(img, 90)
if rotated then
    print("图片旋转成功")
end
```

### 23. Grayscale - 灰度化图片

**函数**: `Grayscale(img)`

**描述**: 将图片灰度化

**参数**:
- `img` (userdata): image.NRGBA 对象

**返回值**:
- 返回灰度化后的 image.Gray 对象

**示例**:
```lua
local gray = Grayscale(img)
if gray then
    print("图片灰度化成功")
end
```

### 24. ApplyThreshold - 阈值化图片

**函数**: `ApplyThreshold(img, threshold, maxVal, typ)`

**描述**: 将图片阈值化

**参数**:
- `img` (userdata): image.NRGBA 对象
- `threshold` (number): 阈值
- `maxVal` (number): 阈值化后的最大值
- `typ` (string): 阈值化类型（"BINARY", "BINARY_INV", "TRUNC", "TOZERO", "TOZERO_INV"）

**返回值**:
- 返回阈值化后的 image.Gray 对象

**示例**:
```lua
local thresholded = ApplyThreshold(img, 128, 255, "BINARY")
if thresholded then
    print("图片阈值化成功")
end
```

### 25. ApplyAdaptiveThreshold - 自适应阈值化图片

**函数**: `ApplyAdaptiveThreshold(img, maxValue, adaptiveMethod, thresholdType, blockSize, C)`

**描述**: 将图像进行自适应阈值化处理

**参数**:
- `img` (userdata): image.NRGBA 对象
- `maxValue` (number): 阈值化后的最大值
- `adaptiveMethod` (string): 自适应方法（"MEAN_C", "GAUSSIAN_C"）
- `thresholdType` (string): 阈值化类型（"BINARY", "BINARY_INV"）
- `blockSize` (number): 计算阈值的区域大小
- `C` (number): 常数值，用于调整计算出的阈值

**返回值**:
- 返回自适应阈值化后的 image.Gray 对象

**示例**:
```lua
local adaptive = ApplyAdaptiveThreshold(img, 255, "GAUSSIAN_C", "BINARY", 11, 2)
if adaptive then
    print("图片自适应阈值化成功")
end
```

### 26. ApplyBinarization - 二值化图片

**函数**: `ApplyBinarization(img, threshold)`

**描述**: 将图像进行二值化处理

**参数**:
- `img` (userdata): image.NRGBA 对象
- `threshold` (number): 阈值

**返回值**:
- 返回二值化后的 image.Gray 对象

**示例**:
```lua
local binary = ApplyBinarization(img, 128)
if binary then
    print("图片二值化成功")
end
```

## 使用示例

### 完整的 OCR 识别示例

```lua
-- 加载 OCR 模型
if PaddleOcr.loadModel(true) then
    print("OCR 模型加载成功")
    
    -- 截取屏幕
    local bitmap = LuaEngine.snapShot(0, 0, 500, 500)
    
    -- 进行 OCR 识别
    local jsonStr = PaddleOcr.detect(bitmap)
    if jsonStr then
        local results = json.decode(jsonStr)
        for i, item in ipairs(results) do
            print(string.format("文字: %s, 坐标: (%d, %d), 置信度: %.2f",
                item.label, item.X, item.Y, item.Score))
        end
    end
end
```

### 颜色查找示例

```lua
-- 在指定区域查找红色
local x, y = FindColor(0, 0, 500, 500, "FF0000", 0.9, 0)
if x >= 0 and y >= 0 then
    print(string.format("找到红色在: (%d, %d)", x, y))
    -- 点击该位置
    touchDown(x, y)
    touchUp(x, y)
end
```

### 图片处理示例

```lua
-- 读取图片
local img = ReadFromPath("/sdcard/test.png")

-- 调整大小
local resized = Resize(img, 800, 600)

-- 灰度化
local gray = Grayscale(resized)

-- 二值化
local binary = ApplyBinarization(gray, 128)

-- 保存处理后的图片
Save(binary, "/sdcard/processed.png", 90)
```

## 注意事项

1. **内存管理**: Go 的垃圾回收会自动处理 image.NRGBA 对象的内存释放，不需要手动调用 `LuaEngine.releaseBmp`
2. **颜色格式**: 颜色字符串使用 "RRGGBB" 格式，支持偏色，如 "FF0000-101010"
3. **相似度**: 相似度取值范围 0.1-1.0，值越高表示颜色要求越精确
4. **查找方向**: 查找方向参数控制颜色查找的顺序，影响性能和结果
5. **OCR 模型**: AutoGo 默认使用 v2 模型，不支持自定义模型加载
