package jlog

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
	"runtime"
)
var (
	logFilename = "logrus.log"
)

const (
	logLineNumber = true
	logCaller = true
	logStack = true
	logStackCount = 20

	disableGLog = false
)

type ContextHook struct {
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
		if pc, file, line, ok := runtime.Caller(skip+i); ok {

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

	/*entry.Data["line"] = lineContent
	if logCaller {
		entry.Data["caller"] = callerContent
	}*/

	entry.Data["stack"] = stackContent

	return nil
}



func Init(logfile string) {
	logFilename = logfile

	flag.Parse() // glog need log path so we parse when init

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
		TimestampFormat:"2006-01-02 15:04:05",
	})
	logrus.AddHook(ContextHook{})

	logPath := path.Join(flag.Lookup("log_dir").Value.String(), logFilename)
	fmt.Println("Log would be written to: ", logPath)
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     28, //days
		Compress:   true, // disabled by default
	})
}

func Flush() {
	glog.Flush()
}

type jloggerT struct {
}

var jloggerInstance jloggerT

func Error(args ...interface{}) {
	jloggerInstance.error(args...)
}

func Errorf(format string, args ...interface{}) {
	jloggerInstance.errorf(format, args...)
}

func Errorln(args ...interface{}) {
	jloggerInstance.errorln(args...)
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

func Fatal(args ...interface{}) {
	jloggerInstance.fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	jloggerInstance.fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	jloggerInstance.fatalln(args...)
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

func (j *jloggerT) info(args ...interface{}) {
	logrus.Info(args...)
	if !disableGLog {
		glog.Info(args...)
	}
}

func (j *jloggerT) infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
	if !disableGLog {
		glog.Infof(format, args...)
	}
}

func (j *jloggerT) infoln(args ...interface{}) {
	logrus.Infoln(args...)
	if !disableGLog {
		glog.Infoln(args...)
	}
}

func (j *jloggerT) fatal(args ...interface{}) {
	logrus.Fatal(args...)
	if !disableGLog {
		glog.Fatal(args...)
	}
}

func (j* jloggerT) error(args ...interface{}) {
	logrus.Error(args...)
	if !disableGLog {
		glog.Error(args...)
	}
}

func (j* jloggerT) errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
	if !disableGLog {

		glog.Errorf(format, args...)
	}
}

func (j* jloggerT) errorln(args ...interface{}) {
	logrus.Errorln(args...)
	if !disableGLog {
		glog.Errorln(args...)
	}
}

func (j* jloggerT) fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
	if !disableGLog {
		glog.Fatalf(format, args...)
	}
}

func (j* jloggerT) fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
	if !disableGLog {
		glog.Fatalln(args...)
	}
}

func (j* jloggerT) warning(args ...interface{}) {
	logrus.Warning(args...)
	if !disableGLog {
		glog.Warning(args...)
	}
}

func (j* jloggerT) warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
	if !disableGLog {
		glog.Warningf(format, args...)
	}
}

func (j* jloggerT) warningln(args ...interface{}) {
	logrus.Warningln(args...)
	if !disableGLog {
		glog.Warningln(args...)
	}
}
