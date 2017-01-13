package bee

import (
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
)

var (
	Engine   *xorm.Engine
	OrderReg = regexp.MustCompile(`-([^,]+)`)
)

func SetDb(driver, datasource string) {
	logger := logs.GetLogger()
	var err error
	Engine, err = xorm.NewEngine(driver, datasource)
	if err != nil {
		logger.Fatalln("open database false", err)
	}
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)
}

const (
	applicationJSON = "application/json"
	applicationXML  = "application/xml"
	textXML         = "text/xml"
)

type apimsg struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ApiController struct {
	beego.Controller
	Db            *xorm.Session
	methodMapping map[string]func() (interface{}, error)
}

func (c *ApiController) Prepare() {
	c.Db = Engine.NewSession()
}

func (c *ApiController) Finish() {
	c.Db.Close()
	c.ServeFormatted()
}

// Init generates default values of controller operations.
func (c *ApiController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	c.methodMapping = make(map[string]func() (interface{}, error))
}

// HandlerFunc call function with the name
func (c *ApiController) HandlerFunc(fnname string) bool {
	if v, ok := c.methodMapping[fnname]; ok {
		var err error
		c.Data["data"], err = v()
		if err != nil {
			c.Data["msg"] = err.Error()
		}
		return true
	}
	return false
}

// Mapping the method to function
func (c *ApiController) Mapping(method string, fn func() (interface{}, error)) {
	c.methodMapping[method] = fn
}

func (c *ApiController) ServeFormatted() {
	accept := c.Ctx.Input.Header("Accept")
	var msg string
	if v := c.Data["msg"]; v != nil {
		msg = v.(string)
	}
	switch accept {
	case applicationJSON:
		c.Data["json"] = apimsg{msg, c.Data["data"]}
		c.ServeJSON()
	case applicationXML, textXML:
		c.Data["xml"] = apimsg{msg, c.Data["data"]}
		c.ServeXML()
	default:
		c.Data["json"] = apimsg{msg, c.Data["data"]}
		c.ServeJSON()
	}
}
