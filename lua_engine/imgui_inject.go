package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/imgui"
	"github.com/Dasongzi1366/AutoGo/utils"
	lua "github.com/yuin/gopher-lua"
)

// injectImguiMethods 注入 imgui HUD 库方法
func injectImguiMethods(e *LuaEngine) {
	L := e.GetState()

	imguiTable := L.NewTable()

	L.SetField(imguiTable, "init", L.NewFunction(func(L *lua.LState) int {
		imgui.Init()
		return 0
	}))

	L.SetField(imguiTable, "close", L.NewFunction(func(L *lua.LState) int {
		imgui.Close()
		return 0
	}))

	L.SetField(imguiTable, "alert", L.NewFunction(func(L *lua.LState) int {
		title := L.CheckString(1)
		content := L.CheckString(2)
		btn1Text := ""
		btn2Text := ""
		if L.GetTop() >= 3 {
			btn1Text = L.CheckString(3)
		}
		if L.GetTop() >= 4 {
			btn2Text = L.CheckString(4)
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		L.Push(lua.LNumber(result))
		return 1
	}))

	L.SetField(imguiTable, "toast", L.NewFunction(func(L *lua.LState) int {
		message := L.CheckString(1)
		utils.Toast(message)
		return 0
	}))

	L.SetField(imguiTable, "drawRect", L.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		thickness := float32(1.0)
		if L.GetTop() >= 6 {
			thickness = float32(L.CheckNumber(6))
		}
		color := uint32(0xFFFFFFFF)
		if len(colorStr) > 0 {
			for i := 0; i < len(colorStr) && i < 8; i++ {
				c := colorStr[i]
				var val byte
				if c >= '0' && c <= '9' {
					val = byte(c - '0')
				} else if c >= 'a' && c <= 'f' {
					val = byte(c - 'a' + 10)
				} else if c >= 'A' && c <= 'F' {
					val = byte(c - 'A' + 10)
				}
				if i%2 == 0 {
					color |= uint32(val) << (28 - i*4)
				} else {
					color |= uint32(val) << (28 - i*4 + 4)
				}
			}
		}
		imgui.DrawRect(x1, y1, x2, y2, color, thickness)
		return 0
	}))

	L.SetGlobal("imgui", imguiTable)
}
