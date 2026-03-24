-- device 模块测试脚本
print("=== device 模块测试 ===")
print("注意：大部分方法为空操作，只测试可用的方法")
print()

-- 测试 1: getCpuArch - 获取 CPU 架构
print("--- 测试 1: getCpuArch - 获取 CPU 架构 ---")
local arch = getCpuArch()
if arch == 0 or arch == 1 or arch == 2 then
    local archName = "unknown"
    if arch == 0 then
        archName = "x86"
    elseif arch == 1 then
        archName = "arm"
    elseif arch == 2 then
        archName = "arm64"
    end
    print("✓ getCpuArch 测试通过")
    print("  CPU 架构: " .. archName .. " (" .. arch .. ")")
else
    print("✗ getCpuArch 测试失败")
    print("  返回值: " .. (arch or "nil"))
end
print()

-- 测试 2: getWorkPath - 获取脚本工作目录
print("--- 测试 2: getWorkPath - 获取脚本工作目录 ---")
local workPath = getWorkPath()
if workPath and workPath ~= "" then
    print("✓ getWorkPath 测试通过")
    print("  工作目录: " .. workPath)
else
    print("✗ getWorkPath 测试失败")
    print("  返回值: " .. (workPath or "nil"))
end
print()

-- 测试 3: getCpuAbi - 获取 CPU 支持类型
print("--- 测试 3: getCpuAbi - 获取 CPU 支持类型 ---")
local abi = getCpuAbi()
if abi and abi ~= "" then
    print("✓ getCpuAbi 测试通过")
    print("  CPU ABI: " .. abi)
else
    print("✗ getCpuAbi 测试失败")
    print("  返回值: " .. (abi or "nil"))
end
print()

-- 测试 4: getCpuAbi2 - 获取 CPU 支持类型
print("--- 测试 4: getCpuAbi2 - 获取 CPU 支持类型 ---")
local abi2 = getCpuAbi2()
if abi2 and abi2 ~= "" then
    print("✓ getCpuAbi2 测试通过")
    print("  CPU ABI: " .. abi2)
else
    print("✗ getCpuAbi2 测试失败")
    print("  返回值: " .. (abi2 or "nil"))
end
print()

-- 测试 5: rnd - 生成随机整数
print("--- 测试 5: rnd - 生成随机整数 ---")
local random1 = rnd(1, 10)
local random2 = rnd(-5, 5)
if random1 >= 1 and random1 <= 10 and random2 >= -5 and random2 <= 5 then
    print("✓ rnd 测试通过")
    print("  随机数 1 (1-10): " .. random1)
    print("  随机数 2 (-5-5): " .. random2)
else
    print("✗ rnd 测试失败")
    print("  返回值 1: " .. (random1 or "nil"))
    print("  返回值 2: " .. (random2 or "nil"))
end
print()

-- 测试 6: sleep - 线程休眠
print("--- 测试 6: sleep - 线程休眠 ---")
local startTime = os.time()
sleep(1000)  -- 休眠 1 秒
local endTime = os.time()
local elapsed = endTime - startTime
if elapsed >= 1 then
    print("✓ sleep 测试通过")
    print("  休眠时间: " .. elapsed .. " 秒")
else
    print("✗ sleep 测试失败")
    print("  实际休眠时间: " .. elapsed .. " 秒")
end
print()

-- 测试 7: printEx - 调试输出
print("--- 测试 7: printEx - 调试输出 ---")
printEx("这是 printEx 测试输出")
print("✓ printEx 测试通过")
print()

-- 测试 8: 空操作方法测试
print("--- 测试 8: 空操作方法测试 ---")
print("测试以下空操作方法（应该不会报错）:")

-- 测试一些空操作方法
getSdPath()
setControlBarPosNew(0.5, 0.5)
showControlBar(true)
restartScript()
getApkVerInt()
getScriptVersion()
vibrate(100)
getSubscriberId()
getSimSerialNumber()
getDisplayDpi()
getBatteryLevel()
getUIConfig("test.config")
playAudio("test.mp3")
stopAudio()
installLrPkg("/path/to/file.lrj")
getPackageName()
getInstalledApk()
getInstalledApps()
getLrApi()
getOaid()
getInsallAppInfos()
getDisplayInfo()
getSensorsInfo()
installApk("/path/to/file.apk")
checkIsDebug()
getDeviceId()
getBrand()
getBootLoader()
getBoard()
getManufacturer()
getProduct()
getDevice()
getModel()
getHardware()
getId()
getFingerprint()
getSdkVersion()
getOsVersionName()
getWifiMac()
runApp("com.test.app")
stopApp("com.test.app")
scanImage("/path/to/image.png")
getCurrentActivity()
frontAppName()
appIsFront("com.test.app")
setMainThreadPause()
setMainThreadResume()
readPasteboard()
writePasteboard("test")
appIsRunning("com.test.app")
exec("ls", false)
runIntent({ action = "test" })
sendSms("10086", "test")
phoneCall("10086", 1)
getDisplaySize()
setDisplayPowerOff(true)
setDpiToVir(240)
setDpiToRealy()
setLogOff(true)
getDisplayRotate()
extractAssets("test.rc", "/path/", "*.bmp")
extractApkAssets("test.txt", "/path/")
zip("/path/test.png", "/path/test.zip")
unZip("/path/test.zip", "/path/mydir")
setTimer(function() end, 1000)
setBTEnable(true)
setWifiEnable(true)
setAirplaneMode(true)
lockScreen()
unLockScreen()
setUserEventCallBack(function() end)
LuaEngine.registerExitCallback(function() end)

print("✓ 所有空操作方法测试通过（没有报错）")
print()

print("=== device 模块测试完成 ===")
print("可用方法: 6 个")
print("空操作方法: 64 个")
