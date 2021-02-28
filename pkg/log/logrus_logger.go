package log

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

var _ Logger = (*logrusLogger)(nil)

var Levels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

// logrusLogger .
type logrusLogger struct {
	entry *logrus.Entry
	pool  *sync.Pool
	level logrus.Level
}

type Option func(*logrusLogger)

// NewLogrusLogger returns a log.Logger that sends log events to a logrus.Logger.
func NewLogrusLogger(w io.Writer, opts ...Option) Logger {
	logger := logrus.New()
	logger.SetOutput(w)
	// logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true}
	e := logrus.NewEntry(logger)

	l := &logrusLogger{
		entry: e,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	for _, fn := range opts {
		fn(l)
	}

	return l
}

// WithLevel configures a logrus logger to log at level for all events.
func WithLevel(level logrus.Level) Option {
	return func(c *logrusLogger) {
		c.level = level
	}
}

func (l *logrusLogger) Print(pairs ...interface{}) {
	if len(pairs) == 0 {
		return
	}
	if len(pairs)%2 != 0 {
		pairs = append(pairs, "")
	}
	buf := l.pool.Get().(*bytes.Buffer)
	for i := 0; i < len(pairs); i += 2 {
		fmt.Fprintf(buf, "%s=%v ", pairs[i], pairs[i+1])
	}

	switch l.level {
	case logrus.InfoLevel:
		l.entry.Info(buf.String())
	case logrus.ErrorLevel:
		l.entry.Error(buf.String())
	case logrus.DebugLevel:
		l.entry.Debug(buf.String())
	case logrus.WarnLevel:
		l.entry.Warn(buf.String())
	default:
		l.entry.Info(buf.String())
	}

	buf.Reset()
	l.pool.Put(buf)
}
