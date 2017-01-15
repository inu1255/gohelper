package bee

import (
	"regexp"

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

type ApiMsg struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type SearchData struct {
	Total   int64         `json:"total"`
	Content []interface{} `json:"content"`
}
