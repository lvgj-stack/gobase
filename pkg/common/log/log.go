package log

import (
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//只能输出结构化日志，但是性能要高于 SugaredLogger
var logger *zap.Logger

//可以输出 结构化日志、非结构化日志。性能茶语 zap.Logger，具体可见上面的的单元测试
var sugarLogger *zap.SugaredLogger

func init() {
	if err := InitLog("./info.log", "./error.log", zap.InfoLevel); err != nil {
		panic(err)
	}
}

// 初始化日志 logger
func InitLog(logPath, errPath string, logLevel zapcore.Level) error {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",                       //结构化（json）输出：msg的key
		LevelKey:     "level",                     //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      "ts",                        //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    "file",                      //结构化（json）输出：打印日志的文件对应的Key
		EncodeLevel:  zapcore.CapitalLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: zapcore.ShortCallerEncoder,  //采用短文件路径编码输出（test/main.go:14 ）
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, //输出的时间格式
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		}, //
	}
	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})
	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})
	// 获取io.Writer的实现
	infoWriter := getWriter(logPath)
	warnWriter := getWriter(errPath)
	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			logLevel,
		), //同时将日志输出到控制台，NewJSONEncoder 是结构化输出
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			zapcore.AddSync(infoWriter),
			infoLevel,
		), //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			zapcore.AddSync(warnWriter),
			warnLevel,
		), //warn及以上写入errPath
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	sugarLogger = logger.Sugar()
	return nil
}

func Info(msg string, fields ...interface{}) {
	field := zapFields(fields...)
	logger.Info(msg, field...)
}

func Warn(msg string, fields ...interface{}) {
	field := zapFields(fields...)
	logger.Warn(msg, field...)
}
func Error(msg string, fields ...interface{}) {
	field := zapFields(fields...)
	logger.Error(msg, field...)
}

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,    //最大M数，超过则切割
		MaxBackups: 5,     //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     30,    //保存30天
		Compress:   false, //是否压缩
	}
}

func zapFields(args ...interface{}) []zap.Field {
	if len(args) == 1 {
		fs, ok := args[0].([]interface{})
		if ok {
			args = fs
		}
	}
	fields := make([]zap.Field, 0, len(args)/2+1)
	for i := 0; i < len(args)-1; i += 2 {
		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			fields = append(fields, zap.Any("exceeds", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			fields = append(fields, zap.Any("invalidKey", val))
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
	}
	return fields
}
