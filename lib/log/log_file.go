package log

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFileAdapter create a logger adapter for file logger.
// eg: { "filepath": "../logs/example.log", "maxsize": 100, "maxage": 30, "maxbackups": 10, "compress": true }
// remark: 	filepath    日志文件路径, 为 empty 使用将采用 os.TempDir() 临时目录
//			maxsize    每个日志文件保存的最大尺寸，默认为 100。 单位：M
//			maxbackups 日志文件最多保存多少个备份
//			maxage     文件最多保存多长时间  单位：天
//			compress   是否压缩。默认不压缩文件。
func NewFileAdapter(cfg string) Instance {
	var conf map[string]interface{}
	if err := json.Unmarshal([]byte(cfg), &conf); err != nil {
		Panic(err)
	}

	c := Config{}
	if filepath, ok := conf["filepath"]; ok {
		c.Filepath = mustOtoa(filepath)
	}
	if maxsize, ok := conf["maxsize"]; ok {
		c.MaxSize = mustOtoi(maxsize)
	}
	if maxage, ok := conf["maxage"]; ok {
		c.MaxAge = mustOtoi(maxage)
	}
	if maxbackups, ok := conf["maxbackups"]; ok {
		c.MaxBackups = mustOtoi(maxbackups)
	}
	if compress, ok := conf["compress"]; ok {
		c.Compress = mustOtob(compress)
	}

	return NewFileAdapter2(c)
}

// NewFileAdapter2 create a logger adapter for file logger.
func NewFileAdapter2(cfg Config) Instance {
	logger := lumberjack.Logger{}

	logger.Filename = cfg.Filepath
	logger.MaxSize = cfg.MaxSize
	logger.MaxAge = cfg.MaxAge
	logger.MaxBackups = cfg.MaxBackups
	logger.Compress = cfg.Compress

	return func() zapcore.Core {
		return zapcore.NewCore(
			zapcore.NewJSONEncoder(DefaultEncoderConfig),
			zapcore.AddSync(&logger),
			DefaultLevelEnablerFunc(),
		)
	}
}

func mustOtoa(obj interface{}) string {
	switch v := obj.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		Panic(fmt.Errorf(`convert fail: can not convert "%v" to string`, obj))
	}

	return ""
}

func mustOtoi(obj interface{}) (ret int) {
	switch v := obj.(type) {
	case string:
		ret = mustAtoi(v)
	case []byte:
		ret = mustAtoi(string(v))
	case int:
		ret = v
	case int64:
		ret = int(v)
	default:
		Panic(fmt.Errorf(`convert fail: can not convert "%v" to int`, obj))
	}

	return ret
}

func mustOtob(obj interface{}) (ret bool) {
	switch v := obj.(type) {
	case string:
		ret = mustAtob(v)
	case []byte:
		ret = mustAtob(string(v))
	case bool:
		ret = v
	default:
		Panic(fmt.Errorf(`convert fail: can not convert "%v" to bool`, obj))
	}

	return ret
}

func mustAtoi(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		Panic(err)
	}

	return value
}

func mustAtob(str string) bool {
	value, err := strconv.ParseBool(str)
	if err != nil {
		Panic(err)
	}

	return value
}
