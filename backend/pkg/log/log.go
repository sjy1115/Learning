package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Log struct {
	Level         string `yaml:"level" json:"level"`
	LogPath       string `yaml:"log_path" json:"log_path"`
	DockerLogPath string `yaml:"docker_log_path" json:"docker_log_path"`
}

func (l *Log) Load() {
	l.Level = "DEBUG"
	l.LogPath = "./learning.log"
	l.DockerLogPath = "/workspace/logs/learning.log"
}

func InitLog(level, logPath string) {
	switch level {
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_path": logPath,
		}).Error("Failed to open log file")
		os.Exit(1)
	}

	logrus.SetOutput(file)
}
