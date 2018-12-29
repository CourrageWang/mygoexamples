package main
import (
	"encoding/json"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	// 日志地址 "out.log" 自定义
	lp := Conf.Common.LogPath
	// 日志级别 DEBUG,ERROR, INFO
	lv := Conf.Common.LogLevel
	// 是否 DEBUG
	isDebug := true
	if Conf.Common.IsDebug != true {
		isDebug = false
	}
	initLogger(lp, lv, isDebug)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}

func initLogger(lp string, lv string, isDebug bool) {
	var js string
	if isDebug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, lv)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, lv, lp, lp)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}

func main() {
	InitLogger("out.log", "DEBUG", false)
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	Log.Info("info:", zap.String("s", s[0]))
	Log.Error("info:", zap.String("s", s[1]))
	Log.Debug("info:", zap.String("s", s[2]))
	Log.Fatal("info:", zap.String("s", s[3]))

}
