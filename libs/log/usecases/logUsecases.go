package usecases

import (
	"default-repo/helpers"
	"default-repo/libs/log/models"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var LogEntry LogEntryInterface

const jsonFild = "jsonField"

type LogEntryInterface interface {
	OnInfo(args models.LogInfo)
	OnDebug(args models.LogDebug)
	OnError(args models.LogError)
	OnExternal(args models.LogExternal)
	OnQuery(args models.LogQuery)
}

type logEntry struct {
	log        *logrus.Logger
	defaultLog models.LogDefault
}

func NewLogEntry(level logrus.Level, defaultLog models.LogDefault) LogEntryInterface {
	log := logrus.New()
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableTimestamp:  true,
		PrettyPrint:       viper.GetBool("logRus.jsonFormatter"),
		DisableHTMLEscape: viper.GetBool("logRus.htmlEscape"),
	})
	log.SetOutput(os.Stdout)
	LogEntry = &logEntry{
		log:        log,
		defaultLog: defaultLog,
	}
	return LogEntry
}

func (l *logEntry) OnInfo(args models.LogInfo) {
	// set default log
	args.LogDefault = l.defaultLog
	args.FileLocation = helpers.SetFileLocation()
	args.Level = "info"

	l.log.WithField(jsonFild, args).Info()
}

func (l *logEntry) OnDebug(args models.LogDebug) {
	// set default log
	args.LogDefault = l.defaultLog
	args.FileLocation = helpers.SetFileLocation()
	args.Level = "debug"

	l.log.WithField(jsonFild, args).Debug()
}

func (l *logEntry) OnError(args models.LogError) {
	// set default log
	args.LogDefault = l.defaultLog
	args.FileLocation = helpers.SetFileLocation()
	args.Level = "error"

	l.log.WithField(jsonFild, args).Error()
}

func (l *logEntry) OnExternal(args models.LogExternal) {
	// set default log
	args.LogDefault = l.defaultLog
	args.FileLocation = helpers.SetFileLocation()
	args.Level = "info"

	l.log.WithField(jsonFild, args).Info()
}

func (l *logEntry) OnQuery(args models.LogQuery) {
	// set default log
	args.LogDefault = l.defaultLog
	args.FileLocation = helpers.SetFileLocation()
	args.Level = "info"

	l.log.WithField(jsonFild, args).Info()
}
