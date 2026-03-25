package safe_models

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/app"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/coroutine"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/device"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/dotocr"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/files"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/http"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/images"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/ime"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/media"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/motion"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/opencv"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/plugin"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/ppocr"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/rhino"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/storages"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/system"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/uiacc"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/utils"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo/yolo"
)

// SafeModules 包含安全的模块（排除 console、hud、imgui、vdisplay）
// 这四个模块在 Android16 下会出现不安全的内存访问报错
var SafeModules = []model.Module{
	&app.AppModule{},
	&coroutine.CoroutineModule{},
	&device.DeviceModule{},
	&dotocr.DotocrModule{},
	&files.FilesModule{},
	&http.HttpModule{},
	&images.ImagesModule{},
	&ime.ImeModule{},
	&media.MediaModule{},
	&motion.MotionModule{},
	&opencv.OpencvModule{},
	&plugin.PluginModule{},
	&ppocr.PpocrModule{},
	&rhino.RhinoModule{},
	&storages.StoragesModule{},
	&system.SystemModule{},
	&uiacc.UiaccModule{},
	&utils.UtilsModule{},
	&yolo.YoloModule{},
}
