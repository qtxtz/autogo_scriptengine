package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/imgui"
	"github.com/Dasongzi1366/AutoGo/utils"
	"github.com/dop251/goja"
	"strconv"
)

// injectImguiMethods 注入 imgui HUD 库方法
func injectImguiMethods(e *JSEngine) {
	vm := e.vm

	imguiObj := vm.NewObject()

	imguiObj.Set("init", func(call goja.FunctionCall) goja.Value {
		imgui.Init()
		return goja.Undefined()
	})

	imguiObj.Set("close", func(call goja.FunctionCall) goja.Value {
		imgui.Close()
		return goja.Undefined()
	})

	imguiObj.Set("alert", func(call goja.FunctionCall) goja.Value {
		title := call.Argument(0).String()
		content := call.Argument(1).String()
		btn1Text := ""
		btn2Text := ""
		if len(call.Arguments) >= 3 {
			btn1Text = call.Argument(2).String()
		}
		if len(call.Arguments) >= 4 {
			btn2Text = call.Argument(3).String()
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		return vm.ToValue(result)
	})

	imguiObj.Set("toast", func(call goja.FunctionCall) goja.Value {
		message := call.Argument(0).String()
		utils.Toast(message)
		return goja.Undefined()
	})

	imguiObj.Set("drawRect", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		colorStr := call.Argument(4).String()
		thickness := float32(1.0)
		if len(call.Arguments) >= 6 {
			thickness = float32(call.Argument(5).ToFloat())
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
		return goja.Undefined()
	})

	vm.Set("imgui", imguiObj)
}

func parseColorString(colorStr string) uint32 {
	if len(colorStr) == 0 {
		return 0xFFFFFFFF
	}
	if colorStr[0] == '#' {
		colorStr = colorStr[1:]
	}
	color, err := strconv.ParseUint(colorStr, 16, 32)
	if err != nil {
		return 0xFFFFFFFF
	}
	return uint32(color)
}
