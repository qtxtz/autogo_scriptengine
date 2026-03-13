package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/plugin"
	lua "github.com/yuin/gopher-lua"
)

func injectPluginMethods(engine *LuaEngine) {

	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)

	registerPluginLuaFunctions(engine)
}

func registerPluginLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("plugin_loadApk", func(L *lua.LState) int {
		apkPath := L.CheckString(1)
		cl := plugin.LoadApk(apkPath)
		ud := L.NewUserData()
		ud.Value = cl
		L.Push(ud)
		return 1
	})
}
