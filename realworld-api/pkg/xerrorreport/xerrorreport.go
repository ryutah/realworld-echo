package xerrorreport

import (
	"context"
	"fmt"

	"github.com/ryutah/realworld-echo/realworld-api/pkg/xlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Service string

type Version string

type ErrorReporter struct {
	Service ServiceContext
}

func NewErrorReporter(service Service, version Version) *ErrorReporter {
	return &ErrorReporter{
		Service: ServiceContext{
			Service: service,
			Version: version,
		},
	}
}

// ServiceContext is a struct for logging the service context of an error.
// It is used to nest the service context under the "serviceContext" key in the JSON payload.
//
//	e.g.:
//		{
//		  "serviceContext": {
//		    "service": "serviceName",
//		    "version": "version",
//		  }
//		}
type ServiceContext struct {
	Service Service
	Version Version
}

var _ zapcore.ObjectMarshaler = (*ServiceContext)(nil)

func (s ServiceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("service", string(s.Service))
	enc.AddString("version", string(s.Version))
	return nil
}

// ErrorContext is a struct for logging the context of an error.
// It is used to nest the context under the "context" key in the JSON payload.
//
//	e.g.:
//		{
//		  "context": {
//		    "user": "user",
//		    "location": {/* location objects */}
//		  }
//		}
type ErrorContext struct {
	User     string
	Location Location
}

var _ zapcore.ObjectMarshaler = (*ErrorContext)(nil)

func (e ErrorContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("user", e.User)
	return enc.AddObject("location", e.Location)
}

// Location is a struct for logging the location of an error.
// It is used to nest the location under the "location" key in the JSON payload.
//
//	e.g.:
//		{
//		  "location": {
//		    "filePath": "foo.go",
//		    "lineNumber": 10,
//		    "functionName": "foo()"
//		  }
//		}
type Location struct {
	File     string
	Line     int
	Function string
}

var _ zapcore.ObjectMarshaler = (*Location)(nil)

func (l Location) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("filePath", l.File)
	enc.AddInt("lineNumber", l.Line)
	enc.AddString("functionName", l.Function)
	return nil
}

// Report reports an error to Operations Suite ErrorReporting.
// It logs the error to Cloud Logging Error Reporting as an Alert.
//
//	see: https://cloud.google.com/error-reporting/docs/formatting-error-messages?hl=ja#reported-error-example
func (e *ErrorReporter) Report(ctx context.Context, err error, errContext ErrorContext) {
	xlog.Alert(
		ctx,
		fmt.Sprintf("%+v", err),
		zap.String("@type", "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent"),
		zap.Object("serviceContext", e.Service),
		zap.Object("context", errContext),
	)
}
