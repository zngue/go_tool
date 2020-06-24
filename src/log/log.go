package log

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)
var Log *zap.Logger
func  init()  {
	Log ,_=zap.NewProduction()
}
func LogInfo(message interface{})  {
	defer func() {
		if r:=recover();r!=nil{
			LogInfo(r)
		}
	}()
	go Log.Info(FormatLog(message))
}
func FormatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
		} else {
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

