package log1

import (
	//"fmt"
	"os"
	logger "github.com/donnie4w/go-logger/logger"
	)

func test() {
	logger.SetRollingFile("log", "mylog.log", 10, 5, logger.KB)
	logger.SetRollingDaily("log", "mylog.log")
	logger.SetLevel(logger.INFO)

	logger.Info("program start.PID[", os.Getpid(), "]")
}