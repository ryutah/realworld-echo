package xerrorreport

import (
	"context"
	"fmt"

	"github.com/ryutah/realworld-echo/pkg/xlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ErrorReporter struct {
	Service ServiceContext
}

func NewErrorReporter(service, version string) *ErrorReporter {
	return &ErrorReporter{
		Service: ServiceContext{
			Service: service,
			Version: version,
		},
	}
}

type ServiceContext struct {
	Service string
	Version string
}

var _ zapcore.ObjectMarshaler = (*ServiceContext)(nil)

func (s ServiceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("service", s.Service)
	enc.AddString("version", s.Version)
	return nil
}

type ErrorContext struct {
	User     string
	Location Location
}

var _ zapcore.ObjectMarshaler = (*ErrorContext)(nil)

func (e ErrorContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("user", e.User)
	return enc.AddObject("location", e.Location)
}

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

func (e *ErrorReporter) Report(ctx context.Context, err error, errContext ErrorContext) {
	xlog.Alert(
		ctx,
		fmt.Sprintf("%+v", err),
		zap.String("@type", "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent"),
		zap.Object("serviceContext", e.Service),
		zap.Object("context", errContext),
	)
}
