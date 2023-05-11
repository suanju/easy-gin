package logrus

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

//JsonFormatter 自定义json 解析
type JsonFormatter struct {
	logrus.JSONFormatter
}

func (f *JsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 构造 JSON 数据
	info := &JsonInfo{
		Time:  entry.Time.Format("2006.01.02 15:04:05"),
		Level: entry.Level.String(),
		Msg:   entry.Message,
	}
	//只有等级匹配进行打印调用者信息
	if entry.Level == logrus.ErrorLevel || entry.Level == logrus.WarnLevel || entry.Level == logrus.DebugLevel || entry.Level == logrus.PanicLevel {
		info.File = entry.Caller.File
		info.Function = entry.Caller.Function
	}
	lineBreak := "\n"
	jsonData, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	formattedMsg := string(jsonData) + lineBreak
	return []byte(formattedMsg), nil
}
