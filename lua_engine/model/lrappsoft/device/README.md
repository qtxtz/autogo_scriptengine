# device 模块（设备方法）

## 概述

device 模块提供了懒人精灵的设备相关方法。由于 AutoGo 运行在 Android 环境，大部分方法已经通过 AutoGo API 实现。

## 重要说明

### ✅ 可用功能

以下功能在 AutoGo 中**可用**（通过 AutoGo API 实现）：

1. **getCpuArch** - 获取 CPU 架构
2. **getWorkPath** - 获取脚本工作目录
3. **getCpuAbi / getCpuAbi2** - 获取 CPU 支持类型
4. **rnd** - 生成随机整数
5. **sleep** - 线程休眠
6. **printEx** - 调试输出
7. **vibrate** - 设备震动
8. **getDisplayDpi** - 获取设备 DPI
9. **getBatteryLevel** - 获取设备电量
10. **playAudio** - 播放音频文件
11. **getPackageName** - 获取当前应用包名
12. **getInstalledApk** - 获取已安装 APK
13. **getInstalledApps** - 获取已安装应用信息
14. **installApk** - 安装 APK
15. **getDisplayInfo** - 获取设备分辨率等详细信息
16. **getBrand** - 获取设备品牌
17. **getBootLoader** - 获取 Bootloader 版本号
18. **getBoard** - 获取主板编号
19. **getProduct** - 获取产品代号
20. **getDevice** - 获取设备别名
21. **getModel** - 获取设备型号
22. **getHardware** - 获取硬件序列号
23. **getId** - 获取 id 修订号
24. **getFingerprint** - 获取编译指纹
25. **getSdkVersion** - 获取 SDK 版本
26. **getOsVersionName** - 获取系统版本名
27. **getWifiMac** - 获取 WiFi MAC 地址
28. **runApp** - 打开指定应用
29. **stopApp** - 关闭指定应用
30. **scanImage** - 刷新图片到媒体库
31. **getCurrentActivity** - 获取最顶层 Activity 名称
32. **frontAppName** - 获取前台应用包名
33. **appIsFront** - 判断指定包名是否为前台运行
34. **getDisplaySize** - 获取屏幕分辨率
35. **setDisplayPowerOff** - 设置息屏运行
36. **getDisplayRotate** - 获取屏幕旋转方向
37. **lockScreen** - 保持屏幕常亮
38. **appIsRunning** - 判断 app 是否正在运行
39. **runIntent** - 构建并发送 Intent
40. **sendSms** - 发送短信
41. **zip** - 压缩文件或文件夹
42. **unZip** - 解压文件
43. **readPasteboard** - 读取剪贴板内容
44. **writePasteboard** - 写入剪贴板内容

### ⚠️ 部分可用功能

以下功能在 AutoGo 中**部分可用**，但需要注意使用限制：

1. **getSdPath** - 获取 SD 卡路径
   - 返回值：`/sdcard`
   - **使用危害**：直接使用 `/sdcard` 可能导致权限问题或路径错误
   - **建议**：使用 `getWorkPath()` 获取正确的工作目录
   - **原因**：AutoGo 无法获取真实的 SD 卡路径，返回固定值仅用于兼容性

2. **restartScript** - 重启当前脚本
   - 返回值：无
   - **使用方法**：调用 `restartScript()` 来重启脚本
   - **前提条件**：需要在引擎初始化时配置 `OnExit = ExitActionRestart`
   - **强制退出**：使用 `os.exit(-1)` 可以强制退出，不执行任何退出动作
   - **正常退出**：使用 `os.exit(0)` 或 `os.exit(其他值)` 会执行配置的退出动作
   - **异常退出**：脚本异常退出时始终打印错误日志
   - **示例**：
     ```lua
     -- 在 Go 代码中初始化引擎时设置
     config := lua_engine.DefaultConfig()
     config.OnExit = lua_engine.ExitActionRestart
     engine := lua_engine.New(config)
     
     -- 在 Lua 脚本中调用
     restartScript()  -- 这会触发重启
     ```

### ⚠️ 空操作方法

以下方法在 AutoGo 中为**空操作**（不执行任何操作），但可以正常调用以保持代码兼容性：

## 空操作方法详细说明

以下是所有 33 个空操作方法及其不可实现的原因：

### Android 系统相关（15 个）

1. **getApkVerInt** - 获取 APK 壳子版本
   - **原因**：AutoGo 不是懒人精灵，无法获取 APK 壳子版本
   - **说明**：这是懒人精灵特有的功能，用于获取脚本引擎的版本号

2. **getScriptVersion** - 获取脚本版本号
   - **原因**：AutoGo 不支持脚本版本号管理
   - **说明**：懒人精灵支持脚本版本管理，AutoGo 不需要此功能

3. **getSubscriberId** - 获取 IMSI（国际移动用户识别码）
   - **原因**：需要 `READ_PHONE_STATE` 权限
   - **说明**：获取 SIM 卡的 IMSI 需要 Android 高级权限，AutoGo 默认不请求此权限

4. **getSimSerialNumber** - 获取 SIM 序列号
   - **原因**：需要 `READ_PHONE_STATE` 权限
   - **说明**：获取 SIM 卡序列号需要 Android 高级权限，AutoGo 默认不请求此权限

5. **installLrPkg** - 安装 lrj 更新包
   - **原因**：懒人精灵特有功能
   - **说明**：AutoGo 不支持懒人精灵的更新包格式（.lrj）

6. **getLrApi** - 获取 FFI API
   - **原因**：AutoGo 不支持 FFI
   - **说明**：懒人精灵支持 FFI（外部函数接口），AutoGo 不需要此功能

7. **getOaid** - 获取 OAID（开放匿名设备标识符）
   - **原因**：需要特定权限和厂商支持
   - **说明**：OAID 是中国广告联盟推出的设备标识符，需要厂商支持和特定权限

8. **getInsallAppInfos** - 获取所有已安装 app 的详细信息
   - **原因**：需要 `QUERY_ALL_PACKAGES` 权限
   - **说明**：获取所有已安装应用的详细信息需要 Android 高级权限，AutoGo 默认不请求此权限

9. **getSensorsInfo** - 获取传感器信息
   - **原因**：AutoGo 不支持传感器访问
   - **说明**：AutoGo 主要用于自动化操作，不需要传感器信息

10. **checkIsDebug** - 检测是否为调试模式
    - **原因**：AutoGo 不支持调试状态检测
    - **说明**：懒人精灵支持调试模式检测，AutoGo 不需要此功能

11. **getDeviceId** - 获取设备 ID
    - **原因**：需要 `READ_PHONE_STATE` 权限
    - **说明**：获取设备 ID 需要 Android 高级权限，AutoGo 默认不请求此权限

12. **getManufacturer** - 获取制造商代号
    - **原因**：需要特定权限
    - **说明**：获取设备制造商信息需要特定权限

13. **setDpiToVir** - 切换虚拟分辨率
    - **原因**：AutoGo 不支持虚拟分辨率切换
    - **说明**：懒人精灵支持虚拟分辨率以适配不同设备，AutoGo 不需要此功能

14. **setDpiToRealy** - 恢复真实分辨率
    - **原因**：AutoGo 不支持真实分辨率恢复
    - **说明**：与 setDpiToVir 配合使用，AutoGo 不需要此功能

15. **getRunEnvType** - 获取运行环境类型
    - **原因**：AutoGo 不支持运行环境类型检测
    - **说明**：懒人精灵支持多种运行环境（root、handle、accessibility），AutoGo 不需要此功能

### 悬浮窗和 UI 功能（4 个）

16. **setControlBarPosNew** - 设置悬浮窗位置
    - **原因**：AutoGo 不支持悬浮窗
    - **说明**：懒人精灵提供悬浮窗控制，AutoGo 不需要此功能

17. **showControlBar** - 显示或隐藏悬浮按钮
    - **原因**：AutoGo 不支持悬浮窗
    - **说明**：懒人精灵提供悬浮窗控制，AutoGo 不需要此功能

18. **getUIConfig** - 获取 UI 配置
    - **原因**：AutoGo 不支持 UI 配置
    - **说明**：懒人精灵支持自定义 UI，AutoGo 不需要此功能

19. **setUserEventCallBack** - 设置用户事件回调
    - **原因**：AutoGo 不支持悬浮窗
    - **说明**：懒人精灵的悬浮窗事件回调，AutoGo 不需要此功能

### 文件系统操作（Android 特有）（2 个）

20. **extractAssets** - 将打包资源释放到目录
    - **原因**：懒人精灵特有功能
    - **说明**：懒人精灵支持将资源打包到脚本中，AutoGo 不需要此功能

21. **extractApkAssets** - 将指定资源释放到目录
    - **原因**：懒人精灵特有功能
    - **说明**：懒人精灵支持从 APK 中提取资源，AutoGo 不需要此功能

### 权限相关功能（3 个）

22. **setRootEnvMode** - 设置 root 模式
    - **原因**：AutoGo 不支持 root 模式设置
    - **说明**：懒人精灵支持 root 模式以获得更高权限，AutoGo 不需要此功能

23. **setHandleEnvMode** - 设置 handle 模式
    - **原因**：AutoGo 不支持 handle 模式设置
    - **说明**：懒人精灵支持 handle 模式，AutoGo 不需要此功能

24. **setAccessibilityEnvMode** - 设置无障碍模式
    - **原因**：AutoGo 不支持无障碍模式设置
    - **说明**：懒人精灵支持无障碍模式，AutoGo 不需要此功能

### 多线程和回调功能（6 个）

25. **setMainThreadPause** - 暂停主线程运行
    - **原因**：AutoGo 不支持多线程控制
    - **说明**：懒人精灵支持多线程控制，AutoGo 使用单线程模型

26. **setMainThreadResume** - 恢复主线程运行
    - **原因**：AutoGo 不支持多线程控制
    - **说明**：懒人精灵支持多线程控制，AutoGo 使用单线程模型

27. **setTimer** - 定时执行指定函数
    - **原因**：AutoGo 不支持定时器
    - **说明**：可以使用 Lua 的 `os.time()` 和 `sleep()` 实现类似功能

28. **setStopCallBack** - 设置脚本结束回调
    - **原因**：AutoGo 不支持脚本结束回调
    - **说明**：懒人精灵支持脚本结束回调，AutoGo 不需要此功能

29. **LuaEngine.registerExitCallback** - 注册退出回调
    - **原因**：AutoGo 不支持 Java 层回调
    - **说明**：懒人精灵支持 Java 层回调，AutoGo 不需要此功能

30. **LuaEngine.setSnapCacheBitmap** - 设置截图缓存图片
    - **原因**：AutoGo 不支持截图缓存
    - **说明**：懒人精灵支持截图缓存以提高性能，AutoGo 不需要此功能

31. **LuaEngine.setSnapCacheTime** - 设置截图缓存时长
    - **原因**：AutoGo 不支持截图缓存
    - **说明**：懒人精灵支持截图缓存以提高性能，AutoGo 不需要此功能

### 插件和通知功能（3 个）

32. **setPluginEventCallBack** - 设置插件事件回调
    - **原因**：AutoGo 不支持插件系统
    - **说明**：懒人精灵支持插件系统，AutoGo 不需要此功能

33. **sendPluginEvent** - 发送插件事件
    - **原因**：AutoGo 不支持插件系统
    - **说明**：懒人精灵支持插件系统，AutoGo 不需要此功能

34. **setNotifyEventCallBack** - 设置系统通知事件回调
    - **原因**：AutoGo 不支持系统通知
    - **说明**：懒人精灵支持系统通知监听，AutoGo 不需要此功能

### 屏幕控制功能（1 个）

35. **unLockScreen** - 释放常亮状态
    - **原因**：AutoGo 不支持释放常亮状态
    - **说明**：AutoGo 的 KeepScreenOn 方法无法释放，需要重启应用

### 其他功能（5 个）

36. **setLogOff** - 关闭或开启日志输出
    - **原因**：AutoGo 不支持日志控制
    - **说明**：AutoGo 的日志输出由系统控制

37. **setBTEnable** - 打开或关闭蓝牙
    - **原因**：需要 `BLUETOOTH_ADMIN` 权限
    - **说明**：控制蓝牙需要 Android 高级权限，AutoGo 默认不请求此权限

38. **setWifiEnable** - 打开或关闭 WiFi
    - **原因**：需要 `CHANGE_WIFI_STATE` 权限
    - **说明**：控制 WiFi 需要 Android 高级权限，AutoGo 默认不请求此权限

39. **setAirplaneMode** - 打开或关闭飞行模式
    - **原因**：需要系统级权限
    - **说明**：控制飞行模式需要系统级权限，普通应用无法实现

40. **exec** - 以最高权限执行命令
    - **原因**：需要 root 权限
    - **说明**：执行 shell 命令需要 root 权限，AutoGo 默认不请求此权限

41. **phoneCall** - 拨号或直接打出电话
    - **原因**：需要 `CALL_PHONE` 权限
    - **说明**：拨打电话需要 Android 高级权限，AutoGo 默认不请求此权限

### Lua 标准库已有功能（1 个）

42. **xpcall** - 保护模式调用函数
    - **原因**：Lua 标准库已提供
    - **说明**：Lua 标准库已提供 xpcall 函数，无需额外实现

## 总结

device 模块提供了懒人精灵的设备相关方法，通过 AutoGo API 实现了大部分功能：

- **可用方法**：44 个（通过 AutoGo API 实现）
- **部分可用方法**：2 个（有使用限制）
- **空操作方法**：33 个（由于平台限制或权限问题）
- **标准库已有**：1 个（xpcall）
- **总计**：78 个方法

## 使用示例

### 示例 1：获取系统信息

```lua
-- 获取 CPU 架构
local arch = getCpuArch()
print("CPU 架构: " .. arch)

-- 获取 CPU 支持类型
local abi = getCpuAbi()
print("CPU ABI: " .. abi)

-- 获取工作目录
local workPath = getWorkPath()
print("工作目录: " .. workPath)

-- 获取设备品牌
local brand = getBrand()
print("设备品牌: " .. brand)

-- 获取设备型号
local model = getModel()
print("设备型号: " .. model)
```

### 示例 2：使用随机数和休眠

```lua
-- 生成随机数
for i = 1, 5 do
    local num = rnd(1, 100)
    print("随机数: " .. num)
    sleep(1000)  -- 休眠 1 秒
end
```

### 示例 3：应用管理

```lua
-- 打开应用
runApp("com.tencent.mm")

-- 获取前台应用包名
local frontApp = frontAppName()
print("前台应用: " .. frontApp)

-- 判断应用是否在前台
if appIsFront("com.tencent.mm") then
    print("微信在前台")
else
    print("微信不在前台")
end

-- 关闭应用
stopApp("com.tencent.mm")
```

### 示例 4：剪贴板操作

```lua
-- 写入剪贴板
writePasteboard("Hello, AutoGo!")

-- 读取剪贴板
local text = readPasteboard()
print("剪贴板内容: " .. text)
```

### 示例 5：调试输出

```lua
-- 使用 printEx 进行原生输出
printEx("这是原生输出")
printEx("调试信息: " .. os.time())
```

## 迁移指南

### 从懒人精灵迁移到 AutoGo

1. **保留可用的功能**：
   - 所有通过 AutoGo API 实现的功能（44 个）
   - 包括设备信息、应用管理、屏幕控制、剪贴板等

2. **谨慎使用部分可用功能**：
   - `getSdPath()` - 建议使用 `getWorkPath()` 替代
   - `restartScript()` - 建议使用 `os.exit()` 替代

3. **移除不可用的功能**：
   - 所有空操作方法（33 个）
   - 这些方法可以正常调用但不执行任何功能

4. **使用替代方案**：
   - 文件操作：使用 Lua 标准库的 `io` 和 `os` 模块
   - 网络请求：使用 `http` 模块
   - 图像处理：使用 `cv` 模块
   - 时间功能：使用 `time` 模块
   - 定时器：使用 `os.time()` 和 `sleep()` 组合实现
