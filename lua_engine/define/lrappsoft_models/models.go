package lrappsoft_models

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/accessibility"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/console"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/crypt"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/cv"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/device"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/dynamicui"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/extension"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/ffi"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/gobridge"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/http"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/image"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/imgui"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/json"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/lfs"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/math"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/network"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/node"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/string_module"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/time"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/touch"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/ui"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/virtualscreen"
)

// LrappsoftModules 包含所有 lrappsoft 模块的数组
var LrappsoftModules = []model.Module{
	&accessibility.AccessibilityModule{},
	&console.ConsoleModule{},
	&crypt.CryptModule{},
	&cv.CvModule{},
	&device.DeviceModule{},
	&dynamicui.DynamicUIModule{},
	&extension.ExtensionModule{},
	&ffi.FfiModule{},
	&gobridge.GoBridgeModule{},
	&http.HttpModule{},
	&image.ImageModule{},
	&imgui.ImGuiModule{},
	&json.JsonModule{},
	&lfs.LfsModule{},
	&math.MathModule{},
	&network.NetworkModule{},
	&node.NodeModule{},
	&string_module.StringModule{},
	&time.TimeModule{},
	&touch.TouchModule{},
	&ui.UIModule{},
	&virtualscreen.VirtualScreenModule{},
}
