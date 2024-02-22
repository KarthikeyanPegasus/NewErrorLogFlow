package errorLogger

import (
	"encoding/json"
	"go.uber.org/zap"
	"runtime"
	"strconv"
	"strings"
)

type Error struct {
	ErrorCode   int    `json:"code"`
	Reason      string `json:"reason"`      // reason is for User
	Description string `json:"description"` // description is for Developer
	Information Info   `json:"info"`
}

type Info struct {
	AppName    string `json:"app_name"`
	Event      string `json:"event"`
	StackTrace string `json:"stack_trace"`
}

type UserError struct {
	ErrorCode int    `json:"code"`
	Reason    string `json:"reason"`
}

func (e *Error) Error() string {
	return strconv.Itoa(e.ErrorCode) + " " + e.Description
}

func Unmarshal(body string, error *Error) error {
	return json.Unmarshal([]byte(body), error)
}

func MarshalCustomError(error *Error) string {
	body, _ := json.Marshal(error)
	return string(body)
}

func MarshalUserError(error *UserError) string {
	body, _ := json.Marshal(error)
	return string(body)
}

//

type Logger struct {
	l *zap.Logger
}

func NewLogger(l *zap.Logger) *Logger {
	return &Logger{l: l}
}

func getCustomReason(appname, event, reason string, errorCode int) string {
	//	"{Appname} - unable to {Event}: {ErrorCode} {Reason}"
	return appname + " - unable to " + event + ": " + strconv.Itoa(errorCode) + " " + reason
}

func (l *Logger) LogError(err *Error) {
	l.l.Error(getCustomReason(err.Information.AppName, err.Information.Event, err.Description, err.ErrorCode),
		zap.String("description", err.Description),
		zap.String("app_name", err.Information.AppName),
		zap.String("stack_trace", err.Information.StackTrace))
}

func NewError(errorCode int, reason, description, appName, event string) *Error {
	stackTrace := [4096]byte{}
	runtime.Stack(stackTrace[:], false)
	stackTraceStr := strings.ReplaceAll(string(stackTrace[:]), "\u0000", "")
	return &Error{
		ErrorCode:   errorCode,
		Reason:      reason,
		Description: description,
		Information: Info{
			AppName:    appName,
			StackTrace: stackTraceStr,
			Event:      event,
		},
	}
}
