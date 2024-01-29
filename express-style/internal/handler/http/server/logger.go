package server

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"io"
)

var _ log.AllLogger = (*customLogger)(nil)

type customLogger struct {
	stdLog *log.Logger
}

func (c customLogger) Trace(v ...interface{}) {

	//TODO implement me
	panic("implement me")
}

func (c customLogger) Debug(v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Info(v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Warn(v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Error(v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Fatal(v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Panic(v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Tracef(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Debugf(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Infof(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Warnf(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Errorf(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Fatalf(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Panicf(format string, v ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Tracew(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Debugw(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Infow(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Warnw(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Errorw(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) Panicw(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) SetLevel(level log.Level) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) SetOutput(writer io.Writer) {
	//TODO implement me
	panic("implement me")
}

func (c customLogger) WithContext(ctx context.Context) log.CommonLogger {
	//TODO implement me
	panic("implement me")
}
