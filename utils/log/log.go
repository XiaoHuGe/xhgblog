package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"xhgblog/utils/setting"
)

var Logrus *logrus.Logger

func SetUp() {
	Logrus = logrus.New()
	Logrus.SetReportCaller(true)

	Logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   true,
		FullTimestamp:   true,
	})
	// Can be any io.Writer, see below for File example
	//Log.SetOutput(os.Stdout)
	logPath := setting.AppSetting.Application.LogPath
	fmt.Println("logPath", logPath)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logrus.Out = file
	} else {
		Logrus.Info("Failed to log to file, using default stderr", err)
		Logrus.Out = os.Stdout
	}
	// Only log the warning severity or above.
	Logrus.SetLevel(logrus.DebugLevel)
}
