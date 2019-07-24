package jlog

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	FatalLevel   = "fatal"
	ErrorLevel   = "error"
	WarningLevel = "warning"
	InfoLevel    = "info"
	DebugLevel   = "debug"
	TraceLevel   = "trace"

	logLineNumber = true
	logCaller     = true
	logStack      = true
	logStackCount = 20
)

type LogConfig struct {
	logFilename  string
	logDirectory string
	logLevel     string
	logStdout    bool
	serviceLabel string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		logLevel:  TraceLevel,
		logStdout: true,
	}
}

func (c *LogConfig) SetLogFileOutput(logDirectory, logFilename string) *LogConfig {
	c.logDirectory = logDirectory
	c.logFilename = logFilename
	c.logStdout = false
	return c
}

func (c *LogConfig) SetLogLevel(logLevel string) *LogConfig {
	c.logLevel = logLevel
	return c
}

func (c *LogConfig) SetLogStdout(logStdout bool) *LogConfig {
	c.logStdout = logStdout
	return c
}

func (c* LogConfig) SetServiceLabel(serviceLabel string) *LogConfig {
	c.serviceLabel = serviceLabel
	return c
}

func Init(config *LogConfig) {
	if !config.logStdout && (config.logFilename == "" || config.logDirectory == "" || config.logLevel == "") {
		panic(fmt.Sprintf("InvalidConfig|logFilename:%s|logDirectory:%s|logLevel:%s",
			config.logFilename, config.logDirectory, config.logLevel))
	}
	logLevel, err := logrus.ParseLevel(config.logLevel)
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	filename := path.Join(config.logDirectory, config.logFilename)
	if config.logStdout {
		logrus.SetOutput(os.Stdout)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
		fmt.Println("Log would be written to: stdout")
	} else {
		logrus.SetOutput(&lumberjack.Logger{
			Filename: filename,
			// When log file size achieves to MaxSize, logrus will compression that.
			MaxSize: 100, // megabytes
			// When backups count larger than MaxBackups, logrus will remove the oldest backups.
			MaxBackups: 100,
			MaxAge:     60,   //days
			Compress:   true, // disabled by default
		})
		fmt.Println("Log would be written to: ", filename)
	}
	logrus.SetLevel(logLevel)
	logrus.AddHook(ContextHook{config})
}

type ContextHook struct {
	logConfig *LogConfig
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {

	skip := 6
FindCaller:
	for i := skip; i < 20; i++ {
		if _, file, _, ok := runtime.Caller(i); ok {
			filename := path.Base(file)
			// fmt.Println("", i, filename)
			switch filename {
			case "exported.go", "logger.go", "entry.go": // We locate these files, then next frame is what we want
			default:
				skip = i
				break FindCaller
			}
		}
	}

	var lineContent, callerContent, stackContent string
	for i := 0; i < logStackCount; i++ {
		if pc, file, line, ok := runtime.Caller(skip + i); ok {
			curStackContent := ""
			if logCaller {
				if len(callerContent) > 0 {
					callerContent += ", "
				}
				funcName := runtime.FuncForPC(pc).Name()
				curContent := path.Base(funcName)
				callerContent += curContent
				curStackContent += curContent + " "
			}

			if logLineNumber {
				if len(lineContent) > 0 {
					lineContent += ", "
				}
				curContent := fmt.Sprintf("%s:%v", path.Base(file), line)
				lineContent += fmt.Sprintf("(%s)", curContent)
				curStackContent += curContent
			}

			curStackContent = fmt.Sprintf("(%s)", curStackContent)
			if len(stackContent) > 0 {
				stackContent += " "
			}
			stackContent += curStackContent
		}
		if !logStack {
			break
		}
	}

	//entry.Data["line"] = lineContent
	//if logCaller {
	//	entry.Data["caller"] = callerContent
	//}

	entry.Data["stack"] = stackContent
	entry.Data["type"] = "jlog"
	if hook.logConfig != nil && len(hook.logConfig.serviceLabel) > 0 {
		entry.Data["service"] = hook.logConfig.serviceLabel
	}

	return nil
}

type jloggerT struct {
}

var jloggerInstance jloggerT

func Trace(args ...interface{}) {
	jloggerInstance.trace(args...)
}

func Tracef(format string, args ...interface{}) {
	jloggerInstance.tracef(format, args...)
}

func Traceln(args ...interface{}) {
	jloggerInstance.traceln(args...)
}

func Debug(args ...interface{}) {
	jloggerInstance.debug(args...)
}

func Debugf(format string, args ...interface{}) {
	jloggerInstance.debugf(format, args...)
}

func Debugln(args ...interface{}) {
	jloggerInstance.debugln(args...)
}

func Info(args ...interface{}) {
	jloggerInstance.info(args...)
}

func Infof(format string, args ...interface{}) {
	jloggerInstance.infof(format, args...)
}

func Infoln(args ...interface{}) {
	jloggerInstance.infoln(args...)
}

func Warning(args ...interface{}) {
	jloggerInstance.warning(args...)
}

func Warningf(format string, args ...interface{}) {
	jloggerInstance.warningf(format, args...)
}

func Warningln(args ...interface{}) {
	jloggerInstance.warningln(args...)
}

func Error(args ...interface{}) {
	jloggerInstance.error(args...)
}

func Errorf(format string, args ...interface{}) {
	jloggerInstance.errorf(format, args...)
}

func Errorln(args ...interface{}) {
	jloggerInstance.errorln(args...)
}

func Fatal(args ...interface{}) {
	jloggerInstance.fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	jloggerInstance.fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	jloggerInstance.fatalln(args...)
}

func (j *jloggerT) trace(args ...interface{}) {
	logrus.Trace(args...)
}

func (j *jloggerT) tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

func (j *jloggerT) traceln(args ...interface{}) {
	logrus.Traceln(args...)
}

func (j *jloggerT) debug(args ...interface{}) {
	logrus.Debug(args...)
}

func (j *jloggerT) debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (j *jloggerT) debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

func (j *jloggerT) info(args ...interface{}) {
	logrus.Info(args...)
}

func (j *jloggerT) infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (j *jloggerT) infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

func (j *jloggerT) warning(args ...interface{}) {
	logrus.Warning(args...)
}

func (j *jloggerT) warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func (j *jloggerT) warningln(args ...interface{}) {
	logrus.Warningln(args...)
}

func (j *jloggerT) error(args ...interface{}) {
	logrus.Error(args...)
}

func (j *jloggerT) errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (j *jloggerT) errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

func (j *jloggerT) fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (j *jloggerT) fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (j *jloggerT) fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}
