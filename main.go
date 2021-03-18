package main

import (
	"flag"
	"fmt"

	"goframe/config"
	"goframe/driver/mysql"
	"goframe/module"
	"goframe/reload"
	"goframe/route"

	"github.com/imkouga/gocore/cfg/conf"
	"github.com/imkouga/gocore/gc"
	"github.com/imkouga/gocore/http/httpserver"
	"github.com/imkouga/gocore/loger"
)

const (
	server  = "goframe"
	version = "v1.0.0"
)

var (
	logLevel   int
	configFile string
)

func init() {
	flag.IntVar(&logLevel, "l", loger.LevelInfo, "Set log level Trace = 0, Debug = 1, Info = 2, Warn = 3, Error = 4, Critical = 5")
	flag.StringVar(&configFile, "c", "./etc/goframe.conf", "configer config file")
	flag.Parse()

	loger.SetLevel(logLevel)
	loger.SetLogFullFileLine(server)

}

func Init() {

	if err := config.Init(configFile); nil != err {
		loger.Critical(err)
	}

	if err := mysql.Init(); nil != err {
		loger.Critical(err)
	}

	if err := reload.Init(); nil != err {
		loger.Error(err)
	}

	if err := route.Init(); nil != err {
		loger.Critical(err)
	}

	if err := module.Init(); nil != err {
		loger.Critical(err)
	}
}

func StartHttpServer() {

	addr := conf.GetValueByStringCarryDefault("server", "http_host", "0.0.0.0")
	port := conf.GetValueByStringCarryDefault("server", "http_port", "0")

	if err := httpserver.HttpServerStart(fmt.Sprintf("%s:%s", addr, port)); nil != err {
		loger.Critical("main.StartHttpServer: ", err.Error())
		return
	}
	loger.Info("goframe http server running ~~")
	return
}

func main() {

	Init()
	StartHttpServer()

	loger.Info("goframe server is running ~~~")

	gc.AutoGC()
}
