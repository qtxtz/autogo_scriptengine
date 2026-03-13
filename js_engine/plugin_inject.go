package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/plugin"
	"github.com/dop251/goja"
)

func injectPluginMethods(engine *JSEngine) {
	vm := engine.GetVM()

	pluginObj := vm.NewObject()
	vm.Set("plugin", pluginObj)

	pluginObj.Set("loadApk", func(call goja.FunctionCall) goja.Value {
		apkPath := call.Argument(0).String()
		cl := plugin.LoadApk(apkPath)
		return vm.ToValue(cl)
	})

	// 注册方法到文档
	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)
}
