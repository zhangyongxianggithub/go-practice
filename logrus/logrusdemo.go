package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main1() {
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")
}

// 创建一个新的logger实例。可以创建任意多个。
var log = logrus.New()

func main() {
	// 设置日志输出为os.Stdout
	log.Out = os.Stdout

	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	log.WithTime(time.Now()).WithFields(logrus.Fields{
		"animal": "dog",
		"size":   10,
	}).Info("一群舔狗出现了。")
	log.WithFields(logrus.Fields{
		"animal": "dog",
		"size":   10,
	}).Info("一群舔狗出现了。")
}
