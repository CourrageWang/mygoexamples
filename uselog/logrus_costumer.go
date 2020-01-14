package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

//参考：https://blog.csdn.net/sureSand/article/details/82909121

type lineHook struct {
	W      LoggerInterface
	field  string
	skip   int //遍历调用栈使用开始的索引位置
	levels []log.Level
}

func NewConetext(file string, levels ...log.Level) (*lineHook, error) {
	w := NewFileWriter()
	config := fmt.Sprintf(`{"filename":"%s","maxdays":7}`, file)
	err := w.Init(config)
	if err != nil {
		return nil, err
	}
	hook := lineHook{
		W:      w,
		field:  "line",
		skip:   2,
		levels: levels,
	}
	if len(levels) == 0 {
		hook.levels = log.AllLevels
	}
	return &hook, nil
}
//格式化日志输出
func (hook lineHook) Format(entry *log.Entry) ([]byte, error) {
	var s strings.Builder
	result := fmt.Sprintf("%s %s 【%s】: %s",
		entry.Time.Format("2006-01-02 15:04:05"),
		entry.Data[hook.field].(string),
		entry.Level.String(),
		entry.Message)
	s.WriteString(result)
	data := make(map[string]interface{}, len(entry.Data)-1)
	for key, value := range entry.Data {
		if key == hook.field {
			continue
		}
		data[key] = value
	}
	d, err := json.Marshal(data)
	if nil != err {
		d = []byte("")
	}
	s.WriteString(fmt.Sprintf("\n    data:%s\n", string(d)))
	return []byte(s.String()), nil
}

//定义那些等级的日志触发hook机制
func (hook lineHook) Levels() []log.Level {
	return log.AllLevels
}

//将异常日志写入到日志文件中
func (hook lineHook) Fire(entry *log.Entry) error {
	entry.Data[hook.field] = findCaller(hook.skip)

	message, err := hook.Format(entry)
	if nil != err {
		message = []byte("日志出现错误")
	}
	switch entry.Level {
	case log.PanicLevel:
		fallthrough
	case log.FatalLevel:
		fallthrough
	case log.ErrorLevel:
		return hook.W.WriteMsg(string(message), LevelError)
	case log.WarnLevel:
		return hook.W.WriteMsg(string(message), LevelWarn)
	case log.InfoLevel:
		return hook.W.WriteMsg(string(message), LevelInfo)
	case log.DebugLevel:
		return hook.W.WriteMsg(string(message), LevelDebug)
	default:
		return nil
	}
	return nil
}

func findCaller(skip int) string {
	file := ""
	line := 0
	var pc uintptr
	for i := 0; i < 11; i++ {
		file, line, pc = getCaller(skip + i)
		//过滤掉所有logrus包，得到生成代码信息
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	fullFnName := runtime.FuncForPC(pc)
	fnName := ""
	if fullFnName != nil {
		fnNameStr := fullFnName.Name()
		//获取函数名
		parts := strings.Split(fnNameStr, ".")
		fnName = parts[len(parts)-1]
	}
	return fmt.Sprintf("%s:%d:%s()", file, line, fnName)
}

func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}
	n := 0
	//获取包名
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, pc
}
