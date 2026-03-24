# lua_engine.go


### var (
GetLuaEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）

### func GetLuaEngine() *LuaEngine {
GetEngine 获取默认引擎实例（GetLuaEngine 的别名）

### func GetEngine() *LuaEngine {
NewLuaEngine 创建新的引擎实例
config: 引擎配置，传入 nil 使用默认配置

### func NewLuaEngine(config *EngineConfig) *LuaEngine {
NewEngine 创建新的引擎实例（NewLuaEngine 的别名）

### func NewEngine(config *EngineConfig) *LuaEngine {

### func (e *LuaEngine) init() {
初始化模块缓存
设置模块搜索路径
注册自定义 require 函数，支持从 embed.FS 加载模块
setupPackagePath 设置 Lua 的模块搜索路径

### func (e *LuaEngine) setupPackagePath() {
获取当前的 package.path
添加配置的搜索路径
设置新的 package.path

### func (e *LuaEngine) GetState() *lua.LState {

### func (e *LuaEngine) registerCoreFunctions() {
registerCustomRequire 注册自定义 require 函数，支持从 embed.FS 加载模块

### func (e *LuaEngine) registerCustomRequire() {
保存原始 require 函数
注册自定义 require 函数
检查缓存
尝试从 embed.FS 加载模块
缓存模块
如果 embed.FS 中没有找到，使用原始 require
返回 nil 表示模块未找到
loadModuleFromFS 从 embed.FS 加载模块

### func (e *LuaEngine) loadModuleFromFS(L *lua.LState, moduleName string) (lua.LValue, bool) {
尝试不同的路径模式
在所有搜索路径中查找

### var fullPath string
如果搜索路径是 "."，直接使用文件名（表示文件系统根目录）
使用 LoadString 加载代码，然后调用以获取返回值
调用函数并获取返回值
获取模块的返回值（模块应该返回一个 table）

### func (e *LuaEngine) consoleLogLua(L *lua.LState) int {

### func (e *LuaEngine) consoleErrorLua(L *lua.LState) int {

### func (e *LuaEngine) injectAllMethods() {
检查白名单
检查黑名单
检查模块是否可用
注册模块
InjectModule 注入指定模块的方法
module: 模块名称，支持: app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui

### func (e *LuaEngine) InjectModule(moduleName string) {
InjectModules 注入多个模块的方法

### func (e *LuaEngine) InjectModules(modules []string) {
GetAvailableModules 获取所有可用模块列表

### func (e *LuaEngine) GetAvailableModules() []string {
RegisterModule 注册一个或多个模块到当前引擎实例
用户可以在自己的代码中调用此方法来注册需要的模块
支持可变长参数，可以一次注册多个模块

### func (e *LuaEngine) RegisterModule(modules ...model.Module) {
InjectAllMethods 注入所有方法（公开方法，允许手动调用）

### func (e *LuaEngine) InjectAllMethods() {

### func (e *LuaEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
ExecuteString 执行 Lua 代码字符串（实例方法）
script: 要执行的 Lua 代码
searchPaths: 可选参数，添加模块搜索路径（用于 require）
支持脚本退出后的动作：
- os.exit(0): 正常退出，执行配置的退出动作（重启/自定义/无动作）
- os.exit(-1): 强制退出，不执行任何退出动作
- os.exit(其他值): 正常退出，执行配置的退出动作
脚本异常退出时始终打印日志

### func (e *LuaEngine) ExecuteString(script string, searchPaths ...string) error {
如果脚本异常退出，打印错误日志
如果跳过退出动作（os.exit(-1)），直接返回
根据配置的退出动作执行相应操作
无动作，直接退出
重启脚本
继续循环，重新执行脚本
执行自定义退出动作
executeStringOnce 执行一次 Lua 代码字符串

### func (e *LuaEngine) executeStringOnce(script string, searchPaths ...string) error {
如果提供了搜索路径，添加到 package.path
注册特殊的 os.exit 函数，用于控制退出动作
addSearchPaths 添加模块搜索路径

### func (e *LuaEngine) addSearchPaths(paths ...string) {
只更新 config.SearchPaths，不更新 package.path
因为我们使用自定义的 require 函数，它会从 config.SearchPaths 和文件系统中加载模块
避免重复
ExecuteFile 执行 Lua 文件
path: 要执行的 Lua 文件路径
支持脚本退出后的动作：
- os.exit(0): 正常退出，执行配置的退出动作（重启/自定义/无动作）
- os.exit(-1): 强制退出，不执行任何退出动作
- os.exit(其他值): 正常退出，执行配置的退出动作
脚本异常退出时始终打印日志

### func (e *LuaEngine) ExecuteFile(path string) error {
如果脚本异常退出，打印错误日志
如果跳过退出动作（os.exit(-1)），直接返回
根据配置的退出动作执行相应操作
无动作，直接退出
重启脚本
继续循环，重新执行脚本
执行自定义退出动作
executeFileOnce 执行一次 Lua 文件

### func (e *LuaEngine) executeFileOnce(path string) error {
如果配置了文件系统，从文件系统读取并执行
读取文件内容
自动检测文件所在目录并添加到搜索路径
注册特殊的 os.exit 函数，用于控制自动重启
执行文件内容
自动检测文件所在目录并添加到搜索路径
注册特殊的 os.exit 函数，用于控制自动重启
addSearchPathsFromPath 从文件路径中提取目录并添加到搜索路径

### func (e *LuaEngine) addSearchPathsFromPath(path string) {
提取目录（去掉文件名）

### var dir string
如果路径中没有目录分隔符，说明是相对路径
如果配置了文件系统，使用当前目录（.）
否则使用空字符串

### func (e *LuaEngine) Close() {
registerExitControl 注册特殊的 os.exit 函数，用于控制退出动作
os.exit(0) - 正常退出，执行配置的退出动作（重启/自定义/无动作）
os.exit(-1) - 强制退出，不执行任何退出动作
os.exit(其他值) - 正常退出，执行配置的退出动作

### func (e *LuaEngine) registerExitControl() {
获取 os 表
保存原始的 os.exit 函数
注册新的 os.exit 函数
如果退出码为 -1，跳过退出动作
调用原始的 os.exit 函数
Restart 重启 Lua 引擎
关闭当前状态并重新初始化，保留模块缓存

### func (e *LuaEngine) Restart() error {
保存模块缓存
关闭当前状态
重新初始化状态
恢复模块缓存
重新设置模块搜索路径
重新注册核心函数
重新注册自定义 require 函数

### func (e *LuaEngine) GetRegistry() *MethodRegistry {
ExecuteString 执行 Lua 代码字符串（全局函数）
script: 要执行的 Lua 代码
searchPaths: 可选参数，添加模块搜索路径（用于 require）

### func ExecuteString(script string, searchPaths ...string) error {

### func ExecuteFile(path string) error {

### func Close() {

### func (e *LuaEngine) registerMethodLua(L *lua.LState) int {

### func (e *LuaEngine) unregisterMethodLua(L *lua.LState) int {

### func (e *LuaEngine) listMethodsLua(L *lua.LState) int {

### func (e *LuaEngine) overrideMethodLua(L *lua.LState) int {

### func (e *LuaEngine) restoreMethodLua(L *lua.LState) int {

### func (e *LuaEngine) sleepLua(L *lua.LState) int {