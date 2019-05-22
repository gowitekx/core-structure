package configs

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

const (
	REQUESTID = "requestID"
	LOGDATA   = "logdata"
	ERROR     = "error"
	INFO      = "info"
	WARN      = "warn"
	DEBUG     = "debug"
)

var NUMBERRUNES = []rune("0123456789")

type Logdata struct {
	logger *log.Logger
}

var Ld Logdata

func init() {
	Ld.logger = log.New()
	Ld.logger.SetLevel(log.ErrorLevel)
	Ld.logger.Formatter = &log.TextFormatter{}

}

//Logger with fields
func (ld Logdata) Logger(ctx context.Context, errortype string, args ...interface{}) {
	currentTime := time.Now()
	var filename string = Config.Logfile + currentTime.Format("20060102") + ".log"
	//Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	ld.logger.SetOutput(f)
	var depth = 1
	var requestid string
	ctxRequestID := ctx.Value(REQUESTID)
	//Get Caller function
	function, file, line, _ := runtime.Caller(depth)
	functionObject := runtime.FuncForPC(function)
	switch errortype {
	case "error":
		ld.logger.WithFields(log.Fields{
			"requestid": ctxRequestID,
			"file":      file,
			"function":  functionObject.Name(),
			"line":      line,
		}).Error(args...)
	case "info":
		ld.logger.WithFields(log.Fields{
			"requestid": requestid,
			"file":      file,
			"function":  functionObject.Name(),
			"line":      line,
		}).Info(args...)
	case "warn":
		ld.logger.WithFields(log.Fields{
			"requestid": requestid,
			"file":      file,
			"function":  functionObject.Name(),
			"line":      line,
		}).Warn(args...)
	case "debug":
		ld.logger.WithFields(log.Fields{
			"requestid": requestid,
			"file":      file,
			"function":  functionObject.Name(),
			"line":      line,
		}).Debug(args...)
	default:
		ld.logger.WithFields(log.Fields{
			"requestid": requestid,
			"file":      file,
			"function":  functionObject.Name(),
			"line":      line,
		}).Error(args...)
	}

}

//WithRequestID returns a context which knows its request ID
func WithRequestID(ctx context.Context) context.Context {
	//Generate UUID
	u2 := uuid.NewV4()
	// fmt.Println(u2)
	return context.WithValue(ctx, REQUESTID, u2)
}
