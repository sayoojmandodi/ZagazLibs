package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type logging struct {
	logger *logrus.Logger
}

// Log will be the instance to access logging struct
var Log logging

// NewLog ...
func NewLog() {
	Log = logging{}
	Log.logger = logrus.New()
}

// InitLog ...
func InitLog(appMode, filepath string) error {
	if appMode != "debug" {
		f, err := os.OpenFile(
			filepath,
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0666,
		)
		if err != nil {
			return err
		}
		Log.logger.SetOutput(f)
	} else {
		Log.logger.Level = logrus.DebugLevel
	}
	Log.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: true,
		DisableSorting:            false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableLevelTruncation:    false,
	})
	return nil
}

func (l *logging) Debug(message string, field ...map[string]interface{}) {

	fields := make(map[string]interface{}, len(field))
	for _, val := range field {
		for index2, val2 := range val {
			fields[index2] = val2
		}
	}
	l.logger.WithFields(fields).Error(message)
}

func (l *logging) Warning(message string, field ...map[string]interface{}) {

	fields := make(map[string]interface{}, len(field))
	for _, val := range field {
		for index2, val2 := range val {
			fields[index2] = val2
		}
	}
	l.logger.WithFields(fields).Warning(message)
}

func (l *logging) Error(message string, field ...map[string]interface{}) {

	fields := make(map[string]interface{}, len(field))
	for _, val := range field {
		for index2, val2 := range val {
			fields[index2] = val2
		}
	}
	l.logger.WithFields(fields).Error(message)
}

func (l *logging) Info(message string, field ...map[string]interface{}) {

	fields := make(map[string]interface{}, len(field))
	for _, val := range field {
		for index2, val2 := range val {
			fields[index2] = val2
		}
	}
	l.logger.WithFields(fields).Info(message)
}
