package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/console"
	"github.com/dop251/goja"
)

func injectConsoleMethods(engine *JSEngine) {
	vm := engine.GetVM()

	consoleObj := vm.NewObject()
	vm.Set("console", consoleObj)

	consoleObj.Set("new", func(call goja.FunctionCall) goja.Value {
		c := console.New()
		return vm.ToValue(c)
	})

	consoleObj.Set("println", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		var args []any
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		c.Println(args...)
		return goja.Undefined()
	})

	consoleObj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		size := int(call.Argument(1).ToInteger())
		c.SetTextSize(size)
		return goja.Undefined()
	})

	consoleObj.Set("setTextColor", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		color := call.Argument(1).String()
		c.SetTextColor(color)
		return goja.Undefined()
	})

	consoleObj.Set("setWindowSize", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		width := int(call.Argument(1).ToInteger())
		height := int(call.Argument(2).ToInteger())
		c.SetWindowSize(width, height)
		return goja.Undefined()
	})

	consoleObj.Set("setWindowPosition", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		x := int(call.Argument(1).ToInteger())
		y := int(call.Argument(2).ToInteger())
		c.SetWindowPosition(x, y)
		return goja.Undefined()
	})

	consoleObj.Set("setWindowColor", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		color := call.Argument(1).String()
		c.SetWindowColor(color)
		return goja.Undefined()
	})

	consoleObj.Set("show", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Show()
		return goja.Undefined()
	})

	consoleObj.Set("hide", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Hide()
		return goja.Undefined()
	})

	consoleObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Clear()
		return goja.Undefined()
	})

	consoleObj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		visible := c.IsVisible()
		return vm.ToValue(visible)
	})

	consoleObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Destroy()
		return goja.Undefined()
	})

	engine.RegisterMethod("console.new", "创建控制台对象", console.New, true)
}
