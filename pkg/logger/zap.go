package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"

	ModeDev  = "dev"
	ModeProd = "prod"
)

type Config struct {
	Level string
	Mode  string
}

type Logger struct {
	sl *zap.SugaredLogger
}

func New(level, mode string) (ILogger, error) {
	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	var encoder zapcore.Encoder

	if mode == ModeDev {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(logLevel),
	)

	z := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &Logger{
		sl: z.Sugar(),
	}, nil
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sl.Debugw(msg, keysAndValues...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sl.Infow(msg, keysAndValues...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sl.Warnw(msg, keysAndValues...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sl.Errorw(msg, keysAndValues...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sl.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Sync() error {
	if l.sl != nil {
		return l.sl.Sync()
	}
	return nil
}

func (l *Logger) With(args ...interface{}) ILogger {
	newLogger := l.sl.With(args...)

	return &Logger{sl: newLogger}
}

