package logging

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	logFormatText    = "text"
	logFormatJSON    = "json"
	logFormatDefault = logFormatJSON
)

var (
	logger      = logrus.New()
	requestID   string
	serviceName string
)

func init() {
	logFormat := os.Getenv("MAILING_SERVICE_LOG_FORMAT")

	if logFormat == "" {
		fmt.Printf("Logging value not defined, setting default logging format: '%s'\n", logFormatDefault)
		logFormat = logFormatDefault
	}

	if logFormat != logFormatJSON || logFormat != logFormatText {
		fmt.Printf("Unsupported Logging format: %s, setting default logging format: %s\n", logFormat, logFormatDefault)
	}

	if logFormat == logFormatJSON {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(UTCFormatter{
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
				PrettyPrint:     false,
			},
		})
	} else {
		// Log as text.
		logger.SetFormatter(UTCFormatter{
			Formatter: &logrus.TextFormatter{
				DisableColors:   false,
				TimestampFormat: time.RFC3339,
				FullTimestamp:   true,
			},
		})
	}

	// Default log level
	logger.SetLevel(logrus.DebugLevel)

	evnLogLevel := os.Getenv("MAILING_SERVICE_LOG_LEVEL")

	switch strings.ToLower(evnLogLevel) {
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	}
	fmt.Printf("Logging configured with level: %s, format: %s\n", logger.GetLevel(), logFormat)
}

// Info log messsage
func Info(logDetails LogDetails) {
	DefaultWithFields(LogDetails{
		Code:    logDetails.Code,
		Details: logDetails.Details,
	}).Info(logDetails.Message)
}

// Infof allows you to log formatted strings.
func Infof(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	DefaultWithFields(LogDetails{}).Info(str)
}

// Error log message with fields
func Error(errLog LogDetails) {
	if errLog.Code == "" {
		logger.Error("Error code is required for method Error in logger")
		logger.Panic(errLog.Error)
	} else if errLog.Message == "" {
		logger.Error("Error message is required for method Error in logger")
		logger.Panic(errLog.Error)
	}

	var stack errors.StackTrace
	error, ok := errLog.Error.(stackTracer)
	if ok {
		stack = error.StackTrace()
	}
	DefaultWithFields(errLog).WithFields(logrus.Fields{
		"err":        errLog.Error,
		"stackTrace": fmt.Sprint(stack),
	}).Error(errLog.Message)
}

// Errorf allows you to log formatted strings.
func Errorf(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	DefaultWithFields(LogDetails{}).Error(str)
}

// Fatal log message
func Fatal(logDetials LogDetails) {
	DefaultWithFields(LogDetails{
		Code:    logDetials.Code,
		Details: logDetials.Details,
	}).Fatal(logDetials.Message)
}

// Fatalf allows you to log formatted strings.
func Fatalf(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	DefaultWithFields(LogDetails{}).Fatal(str)
}

// Trace returns the source code line and function name (of the calling function)
func Trace() (line string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return fmt.Sprintf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

//DefaultWithFields adds default fields to the logger
func DefaultWithFields(logDetails LogDetails) *logrus.Entry {
	pc, _, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	funcName := ""
	if ok && details != nil {
		//fmt.Printf("called from %s\n", details.Name())
		funcName = details.Name()
	}
	return logger.WithFields(logrus.Fields{
		"X-REQUEST-ID": GetServiceRequestID(),
		"functionName": funcName,
		"serviceName":  serviceName,
		"line":         line,
		"code":         logDetails.Code,
		"details":      logDetails.Details,
	})
}

//GetServiceRequestID ..
func GetServiceRequestID() string {
	if requestID == "" {
		requestID = GetRequestID(nil)
		return requestID
	}
	return requestID
}

// GetRequestID is function to generate uuid as request id if client doesn't pass X-REQUEST-ID request header
func GetRequestID(requestIDParams *string) string {
	log.Debug("entering func getRequestID")
	//generate UUID as request ID if it doesn't exist in request header
	if requestIDParams == nil || *requestIDParams == "" {
		theUUID, err := uuid.NewUUID()
		newUUID := ""
		if err == nil {
			newUUID = theUUID.String()
		} else {
			newUUID = GenerateUUID()
		}
		requestIDParams = &newUUID
	}
	return *requestIDParams
}

// GenerateUUID is function to generate our own uuid if the google uuid throws error
func GenerateUUID() string {
	log.Info("entering func generateUUID")
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Error(Trace(), err)
		return ""
	}
	theUUID := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return theUUID
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
