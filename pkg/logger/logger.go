package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return ""
	}
}

//NewLogger init logger.
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

//Clone a copy of Logger
func (l *Logger) Clone() *Logger {
	nl := *l
	return &nl
}

//WithFields 设置当前字段.
func (l *Logger) WithFields(f Fields) *Logger {
	nl := l.Clone()
	if nl.fields == nil {
		nl.fields = make(Fields)
	}
	for k, v := range f {
		nl.fields[k] = v
	}
	return nl
}

//WithContex 设置当前日志内容
func (l *Logger) WithContex(ctx context.Context) *Logger {
	nl := l.Clone()
	nl.ctx = ctx
	return nl
}

//WithCaller 设置当前调用栈信息.
func (l *Logger) WithCaller(skip int) *Logger {
	nl := l.Clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		nl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return nl
}

//WithCallersFrames 设置当前所有调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frams := runtime.CallersFrames(pcs[:depth])
	for frame, more := frams.Next(); more; frame, more = frams.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}

	nl := l.Clone()
	nl.callers = callers
	return nl
}

//JSONFormat .
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

//WithTrace 打印trace追踪信息
func (l *Logger) WithTrace() *Logger {
	// ginCtx, ok := l.ctx.(*gin.Context)
	// if ok {
	// 	return l.WithFields(Fields{
	// 		"trace_id": ginCtx.MustGet("X-Trace-ID"),
	// 		"span_id":  ginCtx.MustGet("X-Span-ID"),
	// 	})
	// }
	return l
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
		break
	case LevelInfo:
		l.newLogger.Print(content)
		break
	case LevelWarn:
		l.newLogger.Print(content)
		break
	case LevelError:
		l.newLogger.Print(content)
		break
	case LevelFatal:
		l.newLogger.Fatal(content) //fatal会导致线程退出
		break
	case LevelPanic:
		l.newLogger.Panic(content)
		break
	}
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.WithContex(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}
