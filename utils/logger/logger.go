package logger

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type StandardLogger struct {
	logPackage *logrus.Logger
}

type FieldMap map[string]interface{}

type Event struct {
	level   int
	Message string
	Fields  FieldMap
	args    []interface{}
	stdLog  *StandardLogger
}

type LogEvent interface {
	Log(string)
	Logf(string, ...interface{})
	AddField(string, interface{}) LogEvent
}

type Logger interface {
	NewEvent(int) LogEvent
	ParseLevel(string) int
}

const (
	Fatal = iota + 1
	Error
	Warn
	Info
	Debug
)

var loggerLevel int

var Log *StandardLogger

func init() {
	Log = NewLogger()
	loggerLevel = Error
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	baseLogger := logrus.New()
	standardLogger := &StandardLogger{baseLogger}
	standardLogger.logPackage.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	standardLogger.logPackage.SetOutput(os.Stdout)
	standardLogger.logPackage.SetLevel(logrus.DebugLevel)
	return standardLogger
}

// SetJSON changes output to machine parsable JSON format
func SetJSON() {
	Log.logPackage.SetFormatter(&logrus.JSONFormatter{})
}

func SetLevel(level int) {
	if level < Fatal || level > Debug {
		return
	}
	loggerLevel = level
}

// WriteLogOnPath writes log on given path, log files will be rotate by constant time cycle
func (l *StandardLogger) WriteLogOnPath(path string) {
	writer, _ := rotateLogs.New(
		path+".%Y%m%d%H%M",
		rotateLogs.WithLinkName(path),
		rotateLogs.WithMaxAge(48*time.Hour),
		rotateLogs.WithRotationTime(time.Hour),
	)
	l.logPackage.SetOutput(writer)
}

func (event *Event) logEvent() {
	switch event.level {
	case Fatal:
		event.fatal()
	case Error:
		event.error()
	case Warn:
		event.warn()
	case Info:
		event.info()
	case Debug:
		event.debug()
	default:
		return
	}
}

func (event Event) fatal() {
	if loggerLevel < Fatal {
		return
	}
	l := event.stdLog
	if event.Fields != nil {
		if event.args != nil {
			l.logPackage.WithFields(logrus.Fields(event.Fields)).Fatalf(event.Message, event.args...)
		}
		l.logPackage.WithFields(logrus.Fields(event.Fields)).Fatal(event.Message)
	} else {
		if event.args != nil {
			l.logPackage.Fatalf(event.Message, event.args...)
		} else {
			l.logPackage.Fatal(event.Message)
		}
	}
}

func (event Event) error() {
	if loggerLevel < Error {
		return
	}
	l := event.stdLog
	if event.Fields != nil {
		if event.args != nil {
			l.logPackage.WithFields(logrus.Fields(event.Fields)).Errorf(event.Message, event.args...)
		}
		l.logPackage.WithFields(logrus.Fields(event.Fields)).Error(event.Message)
	} else {
		if event.args != nil {
			l.logPackage.Errorf(event.Message, event.args...)
		} else {
			l.logPackage.Error(event.Message)
		}
	}
}

func (event Event) warn() {
	l := event.stdLog
	if event.Fields != nil {
		if event.args != nil {
			l.logPackage.WithFields(logrus.Fields(event.Fields)).Warnf(event.Message, event.args...)
		}
		l.logPackage.WithFields(logrus.Fields(event.Fields)).Warn(event.Message)
	} else {
		if event.args != nil {
			l.logPackage.Warnf(event.Message, event.args...)
		} else {
			l.logPackage.Warn(event.Message)
		}
	}
}

func (event Event) info() {
	if loggerLevel < Warn {
		return
	}
	l := event.stdLog
	if event.Fields != nil {
		if event.args != nil {
			l.logPackage.WithFields(logrus.Fields(event.Fields)).Infof(event.Message, event.args...)
		}
		l.logPackage.WithFields(logrus.Fields(event.Fields)).Info(event.Message)
	} else {
		if event.args != nil {
			l.logPackage.Infof(event.Message, event.args...)
		} else {
			l.logPackage.Info(event.Message)
		}
	}
}

func (event Event) debug() {
	if loggerLevel < Debug {
		return
	}
	l := event.stdLog
	if event.Fields != nil {
		if event.args != nil {
			l.logPackage.WithFields(logrus.Fields(event.Fields)).Debugf(event.Message, event.args...)
		}
		l.logPackage.WithFields(logrus.Fields(event.Fields)).Debug(event.Message)
	} else {
		if event.args != nil {
			l.logPackage.Debugf(event.Message, event.args...)
		} else {
			l.logPackage.Debug(event.Message)
		}
	}
}

func (l *StandardLogger) NewEvent(lvl int) LogEvent {
	if lvl < Fatal || lvl > Debug {
		return nil
	}
	event := new(Event)
	event.level = lvl
	event.stdLog = l
	return event
}

func (event *Event) Log(message string) {
	event.Message = message
	event.logEvent()
}

func (event *Event) Logf(message string, args ...interface{}) {
	event.Message = message
	event.args = args
	event.logEvent()
}

func (event *Event) AddField(field string, value interface{}) LogEvent {
	if event.Fields == nil {
		event.Fields = make(FieldMap)
	}
	event.Fields[field] = value
	return event
}

//TODO Adicionar um hash no StandadLogger para evitar comparação de string
func (StandardLogger) ParseLevel(level string) int {
	level = strings.ToLower(level)
	switch level {
	case "fatal":
		return Fatal
	case "error":
		return Error
	case "warn":
		return Warn
	case "info":
		return Info
	case "debug":
		return Debug
	default:
		return Error
	}
}

func SetFormat(format string) {
	format = strings.ToLower(format)
	switch format {
	case "terminal":
		return
	case "json":
		SetJSON()
		return
	default:
		return
	}
}
