# js_engine.go


### var (
GetJSEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）

### func GetJSEngine() *JSEngine {
GetEngine 获取默认引擎实例（GetJSEngine 的别名）

### func GetEngine() *JSEngine {
NewJSEngine 创建新的引擎实例
config: 引擎配置，传入 nil 使用默认配置

### func NewJSEngine(config *EngineConfig) *JSEngine {
NewEngine 创建新的引擎实例（NewJSEngine 的别名）

### func NewEngine(config *EngineConfig) *JSEngine {

### func (e *JSEngine) init() {

### func (e *JSEngine) GetVM() *goja.Runtime {

### func (e *JSEngine) registerCoreFunctions() {
注册 require 功能

### func (e *JSEngine) consoleLogJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) consoleErrorJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) loadJS(call goja.FunctionCall) goja.Value {
读取文件内容

### var content []byte

### var err error
使用配置的文件系统
使用操作系统的文件系统
从路径中提取目录和文件名
保存当前的 __dirname 和 __filename
设置新的 __dirname 和 __filename
执行文件内容
恢复原来的 __dirname 和 __filename

### func (e *JSEngine) injectAllMethods() {
检查白名单
检查黑名单
检查模块是否可用
注册模块
InjectModule 注入指定模块的方法
module: 模块名称，支持: app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui

### func (e *JSEngine) InjectModule(moduleName string) {
InjectModules 注入多个模块的方法

### func (e *JSEngine) InjectModules(modules []string) {
GetAvailableModules 获取所有可用模块列表

### func (e *JSEngine) GetAvailableModules() []string {
RegisterModule 注册一个或多个模块到当前引擎实例
用户可以在自己的代码中调用此方法来注册需要的模块
支持可变长参数，可以一次注册多个模块

### func (e *JSEngine) RegisterModule(modules ...model.Module) {
InjectAllMethods 注入所有方法（公开方法，允许手动调用）

### func (e *JSEngine) InjectAllMethods() {

### func (e *JSEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
ExecuteString 执行 JavaScript 代码字符串
script: 要执行的 JavaScript 代码
dir: 可选参数，指定 __dirname（用于 require），如果为空则使用默认值 "scripts"
支持脚本退出后的动作：
- process.exit(0): 正常退出，执行配置的退出动作（重启/自定义/无动作）
- process.exit(-1): 强制退出，不执行任何退出动作
- process.exit(其他值): 正常退出，执行配置的退出动作

脚本异常退出时始终打印日志

### func (e *JSEngine) ExecuteString(script string, dir ...string) error {
如果脚本异常退出，打印错误日志
如果跳过退出动作（process.exit(-1)），直接返回
根据配置的退出动作执行相应操作
无动作，直接退出
重启脚本
继续循环，重新执行脚本
执行自定义退出动作
executeStringOnce 执行一次 JavaScript 代码字符串

### func (e *JSEngine) executeStringOnce(script string, dir ...string) error {
如果配置了文件系统且 __dirname 未设置，设置 __dirname
注册特殊的 process.exit 函数，用于控制退出动作
使用 IIFE 包装脚本，避免全局作用域污染

### func (e *JSEngine) ExecuteFile(path string) error {
使用 load 函数来加载文件
registerExitControl 注册特殊的 process.exit 函数，用于控制退出动作
process.exit(0) - 正常退出，执行配置的退出动作（重启/自定义/无动作）
process.exit(-1) - 强制退出，不执行任何退出动作
process.exit(其他值) - 正常退出，执行配置的退出动作

### func (e *JSEngine) registerExitControl() {
获取 process 对象
如果 process 对象不存在，创建一个
如果 process 不是对象，无法注册退出控制
保存原始的 exit 函数
注册新的 exit 函数
如果退出码为 -1，跳过退出动作
调用原始的 exit 函数

### func (e *JSEngine) Close() {

### func (e *JSEngine) GetRegistry() *MethodRegistry {
ExecuteString 执行 JavaScript 代码字符串（全局函数）
script: 要执行的 JavaScript 代码
dir: 可选参数，指定 __dirname（用于 require），如果为空则使用默认值 "scripts"

### func ExecuteString(script string, dir ...string) error {

### func ExecuteFile(path string) error {

### func Close() {

### func (e *JSEngine) registerMethodJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) unregisterMethodJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) listMethodsJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) overrideMethodJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) restoreMethodJS(call goja.FunctionCall) goja.Value {

### func (e *JSEngine) sleepJS(call goja.FunctionCall) goja.Value {