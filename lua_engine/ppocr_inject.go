package lua_engine

import (
	"image"

	"github.com/Dasongzi1366/AutoGo/ppocr"
	lua "github.com/yuin/gopher-lua"
)

func injectPpocrMethods(engine *LuaEngine) {

	engine.RegisterMethod("ppocr.new", "创建一个新的PPOCR实例", ppocr.New, true)

	registerPpocrLuaFunctions(engine)
}

func registerPpocrLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("ppocr_new", func(L *lua.LState) int {
		version := L.CheckString(1)
		p := ppocr.New(version)
		ud := L.NewUserData()
		ud.Value = p
		L.Push(ud)
		return 1
	})

	state.Register("ppocr_ocr", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		p := ud.Value.(*ppocr.Ppocr)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		colorStr := L.CheckString(6)
		displayId := 0
		if L.GetTop() >= 7 {
			displayId = L.CheckInt(7)
		}
		results := p.Ocr(x1, y1, x2, y2, colorStr, displayId)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_ocrFromImage", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		p := ud.Value.(*ppocr.Ppocr)
		imgUd := L.CheckUserData(2)
		img, ok := imgUd.Value.(*image.NRGBA)
		if !ok {
			L.Push(lua.LNil)
			return 1
		}
		colorStr := L.CheckString(3)
		results := p.OcrFromImage(img, colorStr)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_ocrFromBase64", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		p := ud.Value.(*ppocr.Ppocr)
		b64 := L.CheckString(2)
		colorStr := L.CheckString(3)
		results := p.OcrFromBase64(b64, colorStr)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_ocrFromPath", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		p := ud.Value.(*ppocr.Ppocr)
		path := L.CheckString(2)
		colorStr := L.CheckString(3)
		results := p.OcrFromPath(path, colorStr)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_close", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		p := ud.Value.(*ppocr.Ppocr)
		p.Close()
		return 0
	})
}

func pushResultsToLua(L *lua.LState, results []ppocr.Result) {
	tbl := L.NewTable()
	for i, result := range results {
		item := L.NewTable()
		L.SetField(item, "X", lua.LNumber(result.X))
		L.SetField(item, "Y", lua.LNumber(result.Y))
		L.SetField(item, "宽", lua.LNumber(result.Width))
		L.SetField(item, "高", lua.LNumber(result.Height))
		L.SetField(item, "标签", lua.LString(result.Label))
		L.SetField(item, "精度", lua.LNumber(result.Score))
		L.SetField(item, "CenterX", lua.LNumber(result.CenterX))
		L.SetField(item, "CenterY", lua.LNumber(result.CenterY))
		L.SetTable(tbl, lua.LNumber(i+1), item)
	}
	L.Push(tbl)
}
