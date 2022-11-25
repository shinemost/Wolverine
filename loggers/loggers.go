package loggers

import (
	"net/http"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func MainMethod() {
	initLogger()
	// 在程序退出之前把缓冲区里的文件都刷到硬盘上
	defer logger.Sync()

	// for i := 0; i < 10000000; i++ {
	httpGet("https://www.baidu.com")
	httpGet("https://www.google.com")

	// logger.Info("i:")
	// }

}

func initLogger() {
	// 这里还可以New Develp环境的日志 只是输出格式不太一样
	// logger, _ = zap.NewProduction()
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	logger = zap.New(core, zap.AddCaller())
}

// 返回一个json编码的格式 一般情况下都用这个
func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(config)
}

// 就是将日志写到何处
func getLogWriter() zapcore.WriteSyncer {
	// 这里我们每次都是重新打开文件，你也可以用open和append来追加文件，免的每次写入文件 都会覆盖掉之前的
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching url..", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("Success..", zap.String("statusCode", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}

func getLogWriterByJack() zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,   // 单位是mb
		MaxBackups: 5,    // 备份数量 切割文件之前 会把文件做备份
		MaxAge:     30,   // 备份天数
		Compress:   true, //默认不压缩
	}
	return zapcore.AddSync(logger)
}
